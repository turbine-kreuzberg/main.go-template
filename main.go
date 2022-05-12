package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	gitHash string
	gitRef  string
	message = &cli.StringFlag{
		Name:  "message",
		Value: "Hello, %v.",
		Usage: "Welcomingmessage",
	}
	newLine = &cli.BoolFlag{
		Name:  "newLine",
		Value: true,
		Usage: "Add a new line at the end of a line.",
	}
	app = &cli.App{
		Name:   "Application name",
		Usage:  "A short discription of this application.",
		Action: run,
		Flags: []cli.Flag{
			message,
			newLine,
		},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run the application.",
				Action: run,
			},
			{
				Name:   "version",
				Usage:  "Print the version.",
				Action: version,
			},
		},
	}
)

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	newLine := c.Bool(newLine.Name)
	names := c.Args().Slice()

	for _, name := range names {
		fmt.Printf(c.String(message.Name), name)
		if newLine {
			fmt.Println()
		}
	}

	return nil
}

func version(c *cli.Context) error {
	_, err := fmt.Printf("version: %s\ngit commit: %s\n", gitRef, gitHash)
	if err != nil {
		return err
	}

	return nil
}
