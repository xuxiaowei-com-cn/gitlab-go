package projects

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列出所有项目 https://docs.gitlab.cn/jh/api/projects.html#%E5%88%97%E5%87%BA%E6%89%80%E6%9C%89%E9%A1%B9%E7%9B%AE
func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "列出所有项目",
		Flags: append(flag.Common(), flag.Owned(true), flag.Sort(), flag.Page(), flag.PerPage(),
			flag.PrintJson(), flag.PrintTime(), flag.Search(), flag.SearchNamespaces(),
			flag.OrderBy(OrderByUsage)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var owned = context.Bool(constant.Owned)
			var sort = context.String(constant.Sort)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var search = context.String(constant.Search)
			var searchNamespaces = context.Bool(constant.SearchNamespaces)
			var orderBy = context.String(constant.OrderBy)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

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
				Owned:            &owned,
				OrderBy:          &orderBy,
			}
			projects, response, err := gitClient.Projects.ListProjects(opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			//name := strings.Replace(project.PathWithNamespace, "xuxiaowei-com-cn/", "", -1)
			//fmt.Printf("git submodule add -b main ../%s.git %s\n", name, name)

			if printJson {
				if printTime {
					for _, project := range projects {
						jsonData, err := json.Marshal(project)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, project := range projects {
						jsonData, err := json.Marshal(project)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, project := range projects {
						log.Printf("ID: %d\n", project.ID)
						log.Printf("Path: %s\n", project.Path)
						log.Printf("Name: %s\n", project.Name)

						fmt.Println("")
					}
				} else {
					for _, project := range projects {
						fmt.Printf("ID: %d\n", project.ID)
						fmt.Printf("Path: %s\n", project.Path)
						fmt.Printf("Name: %s\n", project.Name)

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
