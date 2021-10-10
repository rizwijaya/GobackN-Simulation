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
		m, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println("Menerima Pesan " + m)

		//Jika user mengetikan "KELUAR" maka tutup koneksi ke server
		if strings.TrimSpace(string(m)) == "KELUAR" {
			fmt.Println("Menutup koneksi client ke server...")
			return
		}

		k, err2 := bufio.NewReader(c).ReadString('\n')
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		//fmt.Println("Menerima Panjang " + k)

		k2, err3 := strconv.ParseUint(strings.TrimSpace(string(k)), 10, 64)
		if err3 != nil {
			fmt.Println(err3)
			return
		}

		var i int = 0
		for i != int(k2) {
			f := rand.Intn(1-0+1) + 0
			bufio.NewReader(c).ReadString('\n')
			//fmt.Println("Generate ACK F: " + strconv.Itoa(f))
			//Mengirimkan Acknowledgement
			if f == 0 {
				b := "ACK Lost"
				fmt.Println("Paket Lost")
				fmt.Println(" Discard Paket " + strconv.Itoa(i) + " sampai " + strconv.Itoa(int(k2)))
				//var lo int = i
				// for lo <= int(k2) {
				// 	fmt.Println(" Discard Paket " + strconv.Itoa(lo))
				// 	lo = lo + 1
				// }
				fmt.Fprintf(c, b+"\n") //Kirim Paket
			} else if f == 1 {
				b := "ACK " + strconv.Itoa(i)
				fmt.Println("Menerima Paket " + strconv.Itoa(i) + "\n Mengirimkan Acknowledgement " + strconv.Itoa(i))
				fmt.Fprintf(c, b+"\n")
				i = i + 1
			}
		}
		fmt.Print("Pesan Diterima :", m)
	}
}
