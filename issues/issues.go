package issues

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Issues 议题 API https://docs.gitlab.cn/jh/api/issues.html
func Issues() *cli.Command {
	return &cli.Command{
		Name:    "issue",
		Aliases: []string{"issues"},
		Usage:   "议题 API，中文文档：https://docs.gitlab.cn/jh/api/issues.html",
		Flags:   append(flag.Common(), flag.Page(), flag.PerPage()),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出议题",
				Flags: append(flag.CommonTokenRequired(), flag.Page(), flag.PerPage()),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListIssuesOptions{}
					issues, response, err := gitClient.Issues.ListIssues(opt)
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}

					fmt.Println("")

					for index, issue := range issues {
						jsonData, err := json.Marshal(issue)
						if err != nil {
							panic(err)
						}
						log.Printf("Index: %d: \n%s\n", index, string(jsonData))
						fmt.Println("")
					}

					return nil
				},
			},
		},
	}
}
