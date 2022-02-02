package data

import "go-grpc/grpc_application"

var (
	DataSet []*grpc_application.DataSet = []*grpc_application.DataSet{
		{
			Idx: 1,
			Name: "one",
		},
		{
			Idx: 2,
			Name: "two",
		},
		{
			Idx: 3,
			Name: "three",
		},
	}
)