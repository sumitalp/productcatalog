Product Catalogue
=================

## Author
__Ahsanuzzaman Khan__

This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Echo** including CRUD operations, authentication, routing, pagination, and more.

## Getting started

### Install SQLite3

### Install Golang (go1.13+)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install)
Also make sure you have installed a go1.13+ version.

### Environment Config

make sure your ~/.*shrc have those variable:

```bash
➜  echo $GOPATH
/Users/ahsan/go
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/Users/ahsan/test//bin:/usr/local/go/bin
```

For more info and detailed instructions please check this guide: [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Install dep
https://golang.github.io/dep/docs/installation.html

### Clone the repository

Clone this repository:

```bash
➜ git clone https://github.com/sumitalp/productcatalog.git
```

Or simply use the following command which will handle cloning the repo:

```bash
➜ go get -u -v github.com/sumitalp/productcatalog
```

Switch to the repo folder

```bash
➜ cd $GOPATH/src/github.com/sumitalp/productcatalog
```

### Install dependencies

```bash
➜ go mod download
```

### Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```

### Tests

```bash
➜ go test -v -race ./...
```
