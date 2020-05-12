package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		go timeZoneHandle(s[0], s[1])
	}
	time.Sleep(1 * time.Minute)
}
func timeZoneHandle(timeZone string, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
