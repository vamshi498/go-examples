# go-examples.
Collection of golang programs


# buffered and unbuffered channels
```go
c := make(chan int) // unbuffered channel
c := make(chan int,2) // buffered channel with capacity 2
```
Buffered channel are async in nature.
Unbuffered channel are synchronous in nature.   
Both the sender and receiver should be ready for communication. 
If receiver is not ready sender will block untill receiver is ready