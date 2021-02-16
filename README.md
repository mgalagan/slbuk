## Server
"server" folder contains code of GRPC server 

### Params:
Use --server-address to set server endpoint. Default is "localhost:9001".

### Example usage:
```bash
cd server; go run main.go --server-address="localhost:9001"
```

## gRPC stream client
"client" folder contains example code of GRPC stream client

### Params:
Use --server-address to set grpc server address. Default is "localhost:9001".

### Example usage:
```bash
cd client; go run main.go --server-address="localhost:9001"
```

## protoc-gen-grpc-gateway gRPC to REST gateway
"gateway" folder contains simple gRPC to REST gateway code.

### Params:
Use --gateway-address to set gateway endpoint. Default is "localhost:9002"
Use --grpc-server-address to set grpc server address. Default is "localhost:9001".

### Example usage:
```bash
cd gateway; go run main.go --gateway-address="localhost:9002" --grpc-server-address="localhost:9001"
```

## Swagger gateway
"gateway" folder contains swagger gRPC to REST gateway code.

### Params:
Use --host to set gateway host. Default is "localhost"
Use --port to set gateway host. Default is "9002"
Use --grpc-server-address to set grpc server address. Default is "localhost:9001".

### Example usage:
```bash
cd swgate; go run main.go --host="localhost" --port="9002" --grpc-server-address="localhost:9001"
```
