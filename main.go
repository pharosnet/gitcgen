package main

import (
	"fmt"
	"github.com/pharosnet/gitcgen/action"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	workTreeFlag = "work_tree"
	outputFlag   = "output"
	shortFlag    = "short"
)

var initShowFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     workTreeFlag,
		Aliases:  []string{"w"},
		Usage:    "git project's worktree path, default is ./",
		Required: false,
		Value:    "./",
	},
}

var initGenFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     workTreeFlag,
		Aliases:  []string{"w"},
		Usage:    "git project's worktree path, default is ./",
		Required: false,
		Value:    "./",
	},
	&cli.StringFlag{
		Name:     outputFlag,
		Aliases:  []string{"o"},
		Required: true,
		Value:    "",
		Usage:    "generate go file path, such as ./foo/bar.go",
	},
	&cli.BoolFlag{
		Name:     shortFlag,
		Aliases:  []string{"s"},
		Required: false,
		Value:    true,
		Usage:    "short git commit id",
	},
}

func main() {
	app := cli.NewApp()
	app.Version = "v1.2.0"
	app.Usage = "Automatically generate git commit id code file for Go."
	app.Commands = []*cli.Command{
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "show latest git commit id",
			Action:  show,
			Flags:   initShowFlags,
		},
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "generate a go file contains git commit id",
			Action:  generate,
			Flags:   initGenFlags,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func show(c *cli.Context) (err error) {

	wtp := c.String(workTreeFlag)

	if wtp == "" {
		err = fmt.Errorf("%s's value is empty", workTreeFlag)
		return
	}

	err = action.Show(wtp)

	return
}

func generate(c *cli.Context) (err error) {

	wtp := c.String(workTreeFlag)

	if wtp == "" {
		err = fmt.Errorf("%s's value is empty", workTreeFlag)
		return
	}

	output := c.String(outputFlag)

	if output == "" {
		err = fmt.Errorf("%s's value is empty", outputFlag)
		return
	}

	short := c.Bool(shortFlag)

	err = action.Generate(wtp, short, output)

	return
}
