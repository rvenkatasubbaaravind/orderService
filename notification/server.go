package notification

import (
	"context"
	"io"
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

func (n NotifcationServer) SendStatus(stream pb.Notification_SendStatusServer) error {
	for {
		email, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.NotificationStatus{Message: "Done processing the messages"})
		}
		if err != nil {
			log.Println("Error while receiving from stream", err)
			return err
		}
		n.SendEmailNotification(context.Background(), email)
	}
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
