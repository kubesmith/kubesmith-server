package repos

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
)

type GetAllReposHandler struct {
	search string
}

func (h *GetAllReposHandler) filterRepos(repos []fixtures.Repository, text string) []fixtures.Repository {
	matches := []fixtures.Repository{}
	search := regexp.MustCompile(fmt.Sprintf("%s", strings.ToLower(text)))

	for _, repo := range repos {
		if search.MatchString(repo.Name) {
			matches = append(matches, repo)
		}
	}

	return matches
}

func (h *GetAllReposHandler) Process() ([]fixtures.Repository, error) {
	repos := fixtures.GetRepos()

	if h.search != "" {
		return h.filterRepos(repos, h.search), nil
	}

	return repos, nil
}

func GetAllRepos(server *factory.ServerFactory, c *gin.Context) {
	search, _ := c.GetQuery("search")

	handler := GetAllReposHandler{
		search,
	}

	repos, err := handler.Process()
	if err != nil {
		c.Status(500)
	} else {
		c.JSON(200, repos)
	}
}
