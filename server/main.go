package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	mysql "grpc/db"
	vehicle "grpc/protos"
	"log"
	"net"
)

var db *gorm.DB

type statusMessage struct {
	status  string
	message string
}

type vehicleServer struct {
	vehicle.VehicleServer
}

func init() {
	db = mysql.Connect()
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

func (s *vehicleServer) InsertVehicle(ctx context.Context, req *vehicle.VehicleMessage) (*vehicle.StatusMessage, error) {

	result := db.Create(&req)

	if result.RowsAffected > 0 {
		return &vehicle.StatusMessage{
			Status:  "200",
			Message: "Success",
		}, nil
	}

	return &vehicle.StatusMessage{
		Status:  "500",
		Message: "Failure",
	}, nil
}

func (s *vehicleServer) UpdateVehicle(ctx context.Context, req *vehicle.VehicleMessage) (*vehicle.StatusMessage, error) {

	result := db.Model(&req).Where("vehicle_id = ?", req.VehicleId).Updates(map[string]interface{}{"vehicle_name": req.VehicleName, "vehicle_number": req.VehicleNumber, "vehicle_vin_number": req.VehicleVinNumber, "vehicle_serial_number": req.VehicleSerialNumber})

	if result.RowsAffected == 0 {
		return &vehicle.StatusMessage{
			Status:  "500",
			Message: "Failure",
		}, nil
	} else {
		return &vehicle.StatusMessage{
			Status:  "200",
			Message: "Success",
		}, nil
	}
}

func (s *vehicleServer) DeleteVehicle(ctx context.Context, req *vehicle.GetVehicleRequest) (*vehicle.ListVehiclesResponse, error) {

	id := req.GetVehicleId()

	var vehicleMessage *vehicle.VehicleMessage

	db.Where("vehicle_id = ?", id).Delete(&vehicleMessage)

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

func (s *vehicleServer) InsertGeoDatas(ctx context.Context, req *vehicle.GeoDatas) (*vehicle.StatusMessage, error) {

	aaa := &req
	fmt.Println(aaa)

	//result := db.Create(&req)
	//
	//if result.RowsAffected > 0 {
	//	return &vehicle.StatusMessage{
	//		Status:  "200",
	//		Message: "Success",
	//	}, nil
	//}

	return &vehicle.StatusMessage{
		Status:  "500",
		Message: "Failure",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	vehicle.RegisterVehicleServer(grpcServer, &vehicleServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
