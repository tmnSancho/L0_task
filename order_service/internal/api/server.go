package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartService(orderService orderService) {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/start.html"))
		tmpl.Execute(writer, nil)
	})

	r.HandleFunc("/order", func(writer http.ResponseWriter, request *http.Request) {
		orderUid := request.FormValue("order_uid")

		order := orderService.GetOrderById(orderUid)

		if order.OrderUID == "" {
			tempStruct := struct {
				UidInput string
			}{
				UidInput: orderUid,
			}
			tmpl := template.Must(template.ParseFiles("templates/error.html"))
			tmpl.Execute(writer, tempStruct)
		} else {
			tmpl := template.Must(template.ParseFiles("templates/order.html"))
			tmpl.Execute(writer, order)
		}
	})

	http.Handle("/", r)

	log.Println("Succsess start server")
	if err := http.ListenAndServe("127.0.0.1:8888", nil); err != nil {
		log.Fatalf("Server failed to start. Error: %s\n", err)
	}
}
