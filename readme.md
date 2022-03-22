# yzhanGoRun
Let's Go
## Setup
```shell
# Download form https://go.dev/dl/
wget --no-check-certificate https://go.dev/dl/go1.18.linux-amd64.tar.gz
# tar
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
vim ~/.zshrc
# Set PATH
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go/
export GOPATH=~/go/
```
## Check
```shell
go version
# go version gox.xx linux/amd64
```
## Hello World
### Create
```shell
mkdir hello
cd hello
go mod init hello
vim hello.go
```
### Edit
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
### Run
```shell
cd hello
go run .
# Hello, World!
```