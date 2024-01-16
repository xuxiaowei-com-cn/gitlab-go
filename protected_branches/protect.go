package protected_branches

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

// Protect 保护仓库分支 https://docs.gitlab.cn/jh/api/protected_branches.html#%E4%BF%9D%E6%8A%A4%E4%BB%93%E5%BA%93%E5%88%86%E6%94%AF
func Protect() *cli.Command {
	return &cli.Command{
		Name:  "protect",
		Usage: "保护仓库分支",
		Flags: append(flag.CommonTokenRequired(), flag.Id(true),
			flag.BranchName(true), flag.PushAccessLevel(), flag.MergeAccessLevel(),
			flag.UnprotectAccessLevel(), flag.AllowForcePush(), flag.CodeOwnerApprovalRequired(),
			flag.PrintJson(), flag.PrintTime()),
		Action: func(context *cli.Context) error {
			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var id = context.String(constant.Id)

			var name = context.String(constant.BranchName)
			var pushAccessLevel = context.Int(constant.PushAccessLevel)
			var mergeAccessLevel = context.Int(constant.MergeAccessLevel)
			var unprotectAccessLevel = context.Int(constant.UnprotectAccessLevel)
			var allowForcePush = context.Bool(constant.AllowForcePush)
			var codeOwnerApprovalRequired = context.Bool(constant.CodeOwnerApprovalRequired)

			var printJson = context.Bool(constant.PrintJson)
			var printTime = context.Bool(constant.PrintTime)

			return ProtectRepositoryBranches(baseUrl, token, id, name,
				pushAccessLevel, mergeAccessLevel, unprotectAccessLevel, allowForcePush, codeOwnerApprovalRequired,
				printJson, printTime, false)
		},
	}
}

func ProtectRepositoryBranches(baseUrl string, token string, id interface{}, name string,
	pushAccessLevel int, mergeAccessLevel int, unprotectAccessLevel int, allowForcePush bool, codeOwnerApprovalRequired bool,
	printJson bool, printTime bool, allowFailure bool) error {
	gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return err
	}

	var pushAccessLevelValue gitlab.AccessLevelValue
	switch pushAccessLevel {
	case 0, 5, 10, 20, 30, 40, 50, 60:
		pushAccessLevelValue = gitlab.AccessLevelValue(pushAccessLevel)
	default:
		return errors.New(fmt.Sprintf("pushAccessLevel：%d 不合法", pushAccessLevel))
	}

	var mergeAccessLevelValue gitlab.AccessLevelValue
	switch mergeAccessLevel {
	case 0, 5, 10, 20, 30, 40, 50, 60:
		mergeAccessLevelValue = gitlab.AccessLevelValue(mergeAccessLevel)
	default:
		return errors.New(fmt.Sprintf("mergeAccessLevel：%d 不合法", mergeAccessLevel))
	}

	var unprotectAccessLevelValue gitlab.AccessLevelValue
	switch unprotectAccessLevel {
	case 0, 5, 10, 20, 30, 40, 50, 60:
		unprotectAccessLevelValue = gitlab.AccessLevelValue(unprotectAccessLevel)
	default:
		return errors.New(fmt.Sprintf("unprotectAccessLevel：%d 不合法", unprotectAccessLevel))
	}

	opt := &gitlab.ProtectRepositoryBranchesOptions{
		Name:                 &name,
		PushAccessLevel:      &pushAccessLevelValue,
		MergeAccessLevel:     &mergeAccessLevelValue,
		UnprotectAccessLevel: &unprotectAccessLevelValue,
		AllowForcePush:       &allowForcePush,
		//AllowedToPush:             &allowedToPush,
		//AllowedToMerge:            &allowedToMerge,
		//AllowedToUnprotect:        &allowedToUnprotect,
		CodeOwnerApprovalRequired: &codeOwnerApprovalRequired,
	}

	protectedBranch, response, err := gitClient.ProtectedBranches.ProtectRepositoryBranches(id, opt)
	if err != nil {
		if allowFailure {
			return nil
		}
		return err
	}
	log.Printf("Response StatusCode: %d\n", response.Response.StatusCode)

	fmt.Println("")

	if printJson {
		if printTime {
			jsonData, err := json.Marshal(protectedBranch)
			if err != nil {
				panic(err)
			}

			log.Printf("\n%s\n", string(jsonData))
			fmt.Println("")

		} else {
			jsonData, err := json.Marshal(protectedBranch)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s\n", string(jsonData))
			fmt.Println("")

		}
	} else {
		if printTime {
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

		} else {
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

	return nil
}
