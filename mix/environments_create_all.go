package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/environments"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"log"
)

// EnvironmentsCreateAll 所有项目创建新环境
func EnvironmentsCreateAll() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "所有项目创建新环境",
		Flags: append(flag.CommonTokenRequired(), flag.Owned(true),
			flag.EnvName(false), flag.EnvExternalUrl(), flag.EnvTier(), flag.AllowFailure(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)
			var name = context.String(constant.EnvName)
			var externalURL = context.String(constant.EnvExternalUrl)
			var tier = context.String(constant.EnvTier)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)
			var allowFailure = context.Bool(constant.AllowFailure)

			projectList, err := projects.ListProjects(owned, token, baseUrl, 1, 100)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = environments.CreateEnvironment(baseUrl, token, project.ID, name, externalURL, tier, printJson, printTime, allowFailure)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
