package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Delete 删除（混合命令，多接口命令）
func Delete() *cli.Command {
	return &cli.Command{
		Name:    "mix-delete",
		Aliases: []string{"mix-rm"},
		Usage:   "删除（混合命令，多接口命令）",
		Flags:   append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.Id(false), flag.IIdRange(false)),
		Subcommands: []*cli.Command{
			DeleteArtifacts(),
			DeleteAllArtifacts(),
			DeleteJobs(),
			DeleteAllJobs(),
		},
	}
}
