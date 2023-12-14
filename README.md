# SE07-ATSA

Software Engineering, 7th semester - Advanced Topics in Software Architecture

# Run the system

Start the system

```shell
docker-compose up -d
```

Start the "robots"

```shell
docker-compose -f docker-compose-robots.yaml up -d
```

The frontend will be available [here](http://localhost:3000)

## MQTT

Installing the necessary dependecies for development

- C#:

```shell
dotnet add package MQTTnet
```

- GO:

```shell
go get github.com/eclipse/paho.mqtt.golang
```
