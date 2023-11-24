# Scheduler System

## TODO

- [ ] Containerize
- [ ] Data types

## Go

### Dependencies

[MQTT](https://github.com/eclipse/paho.mqtt.golang)

[Gin](https://github.com/gin-gonic/gin)

[Gorm](https://github.com/go-gorm/gorm)

````shell
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
````

Run dev

````shell
go run .
````

Run tests

````shell
go test -v ./tests
````

Build executable

````shell
go build .
````
