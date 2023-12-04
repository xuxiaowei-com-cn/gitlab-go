package mix

import (
	"errors"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExportAll() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage: "导出所有（混合命令，多接口命令）\n" +
			"已包含：\n" +
			"1. git 仓库\n" +
			"2. wiki 仓库",
		Flags: append(flag.CommonTokenRequired(), flag.Username(true), flag.Owned(true),
			flag.ExportFolder(true), flag.SkipProjectPath(), flag.SkipProjectWikiPath()),
		Action: func(context *cli.Context) error {

			var baseUrl = context.String(constant.BaseUrl)
			var token = context.String(constant.Token)
			var username = context.String(constant.Username)
			var owned = context.Bool(constant.Owned)
			var exportFolder = context.String(constant.ExportFolder)
			var skipProjectPaths = context.StringSlice(constant.SkipProjectPath)
			var skipProjectWikiPaths = context.StringSlice(constant.SkipProjectWikiPath)

			baseURL, err := url.Parse(baseUrl)
			if err != nil {
				return err
			}

			host := baseURL.Host

			for index, skipProjectPath := range skipProjectPaths {
				skipProjectPath = strings.TrimPrefix(skipProjectPath, "/")
				skipProjectPath = strings.TrimSuffix(skipProjectPath, "/")
				skipProjectPaths[index] = skipProjectPath
			}

			for index, skipProjectWikiPath := range skipProjectWikiPaths {
				skipProjectWikiPath = strings.TrimPrefix(skipProjectWikiPath, "/")
				skipProjectWikiPath = strings.TrimSuffix(skipProjectWikiPath, "/")
				skipProjectWikiPaths[index] = skipProjectWikiPath
			}

			projectList, err := projects.ListProjects(owned, token, baseUrl, 1, 100)
			if err != nil {
				return err
			}

			for index, project := range projectList {
				log.Printf("Project Index: %d, WebURL: %s", index, project.WebURL)

				err = Repository(exportFolder, host, username, token, project, skipProjectPaths)
				if err != nil {
					return err
				}

				err = Wiki(exportFolder, host, username, token, project, skipProjectWikiPaths)
				if err != nil {
					return err
				}
			}

			log.Printf("导出完成")

			return nil
		},
	}
}

func Repository(exportFolder string, host string, username string, token string, project *gitlab.Project, skipProjectPaths []string) error {

	c := contains(skipProjectPaths, project.PathWithNamespace)
	if c {
		log.Printf("跳过：%s", project.PathWithNamespace)
		return nil
	}

	gitPath := filepath.Join(exportFolder, "repository", host, project.PathWithNamespace)
	err := os.MkdirAll(gitPath, os.ModePerm)
	if err != nil {
		return err
	}

	HTTPURLToRepo := project.HTTPURLToRepo
	repoUrl, err := url.Parse(HTTPURLToRepo)
	if err != nil {
		return err
	}

	userinfo := url.UserPassword(username, token)
	repoUrl.User = userinfo

	cmd := exec.Command("git", "clone", repoUrl.String(), gitPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	err = GitRemoteSetUrl(gitPath, HTTPURLToRepo)
	if err != nil {
		return err
	}

	return nil
}

func Wiki(exportFolder string, host string, username string, token string, project *gitlab.Project, skipProjectWikiPaths []string) error {

	c := contains(skipProjectWikiPaths, project.PathWithNamespace)
	if c {
		log.Printf("跳过：%s", project.PathWithNamespace)
		return nil
	}

	wikiPath := filepath.Join(exportFolder, "wiki", host, project.PathWithNamespace+".wiki")
	err := os.MkdirAll(wikiPath, os.ModePerm)
	if err != nil {
		return err
	}

	var HTTPURLToRepo string
	if strings.HasSuffix(project.HTTPURLToRepo, ".git") {
		HTTPURLToRepo = strings.TrimSuffix(project.HTTPURLToRepo, ".git") + ".wiki.git"
	} else {
		return errors.New("仓库地址异常: " + project.HTTPURLToRepo)
	}

	repoUrl, err := url.Parse(HTTPURLToRepo)
	if err != nil {
		return err
	}

	userinfo := url.UserPassword(username, token)
	repoUrl.User = userinfo

	cmd := exec.Command("git", "clone", repoUrl.String(), wikiPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	err = GitRemoteSetUrl(wikiPath, HTTPURLToRepo)
	if err != nil {
		return err
	}

	return nil
}

func GitRemoteSetUrl(gitPath string, remoteUrl string) error {
	cmd := exec.Command("git", "remote", "set-url", "origin", remoteUrl)
	cmd.Dir = gitPath
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func contains(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
