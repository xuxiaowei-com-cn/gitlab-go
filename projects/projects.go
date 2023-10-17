package projects

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
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
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.Search(), flag.SearchNamespaces(),
			flag.OrderBy(OrderByUsage)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出所有项目",
				Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.Search(), flag.SearchNamespaces(),
					flag.OrderBy(OrderByUsage)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var sort = context.String(constant.Sort)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var search = context.String(constant.Search)
					var searchNamespaces = context.Bool(constant.SearchNamespaces)
					var orderBy = context.String(constant.OrderBy)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					opt := &gitlab.ListProjectsOptions{
						ListOptions: gitlab.ListOptions{
							Page:    page,
							PerPage: perPage,
						},
						Sort:             &sort,
						Search:           &search,
						SearchNamespaces: &searchNamespaces,
						OrderBy:          &orderBy,
					}
					projects, response, err := gitClient.Projects.ListProjects(opt)
					if err != nil {
						return err
					}
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

					for index, project := range projects {
						log.Printf("Index: %d,\t ID: %d,\t Path: %s,\t Name: %s\n", index, project.ID, project.Path, project.Name)

						//name := strings.Replace(project.PathWithNamespace, "xuxiaowei-com-cn/", "", -1)
						//fmt.Printf("git submodule add -b main ../%s.git %s\n", name, name)
					}

					return nil
				},
			},
		},
	}
}
