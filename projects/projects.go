package projects

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
)

// Projects 项目 API https://docs.gitlab.cn/jh/api/projects.html
func Projects() *cli.Command {
	return &cli.Command{
		Name:    "projects",
		Aliases: []string{"p"},
		Usage:   "项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html",
		Action: func(context *cli.Context) error {
			var token = context.String("token")
			var baseUrl = context.String("baseUrl")
			if baseUrl == "" {
				baseUrl = "https://jihulab.com/api/v4"
			}
			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListProjectsOptions{}
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "token",
				Usage: "your_access_token",
			},
			&cli.StringFlag{
				Name:  "baseUrl",
				Usage: "实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4",
			},
		},
	}
}
