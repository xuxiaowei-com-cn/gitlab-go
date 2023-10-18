package issues

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

const (
	OrderByUsage = "返回根据 created_at、due_date、label_priority、milestone_due、popularity、priority、relative_position、title、updated_at 或 weight 排序的议题。默认值是 created_at。"
	ScopeValue   = "created_by_me"
	ScopeUsage   = "返回满足范围 created_by_me、assigned_to_me 或 all 的议题。默认值是 created_by_me。"
)

// Issues 议题 API https://docs.gitlab.cn/jh/api/issues.html
func Issues() *cli.Command {
	return &cli.Command{
		Name:    "issue",
		Aliases: []string{"issues"},
		Usage:   "议题 API，中文文档：https://docs.gitlab.cn/jh/api/issues.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(), flag.Recursion(),
			flag.AssigneeUsername(), flag.AuthorId(), flag.AuthorUsername(), flag.Confidential(),
			flag.DueDate(), flag.Iids(), flag.In(), flag.IssueType(), flag.IterationId(), flag.Milestone(),
			flag.MilestoneId(), flag.MyReactionEmoji(), flag.OrderBy(OrderByUsage), flag.Scope(ScopeValue, ScopeUsage), flag.Search(),
			flag.Sort(), flag.State(), flag.WithLabelsDetails(),
			flag.CreatedAfter(), flag.CreatedBefore(),
			flag.IssueId(false)),
		Subcommands: []*cli.Command{
			List(),
			Delete(),
		},
	}
}
