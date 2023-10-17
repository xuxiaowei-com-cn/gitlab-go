package access_requests

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// AccessRequests 群组和项目访问请求 API https://docs.gitlab.cn/jh/api/access_requests.html
func AccessRequests() *cli.Command {
	return &cli.Command{
		Name:    "access-request",
		Aliases: []string{"access-requests", "ar"},
		Usage:   "群组和项目访问请求 API，中文文档：https://docs.gitlab.cn/jh/api/access_requests.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			Group(),
			Project(),
		},
	}
}
