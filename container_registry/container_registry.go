package container_registry

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// ContainerRegistry 容器仓库 API https://docs.gitlab.cn/jh/api/container_registry.html
func ContainerRegistry() *cli.Command {
	return &cli.Command{
		Name:    "container-registry",
		Aliases: []string{"cr"},
		Usage:   "容器仓库 API，中文文档：https://docs.gitlab.cn/jh/api/container_registry.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(false)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出仓库内存储库",
				Flags: append(flag.CommonTokenRequired(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
					flag.Id(true)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var printJson = context.Bool(constant.PrintJson)
					var printTime = context.Bool(constant.PrintTime)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}
					opt := &gitlab.ListRegistryRepositoriesOptions{
						ListOptions: gitlab.ListOptions{
							Page:    page,
							PerPage: perPage,
						},
					}

					registryRepositorys, response, err := gitClient.ContainerRegistry.ListProjectRegistryRepositories(id, opt)
					log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)
					if err != nil {
						return err
					}

					fmt.Println("")

					if printJson {
						if printTime {
							for _, registryRepository := range registryRepositorys {
								jsonData, err := json.Marshal(registryRepository)
								if err != nil {
									panic(err)
								}

								log.Printf("\n%s\n", string(jsonData))
								fmt.Println("")
							}
						} else {
							for _, registryRepository := range registryRepositorys {
								jsonData, err := json.Marshal(registryRepository)
								if err != nil {
									panic(err)
								}

								fmt.Printf("%s\n", string(jsonData))
								fmt.Println("")
							}
						}
					} else {
						if printTime {
							for _, registryRepository := range registryRepositorys {
								log.Printf("ID: %d\n", registryRepository.ID)
								log.Printf("Name: %s\n", registryRepository.Name)
								log.Printf("Path: %s\n", registryRepository.Path)
								log.Printf("ProjectID: %d\n", registryRepository.ProjectID)
								log.Printf("Location: %s\n", registryRepository.Location)
								log.Printf("CreatedAt: %s\n", registryRepository.CreatedAt)
								log.Printf("CleanupPolicyStartedAt: %s\n", registryRepository.CleanupPolicyStartedAt)
								log.Printf("TagsCount: %d\n", registryRepository.TagsCount)
								log.Printf("Tags: %s\n", registryRepository.Tags)
								fmt.Println("")
							}
						} else {
							for _, registryRepository := range registryRepositorys {
								fmt.Printf("ID: %d\n", registryRepository.ID)
								fmt.Printf("Name: %s\n", registryRepository.Name)
								fmt.Printf("Path: %s\n", registryRepository.Path)
								fmt.Printf("ProjectID: %d\n", registryRepository.ProjectID)
								fmt.Printf("Location: %s\n", registryRepository.Location)
								fmt.Printf("CreatedAt: %s\n", registryRepository.CreatedAt)
								fmt.Printf("CleanupPolicyStartedAt: %s\n", registryRepository.CleanupPolicyStartedAt)
								fmt.Printf("TagsCount: %d\n", registryRepository.TagsCount)
								fmt.Printf("Tags: %s\n", registryRepository.Tags)
								fmt.Println("")
							}
						}
					}

					return nil
				},
			},
		},
	}
}
