package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Archive 归档
func Archive() *cli.Command {
	return &cli.Command{
		Name:  "mix-archive",
		Usage: "归档（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false)),
		Subcommands: []*cli.Command{
			ArchiveAll(),
		},
	}
}
