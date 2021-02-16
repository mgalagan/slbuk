module github.com/mgalagan/slbuk/gateway

go 1.15

replace github.com/mgalagan/slbuk/grpc => ../grpc

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/mgalagan/slbuk/grpc v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.35.0
)
