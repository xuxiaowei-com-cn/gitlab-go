package job_artifacts

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"io"
	"log"
	"os"
)

// JobsArtifacts 作业产物 API https://docs.gitlab.cn/jh/api/job_artifacts.html
func JobsArtifacts() *cli.Command {
	return &cli.Command{
		Name:    "job-artifact",
		Aliases: []string{"job-artifacts", "ja"},
		Usage:   "作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html",
		Flags:   append(flag.Common(), flag.Id(false), flag.JobId(false), flag.ArtifactsName()),
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "获取（下载）作业产物",
				Flags: append(flag.Common(), flag.Id(true), flag.JobId(true), flag.ArtifactsName()),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)
					var jobId = context.Int(constant.JobId)
					var artifactsName = context.String(constant.ArtifactsName)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					artifactsReader, response, err := gitClient.Jobs.GetJobArtifacts(id, jobId)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					log.Printf("Get Job Artifacts End")

					// 读取构件数据
					artifactsData, err := io.ReadAll(artifactsReader)
					if err != nil {
						return err
					}

					log.Printf("Read All Reader End")

					// 将构件数据保存到文件
					err = os.WriteFile(artifactsName, artifactsData, 0644)
					if err != nil {
						return err
					}

					log.Printf("Write File End")

					return nil
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

					log.Printf("Delete ProjectId: %s, JobId: %d Artifacts\n", id, jobId)
					response, err := gitClient.Jobs.DeleteArtifacts(id, jobId)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					return nil
				},
			},
			{
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
			},
		},
	}
}
