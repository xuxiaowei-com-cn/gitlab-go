package instance_level_ci_variables

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列出所有实例变量 https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html#%E5%88%97%E5%87%BA%E6%89%80%E6%9C%89%E5%AE%9E%E4%BE%8B%E5%8F%98%E9%87%8F
func List() *cli.Command {
	return &cli.Command{Name: "list",
		Usage: "列出所有实例变量",
		Flags: append(flag.CommonTokenRequired(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListInstanceVariablesOptions{
				Page:    page,
				PerPage: perPage,
			}
			variables, response, err := gitClient.InstanceVariables.ListVariables(opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, variable := range variables {
						jsonData, err := json.Marshal(variable)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, variable := range variables {
						jsonData, err := json.Marshal(variable)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, variable := range variables {
						log.Printf("Key: %s\n", variable.Key)
						log.Printf("Username: %s\n", variable.Value)
						log.Printf("Name: %s\n", variable.VariableType)
						log.Printf("State: %t\n", variable.Protected)
						log.Printf("CreatedAt: %t\n", variable.Masked)
						log.Printf("RequestedAt: %t\n", variable.Raw)

						fmt.Println("")
					}
				} else {
					for _, variable := range variables {
						fmt.Printf("Key: %s\n", variable.Key)
						fmt.Printf("Username: %s\n", variable.Value)
						fmt.Printf("Name: %s\n", variable.VariableType)
						fmt.Printf("State: %t\n", variable.Protected)
						fmt.Printf("CreatedAt: %t\n", variable.Masked)
						fmt.Printf("RequestedAt: %t\n", variable.Raw)

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
