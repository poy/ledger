package main

import (
	"fmt"
	"github.com/apoydence/ledger/aggregators"
	"github.com/apoydence/ledger/database"
	"github.com/apoydence/ledger/filters"
	"github.com/apoydence/ledger/transaction"
	"github.com/codegangsta/cli"
	"os"
)

const (
	FileLocation = "file"
	RegexpUsage  = "[PATTERN] [START DATE YYYY/MM/DD] [END DATE YYYY/MM/DD]"
)

func main() {
	app := cli.NewApp()
	app.Name = "ledger"
	app.Usage = "Reconcile your cash"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  FileLocation,
			Usage: "The path to the ledger transaction file",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "regexp",
			Usage:  RegexpUsage,
			Action: regexpFilter,
			Flags: append(app.Flags, cli.StringSliceFlag{
				Name:  "agg",
				Usage: "name of aggregator",
				Value: new(cli.StringSlice),
			}),
		},
		{
			Name:   "aggregators",
			Usage:  "Lists the available aggregators",
			Action: listAggregators,
		},
	}

	app.Run(os.Args)
}

func listAggregators(c *cli.Context) {
	for name := range aggregators.Store() {
		fmt.Println(name)
	}
}

func regexpFilter(c *cli.Context) {
	fileLocation := c.String(FileLocation)
	if len(fileLocation) == 0 {
		fatal(fmt.Sprintf("Missing required '--%s' flag", FileLocation))
	}

	if len(c.Args()) != 3 {
		fatal(fmt.Sprintf("Usage:\nledger regexp %s", RegexpUsage))
	}

	start := loadDate(c.Args().Get(0))
	end := loadDate(c.Args().Get(1))

	filter, err := filters.NewRegexp(c.Args().Get(2))
	if err != nil {
		fatalErr(err)
	}

	db := loadDatabase(fileLocation)

	aggNameSlice := c.StringSlice("agg")
	if len(aggNameSlice) == 0 {
		printResults(db.Query(start, end, filter))
		return
	}

	aggs, err := aggregators.Fetch(aggNameSlice...)
	fatalErr(err)
	results, aggResults := db.Aggregate(start, end, filter, aggs...)
	printResults(results)

	fmt.Println("===============\n")
	for i, aggResult := range aggResults {
		fmt.Printf("%s = $%-6.2f\n", aggNameSlice[i], aggResult)
	}
}

func printResults(results []*transaction.Transaction) {
	for _, t := range results {
		fmt.Println(t)
	}
}

func loadDate(value string) *transaction.Date {
	d := new(transaction.Date)
	extra, err := d.Parse(value)
	if err != nil {
		fatalErr(err)
	}

	if len(extra) > 0 {
		fatal(fmt.Sprintf("Invalid date %s", value))
	}

	return d
}

func loadDatabase(path string) *database.Database {
	file, err := os.Open(path)
	if err != nil {
		fatalErr(err)
	}

	db := database.New()
	reader := transaction.NewReader(file)
	for {
		t, err := reader.Next()
		if err != nil {
			fatalErr(err)
		}
		if t == nil {
			break
		}
		db.Add(t)
	}

	return db
}

func fatalErr(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error())
	os.Exit(1)
}

func fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
