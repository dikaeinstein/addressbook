package command

import (
	"bytes"
	"strings"
	"testing"

	addressbook "github.com/dikaeinstein/addressbook/proto"
)

func TestWritePersonWritesPerson(t *testing.T) {
	buf := new(bytes.Buffer)

	p := addressbook.Person{
		Id:    1234,
		Name:  "Dika Okwa",
		Email: "dokwa@example.com",
		Phones: []*addressbook.Person_PhoneNumber{
			{Number: "07089585754", Type: addressbook.Person_MOBILE},
		},
	}

	writePerson(buf, &p)

	got := buf.String()
	want := `Person ID: 1234
  Name: Dika Okwa
  E-mail address: dokwa@example.com
  Mobile phone #: 07089585754
`
	if got != want {
		t.Errorf("writePerson(%s) => \n\t%q; want %q", p.String(), got, want)
	}
}

func TestListPeopleListPeople(t *testing.T) {
	buf := new(bytes.Buffer)
	ab := addressbook.AddressBook{
		People: []*addressbook.Person{
			{
				Name:  "John Doe",
				Id:    101,
				Email: "john@example.com",
			},
			{
				Name: "Jane Doe",
				Id:   102,
			},
			{
				Name:  "Jack Doe",
				Id:    201,
				Email: "jack@example.com",
				Phones: []*addressbook.Person_PhoneNumber{
					{Number: "555-555-5555", Type: addressbook.Person_WORK},
				},
			},
			{
				Name:  "Jack Buck",
				Id:    301,
				Email: "buck@example.com",
				Phones: []*addressbook.Person_PhoneNumber{
					{Number: "555-555-0000", Type: addressbook.Person_HOME},
					{Number: "555-555-0001", Type: addressbook.Person_MOBILE},
					{Number: "555-555-0002", Type: addressbook.Person_WORK},
				},
			},
			{
				Name:  "Janet Doe",
				Id:    1001,
				Email: "janet@example.com",
				Phones: []*addressbook.Person_PhoneNumber{
					{Number: "555-777-0000"},
					{Number: "555-777-0001", Type: addressbook.Person_HOME},
				},
			},
		},
	}

	listPeople(buf, &ab)

	want := strings.Split(`Person ID: 101
  Name: John Doe
  E-mail address: john@example.com
Person ID: 102
  Name: Jane Doe
Person ID: 201
  Name: Jack Doe
  E-mail address: jack@example.com
  Work phone #: 555-555-5555
Person ID: 301
  Name: Jack Buck
  E-mail address: buck@example.com
  Home phone #: 555-555-0000
  Mobile phone #: 555-555-0001
  Work phone #: 555-555-0002
Person ID: 1001
  Name: Janet Doe
  E-mail address: janet@example.com
  Mobile phone #: 555-777-0000
  Home phone #: 555-777-0001
`, "\n")
	got := strings.Split(buf.String(), "\n")

	if len(got) != len(want) {
		t.Errorf("ListPeople(%s) => \n\t%q has %d lines; want %d",
			ab.String(),
			buf.String(),
			len(got),
			len(want),
		)
	}

	lines := len(got)
	if lines > len(want) {
		lines = len(want)
	}
	for i := 0; i < lines; i++ {
		if want[i] != got[i] {
			t.Errorf("ListPeople(%s) => \n\tline %d %q; want %q",
				ab.String(),
				i,
				got[i],
				want[i],
			)
		}
	}
}
