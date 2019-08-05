package command

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	addressbook "github.com/dikaeinstein/addressbook/proto"
	"github.com/golang/protobuf/proto"
)

type errReader struct {
	io.Reader
	err error
}

func (e *errReader) Read(p []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Reader.Read(p)
	return n, nil
}

// AddPerson adds a person to the address book
type AddPerson struct {
	Input io.Reader
}

// Execute executes the AddPerson command
func (ap *AddPerson) Execute(fName string) {
	// Read the existing address book
	_, err := ioutil.ReadFile(fName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found. Creating a new file.\n", fName)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	book := &addressbook.AddressBook{}

	addr, err := promptForAddress(ap.Input)
	if err != nil {
		log.Fatalln("Error with address:", err)
	}
	book.People = append(book.People, addr)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	f, err := os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.Write(out); err != nil {
		log.Fatalln("Failed to write to address book:", err)
	}
}

func promptForName(rd *bufio.Reader, p *addressbook.Person) error {
	fmt.Println("Enter name: ")

	name, err := rd.ReadString('\n')
	if err != nil {
		return err
	}
	p.Name = strings.TrimSpace(name)

	return nil
}

func promptForID(rd *bufio.Reader, p *addressbook.Person) error {
	fmt.Println("Enter person ID number: ")

	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return err
	}

	return nil
}

func promptForEmail(rd *bufio.Reader, p *addressbook.Person) error {
	fmt.Print("Enter email address (blank for none): ")

	email, err := rd.ReadString('\n')
	if err != nil {
		return err
	}
	p.Email = strings.TrimSpace(email)

	return nil
}

func promptForPhone(rd *bufio.Reader, p *addressbook.Person) error {
	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")

		phone, err := rd.ReadString('\n')
		if err != nil {
			return err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		pn := &addressbook.Person_PhoneNumber{Number: phone}

		fmt.Print("Is this a mobile, home, or work phone? ")
		pType, err := rd.ReadString('\n')
		if err != nil {
			return err
		}
		pType = strings.TrimSpace(pType)

		switch pType {
		case "home":
			pn.Type = addressbook.Person_HOME
		case "mobile":
			pn.Type = addressbook.Person_MOBILE
		case "work":
			pn.Type = addressbook.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q. Using default.\n", pType)
		}

		p.Phones = append(p.Phones, pn)
	}

	return nil
}

func promptForAddress(r io.Reader) (*addressbook.Person, error) {
	er := &errReader{Reader: r}
	rd := bufio.NewReader(er)

	p := &addressbook.Person{}

	promptForID(rd, p)
	promptForName(rd, p)
	promptForEmail(rd, p)
	promptForPhone(rd, p)

	return p, er.err
}
