package ssn

import (
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/karlpokus/srv"
)

func Conf(stdout, stderr *log.Logger) srv.ConfFunc {
	return func(s *srv.Server) error {
		router := httprouter.New()
		router.HandlerFunc("GET", "/ssn/:n", Gen(stdout, stderr))
		s.Router = router
		s.Logger = stdout
		s.Host = "0.0.0.0"
		return nil
	}
}
