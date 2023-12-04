package projects

import (
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
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
		Flags: append(flag.Common(), flag.Sort(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Search(), flag.SearchNamespaces(),
			flag.OrderBy(OrderByUsage)),
		Subcommands: []*cli.Command{
			List(),
		},
	}
}

// ListProjects 列出群组 https://docs.gitlab.cn/jh/api/projects.html#%E5%88%97%E5%87%BA%E6%89%80%E6%9C%89%E9%A1%B9%E7%9B%AE
func ListProjects(owned *bool, token string, baseUrl string, page int, perPage int) ([]*gitlab.Project, error) {
	var results []*gitlab.Project

	gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return nil, err
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
		Owned: owned,
	}

	projects, response, err := gitClient.Projects.ListProjects(opt)
	if err != nil {
		return nil, err
	}

	log.Printf("ListProjects Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

	results = append(results, projects...)

	if len(projects) != 0 {
		projects, err = ListProjects(owned, token, baseUrl, page+1, perPage)
		if err != nil {
			return nil, err
		}

		results = append(results, projects...)
	}

	return results, err
}

func ArchiveProject(token string, baseUrl string, pathWithNamespace string) error {

	gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return err
	}

	_, response, err := gitClient.Projects.ArchiveProject(pathWithNamespace)
	if err != nil {
		return err
	}

	log.Printf("ArchiveProject PathWithNamespace: %s, Response StatusCode: %d\n", pathWithNamespace, response.Response.StatusCode)

	return nil
}
