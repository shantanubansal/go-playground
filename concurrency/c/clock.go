package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3241")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("cl2")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}
