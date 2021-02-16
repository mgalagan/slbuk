module github.com/mgalagan/slbuk/swgate

go 1.15

replace github.com/mgalagan/slbuk/grpc => ../grpc

require (
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/loads v0.20.2
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/spec v0.20.3
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/jessevdk/go-flags v1.4.0
	github.com/mgalagan/slbuk/grpc v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/grpc v1.35.0 // indirect
)
