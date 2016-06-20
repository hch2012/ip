package main

import(
	"net"
	"os"
	"time"
)

func main() {
	conn,err:=net.Dial("udp4",os.Args[1])
	checkErr(err)
	limiter := time.Tick(time.Minute * 3)
	for{
		<-limiter
		conn.Write([]byte("ack"))
	}
}


func checkErr(err error) {
	if err!=nil{
		panic(err)
	}
}