package main

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	mysql "grpc/db"
	vehicle "grpc/protos"
	"log"
	"net"
)

const portNumber = "9000"

var db *gorm.DB

type Vehicle struct {
	VehicleId           string
	VehicleName         string
	VehicleNumber       string
	VehicleVinNumber    string
	VehicleSerialNumber string
}

type vehicleServer struct {
	vehicle.VehicleServer
}

func (s *vehicleServer) GetVehicle(ctx context.Context, req *vehicle.GetVehicleRequest) (*vehicle.GetVehicleResponse, error) {
	id := req.GetVehicleId()

	var message *vehicle.VehicleMessage
	db.Where("vehicle_id = ?", id).Find(&message)

	return &vehicle.GetVehicleResponse{
		VehicleMessage: message,
	}, nil
}

func (s *vehicleServer) ListVehicles(ctx context.Context, req *vehicle.ListVehiclesRequest) (*vehicle.ListVehiclesResponse, error) {

	var result []*vehicle.VehicleMessage

	db.Find(&result)

	messages := make([]*vehicle.VehicleMessage, len(result))
	for i, v := range result {
		messages[i] = v
	}

	return &vehicle.ListVehiclesResponse{
		VehicleMessages: messages,
	}, nil
}

func init() {
	db = mysql.Connect()
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	vehicle.RegisterVehicleServer(grpcServer, &vehicleServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
