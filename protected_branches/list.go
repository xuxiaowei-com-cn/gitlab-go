package protected_branches

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// List 列出受保护的分支 https://docs.gitlab.cn/jh/api/protected_branches.html#%E5%88%97%E5%87%BA%E5%8F%97%E4%BF%9D%E6%8A%A4%E7%9A%84%E5%88%86%E6%94%AF
func List() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "列出受保护的分支",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true), flag.ProtectedBranchesSearch(),
			flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)
			var search = context.String(constant.Search)
			var page = context.Int(constant.Page)
			var perPage = context.Int(constant.PerPage)
			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
			if err != nil {
				return err
			}

			opt := &gitlab.ListProtectedBranchesOptions{
				Search: &search,
				ListOptions: gitlab.ListOptions{
					Page:    page,
					PerPage: perPage,
				},
			}
			protectedBranchs, response, err := gitClient.ProtectedBranches.ListProtectedBranches(id, opt)
			if err != nil {
				return err
			}
			log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

			fmt.Println("")

			if printJson {
				if printTime {
					for _, protectedBranch := range protectedBranchs {
						jsonData, err := json.Marshal(protectedBranch)
						if err != nil {
							panic(err)
						}

						log.Printf("\n%s\n", string(jsonData))
						fmt.Println("")
					}
				} else {
					for _, protectedBranch := range protectedBranchs {
						jsonData, err := json.Marshal(protectedBranch)
						if err != nil {
							panic(err)
						}

						fmt.Printf("%s\n", string(jsonData))
						fmt.Println("")
					}
				}
			} else {
				if printTime {
					for _, protectedBranch := range protectedBranchs {
						log.Printf("ID: %d\n", protectedBranch.ID)
						log.Printf("Name: %s\n", protectedBranch.Name)
						log.Printf("AllowForcePush: %t\n", protectedBranch.AllowForcePush)
						log.Printf("CodeOwnerApprovalRequired: %t\n", protectedBranch.CodeOwnerApprovalRequired)

						log.Printf("PushAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.PushAccessLevels {
							log.Printf("PushAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							log.Printf("PushAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							log.Printf("PushAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							log.Printf("PushAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							log.Printf("PushAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						log.Printf("MergeAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.MergeAccessLevels {
							log.Printf("MergeAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							log.Printf("MergeAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							log.Printf("MergeAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							log.Printf("MergeAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							log.Printf("MergeAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						log.Printf("UnprotectAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.UnprotectAccessLevels {
							log.Printf("UnprotectAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							log.Printf("UnprotectAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							log.Printf("UnprotectAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							log.Printf("UnprotectAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							log.Printf("UnprotectAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						fmt.Println("")
					}
				} else {
					for _, protectedBranch := range protectedBranchs {
						fmt.Printf("ID: %d\n", protectedBranch.ID)
						fmt.Printf("Name: %s\n", protectedBranch.Name)
						fmt.Printf("AllowForcePush: %t\n", protectedBranch.AllowForcePush)
						fmt.Printf("CodeOwnerApprovalRequired: %t\n", protectedBranch.CodeOwnerApprovalRequired)

						fmt.Printf("PushAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.PushAccessLevels {
							fmt.Printf("PushAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							fmt.Printf("PushAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							fmt.Printf("PushAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							fmt.Printf("PushAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							fmt.Printf("PushAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						fmt.Printf("MergeAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.MergeAccessLevels {
							fmt.Printf("MergeAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							fmt.Printf("MergeAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							fmt.Printf("MergeAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							fmt.Printf("MergeAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							fmt.Printf("MergeAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						fmt.Printf("UnprotectAccessLevels: \n")
						for index, branchAccessDescription := range protectedBranch.UnprotectAccessLevels {
							fmt.Printf("UnprotectAccessLevels[%d] ID: %d\n", index, branchAccessDescription.ID)
							fmt.Printf("UnprotectAccessLevels[%d] AccessLevel: %d\n", index, branchAccessDescription.AccessLevel)
							fmt.Printf("UnprotectAccessLevels[%d] AccessLevelDescription: %s\n", index, branchAccessDescription.AccessLevelDescription)
							fmt.Printf("UnprotectAccessLevels[%d] UserID: %d\n", index, branchAccessDescription.UserID)
							fmt.Printf("UnprotectAccessLevels[%d] GroupID: %d\n", index, branchAccessDescription.GroupID)
						}

						fmt.Println("")
					}
				}
			}

			return nil
		},
	}
}
