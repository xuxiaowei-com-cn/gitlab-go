package container_registry

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// ContainerRegistry 容器仓库 API https://docs.gitlab.cn/jh/api/container_registry.html
func ContainerRegistry() *cli.Command {
	return &cli.Command{
		Name:    "container-registry",
		Aliases: []string{"cr"},
		Usage:   "容器仓库 API，中文文档：https://docs.gitlab.cn/jh/api/container_registry.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false), flag.Repository(false), flag.TagName(false)),
		Subcommands: []*cli.Command{
			List(),
			ListTags(),
			GetTag(),
			DeleteTag(),
		},
	}
}
