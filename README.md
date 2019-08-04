# addressbook

[![Build Status](https://travis-ci.com/dikaeinstein/addressbook.svg?branch=master)](https://travis-ci.com/dikaeinstein/addressbook)
[![Coverage Status](https://coveralls.io/repos/github/dikaeinstein/addressbook/badge.svg?branch=master)](https://coveralls.io/github/dikaeinstein/addressbook?branch=master)

A CLI tool built with protocol buffers while following the [Protocol Buffer tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial).

It is used to add contact details to an address book. It uses the protocol buffer to encode and decode contact details from the address book.

## Subcommands

* listPeople - lists the contacts in the address book.
* addPerson - adds a new contact to the address book.

### Usage

./addressbook [sub-command] -file [addressbookfileName]
