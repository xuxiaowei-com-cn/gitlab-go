package pipelines

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Pipelines 流水线 API https://docs.gitlab.cn/jh/api/pipelines.html
func Pipelines() *cli.Command {
	return &cli.Command{
		Name:    "pipeline",
		Aliases: []string{"pipelines", "pl"},
		Usage:   "流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html",
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(), flag.Recursion(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
