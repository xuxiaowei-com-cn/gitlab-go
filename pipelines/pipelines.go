package pipelines

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Pipelines 流水线 API https://docs.gitlab.cn/jh/api/pipelines.html
func Pipelines() *cli.Command {
	return &cli.Command{
		Name:    "pipeline",
		Aliases: []string{"pipelines", "pl"},
		Usage:   "流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html",
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目流水线",
				Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
					flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var sort = context.String(constant.Sort)
					var id = context.String(constant.Id)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var printJson = context.Bool(constant.PrintJson)
					var printTime = context.Bool(constant.PrintTime)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListProjectPipelinesOptions{
						ListOptions: gitlab.ListOptions{
							Page:    page,
							PerPage: perPage,
						},
						Sort: &sort,
					}
					PipelineInfos, response, err := gitClient.Pipelines.ListProjectPipelines(id, opt)
					if err != nil {
						return err
					}
					log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

					fmt.Println("")

					if printJson {
						if printTime {
							for _, pipelineInfo := range PipelineInfos {
								jsonData, err := json.Marshal(pipelineInfo)
								if err != nil {
									panic(err)
								}

								log.Printf("\n%s\n", string(jsonData))
								fmt.Println("")
							}
						} else {
							for _, pipelineInfo := range PipelineInfos {
								jsonData, err := json.Marshal(pipelineInfo)
								if err != nil {
									panic(err)
								}

								fmt.Printf("%s\n", string(jsonData))
								fmt.Println("")
							}
						}
					} else {
						if printTime {
							for _, pipelineInfo := range PipelineInfos {
								log.Printf("ID: %d\n", pipelineInfo.ID)
								log.Printf("IID: %d\n", pipelineInfo.IID)
								log.Printf("ProjectID: %d\n", pipelineInfo.ProjectID)
								log.Printf("Status: %s\n", pipelineInfo.Status)
								log.Printf("CreatedAt: %s\n", pipelineInfo.CreatedAt)

								fmt.Println("")
							}
						} else {
							for _, pipelineInfo := range PipelineInfos {
								fmt.Printf("ID: %d\n", pipelineInfo.ID)
								fmt.Printf("IID: %d\n", pipelineInfo.IID)
								fmt.Printf("ProjectID: %d\n", pipelineInfo.ProjectID)
								fmt.Printf("Status: %s\n", pipelineInfo.Status)
								fmt.Printf("CreatedAt: %s\n", pipelineInfo.CreatedAt)

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
