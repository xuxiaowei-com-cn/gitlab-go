package projects

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Projects 项目 API https://docs.gitlab.cn/jh/api/projects.html
func Projects() *cli.Command {
	return &cli.Command{
		Name:    "project",
		Aliases: []string{"projects", "p"},
		Usage:   "项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html",
		Flags:   append(flag.Common(), flag.Sort()),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出所有项目",
				Flags: append(flag.Common(), flag.Sort()),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var sort = context.String(constant.Sort)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListProjectsOptions{
						Sort: &sort,
					}
					projects, response, err := gitClient.Projects.ListProjects(opt)
					fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}

					for index, project := range projects {
						fmt.Printf("Index: %d,\t ID: %d,\t Path: %s,\t Name: %s\n", index, project.ID, project.Path, project.Name)
					}

					return nil
				},
			},
		},
	}
}
