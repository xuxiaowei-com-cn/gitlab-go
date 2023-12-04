package groups

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

// ListGroups 列出群组 https://docs.gitlab.cn/jh/api/groups.html#%E5%88%97%E5%87%BA%E7%BE%A4%E7%BB%84
func ListGroups(token string, baseUrl string, page int, perPage int) ([]*gitlab.Group, error) {

	var results []*gitlab.Group

	gitClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return nil, err
	}

	opt := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}

	groups, response, err := gitClient.Groups.ListGroups(opt)
	if err != nil {
		return nil, err
	}

	log.Printf("ListGroups Page %d, PerPage: %d, Response StatusCode: %d\n", page, perPage, response.Response.StatusCode)

	results = append(results, groups...)

	if len(groups) != 0 {
		groups, err = ListGroups(token, baseUrl, page+1, perPage)
		if err != nil {
			return nil, err
		}

		results = append(results, groups...)
	}

	return results, nil
}
