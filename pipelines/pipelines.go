package pipelines

import (
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
		Flags:   append(flag.Common(), flag.Sort(), flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目流水线",
				Flags: append(flag.Common(), flag.Sort(), flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var sort = context.String(constant.Sort)
					var id = context.String(constant.Id)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListProjectPipelinesOptions{
						Sort: &sort,
					}
					PipelineInfos, response, err := gitClient.Pipelines.ListProjectPipelines(id, opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					for index, pipelineInfo := range PipelineInfos {
						log.Printf("Index: %d,\t ID: %d,\t IID: %d,\t ProjectID: %d,\t Status: %s,\t CreatedAt: %s\n", index, pipelineInfo.ID, pipelineInfo.IID, pipelineInfo.ProjectID, pipelineInfo.Status, pipelineInfo.CreatedAt)
					}

					return nil
				},
			},
		},
	}
}
