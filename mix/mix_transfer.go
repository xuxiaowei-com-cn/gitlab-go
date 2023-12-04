package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Transfer 转移
func Transfer() *cli.Command {
	return &cli.Command{
		Name:  "mix-transfer",
		Usage: "转移（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false),
			flag.NamespaceSource(false), flag.NamespaceTargetFlag(false), flag.SkipProjectPath()),
		Subcommands: []*cli.Command{
			TransferAll(),
		},
	}
}
