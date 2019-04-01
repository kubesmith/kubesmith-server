package fixtures

type Repository struct {
	ID               int    `json:"id"`
	Provider         string `json:"provider"`
	Name             string `json:"name"`
	URL              string `json:"url"`
	HasRunningBuilds bool   `json:"hasRunningBuilds"`
}

const PROVIDER_BITBUCKET = "bitbucket"
const PROVIDER_GITHUB = "github"

func GetRepos() []Repository {
	return []Repository{
		Repository{
			ID:               136832021,
			Provider:         PROVIDER_GITHUB,
			Name:             "kubesmith/kubesmith-ui",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               245624565,
			Provider:         PROVIDER_GITHUB,
			Name:             "carldanley/website",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               456745674,
			Provider:         PROVIDER_GITHUB,
			Name:             "kubesmith/kubesmith",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               123412355,
			Provider:         PROVIDER_GITHUB,
			Name:             "carldanley/dotfiles",
			URL:              "https://github.com",
			HasRunningBuilds: true,
		},
		Repository{
			ID:               234523454,
			Provider:         PROVIDER_BITBUCKET,
			Name:             "carldanley/homelab",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               234523424,
			Provider:         PROVIDER_GITHUB,
			Name:             "carldanley/go-services",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               789078900,
			Provider:         PROVIDER_GITHUB,
			Name:             "carldanley/radium-cluster",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               678698673,
			Provider:         PROVIDER_BITBUCKET,
			Name:             "carldanley/ghost-personal",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
		Repository{
			ID:               789000900,
			Provider:         PROVIDER_GITHUB,
			Name:             "carldanley/radium-grafana-dashboards",
			URL:              "https://github.com",
			HasRunningBuilds: false,
		},
	}
}
