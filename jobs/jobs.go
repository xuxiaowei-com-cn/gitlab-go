package jobs

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

const (
	ScopeValue = ""
	ScopeUsage = "要显示的作业范围。以下之一或数组：created、pending、running、failed、success、canceled、skipped、waiting_for_resource 或 manual。范围如果未提供，则返回所有作业。"
)

// Jobs 作业 API https://docs.gitlab.cn/jh/api/jobs.html
func Jobs() *cli.Command {
	return &cli.Command{
		Name:    "job",
		Aliases: []string{"jobs", "j"},
		Usage:   "作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.Sort(), flag.Id(false),
			flag.Scope(ScopeValue, ScopeUsage)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目作业",
				Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.Sort(), flag.Id(true),
					flag.Scope(ScopeValue, ScopeUsage)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var scope = context.StringSlice(constant.Scope)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					var bsvs []gitlab.BuildStateValue
					for _, str := range scope {
						bsv := gitlab.BuildStateValue(str)
						bsvs = append(bsvs, bsv)
					}

					opt := &gitlab.ListJobsOptions{
						Scope: &bsvs,
						ListOptions: gitlab.ListOptions{
							Page:    page,
							PerPage: perPage,
						},
					}
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
