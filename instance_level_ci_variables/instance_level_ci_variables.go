package instance_level_ci_variables

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// InstanceLevelCiVariables 实例级 CI/CD 变量 API https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html
func InstanceLevelCiVariables() *cli.Command {
	return &cli.Command{
		Name:    "instance-level-ci-variables",
		Aliases: []string{"instance-level-ci-variable", "ilcv"},
		Usage:   "实例级 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html",
		Flags:   append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
