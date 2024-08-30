package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	reader := bufio.NewReader(os.Stdin)
	for {   fmt.Print("Reverse_shell:~$ ")
			s, _ := reader.ReadString('\n')
			conn.Write([]byte(s))
			//conn.Write([]byte(s+"\n"))
			response := make([]byte, 512000)
			n,_ := conn.Read(response)
			fmt.Print(string(response[:n]))
	}
}

