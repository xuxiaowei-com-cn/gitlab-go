package flag

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

func Common() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    constant.BaseUrl,
			Value:   constant.BaseUrlDefault,
			EnvVars: []string{"CI_API_V4_URL"},
			Usage:   "实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4",
		},
		&cli.StringFlag{
			Name:  constant.Token,
			Usage: "your_access_token",
		},
	}
}

func Sort() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Sort,
		Value: constant.SortDefault,
		Usage: "按照 asc 或者 desc 排序",
	}
}

func Id(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.Id,
		Usage:    "项目 ID 或 URL 编码的路径",
		Required: required,
	}
}

func JobId(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.JobId,
		Usage:    "作业 ID",
		Required: required,
	}
}
