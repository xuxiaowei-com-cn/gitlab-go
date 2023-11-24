package jobs

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/utils"
	"log"
	"sort"
)

// Erase 删除作业 https://docs.gitlab.cn/jh/api/jobs.html#%E5%88%A0%E9%99%A4%E4%BD%9C%E4%B8%9A
func Erase() *cli.Command {
	return &cli.Command{
		Name:  "erase",
		Usage: "删除作业（删除作业产物和作业日志）",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true), flag.JobIdRange(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var jobIdRanges = context.StringSlice(constant.JobIdRange)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			jobIds := utils.Unique(utils.RangeInt(jobIdRanges))
			sort.Ints(jobIds)

			err = EraseJobsPrintln(gitClient, id, jobIds)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func EraseJobsPrintln(gitClient *gitlab.Client, id string, jobIds []int) error {

	// 数字在上方已排序

	for _, jobId := range jobIds {
		job, response, err := gitClient.Jobs.EraseJob(id, jobId)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Delete %s Job %d Response %v：\n%v", id, jobId, response, job)
	}

	return nil
}
