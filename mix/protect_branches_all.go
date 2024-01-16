package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"github.com/xuxiaowei-com-cn/gitlab-go/protected_branches"
	"log"
)

// ProtectBranchesAll 保护所有仓库分支
func ProtectBranchesAll() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "保护所有仓库分支",
		Flags: append(flag.CommonTokenRequired(), flag.Owned(true),
			flag.BranchName(true), flag.PushAccessLevel(), flag.MergeAccessLevel(),
			flag.UnprotectAccessLevel(), flag.AllowForcePush(), flag.CodeOwnerApprovalRequired(),
			flag.AllowFailure(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)

			var name = context.String(constant.BranchName)
			var pushAccessLevel = context.Int(constant.PushAccessLevel)
			var mergeAccessLevel = context.Int(constant.MergeAccessLevel)
			var unprotectAccessLevel = context.Int(constant.UnprotectAccessLevel)
			var allowForcePush = context.Bool(constant.AllowForcePush)
			var codeOwnerApprovalRequired = context.Bool(constant.CodeOwnerApprovalRequired)

			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)
			var allowFailure = context.Bool(constant.AllowFailure)

			projectList, err := projects.ListProjects(owned, token, baseUrl, 1, 100)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = protected_branches.ProtectRepositoryBranches(baseUrl, token, project.ID, name,
					pushAccessLevel, mergeAccessLevel, unprotectAccessLevel, allowForcePush, codeOwnerApprovalRequired,
					printJson, printTime, allowFailure)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
