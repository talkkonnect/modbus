//Test Code with ETT MODVYS RTU Relay4/IN4
//By Suvir Kumar 09/03/2022
package main

import (
	"log"
	"os"
	"time"

	"github.com/talkkonnect/modbus"
)

const (
	rtuDevice = "/dev/ttyAMA0"
)

// func TestRTUClient(t *testing.T) {
// 	// Diagslave does not support broadcast id.
// 	handler := modbus.NewRTUClientHandler(rtuDevice)
// 	handler.SlaveId = 01
// 	ClientTestAll(t, modbus.NewClient(handler))
// }

func main() {
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 01
	handler.Logger = log.New(os.Stdout, "rtu: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()
	var i uint16
	for i = 0; i < 4; i++ {
		modBusSendCommand(handler, i, 0xFF00)
		time.Sleep(1 * time.Second)
		modBusSendCommand(handler, i, 0x0000)
		time.Sleep(1 * time.Second)
	}
	// results, err := client.ReadDiscreteInputs(00, 01)
	// if err != nil || results == nil {
	// 	log.Fatal(err, results)
	// }
	//	modBusSendCommand(01, 0x0000)
}

func modBusSendCommand(handler *modbus.RTUClientHandler, relay uint16, value uint16) {
	client := modbus.NewClient(handler)
	results, err := client.WriteSingleCoil(relay, value)
	if err != nil || results == nil {
		log.Fatal(err, results)
	}
}
