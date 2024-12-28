package module

type Book struct {
	Bid       int    `json:"bid" gorm:"primary_key" form:"bid"`
	Name      string `json:"name" gorm:"size:100;not null" form:"name"`
	Author    string `json:"author" gorm:"size:100;not null" form:"author"`
	Publisher string `json:"publisher" gorm:"size:100;not null" form:"publisher"`
	Intro     string `json:"intro" gorm:"size:255" form:"intro"`
	CoverFile string `json:"cover_file" gorm:"size:255" form:"-"`

	// User
	Email string `json:"email" gorm:"size:100;not null" form:"-"`
}

func NewBook() *Book {
	return &Book{}
}
