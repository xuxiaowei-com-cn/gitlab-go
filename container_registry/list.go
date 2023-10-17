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

// List 列出仓库内存储库 https://docs.gitlab.cn/jh/api/container_registry.html#%E5%88%97%E5%87%BA%E4%BB%93%E5%BA%93%E5%86%85%E5%AD%98%E5%82%A8%E5%BA%93
func List() *cli.Command {
	return &cli.Command{
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

			registryRepositories, response, err := gitClient.ContainerRegistry.ListProjectRegistryRepositories(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, registryRepository := range registryRepositories {
						jsonData, err := json.Marshal(registryRepository)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, registryRepository := range registryRepositories {
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
					for _, registryRepository := range registryRepositories {
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
					for _, registryRepository := range registryRepositories {
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
	}
}
