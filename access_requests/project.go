package access_requests

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Project 为项目列出访问请求 https://docs.gitlab.cn/jh/api/access_requests.html#%E4%B8%BA%E7%BE%A4%E7%BB%84%E6%88%96%E9%A1%B9%E7%9B%AE%E5%88%97%E5%87%BA%E8%AE%BF%E9%97%AE%E8%AF%B7%E6%B1%82
func Project() *cli.Command {
	return &cli.Command{
		Name:    "project",
		Aliases: []string{"projects"},
		Usage:   "为项目列出访问请求",
		Flags:   append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(), flag.Id(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListAccessRequestsOptions{
				Page:    page,
				PerPage: perPage,
			}
			accessRequests, response, err := gitClient.AccessRequests.ListProjectAccessRequests(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, accessRequest := range accessRequests {
						jsonData, err := json.Marshal(accessRequest)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, accessRequest := range accessRequests {
						jsonData, err := json.Marshal(accessRequest)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, accessRequest := range accessRequests {
						log.Printf("ID: %d\n", accessRequest.ID)
						log.Printf("Username: %s\n", accessRequest.Username)
						log.Printf("Name: %s\n", accessRequest.Name)
						log.Printf("State: %s\n", accessRequest.State)
						log.Printf("CreatedAt: %s\n", accessRequest.CreatedAt)
						log.Printf("RequestedAt: %s\n", accessRequest.RequestedAt)
						log.Printf("AccessLevel: %d\n", accessRequest.AccessLevel)

						fmt.Println("")
					}
				} else {
					for _, accessRequest := range accessRequests {
						fmt.Printf("ID: %d\n", accessRequest.ID)
						fmt.Printf("Username: %s\n", accessRequest.Username)
						fmt.Printf("Name: %s\n", accessRequest.Name)
						fmt.Printf("State: %s\n", accessRequest.State)
						fmt.Printf("CreatedAt: %s\n", accessRequest.CreatedAt)
						fmt.Printf("RequestedAt: %s\n", accessRequest.RequestedAt)
						fmt.Printf("AccessLevel: %d\n", accessRequest.AccessLevel)

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
