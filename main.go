package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"

	"github.com/pharosnet/gitcgen/action"
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
	app.Version = "v1.2.1"
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

	p, pErr := getWorkTree(wtp)
	if pErr != nil {
		err = pErr
		return
	}

	err = action.Show(p)

	return
}

func generate(c *cli.Context) (err error) {

	wtp := c.String(workTreeFlag)
	p, pErr := getWorkTree(wtp)
	if pErr != nil {
		err = pErr
		return
	}

	output := c.String(outputFlag)

	if output == "" {
		err = fmt.Errorf("%s's value is empty", outputFlag)
		return
	}

	short := c.Bool(shortFlag)

	err = action.Generate(p, short, output)

	return
}

func getWorkTree(wtp string) (p string, err error) {
	if wtp == "" {
		err = fmt.Errorf("%s's value is empty", workTreeFlag)
		return
	}

	wtp = path.Clean(wtp)

	if _, err0 := os.Stat(wtp); os.IsNotExist(err0) {
		err = fmt.Errorf("dir: %s is not exist", p)
		return
	}

	 p = wtp

	return
}
