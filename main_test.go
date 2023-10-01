package main

import (
	"fmt"
	"github.com/xuxiaowei-com-cn/git-go/buildinfo"
	"testing"
)

func Test_CommitSha(t *testing.T) {
	fmt.Println(buildinfo.CommitSha())
}
