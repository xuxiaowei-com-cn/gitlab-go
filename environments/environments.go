package environments

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Environments 环境 API https://docs.gitlab.cn/jh/api/environments.html
func Environments() *cli.Command {
	return &cli.Command{
		Name:    "environments",
		Aliases: []string{"env"},
		Usage:   "环境 API，中文文档：https://docs.gitlab.cn/jh/api/environments.html",
		Flags: append(flag.Common(), flag.Id(false), flag.EnvName(false), flag.EnvExternalUrl(), flag.EnvTier(),
			flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			List(),
			Create(),
		},
	}
}
