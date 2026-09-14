package main 

import (
	"fmt"
	"log"
	"net"
)


func main() {
	interfaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}
	
	for _, iface := range interfaces {
		fmt.Printf(
			"Index: %d\tName: %s\tMAC: %s\tFlags: %s\n",
			iface.Index,
			iface.Name,
			iface.HardwareAddr,
			iface.Flags,
		)
	}
}
