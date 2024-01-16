package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/project_level_variables"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"log"
)

// ProjectLevelVariablesCreateAll 所有项目创建新变量
func ProjectLevelVariablesCreateAll() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "所有项目创建新变量",
		Flags: append(flag.CommonTokenRequired(), flag.Owned(true),
			flag.VariableKey(true), flag.VariableValue(true), flag.VariableType(), flag.VariableProtected(),
			flag.VariableMasked(), flag.VariableRaw(), flag.VariableEnvironmentScope(), flag.AllowFailure(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)

			var key = context.String(constant.VariableKey)
			var value = context.String(constant.VariableValue)
			var variableType = context.String(constant.VariableType)
			var protected = context.Bool(constant.VariableProtected)
			var masked = context.Bool(constant.VariableMasked)
			var raw = context.Bool(constant.VariableRaw)
			var environmentScope = context.String(constant.VariableEnvironmentScope)

			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)
			var allowFailure = context.Bool(constant.AllowFailure)

			projectList, err := projects.ListProjects(owned, token, baseUrl, 1, 100)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = project_level_variables.CreateVariable(baseUrl, token, project.ID, key, value, protected, masked, raw, environmentScope, variableType,
					printJson, printTime, allowFailure)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
