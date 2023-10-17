package jobs

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
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
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Sort(), flag.Id(false), flag.Scope(ScopeValue, ScopeUsage)),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}

func ListProjectJobs(gitClient *gitlab.Client, pid interface{}, scope []string, page int, perPage int) ([]*gitlab.Job, *gitlab.Response, error) {
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
	jobs, response, err := gitClient.Jobs.ListProjectJobs(pid, opt)
	return jobs, response, err
}
