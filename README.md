# Datadog Demo
A simple api to practice Datadog and Monitoring concepts

## How to start containers
```shell
 DD_API_KEY=key DD_APPLICATION_ID=id docker-compose up -d
```

## How to run app locally
```shell
DATADOG_AGENT_URI=localhost:8125  PORT=8080 MONGO_URI=mongodb://localhost:27017 go run cmd/api/main.go
```