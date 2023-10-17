package job_artifacts

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"io"
	"log"
	"os"
)

// Get 获取（下载）作业产物 https://docs.gitlab.cn/jh/api/job_artifacts.html#%E8%8E%B7%E5%8F%96%E4%BD%9C%E4%B8%9A%E4%BA%A7%E7%89%A9
func Get() *cli.Command {
	return &cli.Command{
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
	}
}
