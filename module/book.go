package module

type Book struct {
	Bid       int    `json:"bid"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Intro     string `json:"intro"`
}
