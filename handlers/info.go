package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Info is used for health checking as well as outputting current app version
func (h Handler) Info() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		err := writeResponse(w, 200, "version", "0.0.1")
		if err != nil {
			log.Println(fmt.Errorf("error writing response: %w", err))
		}
	}
}
