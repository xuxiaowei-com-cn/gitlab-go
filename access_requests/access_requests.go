package access_requests

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// AccessRequests 群组和项目访问请求 API https://docs.gitlab.cn/jh/api/access_requests.html
func AccessRequests() *cli.Command {
	return &cli.Command{
		Name:    "access-request",
		Aliases: []string{"access-requests", "ar"},
		Usage:   "群组和项目访问请求 API，中文文档：https://docs.gitlab.cn/jh/api/access_requests.html",
		Flags:   append(flag.Common(), flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:    "group",
				Aliases: []string{"groups"},
				Usage:   "为群组列出访问请求",
				Flags:   append(flag.Common(), flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListAccessRequestsOptions{}
					accessRequests, response, err := gitClient.AccessRequests.ListGroupAccessRequests(id, opt)
					fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}

					for index, accessRequest := range accessRequests {
						fmt.Printf("Index: %d,\t ID: %d,\t Username: %s,\t Name: %s,\t State: %s,\t CreatedAt: %s,\t RequestedAt: %s,\t AccessLevel: %d\n", index, accessRequest.ID, accessRequest.Username, accessRequest.Name, accessRequest.State, accessRequest.CreatedAt, accessRequest.RequestedAt, accessRequest.AccessLevel)
					}

					return nil
				},
			},
			{
				Name:    "project",
				Aliases: []string{"projects"},
				Usage:   "为项目列出访问请求",
				Flags:   append(flag.Common(), flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListAccessRequestsOptions{}
					accessRequests, response, err := gitClient.AccessRequests.ListProjectAccessRequests(id, opt)
					fmt.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}

					for index, accessRequest := range accessRequests {
						fmt.Printf("Index: %d,\t ID: %d,\t Username: %s,\t Name: %s,\t State: %s,\t CreatedAt: %s,\t RequestedAt: %s,\t AccessLevel: %d\n", index, accessRequest.ID, accessRequest.Username, accessRequest.Name, accessRequest.State, accessRequest.CreatedAt, accessRequest.RequestedAt, accessRequest.AccessLevel)
					}

					return nil
				},
			},
		},
	}
}
