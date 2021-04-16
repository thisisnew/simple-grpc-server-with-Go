package main

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	vehicle "grpc/protos"
	"log"
	"net"
)

const portNumber = "9000"
const dsn = "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var message *vehicle.VehicleMessage
	db.Where("vehicle_id = ?", id).Find(&message)

	return &vehicle.GetVehicleResponse{
		VehicleMessage: message,
	}, nil
}

func (s *vehicleServer) ListVehicles(ctx context.Context, req *vehicle.ListVehiclesRequest) (*vehicle.ListVehiclesResponse, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

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
