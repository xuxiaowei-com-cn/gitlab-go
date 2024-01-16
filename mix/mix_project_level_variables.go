package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// CreateProjectLevelVariables 创建项目级别 CI/CD 变量
func CreateProjectLevelVariables() *cli.Command {
	return &cli.Command{
		Name:  "mix-create-project-level-variables",
		Usage: "创建项目级别 CI/CD 变量（混合命令，多接口命令）",
		Flags: append(flag.Common(),
			flag.VariableKey(false), flag.VariableValue(false), flag.VariableType(), flag.VariableProtected(),
			flag.VariableMasked(), flag.VariableRaw(), flag.VariableEnvironmentScope(),
			flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			ProjectLevelVariablesCreateAll(),
		},
	}
}
