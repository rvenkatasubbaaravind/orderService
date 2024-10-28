package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := oh.Database.Create(&order)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	oh.NotificationClient.SendEmailNotification(ctx, &pb.EmailNotification{Email: order.Email, Text: "Order is succesful"})
	oh.NotificationClient.SendPhoneNotification(ctx, &pb.PhoneNotification{PhoneNo: order.PhoneNo, Text: "Order is succesful"})
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	oh.NotificationClient.SendEmailNotification(ctx, &pb.EmailNotification{Email: order.Email, Text: "Order cancellation is succesful"})
	oh.NotificationClient.SendPhoneNotification(ctx, &pb.PhoneNotification{PhoneNo: order.PhoneNo, Text: "Order cancellation is succesful"})
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

func (oh *OrderHandler) PatchOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Order
	var existingRecord models.Order

	pathVariables := mux.Vars(r)

	existingRecord, status, err := checkIfOrderExists(oh.Database, pathVariables)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.Id != 0 || reqBody.PaymentStatus != "" || reqBody.ProductID != "" || reqBody.Name != "" {
		http.Error(w, "Uneditable fields", http.StatusBadRequest)
		return
	}

	if reqBody.Address != "" {
		existingRecord.Address = reqBody.Address
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
