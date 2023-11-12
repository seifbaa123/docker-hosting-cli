package types

type ProjectConfig struct {
	Id   int    `yaml:"id"`
	Name string `yaml:"name"`
}

type Response struct {
	Message string `json:"message"`
	Image   struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		HasBuild bool   `json:"hasBuild"`
	} `json:"image"`
}
