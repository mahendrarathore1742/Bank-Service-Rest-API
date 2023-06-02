package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname string, lname string, pw string) *Account {

	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Account :=> ", acc.Number)
	fmt.Println("New ID :=> ", acc.ID)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Mahendra", "Singh", "msr123456")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seed the database")
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.Run()
}
