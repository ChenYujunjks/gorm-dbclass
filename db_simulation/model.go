package main

type Author struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"size:255;not null"`
	Address string
	URL     string
	Books   []Book `gorm:"foreignKey:AuthorID"` // 一对多关系
}

type Book struct {
	ISBN     string `gorm:"primaryKey;size:13"`
	Title    string `gorm:"size:255;not null"`
	Year     int
	Price    float64
	AuthorID uint // 外键，关联到 `author` 表的 `id`
}
