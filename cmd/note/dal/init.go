package dal

import "github.com/zhiqinkuang/easy-note/cmd/note/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
