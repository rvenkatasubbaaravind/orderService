package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/order_service/controller"
	"github.com/order_service/notification"
	"github.com/order_service/utils"
)

func main() {
	db, err := utils.GetDatabase()
	if err != nil {
		log.Println("failed connecting to db", err)
		return
	}

	client, err := utils.GetGRPCClient()
	if err != nil {
		log.Println("failed  getting  the grpc client", err)
		return
	}

	oh := controller.OrderHandler{
		Database:           db,
		NotificationClient: client,
	}

	go notification.StartServer()

	router := mux.NewRouter()
	router.HandleFunc("/orders", oh.PostOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", oh.GetOrder).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id}", oh.CancelOrder).Methods(http.MethodDelete)
	router.HandleFunc("/orders/{id}", oh.PatchOrder).Methods(http.MethodPut)

	router.HandleFunc("/orders", oh.GetOrders).Methods(http.MethodGet)

	router.HandleFunc("/notify/{status}", oh.NotifyUsers).Methods(http.MethodGet)

	http.ListenAndServe(":5000", router)
}
