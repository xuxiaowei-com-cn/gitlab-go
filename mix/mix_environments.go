package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

func Environments() *cli.Command {
	return &cli.Command{
		Name:    "mix-create-environments",
		Aliases: []string{"mix-create-environment", "mix-create-env"},
		Usage:   "创建新环境（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false),
			flag.SkipProjectPath(), flag.SkipProjectWikiPath(), flag.AllowFailure(),
			flag.EnvName(false), flag.EnvExternalUrl(), flag.EnvTier(),
			flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			EnvironmentsCreateAll(),
		},
	}
}
