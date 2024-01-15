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

// Create 创建新环境 https://docs.gitlab.cn/jh/api/environments.html#%E5%88%9B%E5%BB%BA%E6%96%B0%E7%8E%AF%E5%A2%83
func Create() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "创建新环境",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true), flag.EnvName(true), flag.EnvExternalUrl(), flag.EnvTier(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var name = context.String(constant.EnvName)
			var externalURL = context.String(constant.EnvExternalUrl)
			var tier = context.String(constant.EnvTier)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.CreateEnvironmentOptions{
				Name:        &name,
				ExternalURL: &externalURL,
			}
			if tier != "" {
				opt.Tier = &tier
			}

			environment, response, err := gitClient.Environments.CreateEnvironment(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					jsonData, err := json.Marshal(environment)
					if err != nil {
						panic(err)
					}

					log.Printf("\n%s\n", string(jsonData))
					fmt.Println("")

				} else {
					jsonData, err := json.Marshal(environment)
					if err != nil {
						panic(err)
					}

					fmt.Printf("%s\n", string(jsonData))
					fmt.Println("")

				}
			} else {
				if printTime {
					log.Printf("ID: %d\n", environment.ID)
					log.Printf("Name: %s\n", environment.Name)
					log.Printf("Slug: %s\n", environment.Slug)
					log.Printf("State: %s\n", environment.State)
					log.Printf("Tier: %s\n", environment.Tier)
					log.Printf("ExternalURL: %s\n", environment.ExternalURL)
					log.Printf("CreatedAt: %s\n", environment.CreatedAt)
					log.Printf("UpdatedAt: %s\n", environment.UpdatedAt)

					fmt.Println("")

				} else {
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

			return nil
		},
	}
}
