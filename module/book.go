package module

type Book struct {
	Bid       int    `json:"bid" gorm:"primary_key"`
	Name      string `json:"name" gorm:"size:100;not null"`
	Author    string `json:"author" gorm:"size:100;not null"`
	Publisher string `json:"publisher" gorm:"size:100;not null"`
	Intro     string `json:"intro" gorm:"size:255"`
	CoverFile string `json:"cover_file" gorm:"size:255"`
}

func NewBook() *Book {
	return &Book{}
}
