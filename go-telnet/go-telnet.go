package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func GetConfig() (string, int, time.Duration) {
	var timeouInt int
	flag.IntVar(&timeouInt, "timeout", 0, "timeout")
	flag.Parse()
	timeout := time.Duration(timeouInt)
	address := flag.Args()[0]
	port, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		log.Fatal("cant parse port")
	}
	return address, port, timeout
}
func myReader(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			log.Fatal("Can't get data from remote server")
		}
		text := scanner.Text()
		fmt.Println(text)
	}
}
func MySender(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Fatal("Exit")
		}
		text := scanner.Text()
		conn.Write([]byte(fmt.Sprintln(text)))
	}
}
func main() {
	address, port, timeout := GetConfig()
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%d", address, port), time.Second*timeout)
	if err != nil {
		log.Fatal("Error while connecting to the remote server")
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go myReader(conn, &wg)
	go MySender(conn, &wg)
	wg.Wait()
}
