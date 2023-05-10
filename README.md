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

#### Build service in docker:
``` bash 
make service-build
```

#### Run service in docker:
``` bash 
make service-start
```

#### Stop service in docker:
``` bash 
make service-stop
```

#### Restart service in docker:
``` bash 
make service-restart
```

