package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

func main() {
	//Konfigurasi awal koneksi ke server
	CONNECT := "127.0.0.1:123"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil { //Jika terdapat eror
		fmt.Println(err)
		return
	}

	for {
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print(">> ")
		// text, _ := reader.ReadString('\n')
		// fmt.Fprintf(c, text+"\n")

		// m := bufio.NewReader(os.Stdin)
		// fmt.Print(">> ")
		// m2, _ := m.ReadString('\n')
		//fmt.Fprintf(c, m2+"\n")
		m, _ := bufio.NewReader(c).ReadString('\n')
		k, _ := bufio.NewReader(c).ReadString('\n')
		k2, _ := strconv.Atoi(k)
		//var a string
		var i int = 0
		//f := rand.Intn(1-0) + 0
		for i != k2 {
			f := rand.Intn(1-0) + 0
			if f == 0 {
				b := "ACK Lost"
				bufio.NewReader(c).ReadString('\n')
				c.Write([]byte(b))
			} else if f == 1 {
				b := "ACK " + strconv.Itoa(i)
				bufio.NewReader(c).ReadString('\n')
				c.Write([]byte(b))
				//a := a + message
				i = i + 1
			}
			fmt.Print("The message received is :", m)
		}
		// message, _ := bufio.NewReader(c).ReadString('\n')
		// fmt.Print("->: " + message)

		//Jika user mengetikan "KELUAR" maka tutup koneksi ke server
		if strings.TrimSpace(string(m)) == "KELUAR" {
			fmt.Println("Menutup koneksi client ke server...")
			return
		}
	}
}
