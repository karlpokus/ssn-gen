package main

import (
	"log"
	"os"

	"ssn"
	"github.com/karlpokus/srv"
	"github.com/julienschmidt/httprouter"
)

func conf(stdout, stderr *log.Logger) srv.ConfFunc {
	return func(s *srv.Server) error {
		router := httprouter.New()
		router.HandlerFunc("GET", "/ssn/:n", ssn.Gen(stdout, stderr))
		s.Router = router
		s.Logger = stdout
		s.Host = "0.0.0.0"
		return nil
	}
}

func main() {
	stdout := log.New(os.Stdout, "server ", log.Ldate|log.Ltime)
	stderr := log.New(os.Stderr, "server ", log.Ldate|log.Ltime)
	s, err := srv.New(conf(stdout, stderr))
	if err != nil {
		stderr.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		stderr.Fatal(err)
	}
	stdout.Println("main exited")
}
