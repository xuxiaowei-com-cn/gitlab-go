package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Export 导出
func Export() *cli.Command {
	return &cli.Command{
		Name:  "mix-export",
		Usage: "导出（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false),
			flag.ExportFolder(false), flag.SkipProjectPath(), flag.SkipProjectWikiPath(), flag.AutoSkipExistFolder()),
		Subcommands: []*cli.Command{
			ExportAll(),
		},
	}
}
