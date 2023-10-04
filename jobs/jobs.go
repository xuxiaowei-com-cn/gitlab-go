package jobs

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

// Jobs 作业 API https://docs.gitlab.cn/jh/api/jobs.html
func Jobs() *cli.Command {
	return &cli.Command{
		Name:    "jobs",
		Aliases: []string{"j"},
		Usage:   "作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html",
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			if baseUrl == "" {
				baseUrl = constant.BaseUrlDefault
			}
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListJobsOptions{}
			jobs, response, err := gitClient.Jobs.ListProjectJobs(id, opt)
			fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
			if err != nil {
				return err
			}

			for index, job := range jobs {
				fmt.Printf("Index: %d,\t ID: %d,\t Name: %s,\t ProjectID: %d,\t Status: %s,\t CreatedAt: %s\n", index, job.ID, job.Name, job.Project.ID, job.Status, job.CreatedAt)
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
				Name:  constant.Sort,
				Value: constant.SortDefault,
				//Usage: "按照 asc 或者 desc（默认：desc）对流水线排序",
			},
			&cli.StringFlag{
				Name:     constant.Id,
				Usage:    "项目 ID 或 URL 编码的路径",
				Required: true,
			},
		},
	}
}
