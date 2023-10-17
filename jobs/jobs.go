package jobs

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Jobs 作业 API https://docs.gitlab.cn/jh/api/jobs.html
func Jobs() *cli.Command {
	return &cli.Command{
		Name:    "job",
		Aliases: []string{"jobs", "j"},
		Usage:   "作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html",
		Flags:   append(flag.Common(), flag.Sort()),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目作业",
				Flags: append(flag.Common(), flag.Sort(), flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListJobsOptions{}
					jobs, response, err := gitClient.Jobs.ListProjectJobs(id, opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					for index, job := range jobs {
						log.Printf("Index: %d,\t ID: %d,\t Name: %s,\t ProjectID: %d,\t Status: %s,\t CreatedAt: %s\n", index, job.ID, job.Name, job.Project.ID, job.Status, job.CreatedAt)
					}

					return nil
				},
			},
		},
	}
}
