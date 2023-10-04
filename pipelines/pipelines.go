package pipelines

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

// Pipelines 流水线 API https://docs.gitlab.cn/jh/api/pipelines.html
func Pipelines() *cli.Command {
	return &cli.Command{
		Name:    "pipelines",
		Aliases: []string{"pl"},
		Usage:   "流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html",
		Action: func(context *cli.Context) error {
			var token = context.String(constant.Token)
			var baseUrl = context.String(constant.BaseUrl)
			var id = context.String(constant.Id)
			if baseUrl == "" {
				baseUrl = constant.BaseUrlDefault
			}
			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListProjectPipelinesOptions{}
			PipelineInfos, response, err := gitClient.Pipelines.ListProjectPipelines(id, opt)
			fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
			if err != nil {
				return err
			}

			for index, pipelineInfo := range PipelineInfos {
				fmt.Printf("Index: %d,\t ID: %d,\t IID: %d,\t ProjectID: %d,\t Status: %s,\t CreatedAt: %s\n", index, pipelineInfo.ID, pipelineInfo.IID, pipelineInfo.ProjectID, pipelineInfo.Status, pipelineInfo.CreatedAt)
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  constant.Token,
				Usage: "your_access_token",
			},
			&cli.StringFlag{
				Name:  constant.BaseUrl,
				Usage: "实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4",
			},
			&cli.StringFlag{
				Name:     constant.Id,
				Usage:    "项目 ID 或 URL 编码的路径",
				Required: true,
			},
		},
	}
}
