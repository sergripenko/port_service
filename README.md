# port_service

### 1.Download and run the app
#### Clone repository:
``` bash 
git clone https://github.com/sergripenko/port_service.git
```

#### Install dependencies:
``` bash 
go mod vendor
```

Put file ports.json in repository.

#### Open project repo and run server:
``` bash 
go run cmd/main.go
```

#### Linters check:
``` bash 
make golangci
```

#### Run tests:
``` bash 
make test
```

#### Build dockerfile:
``` bash 
make build
```

#### Run dockerfile:
``` bash 
make run
```