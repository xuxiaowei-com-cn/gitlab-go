package job_artifacts

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// JobsArtifacts 作业产物 API https://docs.gitlab.cn/jh/api/job_artifacts.html
func JobsArtifacts() *cli.Command {
	return &cli.Command{
		Name:    "job-artifact",
		Aliases: []string{"job-artifacts", "ja"},
		Usage:   "作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html",
		Flags:   append(flag.Common(), flag.Id(false), flag.JobId(false), flag.ArtifactsName()),
		Subcommands: []*cli.Command{
			Get(),
			Download(),
			Delete(),
			DeleteProject(),
		},
	}
}
