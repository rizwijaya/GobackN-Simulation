package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
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
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukan Pesan(Ketik \"KELUAR\" untuk Keluar): ")
		netData, err := reader.ReadString('\n')
		fmt.Fprintf(c, netData+"\n")

		if err != nil {
			fmt.Println(err)
			return
		}

		//Jika user mengetikan "KELUAR" maka tutup server
		if strings.TrimSpace(string(netData)) == "KELUAR" {
			fmt.Println("Menutup server!")
			return
		}

		message := binary(strings.TrimSpace(string(netData)))
		length := len(message)

		fmt.Fprintf(c, strconv.Itoa(length)+"\n")
		var i int = 0
		var j int = 0
		fmt.Println("Masukan Ukuran Window: ")
		fmt.Scanln(&j)
		j = j - 1
		var k int = j

		chars := []rune(message)
		for i != length {
			for i != length-j {
				char := string(chars[i])
				fmt.Fprintf(c, string(char)+"\n")
				b, _ := bufio.NewReader(c).ReadString('\n')
				fmt.Print(b)
				if strings.TrimSpace(string(b)) != "ACK Lost" {
					time.Sleep(1 * time.Second)
					fmt.Print("Acknowledgement Diterima! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + "\nKirim paket selanjutnya\n")
					i = i + 1
					k = k + 1
					time.Sleep(1 * time.Second)
				} else {
					time.Sleep(1 * time.Second)
					fmt.Print("Acknowledgement data bit nya LOST! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + "\nKirim ulang paket yang sama\n")
				}
			}

			for i != length {
				char := string(chars[i])
				fmt.Fprintf(c, string(char)+"\n")
				b, _ := bufio.NewReader(c).ReadString('\n')
				fmt.Print(b)
				if strings.TrimSpace(string(b)) != "ACK Lost" {
					time.Sleep(1 * time.Second)
					fmt.Print("Acknowledgement Diterima! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + "\nKirim paket selanjutnya\n")
					i = i + 1
					time.Sleep(1 * time.Second)
				} else {
					time.Sleep(1 * time.Second)
					fmt.Print("Acknowledgement data bit nya LOST! Sliding window pada rentang " + strconv.Itoa(i+1) + " ke " + strconv.Itoa(k+1) + "\nKirim ulang paket yang sama\n")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}
}
