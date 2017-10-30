package main


import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Pornire server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	ln1, _ := net.Listen("tcp", ":8082")
	conn1, _ := ln1.Accept()

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Locul de la clientul 1:", string(message))
		message1, _ := bufio.NewReader(conn1).ReadString('\n')
		fmt.Print("Locul de la clientul 2:", string(message1))

		if(message == message1){
			conn.Write([]byte("Locul v-a fost alocat cu succes\n"))
			conn1.Write([]byte("Va rugam alegeti alt numar\n"))
			message2, _ := bufio.NewReader(conn1).ReadString('\n')
			fmt.Print("Noul loc de la clientul 2:", string(message2))
			if(message != message2){
				conn1.Write([]byte("Locul v-a fost alocat cu succes\n"))
			}
		}
		if(message != message1){
		conn.Write([]byte("Locul v-a fost alocat cu succes\n"))
		conn1.Write([]byte("Locul v-a fost alocat cu succes\n"))
		}


		//laborator 3
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message Primit:", string(message))
		//
		//newmessage := "";
		//stringSlice := strings.Split(message, " ")
		//
		//for i := 0; i < len(stringSlice); i++ {
		//	cuvant := stringSlice[i]
		//	a := cuvant[:1]
		//	ok := 1
		//	if ok == 1{
		//		newmessage += strings.Replace(stringSlice[i][:1], a, strings.ToUpper(a), -1)
		//		newmessage += stringSlice[i][1:len(stringSlice[i])]
		//		newmessage += " "
		//	}
		//}
		//
		//conn.Write([]byte(newmessage + "\n"))
}
