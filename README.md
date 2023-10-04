# GitLab Go 脚手架

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go get -u github.com/urfave/cli/v2
go get -u github.com/xanzy/go-gitlab
go get -u github.com/xuxiaowei-com-cn/git-go@main
go get -u gopkg.in/yaml.v3
```

### mod

```shell
go mod tidy
```

```shell
go mod download
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
$ go run main.go --help
NAME:
   gitlab-go - 基于 Go 语言开发的 GitLab 命令行工具

USAGE:
   gitlab-go [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
   projects, p        项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html
   pipelines, pl      流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html
   jobs, j            作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html
   job-artifacts, ja  作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

- [projects - 项目 API](https://docs.gitlab.cn/jh/api/projects.html)

```shell
$ go run main.go projects --help
NAME:
   gitlab-go projects - 项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html

USAGE:
   gitlab-go projects command [command options] [arguments...]

COMMANDS:
   list     列出所有项目
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
   --token value     your_access_token
   --sort value      按照 asc 或者 desc 排序 (default: "desc")
   --help, -h        show help
```

- [pipelines - 流水线 API](https://docs.gitlab.cn/jh/api/pipelines.html)

```shell
$ go run main.go pipelines --help
NAME:
   gitlab-go pipelines - 流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html

USAGE:
   gitlab-go pipelines command [command options] [arguments...]

COMMANDS:
   list     列出项目流水线
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
   --token value     your_access_token
   --sort value      按照 asc 或者 desc 排序 (default: "desc")
   --id value        项目 ID 或 URL 编码的路径
   --help, -h        show help
```

- [jobs - 作业 API](https://docs.gitlab.cn/jh/api/jobs.html)

```shell
$ go run main.go jobs --help
NAME:
   gitlab-go jobs - 作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html

USAGE:
   gitlab-go jobs command [command options] [arguments...]

COMMANDS:
   list     列出项目作业
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
   --token value     your_access_token
   --sort value      按照 asc 或者 desc 排序 (default: "desc")
   --help, -h        show help
```

- [job-artifacts - 作业产物 API](https://docs.gitlab.cn/jh/api/job_artifacts.html)

```shell
$ go run main.go job-artifacts --help
NAME:
   gitlab-go job-artifacts - 作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html

USAGE:
   gitlab-go job-artifacts command [command options] [arguments...]

COMMANDS:
   get                   获取作业产物（未完成）
   download, dl          下载产物归档文件（未完成）
   delete, rm            删除作业产物
   delete-project, rm-p  删除项目产物
   help, h               Shows a list of commands or help for one command

OPTIONS:
   --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
   --token value     your_access_token
   --id value        项目 ID 或 URL 编码的路径
   --job-id value    作业 ID
   --help, -h        show help
```

### test

```shell
go test ./... -v
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
        go build -o buildinfo/buildinfo.exe buildinfo/buildinfo.go
        GOOS=windows GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo.exe now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo.exe commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo.exe commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo.exe commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo.exe commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo.exe commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo.exe commitTag) -X main.GoVersion=$(buildinfo/buildinfo.exe goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-windows-amd64.exe .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo.exe buildinfo/buildinfo.go
        GOOS=windows GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo.exe now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo.exe commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo.exe commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo.exe commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo.exe commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo.exe commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo.exe commitTag) -X main.GoVersion=$(buildinfo/buildinfo.exe goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-windows-arm64.exe .
        ```

- Linux
    - amd64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-linux-amd64 .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-linux-arm64 .
        ```

- Darwin
    - amd64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=darwin GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-darwin-amd64 .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=darwin GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState= -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform= -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-darwin-arm64 .
        ```
