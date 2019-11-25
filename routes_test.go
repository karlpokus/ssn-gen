package ssn

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/karlpokus/routest/v2"
	"github.com/karlpokus/srv"
)

func TestFoo(t *testing.T) {
	routest.Test(t, func() http.Handler {
		stdout := log.New(ioutil.Discard, "", 0)
		stderr := log.New(ioutil.Discard, "", 0)
		s, err := srv.New(Conf(stdout, stderr))
		if err != nil {
			t.Errorf("expected nil err, got %s", err)
		}
		return s
	}, []routest.Data{
		{
			"gen ssn invalid",
			"GET",
			"/ssn/foo",
			nil,
			nil,
			400,
			[]byte("foo is not a valid int"),
		},
		{
			"gen ssn valid but rate limited",
			"GET",
			"/ssn/2",
			nil,
			nil,
			429,
			[]byte("Too Many Requests"),
		},
	})
}
