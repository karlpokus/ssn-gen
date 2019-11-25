package main

import (
	"log"
	"os"

	"github.com/karlpokus/srv"
	"ssn"
)

func main() {
	stdout := log.New(os.Stdout, "server ", log.Ldate|log.Ltime)
	stderr := log.New(os.Stderr, "server ", log.Ldate|log.Ltime)
	s, err := srv.New(ssn.Conf(stdout, stderr))
	if err != nil {
		stderr.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		stderr.Fatal(err)
	}
	stdout.Println("main exited")
}
