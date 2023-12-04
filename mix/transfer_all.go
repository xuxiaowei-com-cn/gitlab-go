package mix

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"log"
	"strings"
)

func TransferAll() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "将一个命令空间的项目转移到新的命名空间",
		Flags: append(flag.CommonTokenRequired(), flag.Owned(true), flag.NamespaceSource(true),
			flag.NamespaceTargetFlag(true), flag.SkipProjectPath()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)
			var namespaceSource = context.String(constant.NamespaceSource)
			var namespaceTarget = context.String(constant.NamespaceTarget)
			var skipProjectPaths = context.StringSlice(constant.SkipProjectPath)

			for index, skipProjectPath := range skipProjectPaths {
				skipProjectPath = strings.TrimPrefix(skipProjectPath, "/")
				skipProjectPath = strings.TrimSuffix(skipProjectPath, "/")
				skipProjectPaths[index] = skipProjectPath
			}

			projectList, err := projects.ListNamespaceProjects(owned, token, baseUrl, 1, 100, namespaceSource, skipProjectPaths)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = TransferProject(baseUrl, token, project.PathWithNamespace, namespaceTarget)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}

// TransferProject 将项目转移到新的命名空间 https://docs.gitlab.cn/jh/api/projects.html#%E5%B0%86%E9%A1%B9%E7%9B%AE%E8%BD%AC%E7%A7%BB%E5%88%B0%E6%96%B0%E7%9A%84%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4
func TransferProject(baseUrl string, token string, id string, namespace string) error {

	gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return err
	}

	opt := &gitlab.TransferProjectOptions{
		Namespace: namespace,
	}

	project, response, err := gitClient.Projects.TransferProject(id, opt)
	if err != nil {
		return err
	}

	log.Printf("项目: %s 转移到: %s 结果: %s", id, project.WebURL, response.Status)

	return nil
}
