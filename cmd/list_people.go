package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	addressbook "github.com/dikaeinstein/addressbook/proto"
)

func writePerson(w io.Writer, p *addressbook.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case addressbook.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case addressbook.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case addressbook.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func listPeople(w io.Writer, book *addressbook.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

// ListPeople lists people in the addressbook
type ListPeople struct {
	Output io.Writer
}

// Execute executes the ListPeople command
func (l *ListPeople) Execute(fName string) {
	// Read the existing address book.
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	var book addressbook.AddressBook
	// Unmarshal proto
	if err := proto.Unmarshal(in, &book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	listPeople(l.Output, &book)
}
