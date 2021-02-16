module github.com/mgalagan/slbuk/server

go 1.15

replace github.com/mgalagan/slbuk/grpc => ../grpc

require (
	github.com/mgalagan/slbuk/grpc v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.35.0
)
