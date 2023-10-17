package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"github.com/xuxiaowei-com-cn/gitlab-go/constant"
	"github.com/xuxiaowei-com-cn/gitlab-go/flag"
	"log"
)

const (
	ScopeValue = ""
	ScopeUsage = "要显示的作业范围。以下之一或数组：created、pending、running、failed、success、canceled、skipped、waiting_for_resource 或 manual。范围如果未提供，则返回所有作业。"
)

// Jobs 作业 API https://docs.gitlab.cn/jh/api/jobs.html
func Jobs() *cli.Command {
	return &cli.Command{
		Name:    "job",
		Aliases: []string{"jobs", "j"},
		Usage:   "作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html",
		Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
			flag.Sort(), flag.Id(false), flag.Scope(ScopeValue, ScopeUsage)),
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "列出项目作业",
				Flags: append(flag.Common(), flag.Page(), flag.PerPage(), flag.PrintJson(), flag.PrintTime(),
					flag.Sort(), flag.Id(true), flag.Scope(ScopeValue, ScopeUsage)),
				Action: func(context *cli.Context) error {
					var baseUrl = context.String(constant.BaseUrl)
					var token = context.String(constant.Token)
					var id = context.String(constant.Id)
					var page = context.Int(constant.Page)
					var perPage = context.Int(constant.PerPage)
					var scope = context.StringSlice(constant.Scope)
					var printJson = context.Bool(constant.PrintJson)
					var printTime = context.Bool(constant.PrintTime)

					gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
					if err != nil {
						return err
					}

					jobs, response, err := ListProjectJobs(gitClient, id, scope, page, perPage)
					if err != nil {
						return err
					}
					log.Printf("Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

					fmt.Println("")

					if printJson {
						if printTime {
							for _, job := range jobs {
								jsonData, err := json.Marshal(job)
								if err != nil {
									panic(err)
								}

								log.Printf("\n%s\n", string(jsonData))
								fmt.Println("")
							}
						} else {
							for _, job := range jobs {
								jsonData, err := json.Marshal(job)
								if err != nil {
									panic(err)
								}

								fmt.Printf("%s\n", string(jsonData))
								fmt.Println("")
							}
						}
					} else {
						if printTime {
							for _, job := range jobs {
								log.Printf("ID: %d\n", job.ID)
								log.Printf("Name: %s\n", job.Name)
								log.Printf("Coverage: %f\n", job.Coverage)
								log.Printf("AllowFailure: %t\n", job.AllowFailure)
								log.Printf("CreatedAt: %s\n", job.CreatedAt)
								log.Printf("StartedAt: %s\n", job.StartedAt)
								log.Printf("FinishedAt: %s\n", job.FinishedAt)
								log.Printf("ErasedAt: %s\n", job.ErasedAt)
								log.Printf("Duration: %f\n", job.Duration)
								log.Printf("QueuedDuration: %f\n", job.QueuedDuration)
								log.Printf("ArtifactsExpireAt: %s\n", job.ArtifactsExpireAt)
								log.Printf("TagList: %s\n", job.TagList)
								log.Printf("Ref: %s\n", job.Ref)
								log.Printf("Stage: %s\n", job.Stage)
								log.Printf("Status: %s\n", job.Status)
								log.Printf("FailureReason: %s\n", job.FailureReason)
								log.Printf("Tag: %t\n", job.Tag)
								log.Printf("WebURL: %s\n", job.WebURL)
								log.Printf("Commit.ID: %s\n", job.Commit.ID)
								log.Printf("Commit.ShortID: %s\n", job.Commit.ShortID)
								log.Printf("Commit.Title: %s\n", job.Commit.Title)
								log.Printf("Commit.AuthorName: %s\n", job.Commit.AuthorName)
								log.Printf("Commit.AuthorEmail: %s\n", job.Commit.AuthorEmail)
								log.Printf("Commit.AuthoredDate: %s\n", job.Commit.AuthoredDate)
								log.Printf("Commit.CommitterName: %s\n", job.Commit.CommitterName)
								log.Printf("Commit.CommitterEmail: %s\n", job.Commit.CommitterEmail)
								log.Printf("Commit.CommittedDate: %s\n", job.Commit.CommittedDate)
								log.Printf("Commit.CreatedAt: %s\n", job.Commit.CreatedAt)
								log.Printf("Commit.Message: %s\n", job.Commit.Message)
								log.Printf("Commit.ParentIDs: %s\n", job.Commit.ParentIDs)
								log.Printf("Commit.Stats: %d\n", job.Commit.Stats)
								log.Printf("Commit.Status: %d\n", job.Commit.Status)
								log.Printf("Commit.LastPipeline: %s\n", job.Commit.LastPipeline)
								log.Printf("Commit.ProjectID: %d\n", job.Commit.ProjectID)
								log.Printf("Commit.Trailers: %s\n", job.Commit.Trailers)
								log.Printf("Commit.WebURL: %s\n", job.Commit.WebURL)
								log.Printf("Pipeline.ID: %d\n", job.Pipeline.ID)
								log.Printf("Pipeline.ProjectID: %d\n", job.Pipeline.ProjectID)
								log.Printf("Pipeline.Ref: %s\n", job.Pipeline.Ref)
								log.Printf("Pipeline.Sha: %s\n", job.Pipeline.Sha)
								log.Printf("Pipeline.Status: %s\n", job.Pipeline.Status)
								for index, jobArtifact := range job.Artifacts {
									log.Printf("Artifacts[%d].FileType: %s\n", index, jobArtifact.FileType)
									log.Printf("Artifacts[%d].Filename: %s\n", index, jobArtifact.Filename)
									log.Printf("Artifacts[%d].Size: %d\n", index, jobArtifact.Size)
									log.Printf("Artifacts[%d].FileFormat: %s\n", index, jobArtifact.FileFormat)
								}
								log.Printf("ArtifactsFile.Filename: %s\n", job.ArtifactsFile.Filename)
								log.Printf("ArtifactsFile.Size: %d\n", job.ArtifactsFile.Size)
								log.Printf("Runner.ID: %d\n", job.Runner.ID)
								log.Printf("Runner.Description: %s\n", job.Runner.Description)
								log.Printf("Runner.Active: %t\n", job.Runner.Active)
								log.Printf("Runner.IsShared: %t\n", job.Runner.IsShared)
								log.Printf("Runner.Name: %s\n", job.Runner.Name)

								// log.Printf("Project: %s\n", job.Project)
								// log.Printf("User: %s\n", job.User)

								fmt.Println("")
							}
						} else {
							for _, job := range jobs {
								fmt.Printf("ID: %d\n", job.ID)
								fmt.Printf("Name: %s\n", job.Name)
								fmt.Printf("Coverage: %f\n", job.Coverage)
								fmt.Printf("AllowFailure: %t\n", job.AllowFailure)
								fmt.Printf("CreatedAt: %s\n", job.CreatedAt)
								fmt.Printf("StartedAt: %s\n", job.StartedAt)
								fmt.Printf("FinishedAt: %s\n", job.FinishedAt)
								fmt.Printf("ErasedAt: %s\n", job.ErasedAt)
								fmt.Printf("Duration: %f\n", job.Duration)
								fmt.Printf("QueuedDuration: %f\n", job.QueuedDuration)
								fmt.Printf("ArtifactsExpireAt: %s\n", job.ArtifactsExpireAt)
								fmt.Printf("TagList: %s\n", job.TagList)
								fmt.Printf("Ref: %s\n", job.Ref)
								fmt.Printf("Stage: %s\n", job.Stage)
								fmt.Printf("Status: %s\n", job.Status)
								fmt.Printf("FailureReason: %s\n", job.FailureReason)
								fmt.Printf("Tag: %t\n", job.Tag)
								fmt.Printf("WebURL: %s\n", job.WebURL)
								fmt.Printf("Commit.ID: %s\n", job.Commit.ID)
								fmt.Printf("Commit.ShortID: %s\n", job.Commit.ShortID)
								fmt.Printf("Commit.Title: %s\n", job.Commit.Title)
								fmt.Printf("Commit.AuthorName: %s\n", job.Commit.AuthorName)
								fmt.Printf("Commit.AuthorEmail: %s\n", job.Commit.AuthorEmail)
								fmt.Printf("Commit.AuthoredDate: %s\n", job.Commit.AuthoredDate)
								fmt.Printf("Commit.CommitterName: %s\n", job.Commit.CommitterName)
								fmt.Printf("Commit.CommitterEmail: %s\n", job.Commit.CommitterEmail)
								fmt.Printf("Commit.CommittedDate: %s\n", job.Commit.CommittedDate)
								fmt.Printf("Commit.CreatedAt: %s\n", job.Commit.CreatedAt)
								fmt.Printf("Commit.Message: %s\n", job.Commit.Message)
								fmt.Printf("Commit.ParentIDs: %s\n", job.Commit.ParentIDs)
								fmt.Printf("Commit.Stats: %d\n", job.Commit.Stats)
								fmt.Printf("Commit.Status: %d\n", job.Commit.Status)
								fmt.Printf("Commit.LastPipeline: %s\n", job.Commit.LastPipeline)
								fmt.Printf("Commit.ProjectID: %d\n", job.Commit.ProjectID)
								fmt.Printf("Commit.Trailers: %s\n", job.Commit.Trailers)
								fmt.Printf("Commit.WebURL: %s\n", job.Commit.WebURL)
								fmt.Printf("Pipeline.ID: %d\n", job.Pipeline.ID)
								fmt.Printf("Pipeline.ProjectID: %d\n", job.Pipeline.ProjectID)
								fmt.Printf("Pipeline.Ref: %s\n", job.Pipeline.Ref)
								fmt.Printf("Pipeline.Sha: %s\n", job.Pipeline.Sha)
								fmt.Printf("Pipeline.Status: %s\n", job.Pipeline.Status)
								for index, jobArtifact := range job.Artifacts {
									fmt.Printf("Artifacts[%d].FileType: %s\n", index, jobArtifact.FileType)
									fmt.Printf("Artifacts[%d].Filename: %s\n", index, jobArtifact.Filename)
									fmt.Printf("Artifacts[%d].Size: %d\n", index, jobArtifact.Size)
									fmt.Printf("Artifacts[%d].FileFormat: %s\n", index, jobArtifact.FileFormat)
								}
								fmt.Printf("ArtifactsFile.Filename: %s\n", job.ArtifactsFile.Filename)
								fmt.Printf("ArtifactsFile.Size: %d\n", job.ArtifactsFile.Size)
								fmt.Printf("Runner.ID: %d\n", job.Runner.ID)
								fmt.Printf("Runner.Description: %s\n", job.Runner.Description)
								fmt.Printf("Runner.Active: %t\n", job.Runner.Active)
								fmt.Printf("Runner.IsShared: %t\n", job.Runner.IsShared)
								fmt.Printf("Runner.Name: %s\n", job.Runner.Name)

								// fmt.Printf("Project: %s\n", job.Project)
								// fmt.Printf("User: %s\n", job.User)

								fmt.Println("")
							}
						}
					}

					return nil
				},
			},
		},
	}
}

func ListProjectJobs(gitClient *gitlab.Client, pid interface{}, scope []string, page int, perPage int) ([]*gitlab.Job, *gitlab.Response, error) {
	var bsvs []gitlab.BuildStateValue
	for _, str := range scope {
		bsv := gitlab.BuildStateValue(str)
		bsvs = append(bsvs, bsv)
	}

	opt := &gitlab.ListJobsOptions{
		Scope: &bsvs,
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	jobs, response, err := gitClient.Jobs.ListProjectJobs(pid, opt)
	return jobs, response, err
}
