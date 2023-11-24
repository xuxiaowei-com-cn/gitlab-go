package mix

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/pipelines"
	"github.com/xuxiaowei-com-cn/gitlab-go/utils"
	"log"
	"sort"
)

// DeleteJobs 删除作业产物和作业日志
func DeleteJobs() *cli.Command {
	return &cli.Command{
		Name:    "job",
		Aliases: []string{"jobs"},
		Usage:   "根据项目路径/ID、流水线IID范围删除作业产物和作业日志（混合命令，多接口命令，立即删除）",
		Flags:   append(flag.CommonTokenRequired(), flag.Sort(), flag.Page(), flag.PerPage(), flag.Id(true), flag.IIdRange(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var sortStr = context.String(constant.Sort)
			var id = context.String(constant.Id)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var iidRanges = context.StringSlice(constant.IIdRange)
			fmt.Println(iidRanges)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			iids := utils.Unique(utils.RangeInt(iidRanges))
			sort.Ints(iids)

			fmt.Printf("%d\n", iids)

			iidsLen := len(iids)
			if iidsLen == 0 {
				// 输入数字为空，结束程序
				return nil
			}

			err = DeleteJobsRecursion(gitClient, id, page, perPage, sortStr, iids)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func DeleteJobsRecursion(gitClient *gitlab.Client, id interface{}, page int, perPage int, sort string, iids []int) error {

	pipelineInfos, response, err := pipelines.ListProjectPipelines(gitClient, id, page, perPage, sort)

	if err != nil {
		return err
	}

	log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

	for _, pipelineInfo := range pipelineInfos {
		fmt.Printf("%d\n", pipelineInfo.IID)

		// 数字在上方已排序

		// 最小值
		iidsMin := iids[0]
		// 最大值
		iidsMax := iids[len(iids)-1]

		if iidsMin == pipelineInfo.IID {
			// 等于最小值，删除最小值
			iids = iids[1:]
			err = ExecuteDeleteJobs(gitClient, id, pipelineInfo.ID, 1, 100)
			if err != nil {
				return err
			}
		} else if pipelineInfo.IID == iidsMax {
			// 等于最大值
			iids = iids[:len(iids)-1]
			err = ExecuteDeleteJobs(gitClient, id, pipelineInfo.ID, 1, 100)
			if err != nil {
				return err
			}
		} else if iidsMin < pipelineInfo.IID {
			// 大于最小值
			err = jobsForExecute(&iids, pipelineInfo.IID, gitClient, id, pipelineInfo.ID)
			if err != nil {
				return err
			}
		} else if pipelineInfo.IID < iidsMax {
			// 小于最大值
			err = jobsForExecute(&iids, pipelineInfo.IID, gitClient, id, pipelineInfo.ID)
			if err != nil {
				return err
			}
		} else {
			// 不在合法范围内
			return nil
		}

		iidsLen := len(iids)
		if iidsLen == 0 {
			// 输入数字已处理完成，跳出循环
			return nil
		}

	}

	if len(pipelineInfos) > 0 {
		err := DeleteJobsRecursion(gitClient, id, page+1, perPage, sort, iids)
		if err != nil {
			return err
		}
	}

	return nil
}

func jobsForExecute(iids *[]int, pipelineInfoIId int, gitClient *gitlab.Client, id interface{}, pipelineInfoId int) error {
	for i := 0; i < len(*iids); i++ {
		if (*iids)[i] == pipelineInfoIId {
			fmt.Printf("数组中包含%d\n", pipelineInfoIId)
			*iids = append((*iids)[:i], (*iids)[i+1:]...)
			err := ExecuteDeleteJobs(gitClient, id, pipelineInfoId, 1, 100)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func ExecuteDeleteJobs(gitClient *gitlab.Client, id interface{}, pipelineInfoId int, page int, perPage int) error {
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

		_, response, err = gitClient.Jobs.EraseJob(id, job.ID)
		if err != nil {
			return err
		}
		log.Printf("Delete Project %s Job %d Response StatusCode: %d\n", id, job.ID, response.Response.StatusCode)
	}

	if len(jobs) == perPage {
		err = ExecuteDeleteJobs(gitClient, id, pipelineInfoId, page+1, perPage)
		if err != nil {
			return err
		}
	}

	return nil
}
