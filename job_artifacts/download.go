package job_artifacts

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
)

// Download 下载产物归档文件 https://docs.gitlab.cn/jh/api/job_artifacts.html#%E4%B8%8B%E8%BD%BD%E4%BA%A7%E7%89%A9%E5%BD%92%E6%A1%A3%E6%96%87%E4%BB%B6
func Download() *cli.Command {
	return &cli.Command{
		Name:    "download",
		Aliases: []string{"dl"},
		Usage:   "下载产物归档文件（未完成）",
		Flags:   append(flag.Common(), flag.Id(true), flag.JobId(false)),
		Action: func(context *cli.Context) error {
			fmt.Println("下载产物归档文件")

			return fmt.Errorf("未完成")
		},
	}
}
