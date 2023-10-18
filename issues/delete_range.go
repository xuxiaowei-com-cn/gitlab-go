package issues

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/utils"
	"log"
)

// DeleteRange 删除议题 https://docs.gitlab.cn/jh/api/issues.html#%E5%88%A0%E9%99%A4%E8%AE%AE%E9%A2%98
func DeleteRange() *cli.Command {
	return &cli.Command{
		Name:    "delete-range",
		Aliases: []string{"rm-range"},
		Usage:   "删除议题（范围）",
		Flags:   append(flag.CommonTokenRequired(), flag.Id(true), flag.IssueIdRange(true), flag.Recursion()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var issueIdRanges = context.StringSlice(constant.IssueIdRange)
			fmt.Println(issueIdRanges)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			issueIds := utils.RangeInt(issueIdRanges)

			for _, issueId := range issueIds {
				response, err := gitClient.Issues.DeleteIssue(id, issueId)
				if err != nil {
					log.Printf(err.Error())
				}
				log.Printf("Delete Id %s issueId: %d Response StatusCode: %d\n", id, issueId, response.Response.StatusCode)
			}

			return nil
		},
	}
}
