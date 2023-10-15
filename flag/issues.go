package flag

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

// 此处方法是议题专用方法，如果有重名方法（多个命令使用的方法，将移至 flag.go 中，并使用参数控制提示语）

func AssigneeUsername() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.AssigneeUsername,
		Usage: "对于给定的用户名 username，返回指派给这个用户的议题。与 assignee_id 相似且与其冲突。在免费版中，assignee_username 数组只能包含单个成员，否则将报参数错误提示。",
	}
}

func AuthorId() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.AuthorId,
		Usage: "对于给定的用户 id，返回这个用户创建的议题。与 author_username 冲突。与 scope=all 或 scope=assigned_to_me 配合使用。",
	}
}

func AuthorUsername() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.AuthorUsername,
		Usage: "对于给定的用户名 username，返回这个用户创建的议题。与 author_id 相似且与其冲突。",
	}
}

func Confidential() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.Confidential,
		Value: false,
		Usage: "筛选私密议题与公开议题。",
	}
}

func DueDate() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.DueDate,
		Usage: "返回没有截止日期、已经逾期、本周内逾期、本月内逾期或介于两周前和下个月之间逾期的议题。可接受的值：0（没有截止日期）、any、today、tomorrow、overdue、week、month、next_month_and_previous_two_weeks。",
	}
}

func Iids() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Iids,
		Usage: "返回包含给定 iid 的议题。",
	}
}

func In() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.In,
		Value: "title,description",
		Usage: "修改 search 属性的范围。可以使用 title、description 或使用半角逗号对他们进行连接。默认值是 title,description。",
	}
}

func IssueType() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.IssueType,
		Usage: "筛选议题的类型，可选值为 issue、incident 或 test_case。引入于 13.12 版本。",
	}
}

func IterationId() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.IterationId,
		Usage: "对于给定的迭代 ID，返回与这个迭代关联的议题。使用 None 则返回未与迭代关联的议题。使用 Any 则返回存在关联迭代的议题。引入于 13.6 版本。",
	}
}

func Milestone() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Milestone,
		Usage: "里程碑名称。使用 None 则列出没有里程碑的议题。使用 Any 则列出存在关联里程碑的议题。None 及 Any 的用法将会在未来被弃用，请使用 milestone_id 替代。milestone 与 milestone_id 冲突。",
	}
}

func MilestoneId() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.MilestoneId,
		Usage: "对于给定的时间段（None、Any、Upcoming 或 Started），返回与该时间段里程碑相关联的议题。使用 None 则列出没有里程碑的议题。使用 Any 则列出存在关联里程碑的议题。使用 Upcoming 则列出与未开始里程碑相关联的议题。使用 Started 则列出与已开始里程碑相关联的议题。milestone 和 milestone_id 冲突。引入于 14.3 版本。",
	}
}

func MyReactionEmoji() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.MyReactionEmoji,
		Usage: "对于给定的 emoji，返回用户使用该表情回应的议题。使用 None 则返回没有使用表情回应的议题。使用 Any 则返回使用至少一个表情回应的议题。",
	}
}

func OrderBy() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.OrderBy,
		Value: "created_at",
		Usage: "返回根据 created_at、due_date、label_priority、milestone_due、popularity、priority、relative_position、title、updated_at 或 weight 排序的议题。默认值是 created_at。",
	}
}

func Scope() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Scope,
		Value: "created_by_me",
		Usage: "返回满足范围 created_by_me、assigned_to_me 或 all 的议题。默认值是 created_by_me。",
	}
}

func Search() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Search,
		Usage: "根据 title 和 description 搜索议题。",
	}
}

func SearchNamespaces() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.SearchNamespaces,
		Value: false,
		Usage: "匹配搜索条件时包括上级命名空间。默认为 false。",
	}
}

func State() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.State,
		Value: "all",
		Usage: "返回全部 all 议题或仅返回处于 opened 或 closed 状态的议题。",
	}
}

func WithLabelsDetails() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.WithLabelsDetails,
		Value: false,
		Usage: "若为 true 则返回更详尽的标签信息：:name、:color、:description、:description_html、:text_color。默认值是 false。description_html 属性引入于 12.7 版本。",
	}
}

func CreatedAfter() cli.Flag {
	return &cli.TimestampFlag{
		Name:   constant.CreatedAfter,
		Usage:  "对于给定的时间戳，返回不早于该时间创建的议题。时间戳应符合 ISO 8601 格式（2019-03-15T08:00:00Z）",
		Layout: "2006-01-02T15:04:05Z",
	}
}

func CreatedBefore() cli.Flag {
	return &cli.TimestampFlag{
		Name:   constant.CreatedBefore,
		Usage:  "对于给定的时间戳，返回不晚于该时间创建的议题。时间戳应符合 ISO 8601 格式（2019-03-15T08:00:00Z）。",
		Layout: "2006-01-02T15:04:05Z",
	}
}
