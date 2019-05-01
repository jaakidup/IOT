package main


func main() {
  server := &Server{"My Server"}
  server.Running()
  defer server.Shutdown()

  device := &SamsungDevice{ID: "sam001"}
  server.CallDevice(device)
}




