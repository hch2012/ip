package main

import(
	"net"
	"os"
	"time"
	"fmt"
)

func main() {
	conn,err:=net.Dial("udp4",os.Args[1])
	checkErr(err)
	limiter := time.Tick(time.Minute * 3)
	for{
		<-limiter
		fmt.Println("send------")
		conn.Write([]byte("ack"))
	}
}


func checkErr(err error) {
	if err!=nil{
		panic(err)
	}
}