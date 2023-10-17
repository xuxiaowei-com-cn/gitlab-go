package container_registry

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// DeleteTag 删除仓库里存储库的某个标签 https://docs.gitlab.cn/jh/api/container_registry.html#%E5%88%A0%E9%99%A4%E4%BB%93%E5%BA%93%E9%87%8C%E5%AD%98%E5%82%A8%E5%BA%93%E7%9A%84%E6%9F%90%E4%B8%AA%E6%A0%87%E7%AD%BE
func DeleteTag() *cli.Command {
	return &cli.Command{
		Name:    "delete-tag",
		Aliases: []string{"rm-tag"},
		Usage:   "删除仓库里存储库的某个标签",
		Flags: append(flag.CommonTokenRequired(),
			flag.Id(true), flag.Repository(true), flag.TagName(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var repository = context.Int(constant.Repository)
			var tagName = context.String(constant.TagName)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			response, err := gitClient.ContainerRegistry.DeleteRegistryRepositoryTag(id, repository, tagName)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			return nil
		},
	}
}
