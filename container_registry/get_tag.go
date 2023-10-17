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

// GetTag 获取仓库里存储库的某个标签的详情 https://docs.gitlab.cn/jh/api/container_registry.html#%E8%8E%B7%E5%8F%96%E4%BB%93%E5%BA%93%E9%87%8C%E5%AD%98%E5%82%A8%E5%BA%93%E7%9A%84%E6%9F%90%E4%B8%AA%E6%A0%87%E7%AD%BE%E7%9A%84%E8%AF%A6%E6%83%85
func GetTag() *cli.Command {
	return &cli.Command{
		Name:  "get-tag",
		Usage: "获取仓库里存储库的某个标签的详情",
		Flags: append(flag.CommonTokenRequired(), flag.PrintJson(), flag.PrintTime(),
			flag.Id(true), flag.Repository(true), flag.TagName(true)),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var repository = context.Int(constant.Repository)
			var tagName = context.String(constant.TagName)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			registryRepositoryTag, response, err := gitClient.ContainerRegistry.GetRegistryRepositoryTagDetail(id, repository, tagName)
			if err != nil {
				return err
			}
			log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					jsonData, err := json.Marshal(registryRepositoryTag)
					if err != nil {
						panic(err)
					}

					log.Printf("\n%s\n", string(jsonData))
					fmt.Println("")
				} else {
					jsonData, err := json.Marshal(registryRepositoryTag)
					if err != nil {
						panic(err)
					}

					fmt.Printf("%s\n", string(jsonData))
					fmt.Println("")
				}
			} else {
				if printTime {
					log.Printf("Name: %s\n", registryRepositoryTag.Name)
					log.Printf("Path: %s\n", registryRepositoryTag.Path)
					log.Printf("Location: %s\n", registryRepositoryTag.Location)
					log.Printf("Revision: %s\n", registryRepositoryTag.Revision)
					log.Printf("ShortRevision: %s\n", registryRepositoryTag.ShortRevision)
					log.Printf("Digest: %s\n", registryRepositoryTag.Digest)
					log.Printf("CreatedAt: %s\n", registryRepositoryTag.CreatedAt)
					log.Printf("TotalSize: %d\n", registryRepositoryTag.TotalSize)
					fmt.Println("")
				} else {
					fmt.Printf("Name: %s\n", registryRepositoryTag.Name)
					fmt.Printf("Path: %s\n", registryRepositoryTag.Path)
					fmt.Printf("Location: %s\n", registryRepositoryTag.Location)
					fmt.Printf("Revision: %s\n", registryRepositoryTag.Revision)
					fmt.Printf("ShortRevision: %s\n", registryRepositoryTag.ShortRevision)
					fmt.Printf("Digest: %s\n", registryRepositoryTag.Digest)
					fmt.Printf("CreatedAt: %s\n", registryRepositoryTag.CreatedAt)
					fmt.Printf("TotalSize: %d\n", registryRepositoryTag.TotalSize)
					fmt.Println("")
				}
			}

			return nil
		},
	}
}
