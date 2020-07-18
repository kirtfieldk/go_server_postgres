package lp_routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/keithkfield/pg_api/helpers"
)

type Order struct {
	Person  string `json:"person"`
	Age     int32  `json:"age"`
	Loyalty string `json:"loyalty"`
}
type Company struct {
	Name     string `json:"name"`
	Buisness string `json:"buisness"`
}
type Users struct {
	User_id string `json:"user_id"`
	Name    string `json:"name"`
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
	fmt.Printf("%s", order.Person)

}

/* Write Json */
func AddOrder(w http.ResponseWriter, r *http.Request) {
	o1 := Order{"Keith Kirtfield", 22, "Platinum"}
	o1Json, err := json.Marshal(o1)
	if err != nil {
		fmt.Println("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(o1Json)
}

func UsersRoute(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	user := []Users{
		{User_id: "P840283", Name: "Keith Kirtfield"},
		{User_id: "P840283", Name: "Keith Kirtfield"},
	}
	o1Json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(o1Json)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
