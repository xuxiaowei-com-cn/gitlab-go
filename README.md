# GitLab Go

## 开发命令

### get

```shell
go get -u github.com/urfave/cli/v2
```

### mod

```shell
go mod tidy
```

### run

```shell
go run main.go
```

### run help

```shell
$ go run main.go help
NAME:
   boom - make an explosive entrance

USAGE:
   boom [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### test

```shell
go test -v
```

### build

```shell
go build main.go
```
