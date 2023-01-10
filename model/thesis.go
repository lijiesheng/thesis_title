package model

// 论文题目表
type Thesis struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Size    int    `db:"size" json:"size"`
	Type    string `db:"type" json:"type"`
	SizeInt int    `db:"size_int" json:"size_int"`
}
