package main

import(
	"net"
	"os"
	"strings"
)

func main() {
	addr,err:=net.ResolveUDPAddr("udp4",os.Args[1])
	checkErr(err)
	listener,err:=net.ListenUDP("udp4",addr)
	checkErr(err)
	defer listener.Close()
	var buff =make([]byte,1024)
	dstFile,err := os.Create(os.Args[2])
	checkErr(err)
	for {
		_,addr,err:=listener.ReadFromUDP(buff)
		checkErr(err)
		dstFile.Seek(0,0)
		dstFile.WriteString(strings.Split(addr.String(),":")[0])
	}
}


func checkErr(err error) {
	if err!=nil{
		panic(err)
	}
}