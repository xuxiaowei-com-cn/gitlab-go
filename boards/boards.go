package boards

import (
	"encoding/json"
	"fmt"
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
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目议题板",
				Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
					flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var printJson = context.Bool(constant.PrintJson)
					var printTime = context.Bool(constant.PrintTime)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListIssueBoardsOptions{}
					issueBoards, response, err := gitClient.Boards.ListIssueBoards(id, opt)
					if err != nil {
						return err
					}
					log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

					fmt.Println("")

					if printJson {
						if printTime {
							for _, issueBoard := range issueBoards {
								jsonData, err := json.Marshal(issueBoard)
								if err != nil {
									panic(err)
								}

								log.Printf("\n%s\n", string(jsonData))
								fmt.Println("")
							}
						} else {
							for _, issueBoard := range issueBoards {
								jsonData, err := json.Marshal(issueBoard)
								if err != nil {
									panic(err)
								}

								fmt.Printf("%s\n", string(jsonData))
								fmt.Println("")
							}
						}
					} else {
						if printTime {
							for _, issueBoard := range issueBoards {
								log.Printf("ID: %d\n", issueBoard.ID)
								log.Printf("Name: %s\n", issueBoard.Name)
								log.Printf("Project.ID: %d\n", issueBoard.Project.ID)

								// log.Printf("Project: %s\n", issueBoard.Project)
								// log.Printf("User: %s\n", issueBoard.User)

								fmt.Println("")
							}
						} else {
							for _, issueBoard := range issueBoards {
								fmt.Printf("ID: %d\n", issueBoard.ID)
								fmt.Printf("Name: %s\n", issueBoard.Name)
								fmt.Printf("Project.ID: %d\n", issueBoard.Project.ID)

								// log.Printf("Project: %s\n", issueBoard.Project)
								// log.Printf("User: %s\n", issueBoard.User)

								fmt.Println("")
							}
						}
					}

					return nil
				},
			},
		},
	}
}
