package projects

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

const (
	OrderByUsage = "返回按 id、name、path、created_at、updated_at、last_activity_at 或 similarity 字段排序的项目。repository_size、storage_size、packages_size 或 wiki_size 字段只允许管理员使用。similarity（引入于 14.1 版本）仅在搜索时可用，并且仅限于当前用户所属的项目。默认是created_at。"
)

// Projects 项目 API https://docs.gitlab.cn/jh/api/projects.html
func Projects() *cli.Command {
	return &cli.Command{
		Name:    "project",
		Aliases: []string{"projects", "p"},
		Usage:   "项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html",
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Search(), flag.SearchNamespaces(),
			flag.OrderBy(OrderByUsage)),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}
