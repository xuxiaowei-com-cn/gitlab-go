package main

import (
	"fmt"
	"github.com/xuxiaowei-com-cn/git-go/buildinfo"
	"github.com/xuxiaowei-com-cn/gitlab-go/access_requests"
	"github.com/xuxiaowei-com-cn/gitlab-go/boards"
	"github.com/xuxiaowei-com-cn/gitlab-go/instance_level_ci_variables"
	"github.com/xuxiaowei-com-cn/gitlab-go/issues"
	"github.com/xuxiaowei-com-cn/gitlab-go/job_artifacts"
	"github.com/xuxiaowei-com-cn/gitlab-go/jobs"
	"github.com/xuxiaowei-com-cn/gitlab-go/pipelines"
	"github.com/xuxiaowei-com-cn/gitlab-go/projects"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	Name              = "gitlab-go"
	Description       = "基于 Go 语言开发的 GitLab 命令行工具"
	URL               = "https://github.com/xuxiaowei-com-cn/gitlab-go.git"
	BugReportUrl      = "https://github.com/xuxiaowei-com-cn/gitlab-go/issues"
	OrganizationName  = "徐晓伟工作室"
	OrganizationUrl   = "http://xuxiaowei.com.cn"
	OrganizationEmail = "xuxiaowei@xuxiaowei.com.cn"
	Copyright         = "徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>"
	Author            = "徐晓伟"
	Email             = "xuxiaowei@xuxiaowei.com.cn"
)

var (
	BuildDate          string // 构建时间，如：2023-07-19T12:20:54Z
	Compiler           string // 编译器，如：gc
	GitCommitBranch    string // 提交分支名称
	GitCommitSha       string // 项目为其构建的提交修订 fa3d7990104d7c1f16943a67f11b154b71f6a132
	GitCommitShortSha  string // 项目为其构建的提交修订的前八个字符 fa3d7990
	GitCommitTag       string // 提交标签名称
	GitCommitTimestamp string // ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00
	GitTreeState       string // clean
	GitVersion         string // git 版本号，如：v1.27.4
	GoVersion          string // go 版本号，如：go1.20.6
	Major              string // 主版本，如：1
	Minor              string // 次版本，如：27
	Revision           string // 修订版本，如：4
	Platform           string // 平台，如：linux/amd64
	CiPipelineId       string // 流水线，如：ID8754
	CiJobId            string // 作业ID，如：14468
)

func init() {
	if GitVersion == "" {
		GitVersion = "dev"
	}
}

func main() {
	app := &cli.App{
		Name:      Name,
		Version:   versionInfo(),
		Authors:   []*cli.Author{{Name: Author, Email: Email}},
		Usage:     Description,
		Copyright: Copyright,
		Commands: []*cli.Command{
			access_requests.AccessRequests(),
			boards.Boards(),
			instance_level_ci_variables.InstanceLevelCiVariables(),
			issues.Issues(),
			job_artifacts.JobsArtifacts(),
			jobs.Jobs(),
			pipelines.Pipelines(),
			projects.Projects(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func versionInfo() string {
	var info = buildinfo.Version{
		Name:         Name,
		Description:  Description,
		URL:          URL,
		BugReportUrl: BugReportUrl,
		BuildVersion: buildinfo.BuildVersion{
			BuildDate:          BuildDate,
			Compiler:           Compiler,
			GitCommitSha:       GitCommitSha,
			GitCommitShortSha:  GitCommitShortSha,
			GitCommitTag:       GitCommitTag,
			GitCommitTimestamp: GitCommitTimestamp,
			GitCommitBranch:    GitCommitBranch,
			GitTreeState:       GitTreeState,
			GitVersion:         GitVersion,
			GoVersion:          GoVersion,
			Major:              Major,
			Minor:              Minor,
			Revision:           Revision,
			Platform:           Platform,
			CiPipelineId:       CiPipelineId,
			CiJobId:            CiJobId,
		},
		Organization: buildinfo.Organization{
			Name:  OrganizationName,
			Url:   OrganizationUrl,
			Email: OrganizationEmail,
		},
	}

	if len(os.Args) > 1 && os.Args[1] == "--version" {
		yamlData, err := yaml.Marshal(info)
		if err != nil {
			fmt.Println("版本信息无法转换为 YAML 格式:", err)
			return ""
		}

		return string(yamlData)
	}

	return info.BuildVersion.GitVersion
}
