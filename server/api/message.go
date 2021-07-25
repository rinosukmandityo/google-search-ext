package api

type Search struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
	Domain string `json:"domain"`
	Cite   string `json:"cite"`
}
