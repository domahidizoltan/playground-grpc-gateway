# playground-grpc-gateway

A playground project to learn how to apply easily Http Gateway over gRPC communication.  
The sample service will receive a car object and will respond an echo with the same car object with a timestamp.  

<br/>

## gRPC service  

### Install  

Install locally these tools:
protoc 3.17.3
grpcurl 1.8.2

Install these Go dependencies: 
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

### Test

Optionally generate gRPC files from proto file:
```
protoc \
  --go_out=generated \
  --go_opt=paths=source_relative \
  --go-grpc_out=generated \
  --go-grpc_opt=paths=source_relative car/carservice.proto
```

Start the service:
```
go run .
```

Call the unary echo service:
```
grpcurl \
  -plaintext \
  -proto=car/carservice.proto \
  -H "x-api-key: 123" \
  -d "{\"brand\": \"Toyota\", \"model\": \"Aygo\", \"year\": 2007}" \
  localhost:2000 car.CarService/EchoCar
```

Responds:
```
{
  "car": {
    "brand": "Toyota",
    "model": "Aygo",
    "year": 2007
  },
  "timestamp": "1631037837"
}
```

Call the streaming echo service:
```
grpcurl \
  -plaintext \
  -proto=car/carservice.proto \
  -H "x-api-key: 123" \
  -d "{\"brand\": \"Toyota\", \"model\": \"Aygo\", \"year\": 2007} {\"brand\": \"VW\", \"model\": \"Passat\", \"year\": 2010}" \
  localhost:2000 car.CarService/EchoCars
```

Responds:
```
{
  "car": {
    "brand": "Toyota",
    "model": "Aygo",
    "year": 2007
  },
  "timestamp": "1631038031"
}
{
  "car": {
    "brand": "VW",
    "model": "Passat",
    "year": 2010
  },
  "timestamp": "1631038031"
}
```

<br/>

## gRPC gateway  

### Install  

Install these Go dependencies:  
```
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
```

### Test

Optionally generate gRPC gateway files from proto file:
```
protoc \
  --grpc-gateway_out=generated \
  --grpc-gateway_opt=paths=source_relative,generate_unbound_methods=true car/carservice
```

Start the service:
```
go run .
```

Call the unary echo service:
```
curl \
  -X POST \
  -H "grpc-metadata-x-api-key: 123" \
  -d "{\"brand\": \"Toyota\", \"model\": \"Aygo\", \"year\": 2007}" \
  http://localhost:2020/car.CarService/EchoCar 
```

Responds:
```
{"car":{"brand":"Toyota", "model":"Aygo", "year":2007}, "timestamp":"1631038782"}%     
```

Call the streaming echo service:
```
curl \
  -X POST \
  -H "grpc-metadata-x-api-key: 123" \
  -d "{\"brand\": \"Toyota\", \"model\": \"Aygo\", \"year\": 2007} {\"brand\": \"VW\", \"model\": \"Passat\", \"year\": 2010}" \
  http://localhost:2020/car.CarService/EchoCars
```

Responds:
```
{"result":{"car":{"brand":"Toyota","model":"Aygo","year":2007},"timestamp":"1631038807"}}
{"result":{"car":{"brand":"VW","model":"Passat","year":2010},"timestamp":"1631038807"}}
```
