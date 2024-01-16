package project_level_variables

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// ProjectLevelVariables 项目级别 CI/CD 变量 API https://docs.gitlab.cn/jh/api/project_level_variables.html
func ProjectLevelVariables() *cli.Command {
	return &cli.Command{
		Name:    "project-level-variables",
		Aliases: []string{"project-level-variable", "plv"},
		Usage:   "项目级别 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/project_level_variables.html",
		Flags:   append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
