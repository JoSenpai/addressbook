package main

import (
	"fmt"

	pb "github.com/josenpai/addressbook/proto/tutorial"
)

func main() {

	fmt.Printf("Hello World")

	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

}
