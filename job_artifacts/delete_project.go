package job_artifacts

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// DeleteProject 删除项目产物 https://docs.gitlab.cn/jh/api/job_artifacts.html#%E5%88%A0%E9%99%A4%E9%A1%B9%E7%9B%AE%E4%BA%A7%E7%89%A9
func DeleteProject() *cli.Command {
	return &cli.Command{
		Name:    "delete-project",
		Aliases: []string{"delete-projects", "rm-p"},
		Usage:   "删除项目产物",
		Flags:   append(flag.Common(), flag.Id(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			log.Printf("Delete ProjectId: %s Artifacts\n", id)
			response, err := gitClient.Jobs.DeleteProjectArtifacts(id)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			return nil
		},
	}
}
