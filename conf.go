package ssn

import (
	"log"

	"github.com/karlpokus/srv"
)

func Conf(stdout, stderr *log.Logger) srv.ConfFunc {
	return func(s *srv.Server) error {
		s.Router = routes(stdout, stderr)
		s.Logger = stdout
		s.Host = "0.0.0.0"
		return nil
	}
}
