package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"log"
)

// ArchiveAll 归档所有项目
func ArchiveAll() *cli.Command {
	return &cli.Command{
		Name:  "all",
		Usage: "归档所有项目（混合命令，多接口命令）",
		Flags: append(flag.CommonTokenRequired(), flag.Owned(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)

			projectList, err := projects.ListProjects(&owned, token, baseUrl, 1, 100)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = projects.ArchiveProject(token, baseUrl, project.PathWithNamespace)

				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
