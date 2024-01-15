package environments

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列举环境 https://docs.gitlab.cn/jh/api/environments.html#%E5%88%97%E4%B8%BE%E7%8E%AF%E5%A2%83
func List() *cli.Command {
	return &cli.Command{Name: "list",
		Usage: "列出所有实例变量",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
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

			opt := &gitlab.ListEnvironmentsOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			}
			environments, response, err := gitClient.Environments.ListEnvironments(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, environment := range environments {
						jsonData, err := json.Marshal(environment)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, environment := range environments {
						jsonData, err := json.Marshal(environment)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, environment := range environments {
						log.Printf("ID: %d\n", environment.ID)
						log.Printf("Name: %s\n", environment.Name)
						log.Printf("Slug: %s\n", environment.Slug)
						log.Printf("State: %s\n", environment.State)
						log.Printf("Tier: %s\n", environment.Tier)
						log.Printf("ExternalURL: %s\n", environment.ExternalURL)
						log.Printf("CreatedAt: %s\n", environment.CreatedAt)
						log.Printf("UpdatedAt: %s\n", environment.UpdatedAt)

						fmt.Println("")
					}
				} else {
					for _, environment := range environments {
						fmt.Printf("ID: %d\n", environment.ID)
						fmt.Printf("Name: %s\n", environment.Name)
						fmt.Printf("Slug: %s\n", environment.Slug)
						fmt.Printf("State: %s\n", environment.State)
						fmt.Printf("Tier: %s\n", environment.Tier)
						fmt.Printf("ExternalURL: %s\n", environment.ExternalURL)
						fmt.Printf("CreatedAt: %s\n", environment.CreatedAt)
						fmt.Printf("UpdatedAt: %s\n", environment.UpdatedAt)

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
