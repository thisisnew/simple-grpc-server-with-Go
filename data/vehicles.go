package data

import (
	vehicle "grpc/protos"
)

var Vehicles = []*vehicle.VehicleMessage{
	{
		VehicleId: "1",
		VehicleName: "스파크",
		VehicleNumber: "101가1111",
		VehicleVinNumber: "VN11111111111111111111111",
		VehicleSerialNumber: "SANBON123456789121",
	},
	{
		VehicleId: "2",
		VehicleName: "SM5(LPG)",
		VehicleNumber: "12라5555",
		VehicleVinNumber: "1234154778487457",
		VehicleSerialNumber: "",
	},
	{
		VehicleId: "3",
		VehicleName: "투싼 TL(디젤)",
		VehicleNumber: "11아1212",
		VehicleVinNumber: "1245147148747874",
		VehicleSerialNumber: "",
	},
	{
		VehicleId: "4",
		VehicleName: "아이오닉 EV",
		VehicleNumber: "11하1212",
		VehicleVinNumber: "12120222001211",
		VehicleSerialNumber: "",
	},
	{
		VehicleId: "5",
		VehicleName: "니로 HEV",
		VehicleNumber: "22호1231",
		VehicleVinNumber: "KMHS281HGMU327110",
		VehicleSerialNumber: "833-29191-111",
	},
}