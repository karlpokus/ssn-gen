package ssn

import (
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/julienschmidt/httprouter"
)

func Gen(stdout, stderr *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		n := params.ByName("n")
		i, err := strconv.Atoi(n)
    if err != nil {
			msg := fmt.Sprintf("%s is not a valid int", n)
			stderr.Println(msg)
			http.Error(w, msg, 400)
			return
    }
		stdout.Printf("generated %d ssns", i)
		fmt.Fprintf(w, GenN(i))
	}
}
