package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	"github.com/codegangsta/cli"
	"github.com/poy/ledger/aggregators"
	"github.com/poy/ledger/database"
	"github.com/poy/ledger/filters"
	"github.com/poy/ledger/transaction"
)

const (
	FileLocation = "file"

	ReportName  = "report"
	ReportUsage = "[START DATE YYYY/MM/DD] [END DATE YYYY/MM/DD]"

	FmtName  = "fmt"
	FmtUsage = "[FILE NAME]"

	GroupName  = "group"
	GroupUsage = ReportUsage
)

func main() {
	app := cli.NewApp()
	app.Name = "ledger"
	app.Usage = ReportUsage
	app.Action = report
	app.Flags = buildReportFlags()
	app.Commands = []cli.Command{
		{
			Name:   "aggregators",
			Usage:  "Lists the available aggregators",
			Action: listAggregators,
		},
		{
			Name:   GroupName,
			Usage:  GroupUsage,
			Action: grouping,
			Flags:  buildReportFlags(),
		},
		{
			Name:   FmtName,
			Usage:  FmtUsage,
			Action: fmtFile,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "write",
					Usage: "Write formatted file",
				},
			},
		},
	}

	app.Run(os.Args)
}

func buildFileFlag() cli.Flag {
	return cli.StringFlag{
		Name:  FileLocation,
		Usage: "The path to the ledger transaction file",
	}
}

func buildReportFlags() []cli.Flag {
	var results []cli.Flag
	results = append(results, buildFileFlag())
	results = append(results, buildFilterFlags()...)
	results = append(results, cli.StringSliceFlag{
		Name:  "agg",
		Usage: "name of aggregator",
		Value: new(cli.StringSlice),
	})

	return results
}

func buildFilterFlags() []cli.Flag {
	var results []cli.Flag
	for name := range filters.Store() {
		results = append(results, cli.StringSliceFlag{
			Name:  fmt.Sprintf("filter-%s", name),
			Value: new(cli.StringSlice),
		})
	}
	return results
}

func listAggregators(c *cli.Context) {
	for name := range aggregators.Store() {
		fmt.Println(name)
	}
}

func report(c *cli.Context) {
	if len(c.Args()) != 2 {
		fatal("", c)
	}

	fileLocation := c.String(FileLocation)
	if len(fileLocation) == 0 {
		fatal(fmt.Sprintf("Missing required '--%s' flag\n", FileLocation), c)
	}

	filter := buildFilter(c)

	start := loadDate(c.Args().Get(0))
	end := loadDate(c.Args().Get(1))

	db := new(database.Database)
	loadDatabase(fileLocation, db.Add)
	emitter := func(v string) {
		println(v)
	}

	aggNameSlice := c.StringSlice("agg")
	if len(aggNameSlice) == 0 {
		printResults(db.Query(start, end, filter), emitter)
		return
	}

	aggs, err := aggregators.Fetch(aggNameSlice...)
	fatalErr(err)
	results, aggResults := db.Aggregate(start, end, filter, aggs...)
	printResults(results, emitter)

	fmt.Println("===============")
	fmt.Println()
	for i, aggResult := range aggResults {
		fmt.Printf("%s = %s\n", aggNameSlice[i], aggResult)
	}
}

func grouping(c *cli.Context) {
	if len(c.Args()) != 2 {
		fatal("", c)
	}

	fileLocation := c.String(FileLocation)
	if len(fileLocation) == 0 {
		fatal(fmt.Sprintf("Missing required '--%s' flag\n", FileLocation), c)
	}

	filter := buildFilter(c)

	start := loadDate(c.Args().Get(0))
	end := loadDate(c.Args().Get(1))

	aggNameSlice := c.StringSlice("agg")
	if len(aggNameSlice) == 0 {
		fatal("Missing required '--agg' flag(s)\n", c)
	}

	grouping := database.NewGrouping()
	loadDatabase(fileLocation, grouping.Add)

	aggs, err := aggregators.Fetch(aggNameSlice...)
	fatalErr(err)
	results := grouping.Aggregate(start, end, filter, aggs...)

	for accName, value := range results {
		println(accName)
		for i, r := range value {
			fmt.Printf("\t%s = %s\n", aggNameSlice[i], r)
		}
		println()
	}
}

func fmtFile(c *cli.Context) {
	if len(c.Args()) != 1 {
		fatal(fmt.Sprintf("Usage %s %s\n\nfmt is used to normalize the spacing and sort your ledger file\n", FmtName, FmtUsage), c)
	}

	fileName := c.Args().First()
	results := readTransactions(fileName)

	emitter := func(v string) {
		println(v)
	}

	if c.Bool("write") {
		writer, err := os.Create(fileName)
		fatalErr(err)
		defer writer.Close()
		emitter = func(v string) {
			_, err = writer.Write([]byte(v + "\n"))
			fatalErr(err)
		}
	}

	printResults(results, emitter)
}

func readTransactions(fileName string) []*transaction.Transaction {
	file := openFile(fileName)
	defer file.Close()
	reader := transaction.NewReader(file)
	var results []*transaction.Transaction
	for {
		t, err := reader.Next()
		if err != nil {
			fatalErr(err)
		}

		if t == nil {
			break
		}
		results = append(results, t)
	}
	return results
}

func buildFilter(c *cli.Context) database.Filter {
	var results []database.Filter
	for name, factory := range filters.Store() {
		filterArgs := c.StringSlice(fmt.Sprintf("filter-%s", name))
		if len(filterArgs) <= 0 {
			continue
		}

		for _, arg := range filterArgs {
			builtFilter, err := factory.Generate(arg)
			fatalErr(err)
			results = append(results, builtFilter)
		}
	}

	if len(results) == 0 {
		return nil
	}

	return database.CombineFilters(results...)
}

func printResults(results []*transaction.Transaction, emit func(string)) {
	var width int
	for _, t := range results {
		mw := t.MinimumWidth()
		if mw > width {
			width = mw
		}
	}

	sort.Sort(transaction.TransactionList(results))
	for _, t := range results {
		emit(t.Format(width + 3))
	}
}

func loadDate(value string) time.Time {
	date, extra, err := transaction.ParseDate(value)
	if err != nil {
		fatalErr(err)
	}

	if len(extra) > 0 {
		fatalErr(fmt.Errorf("Invalid date %s", value))
	}

	return date
}

func openFile(path string) io.ReadCloser {
	file, err := os.Open(path)
	fatalErr(err)
	return file
}

func loadDatabase(path string, f func(t ...*transaction.Transaction)) {
	file := openFile(path)

	reader := transaction.NewReader(file)
	for {
		t, err := reader.Next()
		if err != nil {
			fatalErr(err)
		}
		if t == nil {
			break
		}
		f(t)
	}
}

func fatalErr(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error())
	os.Exit(1)
}

func fatal(msg string, c *cli.Context) {
	fmt.Println(msg)
	cli.ShowAppHelp(c)
	os.Exit(1)
}
