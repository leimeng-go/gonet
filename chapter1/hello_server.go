package chapter1

import (
	"flag"
	"fmt"
	"net"
)


func HelloServer(){
	ip:=flag.String("ip","127.0.0.1","输入IP地址")
	port:=flag.Int("port",8081,"输入端口")
	flag.Parse()

	listener,err:=net.Listen("tcp4",fmt.Sprintf("%s:%d",*ip,*port))
	if err!=nil{
		fmt.Printf("created listener failed: %s",err.Error())
		return 
	}
	defer listener.Close()

	conn,err:=listener.Accept()
	if err!=nil{
		fmt.Printf("created connection failed: %s",err.Error())
		return 
	}
	defer conn.Close()
	n,err:=conn.Write([]byte("Hello World"))
	if err!=nil{
		fmt.Printf("conn writed failed: %s",err.Error())
		return 
	}
	fmt.Printf("conn writed %d byte",n)
    
}