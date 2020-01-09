## Golang and MUX router REST API

## Install Golang - Linux

```bash
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

vim ~/.profile

(append :/usr/local/go/bin to PATH)
```

## Install Mux library

```bash
go get github.com/gorilla/mux
```

## Install Firestore library

```bash
go get cloud.google.com/go/firestore
```

## Build

```bash
go build
```

## Run

```bash
go run .
```

```bash
go run *.go
```
