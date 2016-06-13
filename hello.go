package main

import "fmt"
import "net"

func main(){
	fmt.Printf("Hello,World\n中文输出\n")
	fmt.Printf("たなか\n")
	if ln,err :=net.Listen("tcp","7892");err ==nil{
		defer ln.Close()
		for{
			ln.Accept()
			fmt.Println("Receive a Message")
		}
	}
}
