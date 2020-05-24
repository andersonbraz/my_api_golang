# Contêiner MongoDB + API Golang

## Activity

Evolve your Golang service to expose a /info API to HTTP GET requests, where the JSON content that gets returned is obtained from a MongoDB entry.

[GET http://localhost:8030/api/infos](http://localhost:8030/api/infos)

## Step 1 - Create Image

```shell
docker build -t mongo-client:1.0 . -f Dockerfile
```

## Step 2 - Up Containers

```shell
docker-compose -f "docker-compose.yml" up -d --build
```

## Step 3 - Populate Database

```shell
go run dump.go
```

## References
