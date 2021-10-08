package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

func decimalToBinary(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func main() {
	//Konfigurasi awal Server
	PORT := ":123"
	l, err := net.Listen("tcp", PORT) //Listen via TCP
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil { //Jika terdapat eror
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		netData2, _ := strconv.Atoi(netData)
		message := decimalToBinary(netData2)
		length := len(strconv.Itoa(message))

		var i int = 0
		var j int = 0
		//var b string = ""
		fmt.Scanf("Enter the window size -> %d", &j)

		j = j - 1
		var k int = j

		for i != length {
			for i != (length - j) {
				c.Write([]byte(strconv.Itoa(message)))
				b, _ := bufio.NewReader(c).ReadString('\n')
				fmt.Print(b)
				if b != "ACK Lost" {
					time.Sleep(1)
					fmt.Print("Acknowledgement Diterima! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + " Sekarang kirim paket selanjutnya")
					i = i + 1
					k = k + 1
					time.Sleep(1)
				} else {
					time.Sleep(1)
					fmt.Print("Acknowledgement data bit nya LOST! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + " Sekarang kirim ulang paket yang sama")
				}

				for i != length {
					c.Write([]byte(strconv.Itoa(message)))
					b, _ := bufio.NewReader(c).ReadString('\n')
					fmt.Print(b)
					if b != "ACK Lost" {
						time.Sleep(1)
						fmt.Print("Acknowledgement Diterima! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + " Sekarang kirim paket selanjutnya")
						i = i + 1
						time.Sleep(1)
					} else {
						time.Sleep(1)
						fmt.Print("Acknowledgement data bit nya LOST! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + " Sekarang kirim ulang paket yang sama")
						time.Sleep(1)
					}
				}
			}
		}

		//Jika user mengetikan "KELUAR" maka tutup server
		if strings.TrimSpace(string(netData)) == "KELUAR" {
			fmt.Println("Menutup server!")
			return
		}

		//Lakukan hashing sha3 512
		//h := sha3.New512()rite([]byte(string(netData)))
		// hasil := h.Sum(nil)

		//Cetak hasil hashing sha3 - 512
		//fmt.Printf("-> %x\n", hasil)

		//Waktu pengiriman
		// t := time.Now() //Dapatkan waktu sekarang
		// myTime := t.Format(time.RFC3339) + "\n"
		// c.Write([]byte(myTime)) //Cetak waktu

	}
}
