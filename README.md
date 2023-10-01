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
go run main.go help
```

```shell
$ go run main.go help
NAME:
   boom - make an explosive entrance

USAGE:
   boom [global options] command [command options] [arguments...]

VERSION:
   v0.0.1-snapshot

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### test

```shell
go test -v
```

### build

```shell
go build
# GOOS=：设置构建的目标操作系统（darwin | freebsd | linux | windows）
# GOARCH=：设置构建的目标操作系统（386 | amd64 | arm | arm64）
# -v：打印编译过程中的详细信息
# -ldflags：设置在编译时传递给链接器的参数
# -ldflags "-s -w -buildid="
#                           -s: 删除符号表信息，减小可执行文件的大小。
#                           -w: 删除调试信息，使可执行文件在运行时不会打印调试信息。
#                           -buildid=: 删除构建ID，使可执行文件在运行时不会打印构建ID。
# -trimpath：去掉所有包含 go path 的路径
# -o：指定构建后输出的文件名
```

- Windows
    - amd64
        ```shell
        GOOS=windows GOARCH=amd64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-windows-amd64.exe .
        ```
    - arm64
        ```shell
        GOOS=windows GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-windows-arm64.exe .
        ```

- Linux
    - amd64
        ```shell
        GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-linux-amd64 .
        ```
    - arm64
        ```shell
        GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-linux-arm64 .
        ```

- Darwin
    - amd64
        ```shell
        GOOS=darwin GOARCH=amd64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-darwin-amd64 .
        ```
    - arm64
        ```shell
        GOOS=darwin GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o gitlab-go-darwin-arm64 .
        ```
