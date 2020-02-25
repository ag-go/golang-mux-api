## Golang and MUX router REST API

## Install Golang - Linux

```bash
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

vim ~/.profile

(append :/usr/local/go/bin to PATH)
```

## Go Modules Setup

```bash
go mod init gitlab.com/pragmaticreviews/golang-mux-api
```

## Export Environment variable 

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
```

## How to get the private key JSON file:
## From the Firebase Console: Project Overview -> Project Settings -> Service Accounts -> Generate new private key

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
