package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	tests "github.com/iamskp11/mock-tcp-web-server/tests"
)

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		panic(err)
	}
	go tests.Test()
	for {
		fmt.Println("Waiting to accept connection")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Connection accepted")
		// locallAddr := conn.LocalAddr()
		// remoteAddr := conn.RemoteAddr()
		// fmt.Println(locallAddr, remoteAddr)
		go do(conn)
	}
}

func get_method_and_endpoint(s string) (string, string) {
	// fmt.Println(s)
	// fmt.Println(s[0:10])
	// fmt.Println(len(s))
	arr := make([]string, 2)
	ctr := 0
	var curr_string string
	for _, ch := range s {
		if ch == ' ' {
			arr[ctr] = curr_string
			curr_string = ""
			ctr += 1
			if ctr > 1 {
				break
			}
		} else {
			curr_string += string(ch)
		}
	}
	return arr[0], arr[1]
}

func greet_user() string {
	all_greeting_words := []string{"Hello Bro", "Hello Yar", "Hello Bhai"}
	rand_idx := rand.Intn(3)
	return all_greeting_words[rand_idx]
}

func do(conn net.Conn) {
	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		fmt.Println("Some error occured while reading data")
		panic(err)
	}
	s := string(b)
	//fmt.Println("Read complete\n", s)
	time.Sleep(time.Second * 1)
	method, endpoint := get_method_and_endpoint(s)
	fmt.Println("Called ", method, endpoint)
	if method == "GET" && endpoint == "/hello" {
		response := []byte("HTTP/1.1 200 OK\r\n\r\n")
		output := greet_user()
		response = append(response, output...)
		response = append(response, "\r\n"...)
		conn.Write(response)
	} else {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nInvalid URL or Method\r\n"))
	}
	conn.Close()
}
