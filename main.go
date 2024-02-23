package main

import (
	"flag"
	"fmt"
	"github.com/armon/go-socks5"
	"os"
	"strconv"
)

var (
	port     int
	username string
	password string
	conf     = &socks5.Config{}
)

func init() {
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "pwd", "", "password")
	flag.IntVar(&port, "p", 1080, "port on listen, must be greater than 0")
	flag.Parse()
	if username != "" {
		cred := socks5.StaticCredentials{
			username: password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: cred}
		conf.AuthMethods = append(conf.AuthMethods, cator)
	}
}

func main() {
	server, err := socks5.New(conf)
	if err != nil {
		fmt.Println("[error]", err)
		os.Exit(0)
	}
	// Create SOCKS5 proxy on localhost port 1080
	fmt.Printf("listening socks5://%s\n", "0.0.0.0:"+strconv.Itoa(port))
	if err := server.ListenAndServe("tcp", ":"+strconv.Itoa(port)); err != nil {
		fmt.Println("[error]", err)
		os.Exit(0)
	}
}
