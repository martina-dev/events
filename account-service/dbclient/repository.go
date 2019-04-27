package dbclient

import (
	//import the generated protobuf code
	pb "github.com/events/account-service/proto/account"
)

//Repository interface holds all the method declarations that can be performed when writing to a
//Database Instance. You are able to implement any database that conforms to the Repository interface
//and use it for development and testing
type Repository interface {
	Create(*pb.Account) (*pb.Account, error)
}

type Repository struct {
}
