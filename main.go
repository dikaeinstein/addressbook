package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dikaeinstein/addressbook/cmd"
)

func main() {
	// var fName = flag.String("file", "", "-file [filename]")
	listPeopleCommand := flag.NewFlagSet("listPeople", flag.ExitOnError)
	var lName = listPeopleCommand.String("file", "", "-file [filename]")
	addPersonCommand := flag.NewFlagSet("addPerson", flag.ExitOnError)
	var aName = addPersonCommand.String("file", "", "-file [filename]")

	l := cmd.ListPeople{Output: os.Stdout}
	ap := cmd.AddPerson{Input: os.Stdin}

	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s [cmd] -file", os.Args[0])
	}

	switch os.Args[1] {
	case "listPeople":
		listPeopleCommand.Parse(os.Args[2:])
		l.Execute(*lName)
	case "addPerson":
		addPersonCommand.Parse(os.Args[2:])
		ap.Execute(*aName)
	default:
		fmt.Println("Invalid command!")
		os.Exit(2)
	}
}
