package job_artifacts

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Delete 删除作业产物 https://docs.gitlab.cn/jh/api/job_artifacts.html#%E5%88%A0%E9%99%A4%E4%BD%9C%E4%B8%9A%E4%BA%A7%E7%89%A9
func Delete() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"rm"},
		Usage:   "删除作业产物",
		Flags:   append(flag.Common(), flag.Id(true), flag.JobId(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var jobId = context.Int(constant.JobId)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			log.Printf("Delete ProjectId: %s, JobId: %d Artifacts\n", id, jobId)
			response, err := gitClient.Jobs.DeleteArtifacts(id, jobId)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			return nil
		},
	}
}
