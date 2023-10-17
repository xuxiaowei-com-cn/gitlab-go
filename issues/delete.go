package issues

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Delete 删除议题 https://docs.gitlab.cn/jh/api/issues.html#%E5%88%A0%E9%99%A4%E8%AE%AE%E9%A2%98
func Delete() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"rm"},
		Usage:   "删除议题",
		Flags:   append(flag.CommonTokenRequired(), flag.Id(true), flag.IssueId(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var issueId = context.Int(constant.IssueId)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			response, err := gitClient.Issues.DeleteIssue(id, issueId)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			return nil
		},
	}
}
