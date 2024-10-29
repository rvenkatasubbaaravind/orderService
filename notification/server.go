package notification

import (
	"context"
	"log"
	"net"

	pb "github.com/order_service/notification/pb"
	"google.golang.org/grpc"
)

type NotifcationServer struct {
	pb.UnimplementedNotificationServer
}

func (NotifcationServer) SendEmailNotification(context.Context, *pb.EmailNotification) (*pb.NotificationStatus, error) {
	return &pb.NotificationStatus{
		Message: "Success",
	}, nil
}

func (NotifcationServer) SendPhoneNotification(context.Context, *pb.PhoneNotification) (*pb.NotificationStatus, error) {
	return &pb.NotificationStatus{
		Message: "Success",
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, &NotifcationServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
