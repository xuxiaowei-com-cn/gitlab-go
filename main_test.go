package main

import (
	"bytes"
	"fmt"
	"github.com/xuxiaowei-com-cn/git-go/buildinfo"
	"os"
	"strings"
	"testing"
)

func Test_CommitSha(t *testing.T) {
	fmt.Println(buildinfo.CommitSha())
}

func Test_Projects(t *testing.T) {
	var buf bytes.Buffer

	// 不带参数的测试输出
	os.Args = []string{"cmd"}
	main()
	if got := buf.String(); !strings.Contains(got, "") {
		t.Errorf("异常信息：\n%s", got)
	}

	buf.Reset()

	// 使用版本参数测试输出
	os.Args = []string{"cmd", "projects"}
	main()
	if got := buf.String(); !strings.Contains(got, "") {
		t.Errorf("异常信息：\n%s", got)
	}

	// 使用版本参数测试输出
	os.Args = []string{"cmd", "p"}
	main()
	if got := buf.String(); !strings.Contains(got, "") {
		t.Errorf("异常信息：\n%s", got)
	}
}
