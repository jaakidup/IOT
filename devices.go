package main

import "fmt"

type IOTDevice interface {
  Connect()
  // Authenticate()
  // Config()
  // Update()
  Disconnect()
}



type SamsungDevice struct {
  ID string
}
func (sd *SamsungDevice) Connect() {
  fmt.Println("Connecting to device with ID: ", sd.ID)
}
func (sd *SamsungDevice) Disconnect() {
  fmt.Println("Disconnecting from device with ID: ", sd.ID)
}