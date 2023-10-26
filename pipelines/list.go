package pipelines

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列出项目流水线 https://docs.gitlab.cn/jh/api/pipelines.html#%E5%88%97%E5%87%BA%E9%A1%B9%E7%9B%AE%E6%B5%81%E6%B0%B4%E7%BA%BF
func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "列出项目流水线",
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(), flag.Recursion(),
			flag.Id(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var sort = context.String(constant.Sort)
			var id = context.String(constant.Id)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)
			var recursion = context.Bool(constant.Recursion)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			err = ListProjectPipelinesPrintln(gitClient, id, page, perPage, sort, printJson, printTime, recursion)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func ListProjectPipelines(gitClient *gitlab.Client, id interface{}, page int, perPage int, sort string) ([]*gitlab.PipelineInfo, *gitlab.Response, error) {
	opt := &gitlab.ListProjectPipelinesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
		Sort: &sort,
	}
	pipelineInfos, response, err := gitClient.Pipelines.ListProjectPipelines(id, opt)
	return pipelineInfos, response, err
}

func ListProjectPipelinesPrintln(gitClient *gitlab.Client, id interface{}, page int, perPage int, sort string, printJson bool, printTime bool, recursion bool) error {

	pipelineInfos, response, err := ListProjectPipelines(gitClient, id, page, perPage, sort)
	if err != nil {
		return err
	}

	log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

	fmt.Println("")

	if printJson {
		if printTime {
			for _, pipelineInfo := range pipelineInfos {
				jsonData, err := json.Marshal(pipelineInfo)
				if err != nil {
					panic(err)
				}

				log.Printf("\n%s\n", string(jsonData))
				fmt.Println("")
			}
		} else {
			for _, pipelineInfo := range pipelineInfos {
				jsonData, err := json.Marshal(pipelineInfo)
				if err != nil {
					panic(err)
				}

				fmt.Printf("%s\n", string(jsonData))
				fmt.Println("")
			}
		}
	} else {
		if printTime {
			for _, pipelineInfo := range pipelineInfos {
				log.Printf("ID: %d\n", pipelineInfo.ID)
				log.Printf("IID: %d\n", pipelineInfo.IID)
				log.Printf("ProjectID: %d\n", pipelineInfo.ProjectID)
				log.Printf("Status: %s\n", pipelineInfo.Status)
				log.Printf("CreatedAt: %s\n", pipelineInfo.CreatedAt)

				fmt.Println("")
			}
		} else {
			for _, pipelineInfo := range pipelineInfos {
				fmt.Printf("ID: %d\n", pipelineInfo.ID)
				fmt.Printf("IID: %d\n", pipelineInfo.IID)
				fmt.Printf("ProjectID: %d\n", pipelineInfo.ProjectID)
				fmt.Printf("Status: %s\n", pipelineInfo.Status)
				fmt.Printf("CreatedAt: %s\n", pipelineInfo.CreatedAt)

				fmt.Println("")
			}
		}
	}

	if recursion && len(pipelineInfos) > 0 {
		err := ListProjectPipelinesPrintln(gitClient, id, page+1, perPage, sort, printJson, printTime, recursion)
		if err != nil {
			return err
		}
	}

	return nil
}
