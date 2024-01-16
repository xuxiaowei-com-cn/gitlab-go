package project_level_variables

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列出项目变量 https://docs.gitlab.cn/jh/api/project_level_variables.html#%E5%88%97%E5%87%BA%E9%A1%B9%E7%9B%AE%E5%8F%98%E9%87%8F
func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "列出项目变量",
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

			opt := &gitlab.ListProjectVariablesOptions{
				Page:    page,
				PerPage: perPage,
			}
			projectVariables, response, err := gitClient.ProjectVariables.ListVariables(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, projectVariable := range projectVariables {
						jsonData, err := json.Marshal(projectVariable)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, projectVariable := range projectVariables {
						jsonData, err := json.Marshal(projectVariable)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, projectVariable := range projectVariables {
						log.Printf("Key: %s\n", projectVariable.Key)
						log.Printf("Value: %s\n", projectVariable.Value)
						log.Printf("VariableType: %s\n", projectVariable.VariableType)
						log.Printf("Protected: %t\n", projectVariable.Protected)
						log.Printf("Masked: %t\n", projectVariable.Masked)
						log.Printf("Raw: %t\n", projectVariable.Raw)
						log.Printf("EnvironmentScope: %s\n", projectVariable.EnvironmentScope)

						fmt.Println("")
					}
				} else {
					for _, projectVariable := range projectVariables {
						fmt.Printf("Key: %s\n", projectVariable.Key)
						fmt.Printf("Value: %s\n", projectVariable.Value)
						fmt.Printf("VariableType: %s\n", projectVariable.VariableType)
						fmt.Printf("Protected: %t\n", projectVariable.Protected)
						fmt.Printf("Masked: %t\n", projectVariable.Masked)
						fmt.Printf("Raw: %t\n", projectVariable.Raw)
						fmt.Printf("EnvironmentScope: %s\n", projectVariable.EnvironmentScope)

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
