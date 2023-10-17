package issues

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
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
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.AssigneeUsername(), flag.AuthorId(), flag.AuthorUsername(), flag.Confidential(),
			flag.DueDate(), flag.Iids(), flag.In(), flag.IssueType(), flag.IterationId(), flag.Milestone(),
			flag.MilestoneId(), flag.MyReactionEmoji(), flag.OrderBy(OrderByUsage), flag.Scope(ScopeValue, ScopeUsage), flag.Search(),
			flag.Sort(), flag.State(), flag.WithLabelsDetails(),
			flag.CreatedAfter(), flag.CreatedBefore()),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出议题",
				Flags: append(flag.CommonTokenRequired(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
					flag.AssigneeUsername(), flag.AuthorId(), flag.AuthorUsername(), flag.Confidential(),
					flag.DueDate(), flag.Iids(), flag.In(), flag.IssueType(), flag.IterationId(), flag.Milestone(),
					flag.MilestoneId(), flag.MyReactionEmoji(), flag.OrderBy(OrderByUsage), flag.Scope(ScopeValue, ScopeUsage), flag.Search(),
					flag.Sort(), flag.State(), flag.WithLabelsDetails(),
					flag.CreatedAfter(), flag.CreatedBefore()),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var printJson = context.Bool(constant.PrintJson)
					var printTime = context.Bool(constant.PrintTime)

					var assigneeId = context.Int(constant.AssigneeId)
					var assigneeUsername = context.String(constant.AssigneeUsername)
					var authorId = context.Int(constant.AuthorId)
					var authorUsername = context.String(constant.AuthorUsername)
					var confidential = context.Bool(constant.Confidential)
					var createdAfter = context.Timestamp(constant.CreatedAfter)
					var createdBefore = context.Timestamp(constant.CreatedBefore)
					var dueDate = context.String(constant.DueDate)
					// var epicId = context.String(constant.EpicId)
					// var healthStatus = context.String(constant.HealthStatus)
					var iids = context.IntSlice(constant.Iids)
					var in = context.String(constant.In)
					//var issueType = context.String(constant.IssueType)
					var iterationId = context.Int(constant.IterationId)
					// var iterationTitle = context.Int(constant.IterationTitle)
					// var labels = context.StringSlice(constant.Labels)
					var milestone = context.String(constant.Milestone)
					// var milestoneId = context.String(constant.MilestoneId)
					var myReactionEmoji = context.String(constant.MyReactionEmoji)
					// var nonArchived = context.String(constant.NonArchived)
					// var not = context.String(constant.Not)
					var orderBy = context.String(constant.OrderBy)
					var scope = context.String(constant.Scope)
					var search = context.String(constant.Search)
					var sort = context.String(constant.Sort)
					var state = context.String(constant.State)
					// var updatedAfter = context.String(constant.UpdatedAfter)
					// var updatedBefore = context.String(constant.UpdatedBefore)
					// var weight = context.String(constant.Weight)
					var withLabelsDetails = context.Bool(constant.WithLabelsDetails)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}
					opt := &gitlab.ListIssuesOptions{
						WithLabelDetails: &withLabelsDetails,
						Scope:            &scope,
						In:               &in,
						OrderBy:          &orderBy,
						Sort:             &sort,
						Confidential:     &confidential,
						ListOptions: gitlab.ListOptions{
							Page:    page,
							PerPage: perPage,
						},
					}
					if assigneeId != 0 {
						opt.AssigneeID = gitlab.AssigneeID(assigneeId)
					}
					if authorId != 0 {
						opt.AuthorID = &authorId
					}
					if createdAfter != nil {
						opt.CreatedAfter = createdAfter
					}
					if createdBefore != nil {
						opt.CreatedBefore = createdBefore
					}
					if search != "" {
						opt.Search = &search
					}
					if state != "" {
						opt.State = &state
					}
					if milestone != "" {
						opt.Milestone = &milestone
					}
					if authorUsername != "" {
						opt.AuthorUsername = &authorUsername
					}
					if assigneeUsername != "" {
						opt.AssigneeUsername = &assigneeUsername
					}
					if myReactionEmoji != "" {
						opt.MyReactionEmoji = &myReactionEmoji
					}
					if len(iids) > 0 {
						opt.IIDs = &iids
					}
					if dueDate != "" {
						opt.DueDate = &dueDate
					}
					if iterationId != 0 {
						opt.IterationID = &iterationId
					}

					//opt := &gitlab.ListIssuesOptions{
					//	// Labels:              labels,
					//	// NotLabels:           notLabels,
					//	// NotMilestone:        notMilestone,
					//	// AuthorID:       &authorId,
					//	// NotAuthorUsername:   notAuthorUsername,
					//	// NotAuthorID:         notAuthorID,
					//	// AssigneeID:          &assigneeId,
					//	// NotAssigneeID:       notAssigneeID,
					//	// NotAssigneeUsername: notAssigneeUsername,
					//	// NotMyReactionEmoji: notMyReactionEmoji,
					//	// NotIn:         notIn,
					//	// NotSearch:     notSearch,
					//	// CreatedAfter:  createdAfter,
					//	// CreatedBefore: createdBefore,
					//	// UpdatedAfter:  updatedAfter,
					//	// UpdatedBefore: updatedBefore,
					//	Confidential: &confidential,
					//	//IssueType:    &issueType,
					//	IterationID: &iterationId,
					//	ListOptions: gitlab.ListOptions{
					//		Page:    page,
					//		PerPage: perPage,
					//	},
					//}
					issues, response, err := gitClient.Issues.ListIssues(opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					fmt.Println("")

					if printJson {
						if printTime {
							for _, issue := range issues {
								jsonData, err := json.Marshal(issue)
								if err != nil {
									panic(err)
								}

								log.Printf("\n%s\n", string(jsonData))
								fmt.Println("")
							}
						} else {
							for _, issue := range issues {
								jsonData, err := json.Marshal(issue)
								if err != nil {
									panic(err)
								}

								fmt.Printf("%s\n", string(jsonData))
								fmt.Println("")
							}
						}
					} else {
						if printTime {
							for _, issue := range issues {
								log.Printf("ID: %d\n", issue.ID)
								log.Printf("Title: %s\n", issue.Title)
								log.Printf("State: %s\n", issue.State)
								log.Printf("CreatedAt: %s\n", issue.CreatedAt)
								log.Printf("UpdatedAt: %s\n", issue.UpdatedAt)
								log.Printf("ClosedAt: %s\n", issue.ClosedAt)
								log.Printf("DueDate: %s\n", issue.DueDate)
								log.Printf("Author: %s\n", issue.Author.Name)
								log.Printf("Assignee: %s\n", issue.Assignee.Name)
								log.Printf("Labels: %s\n", issue.Labels)
								log.Printf("Milestone: %s\n", issue.Milestone.Title)
								log.Printf("ProjectID: %d\n", issue.ProjectID)
								fmt.Println("")
							}
						} else {
							for _, issue := range issues {
								fmt.Printf("ID: %d\n", issue.ID)
								fmt.Printf("Title: %s\n", issue.Title)
								fmt.Printf("State: %s\n", issue.State)
								fmt.Printf("CreatedAt: %s\n", issue.CreatedAt)
								fmt.Printf("UpdatedAt: %s\n", issue.UpdatedAt)
								fmt.Printf("ClosedAt: %s\n", issue.ClosedAt)
								fmt.Printf("DueDate: %s\n", issue.DueDate)
								fmt.Printf("Author: %s\n", issue.Author.Name)
								fmt.Printf("Labels: %s\n", issue.Labels)
								fmt.Printf("ProjectID: %d\n", issue.ProjectID)
								fmt.Println("")
							}
						}
					}

					return nil
				},
			},
		},
	}
}
