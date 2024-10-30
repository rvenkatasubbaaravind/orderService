package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/gorilla/mux"
	"github.com/order_service/models"
	pb "github.com/order_service/notification/pb"
	"github.com/order_service/utils"
	"gorm.io/gorm"
)

type OrderHandler struct {
	Database           *gorm.DB
	NotificationClient pb.NotificationClient
}

func (oh *OrderHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
	var body models.PostOrder
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.Address = body.Address
	order.Email = body.Email
	order.PaymentStatus = "Prepaid"
	order.Name = body.Name
	order.ProductID = body.ProductID
	order.Status = "NEW"
	order.PhoneNo = body.PhoneNo

	res := oh.Database.Create(&order)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, err := oh.NotificationClient.SendEmailNotification(ctx, &pb.EmailNotification{Email: order.Email, Text: "Order cancellation is succesful"})
		if err != nil {
			log.Println("Failed in delivering email", err.Error())
			return
		}
		log.Println("Status of the email", status.Message)
	}()
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, err := oh.NotificationClient.SendPhoneNotification(ctx, &pb.PhoneNotification{PhoneNo: order.PhoneNo, Text: "Order cancellation is succesful"})
		if err != nil {
			log.Println("Failed in delivering SMS", err.Error())
			return
		}
		log.Println("Status of the sms", status.Message)
	}()
	w.Write([]byte(fmt.Sprintf("orderid %d is success", order.Id)))
}

func (oh *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)

	var order models.Order

	order, status, err := checkIfOrderExists(oh.Database, pathVariables)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = oh.Database.Delete(&order).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, err := oh.NotificationClient.SendEmailNotification(ctx, &pb.EmailNotification{Email: order.Email, Text: "Order cancellation is succesful"})
		if err != nil {
			log.Println("Failed in delivering email", err.Error())
			return
		}
		log.Println("Status of the email", status.Message)
	}()
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, err := oh.NotificationClient.SendPhoneNotification(ctx, &pb.PhoneNotification{PhoneNo: order.PhoneNo, Text: "Order cancellation is succesful"})
		if err != nil {
			log.Println("Failed in delivering SMS", err.Error())
			return
		}
		log.Println("Status of the sms", status.Message)
	}()
	w.Write([]byte("deleted order"))
}

func (oh *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)

	var order models.Order

	order, status, err := checkIfOrderExists(oh.Database, pathVariables)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (oh *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	err := oh.Database.Find(&orders).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (oh *OrderHandler) PatchOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody models.PatchOrder
	var existingRecord models.Order

	pathVariables := mux.Vars(r)

	existingRecord, status, err := checkIfOrderExists(oh.Database, pathVariables)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&existingRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.Email != "" {
		existingRecord.Email = reqBody.Email
	}

	if reqBody.PhoneNo != "" {
		existingRecord.PhoneNo = reqBody.PhoneNo
	}

	oh.Database.Save(&existingRecord)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingRecord)
}

func checkIfOrderExists(db *gorm.DB, pathVariables map[string]string) (order models.Order, code int, err error) {
	order_id, ok := pathVariables["id"]
	if !ok {
		return order, http.StatusBadRequest, errors.New("missing id field")
	}

	order_id_no, err := utils.GetNumber(order_id)
	if err != nil {
		return order, http.StatusBadRequest, err
	}

	err = db.Where("id= ?", order_id_no).First(&order).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return order, http.StatusNotFound, err
	} else if err != nil {
		return order, http.StatusInternalServerError, err
	}
	return
}

func (oh *OrderHandler) NotifyUsers(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	status, ok := pathVariables["status"]
	if !ok {
		http.Error(w, "Missing required path param", http.StatusBadRequest)
		return
	}

	if !slices.Contains([]string{"OOD", "SHIP", "FDB"}, status) {
		http.Error(w, "Not a valid state to act", http.StatusBadRequest)
		return
	}

	var orders []models.Order

	err := oh.Database.Where("status", status).Find(&orders).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getText := func(status string) string {
		switch status {
		case "OOD":
			return "Out For delivery"
		case "SHIP":
			return "Ready to ship"
		case "FDB":
			return "Rate your FeedBack"
		}
		return ""
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		stream, err := oh.NotificationClient.SendStatus(ctx)
		if err != nil {
			log.Println("Failed to get the stream for client", err)
			return
		}
		for _, order := range orders {
			err = stream.Send(&pb.EmailNotification{
				Email: order.Email,
				Text:  getText(status),
			})
			if err != nil {
				log.Printf("Fail to send a notification to %s ", order.Name)
				continue
			}
		}

		message, err := stream.CloseAndRecv()
		if err != nil {
			log.Println("Fail to close and receive the stream")
			return
		}
		log.Println("response of the stream", message)

	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Accepted"))
}
