package boards

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Boards 项目议题板 API https://docs.gitlab.cn/jh/api/boards.html
func Boards() *cli.Command {
	return &cli.Command{
		Name:    "board",
		Aliases: []string{"boards"},
		Usage:   "项目议题板 API，中文文档：https://docs.gitlab.cn/jh/api/boards.html",
		Flags:   append(flag.Common(), flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目议题板",
				Flags: append(flag.Common(), flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListIssueBoardsOptions{}
					issueBoards, response, err := gitClient.Boards.ListIssueBoards(id, opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					for index, issueBoard := range issueBoards {
						jsonData, err := json.Marshal(issueBoard)
						if err != nil {
							panic(err)
						}
						log.Printf("Index: %d: %s\n", index, string(jsonData))
					}

					return nil
				},
			},
		},
	}
}
