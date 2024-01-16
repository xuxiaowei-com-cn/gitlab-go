package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// ProtectBranches 保护仓库分支
func ProtectBranches() *cli.Command {
	return &cli.Command{
		Name:  "mix-protect-branches",
		Usage: "保护仓库分支（混合命令，多接口命令）",
		Flags: append(flag.Common(), flag.Owned(false),
			flag.BranchName(false), flag.PushAccessLevel(), flag.MergeAccessLevel(),
			flag.UnprotectAccessLevel(), flag.AllowForcePush(), flag.CodeOwnerApprovalRequired(),
			flag.PrintJson(), flag.PrintTime()),
		Subcommands: []*cli.Command{
			ProtectBranchesAll(),
		},
	}
}
