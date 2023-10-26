package mix

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/pipelines"
	"log"
)

// DeleteAllArtifacts 删除所有产物
func DeleteAllArtifacts() *cli.Command {
	return &cli.Command{
		Name:    "all-artifact",
		Aliases: []string{"all-artifacts"},
		Usage:   "根据项目路径/ID删除所有产物（混合命令，立即删除）",
		Flags:   append(flag.CommonTokenRequired(), flag.Sort(), flag.Page(), flag.PerPage(), flag.Id(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var sortStr = context.String(constant.Sort)
			var id = context.String(constant.Id)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			err = DeleteAllArtifactsRecursion(gitClient, id, page, perPage, sortStr)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func DeleteAllArtifactsRecursion(gitClient *gitlab.Client, id interface{}, page int, perPage int, sort string) error {

	pipelineInfos, response, err := pipelines.ListProjectPipelines(gitClient, id, page, perPage, sort)

	if err != nil {
		return err
	}

	log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

	for _, pipelineInfo := range pipelineInfos {
		fmt.Printf("%d\n", pipelineInfo.IID)

		err = ExecuteDeleteAllArtifacts(gitClient, id, pipelineInfo.ID, 1, 100)
		if err != nil {
			return err
		}

	}

	if len(pipelineInfos) > 0 {
		err := DeleteAllArtifactsRecursion(gitClient, id, page+1, perPage, sort)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteDeleteAllArtifacts(gitClient *gitlab.Client, id interface{}, pipelineInfoId int, page int, perPage int) error {
	fmt.Printf("执行删除 %d \n", pipelineInfoId)

	opt := &gitlab.ListJobsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}

	jobs, response, err := gitClient.Jobs.ListPipelineJobs(id, pipelineInfoId, opt)
	if err != nil {
		return err
	}

	log.Printf("List Project %s Pipeline %d Jobs Response StatusCode: %d\n", id, pipelineInfoId, response.Response.StatusCode)

	for _, job := range jobs {
		response, err = gitClient.Jobs.DeleteArtifacts(id, job.ID)
		if err != nil {
			return err
		}
		log.Printf("Delete Project %s Job %d Response StatusCode: %d\n", id, job.ID, response.Response.StatusCode)
	}

	if len(jobs) == perPage {
		err = ExecuteDeleteAllArtifacts(gitClient, id, pipelineInfoId, page+1, perPage)
		if err != nil {
			return err
		}
	}

	return nil
}
