package models

type Order struct {
	Id            uint   `json:"id,omitempty" gorm:"unique;primaryKey;autoIncrement"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	PhoneNo       string `json:"phone_no"`
	ProductID     string `json:"product_id"`
	PaymentStatus string `json:"payment_status,omitempty"`
	Status        string `json:"status"`
}

type PostOrder struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	PhoneNo   string `json:"phone_no"`
	ProductID string `json:"product_id"`
}

type PatchOrder struct {
	Email   string `json:"email"`
	PhoneNo string `json:"phone_no"`
}
