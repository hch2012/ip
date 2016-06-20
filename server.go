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
	lastIp:=""
	dstFile,err := os.Open(os.Args[2])
	checkErr(err)
	defer dstFile.Close()
	n,_:=dstFile.Read(buff)
	checkErr(err)
	if n!=0 {
		lastIp=string(buff[0:n])
	}
	for {
		_,addr,err:=listener.ReadFromUDP(buff)
		checkErr(err)

		ip:=strings.Split(addr.String(),":")[0]
		
		if ip==lastIp {
			continue
		}
		dstFile.Close()
		dstFile,err:=os.Create(os.Args[2])
		checkErr(err)
		defer dstFile.Close()
		dstFile.WriteString(ip)
	}
}


func checkErr(err error) {
	if err!=nil{
		panic(err)
	}
}