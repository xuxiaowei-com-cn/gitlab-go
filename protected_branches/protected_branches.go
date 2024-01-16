package protected_branches

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// ProtectedBranches 受保护的分支 API https://docs.gitlab.cn/jh/api/protected_branches.html
func ProtectedBranches() *cli.Command {
	return &cli.Command{
		Name:    "protected-branches",
		Aliases: []string{"pb"},
		Usage:   "受保护的分支 API，中文文档：https://docs.gitlab.cn/jh/api/protected_branches.html",
		Flags: append(flag.Common(), flag.Id(false), flag.ProtectedBranchesSearch(),
			flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
