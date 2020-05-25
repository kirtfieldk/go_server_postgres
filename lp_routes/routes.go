package lp_routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/keithkfield/pg_api/helpers"
)

type Order struct {
	Person  string
	Age     int32
	Loyalty string
}
type Company struct {
	Name     string
	Buisness string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

/* Rrading and recieving JSON */
func SearchRoute(w http.ResponseWriter, r *http.Request) {

	var order Order
	err := helpers.DecodeJsonBody(w, r, &order)
	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, "Person: %+v", order)
	fmt.Printf("%s", order.Person)

}
