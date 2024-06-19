package main

import (
	"flag"
	"fmt"
	"os"

  utils "go/cheatsheet/utils"
)

func main() {
	eFlag := flag.String("e", "", "Create new cheatsheet")
	sFlag := flag.String("s", "", "Search")
	lFlag := flag.Bool("l", false, "List all cheatsheets")

	// Parse the flags
	flag.Parse()
	cheatPath, err := utils.ReadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	switch {
	case *eFlag != "":
		err = utils.OpenEditor(cheatPath.Path + *eFlag + ".md")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *sFlag != "":
		err = utils.PrintFile(cheatPath.Path + *sFlag + ".md")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *lFlag:
		err = utils.ListFiles()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
