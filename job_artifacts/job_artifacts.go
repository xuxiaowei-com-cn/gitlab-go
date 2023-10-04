package job_artifacts

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// JobsArtifacts 作业产物 API https://docs.gitlab.cn/jh/api/job_artifacts.html
func JobsArtifacts() *cli.Command {
	return &cli.Command{
		Name:    "job-artifacts",
		Aliases: []string{"ja"},
		Usage:   "作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html",
		Flags:   append(flag.Common(), flag.Id(false), flag.JobId(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "获取作业产物（未完成）",
				Flags: append(flag.Common(), flag.Id(true), flag.JobId(false)),
				Action: func(context *cli.Context) error {
					fmt.Println("获取作业产物")

					return fmt.Errorf("未完成")
				},
			},
			{
				Name:    "download",
				Aliases: []string{"dl"},
				Usage:   "下载产物归档文件（未完成）",
				Flags:   append(flag.Common(), flag.Id(true), flag.JobId(false)),
				Action: func(context *cli.Context) error {
					fmt.Println("下载产物归档文件")

					return fmt.Errorf("未完成")
				},
			},
			{
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

					fmt.Printf("Delete ProjectId: %s, JobId: %d Artifacts\n", id, jobId)
					response, err := gitClient.Jobs.DeleteArtifacts(id, jobId)
					fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
}
