package chapter1

import (
	"flag"
	"fmt"
	"net"
)

func HelloClient(){
	ip:=flag.String("ip", "127.0.0.1", "请输入IP地址")
	port:=flag.Int("port",8081,"请输入port")
	conn,err:=net.Dial("tcp4", fmt.Sprintf("%s:%d", *ip,*port))
	if err!=nil{
		fmt.Printf("created conn failed: %s",err.Error())
		return 
	}
	defer conn.Close()

	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil{
		fmt.Printf("read failed: %s",err.Error())
		return
	}
	fmt.Printf("read from tcp %d byte,contenct: %s\n",n,string(buf))
}