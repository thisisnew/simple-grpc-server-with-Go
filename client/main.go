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

func (s *vehicleServer) InsertVehicle(ctx context.Context, req *vehicle.VehicleMessage) (*vehicle.GetVehicleResponse, error) {

	result := db.Create(&req)

	if result.RowsAffected > 0 {
		return &vehicle.GetVehicleResponse{
			VehicleMessage: req,
		}, nil
	}

	return &vehicle.GetVehicleResponse{}, nil
}

func (s *vehicleServer) UpdateVehicle(ctx context.Context, req *vehicle.VehicleMessage) (*vehicle.GetVehicleResponse, error) {

	result := db.Model(&req).Where("vehicle_id = ?", req.VehicleId).Updates(map[string]interface{}{"vehicle_name": req.VehicleName, "vehicle_number": req.VehicleNumber, "vehicle_vin_number": req.VehicleVinNumber, "vehicle_serial_number": req.VehicleSerialNumber})

	if result.RowsAffected == 0 {

		return &vehicle.GetVehicleResponse{}, nil
	} else {

		var message *vehicle.VehicleMessage
		db.Where("vehicle_id = ?", req.VehicleId).Find(&message)

		return &vehicle.GetVehicleResponse{
			VehicleMessage: message,
		}, nil
	}
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
