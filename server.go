package main
import (
	"io"
	"fmt"
	"net"
	"strconv"
	"strings"
	"os/exec"
)
func exe(b []byte) string {
	args := strings.Fields(string(b))
	cmd := exec.Command(args[0], args[1:]...)
	output,_ := cmd.Output()
	return string(output)
}
func tinh(b []byte) string {
    a:=[2]int{0,0}
    dem:=0
    var c byte
    for i:=0;i<len(b);i++{
        tmp:=0
        for i<len(b)&&'0'<=b[i]&&b[i]<='9'{
        tmp=tmp*10+int(b[i]-'0')
        i++
     }
	 if dem < 2 {
		a[dem] = tmp
		dem++
	}
	if i < len(b) {
		c = b[i]
	}
}
if c == '+' && dem == 2 {
	return strconv.Itoa(a[0] + a[1])
}
if c == '-' && dem == 2 {
	return strconv.Itoa(a[0] - a[1])
}
if c == '*' && dem == 2 {
	return strconv.Itoa(a[0] * a[1])
}
if c == '/' && dem == 2 {
	return strconv.Itoa(a[0] / a[1])
}
return ""
}
func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
	size, err := conn.Read(b[0:])
	if err == io.EOF {
		fmt.Println("End")
		break
 	}
	 input := string(b[:size])
	 input = input[:len(input)-1]
	//fmt.Print(tinh(b))
	conn.Write([]byte(exe([]byte(input))+"\n"))
 	}
}
func main() {
	listener,_ := net.Listen("tcp", ":1234")
	fmt.Println("Listening on 0.0.0.0:1234")
	for {
	conn,_ := listener.Accept()
	fmt.Println("Received connection")
	go echo(conn)
	}
	//fmt.Println(tinh([]byte("456+8")))
}