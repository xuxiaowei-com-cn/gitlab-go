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

// Create 创建变量 https://docs.gitlab.cn/jh/api/project_level_variables.html#%E5%88%9B%E5%BB%BA%E5%8F%98%E9%87%8F
func Create() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "创建变量",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true),
			flag.VariableKey(true), flag.VariableValue(true), flag.VariableType(), flag.VariableProtected(),
			flag.VariableMasked(), flag.VariableRaw(), flag.VariableEnvironmentScope(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)

			var key = context.String(constant.VariableKey)
			var value = context.String(constant.VariableValue)
			var variableType = context.String(constant.VariableType)
			var protected = context.Bool(constant.VariableProtected)
			var masked = context.Bool(constant.VariableMasked)
			var raw = context.Bool(constant.VariableRaw)
			var environmentScope = context.String(constant.VariableEnvironmentScope)

			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.CreateProjectVariableOptions{
				Key:              &key,
				Value:            &value,
				Protected:        &protected,
				Masked:           &masked,
				Raw:              &raw,
				EnvironmentScope: &environmentScope,
			}
			if variableType == "env_var" {
				opt.VariableType = gitlab.Ptr(gitlab.EnvVariableType)
			} else if variableType == "file" {
				opt.VariableType = gitlab.Ptr(gitlab.FileVariableType)
			}

			projectVariable, response, err := gitClient.ProjectVariables.CreateVariable(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					jsonData, err := json.Marshal(projectVariable)
					if err != nil {
						panic(err)
					}

					log.Printf("\n%s\n", string(jsonData))
					fmt.Println("")

				} else {
					jsonData, err := json.Marshal(projectVariable)
					if err != nil {
						panic(err)
					}

					fmt.Printf("%s\n", string(jsonData))
					fmt.Println("")

				}
			} else {
				if printTime {
					log.Printf("Key: %s\n", projectVariable.Key)
					log.Printf("Value: %s\n", projectVariable.Value)
					log.Printf("VariableType: %s\n", projectVariable.VariableType)
					log.Printf("Protected: %t\n", projectVariable.Protected)
					log.Printf("Masked: %t\n", projectVariable.Masked)
					log.Printf("Raw: %t\n", projectVariable.Raw)
					log.Printf("EnvironmentScope: %s\n", projectVariable.EnvironmentScope)

					fmt.Println("")

				} else {
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

			return nil
		},
	}
}
