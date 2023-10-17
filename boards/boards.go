package boards

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Boards 项目议题板 API https://docs.gitlab.cn/jh/api/boards.html
func Boards() *cli.Command {
	return &cli.Command{
		Name:    "board",
		Aliases: []string{"boards"},
		Usage:   "项目议题板 API，中文文档：https://docs.gitlab.cn/jh/api/boards.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
