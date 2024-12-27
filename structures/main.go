package main

import (
	"fmt"

	"wakisa.com/magazine"
)

func main() {
	address := magazine.Address{Street: "123 Oak St", City: "Omala", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address = address //calling annonymous struct field in magazine package
	fmt.Println(subscriber.Address)

}
