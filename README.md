# timer
basic task manager in Go

go run main.go

## Key points
1. interfacelist has a list of interfaces related to queue and DB which is used in main program
2. /db and /msgqueue gives implementations. So if one decides to add new type of queue like Kafka, it can be added as an implementation
3. Business logic can be tested by stubs calling to interfaces
4. The calling client (main.go) should not have any dependency on implementation of either msgqueue or db