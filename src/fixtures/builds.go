package fixtures

const STATUS_RUNNING = "running"

type Build struct {
	ID                  string          `json:"id"`
	Repository          BuildRepository `json:"repo"`
	Message             string          `json:"message"`
	SHA                 string          `json:"sha"`
	Branch              string          `json:"branch"`
	Reference           string          `json:"reference"`
	DateCreated         int             `json:"dateCreated"`
	DateStarted         int             `json:"dateStarted"`
	DateFinished        int             `json:"dateFinished"`
	CurrentPipelineName string          `json:"currentPipelineName"`
	Status              string          `json:"status"`
	Author              BuildAuthor     `json:"author"`
}

type BuildRepository struct {
	ID int `json:"id"`
}

type BuildAuthor struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GetBuilds() []Build {
	return []Build{
		Build{
			ID: "P8nDg7yPjT5F6t",
			Repository: BuildRepository{
				ID: 456745674,
			},
			Message:             "Merge pull request #1 from some super long PR title that probably should have been shortened but no one wanted to shorten it",
			SHA:                 "ce0f4b6f234fdg4sdfg44tafgwaert366",
			Branch:              "master",
			Reference:           "refs/tags/0.1.14",
			DateCreated:         1554148197,
			DateStarted:         1554148297,
			DateFinished:        0,
			CurrentPipelineName: "Install Dependencies",
			Status:              STATUS_RUNNING,
			Author: BuildAuthor{
				Username: "carldanley",
				Email:    "carldanley@gmail.com",
				Avatar:   "https://avatars3.githubusercontent.com/u/1470571?v=4",
			},
		},
	}
}
