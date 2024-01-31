# README

## Simple App

This is a Go application designed to simulate performance issues. It includes an APM server agent that works with the Elasticsearch, Kibana, and APM server stack.

## Build App

```bash
export GOOS=linux
export GOARCH=amd64
go build main.go
```

## Run App

```bash
./main
```

``` bash
❯ ./main
2024/01/31 15:58:30 Serveur démarré sur http://localhost:8080
```

## Access

```bash
curl http://localhost:8080
```