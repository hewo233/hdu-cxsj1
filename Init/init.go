package Init

import "github.com/hewo233/hdu-cxsj1/db"

func Init() {
	db.ConnectDB()
	db.UpdateDB()
}
