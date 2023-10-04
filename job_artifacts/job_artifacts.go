package job_artifacts

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

// JobsArtifacts 作业产物 API https://docs.gitlab.cn/jh/api/job_artifacts.html
func JobsArtifacts() *cli.Command {
	return &cli.Command{
		Name:    "job-artifacts",
		Aliases: []string{"ja"},
		Usage:   "作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html",
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			if baseUrl == "" {
				baseUrl = constant.BaseUrlDefault
			}
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var jobId = context.Int(constant.JobId)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			fmt.Printf("Delete JobId: %d Artifacts\n", jobId)
			response, err := gitClient.Jobs.DeleteArtifacts(id, jobId)
			fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
			if err != nil {
				return err
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    constant.BaseUrl,
				EnvVars: []string{"CI_API_V4_URL"},
				Usage:   "实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4",
			},
			&cli.StringFlag{
				Name:  constant.Token,
				Usage: "your_access_token",
			},
			&cli.StringFlag{
				Name:  constant.Request,
				Usage: "请求方式",
			},
			&cli.StringFlag{
				Name:  constant.Sort,
				Value: constant.SortDefault,
				//Usage: "按照 asc 或者 desc（默认：desc）对流水线排序",
			},
			&cli.StringFlag{
				Name:     constant.Id,
				Usage:    "项目 ID 或 URL 编码的路径",
				Required: true,
			},
			&cli.StringFlag{
				Name:     constant.JobId,
				Usage:    "作业 ID",
				Required: true,
			},
		},
	}
}
