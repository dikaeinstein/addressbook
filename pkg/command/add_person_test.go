package command

import (
	"bufio"
	"strings"
	"testing"

	addressbook "github.com/dikaeinstein/addressbook/proto"
	"github.com/golang/protobuf/proto"
)

func setup(in string) (*bufio.Reader, addressbook.Person) {
	p := addressbook.Person{}
	return bufio.NewReader(strings.NewReader(in)), p
}

func TestPromptForName(t *testing.T) {
	in := `Example Name
`
	rd, p := setup(in)

	err := promptForName(rd, &p)
	if err != nil {
		t.Fatalf("promptForName(%s) had unexpected error: %v", in, err)
	}
	if p.Name != "Example Name" {
		t.Errorf("promptForName(%s) => %q; want %q", in, p.Name, "Example Name")
	}
}

func TestPromptForID(t *testing.T) {
	in := `1234
`
	rd, p := setup(in)

	err := promptForID(rd, &p)
	if err != nil {
		t.Fatalf("promptForID(%q) had unexpected error: %v", in, err)
	}
	if p.Id != 1234 {
		t.Errorf("promptForID(%q) => %d; want %d", in, p.Id, 1234)
	}
}

func TestPromptForEmail(t *testing.T) {
	in := `dika@example.com
`
	rd, p := setup(in)

	err := promptForEmail(rd, &p)
	if err != nil {
		t.Fatalf("promptForEmail(%q) had unexpected error: %v", in, err)
	}
	if p.Email != "dika@example.com" {
		t.Errorf("promptForEmail(%q) => %q; want %q", in, p.Email, "dika@example.com")
	}
}

func TestPromptForPhone(t *testing.T) {
	in := `123-456-7890
home
222-222-2222
mobile
111-111-1111
work
777-777-7777
unknown

`
	rd, p := setup(in)

	want := []*addressbook.Person_PhoneNumber{
		{Number: "123-456-7890", Type: addressbook.Person_HOME},
		{Number: "222-222-2222", Type: addressbook.Person_MOBILE},
		{Number: "111-111-1111", Type: addressbook.Person_WORK},
		{Number: "777-777-7777", Type: addressbook.Person_MOBILE},
	}

	err := promptForPhone(rd, &p)

	if err != nil {
		t.Fatalf("promptForPhone(%q) had unexpected error: %v", in, err)
	}
	if len(p.Phones) != len(want) {
		t.Errorf("want %d phone numbers, got %d", len(want), len(p.Phones))
	}
	phones := len(p.Phones)
	if phones > len(want) {
		phones = len(want)
	}
	for i := 0; i < phones; i++ {
		if !proto.Equal(p.Phones[i], want[i]) {
			t.Errorf("want phone %q, got %q", *want[i], *p.Phones[i])
		}
	}
}
