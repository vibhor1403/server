
package main

import (
	"net"
	"os"
	"fmt"
	"strings"
)

var store map[string]string

func main() {

	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	store = make(map[string]string)

	setValue("abc", "dhef")
	setValue("123", "23")
	setValue("vib", "gh")
	setValue("ff", "hh")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
		// conn.Close() // we're finished
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		
		splits := strings.Split(string(buf[0:n-1]), " ")
		//fmt.Println(splits[0])
		if splits[0] == "get" {
			val, ok := getValue(splits[1])
			//fmt.Println(splits[1], val)
			if ok == false {
				_, err2 := conn.Write([]byte("Not present"))
				if err2 != nil {
					return
				}
			} else {
			_, err2 := conn.Write([]byte (val) )
				if err2 != nil {
					return
				}
			}
		} else if splits[0] == "set" {
				setValue (splits[1], splits[2])
				_, err2 := conn.Write([]byte ("Value Set") )
				if err2 != nil {
					return
				}

		} else {
			_, err2 := conn.Write([]byte ("Wrong Call") )
			if err2 != nil {
				return
			}

		}
		
	}
}

func getValue(key string) (a string, b bool) {
	a,b = store[key]
	return
}

func setValue(key, value string) {
	store[key] = value
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
