package flag

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
)

const (
	RangeMaxInterval = 10000
)

func Common() []cli.Flag {
	return []cli.Flag{
		BaseUrl(),
		Token(false),
	}
}

func CommonTokenRequired() []cli.Flag {
	return []cli.Flag{
		BaseUrl(),
		Token(true),
	}
}

func PrintJson() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.PrintJson,
		Value: constant.PrintJsonDefault,
		Usage: "打印 JSON",
	}
}

func PrintTime() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.PrintTime,
		Value: constant.PrintTimeDefault,
		Usage: "打印时间",
	}
}

func BaseUrl() cli.Flag {
	return &cli.StringFlag{
		Name:    constant.BaseUrl,
		Value:   constant.BaseUrlDefault,
		EnvVars: []string{"CI_API_V4_URL"},
		Usage:   "实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4",
	}
}

func Token(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.Token,
		Usage:    "your_access_token",
		Required: required,
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

func JobIdRange(required bool) cli.Flag {
	return &cli.StringSliceFlag{
		Name: constant.JobIdRange,
		Usage: "Job ID的范围，支持范围如下：\n\t" +
			"单数：1\n\t" +
			"多个数字（使用英文逗号隔开）：1,2,3,7,8,15\n\t" +
			"支持范围：5-10,\n\t" +
			fmt.Sprintf("支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + %d，数据范围不超过 %d）", RangeMaxInterval, RangeMaxInterval),
		Required: required,
	}
}

func IIdRange(required bool) cli.Flag {
	return &cli.StringSliceFlag{
		Name: constant.IIdRange,
		Usage: "流水线ID的范围，支持范围如下：\n\t" +
			"单数：1\n\t" +
			"多个数字（使用英文逗号隔开）：1,2,3,7,8,15\n\t" +
			"支持范围：5-10,\n\t" +
			fmt.Sprintf("支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + %d，数据范围不超过 %d）", RangeMaxInterval, RangeMaxInterval),
		Required: required,
	}
}

func Repository(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.Repository,
		Usage:    "仓库里存储库的 ID",
		Required: required,
	}
}

func TagName(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.TagName,
		Usage:    "标签的名称",
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

func ArtifactsName() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.ArtifactsName,
		Value: "artifacts.zip",
		Usage: "保存产物名称（保存到系统磁盘的名称）",
	}
}

func Page() cli.Flag {
	return &cli.IntFlag{
		Name:  constant.Page,
		Value: 1,
		Usage: "页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination",
	}
}

func PerPage() cli.Flag {
	return &cli.IntFlag{
		Name:  constant.PerPage,
		Value: 20,
		Usage: "每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination",
	}
}

func OrderBy(usage string) cli.Flag {
	return &cli.StringFlag{
		Name:  constant.OrderBy,
		Value: "created_at",
		Usage: usage,
	}
}

func Scope(value string, usage string) cli.Flag {
	return &cli.StringFlag{
		Name:  constant.Scope,
		Value: value,
		Usage: usage,
	}
}

func IssueId(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.IssueId,
		Usage:    "项目议题的内部 ID",
		Required: required,
	}
}

func Recursion() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.Recursion,
		Usage: "递归",
	}
}

func Owned(required bool) cli.Flag {
	return &cli.BoolFlag{
		Name:     constant.Owned,
		Usage:    "当前用户明确拥有的项目。",
		Required: required,
	}
}

func NamespaceSource(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.NamespaceSource,
		Usage:    "源命名空间。如：用户名、群组名",
		Required: required,
	}
}

func NamespaceTargetFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.NamespaceTarget,
		Usage:    "目标命名空间。如：用户名、群组名",
		Required: required,
	}
}

func IssueIdRange(required bool) cli.Flag {
	return &cli.StringSliceFlag{
		Name: constant.IssueIdRange,
		Usage: "议题ID的范围，支持范围如下：\n\t" +
			"单数：1\n\t" +
			"多个数字（使用英文逗号隔开）：1,2,3,7,8,15\n\t" +
			"支持范围：5-10,\n\t" +
			fmt.Sprintf("支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + %d，数据范围不超过 %d）", RangeMaxInterval, RangeMaxInterval),
		Required: required,
	}
}

func ExportFolder(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.ExportFolder,
		Usage:    "导出文件夹",
		Required: required,
	}
}

func SkipProjectPath() cli.Flag {
	return &cli.StringSliceFlag{
		Name:  constant.SkipProjectPath,
		Usage: "跳过项目路径",
	}
}

func SkipProjectWikiPath() cli.Flag {
	return &cli.StringSliceFlag{
		Name:  constant.SkipProjectWikiPath,
		Usage: "跳过项目wiki路径",
	}
}

func AutoSkipExistFolder() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.AutoSkipExistFolder,
		Usage: "自动跳过已存在的文件夹",
		Value: false,
	}
}

func AllowFailure() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.AllowFailure,
		Usage: "允许失败",
		Value: false,
	}
}

func EnvName(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.EnvName,
		Usage:    "环境名称",
		Required: required,
	}
}

func EnvExternalUrl() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.EnvExternalUrl,
		Usage: "该环境的链接位置",
	}
}

func EnvTier() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.EnvTier,
		Usage: "新环境的层级。允许设置的值为 production， staging， testing， development 和 other",
	}
}

func VariableKey(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.VariableKey,
		Usage:    "变量的 key。不能超过 255 个字符。仅支持 A-Z、a-z、0-9 和 _",
		Required: required,
	}
}

func VariableValue(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     constant.VariableValue,
		Usage:    "变量的 value",
		Required: required,
	}
}

func VariableType() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.VariableType,
		Usage: "变量类型。可用类型为：env_var 和 file",
		Value: "env_var",
	}
}

func VariableProtected() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.VariableProtected,
		Usage: "变量是否受保护。",
		Value: false,
	}
}

func VariableMasked() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.VariableMasked,
		Usage: "变量是否隐藏。",
		Value: false,
	}
}

func VariableRaw() cli.Flag {
	return &cli.BoolFlag{
		Name:  constant.VariableRaw,
		Usage: "变量是否被视为原始字符串。当为 true 时，值中的变量不会扩展",
		Value: false,
	}
}

func VariableEnvironmentScope() cli.Flag {
	return &cli.StringFlag{
		Name:  constant.VariableEnvironmentScope,
		Usage: "变量的 environment_scope。",
		Value: "*",
	}
}
