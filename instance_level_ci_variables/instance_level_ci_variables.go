package instance_level_ci_variables

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// InstanceLevelCiVariables 实例级 CI/CD 变量 API https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html
func InstanceLevelCiVariables() *cli.Command {
	return &cli.Command{
		Name:    "instance-level-ci-variable",
		Aliases: []string{"instance-level-ci-variables", "ilcv"},
		Usage:   "实例级 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html",
		Flags:   append(flag.Common()),
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "列出所有实例变量",
				Flags: flag.CommonTokenRequired(),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListInstanceVariablesOptions{}
					variables, response, err := gitClient.InstanceVariables.ListVariables(opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					for index, variable := range variables {
						log.Printf("Index: %d,\t Key: %s,\t Value: %s,\t VariableType: %s,\t Protected: %t,\t Masked: %t,\t Raw: %t\n", index, variable.Key, variable.Value, variable.VariableType, variable.Protected, variable.Masked, variable.Raw)
					}

					return nil
				},
			},
		},
	}
}
