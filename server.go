package main

import "fmt"
import "time"

type Server struct {
  name string
}

func (s *Server) Running() {
  fmt.Println(s.name , " running")
}

func (s *Server) CallDevice(device IOTDevice) {
  device.Connect()
  time.Sleep(1000 * time.Millisecond)
  device.Disconnect()
}


func(s *Server) Shutdown() {
  fmt.Println(s.name, " shutting down...")
}
