package main 

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"golang.org/x/sys/unix"
)

func htons(v uint16) uint16 {
	return (v<<8)&0xff00 | v >> 8
}


func main() {
	ifaceName := ""

	iface, err := net.InterfaceByName(ifaceName)


	if err != nil {
		log.Fatal(err)
	}

	fd, err := unix.Socket(
		unix.AF_PACKET,
		unix.SOCK_RAW,
		int(htons(unix.ETH_P_ALL)),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer unix.Close(fd)

	addr := &unix.SockaddrLinklayer{
		Protocol: htons(unix.ETH_P_ALL),
		Ifindex: iface.Index,
	}

	if err := unix.Bind(fd, addr); err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 65535)

	for {
		n, _, err := unix.Recvfrom(fd, buffer, 0)
		if err != nil {
			log.Fatal(err)
		}

		packet := buffer[:n]

		fmt.Printf("\nCaptured %d bytes\n", n)
		fmt.Print(hex.Dump(packet))
	}
}
