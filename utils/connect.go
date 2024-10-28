package utils

import (
	"log"

	"github.com/order_service/models"
	pb "github.com/order_service/notification/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() (db *gorm.DB, err error) {
	dsn := "host=localhost user=myusername password=mypassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Order{})
	return
}

func GetGRPCClient() (c pb.NotificationClient, err error) {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c = pb.NewNotificationClient(conn)

	return c, err
}
