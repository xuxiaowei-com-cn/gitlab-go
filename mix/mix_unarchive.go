package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Unarchive 取消归档
func Unarchive() *cli.Command {
	return &cli.Command{
		Name:  "mix-unarchive",
		Usage: "取消归档（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false)),
		Subcommands: []*cli.Command{
			UnarchiveAll(),
		},
	}
}
