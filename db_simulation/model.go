package main

type tigerman struct {
	ID       uint `gorm:"primaryKey"`
	Building string
}

// Customer 表结构
type Customer struct {
	Email   string `gorm:"primaryKey;size:100"` // 主键
	Name    string `gorm:"size:100"`
	Street  string `gorm:"size:100"`
	ZipCode string `gorm:"size:10"`

	Phones []CustomerPhone `gorm:"foreignKey:CustomerEmail;references:Email"` // 通过外键定义与 CustomerPhone 的一对多关系
}

// CustomerPhone 表结构
type CustomerPhone struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"` // 主键
	CustomerEmail string `gorm:"size:100"`                 // 外键字段，关联 Customer 表的 Email 字段
	Phone         string `gorm:"size:20"`

	// 通过 foreignKey 关联到 Customer 表的 Email 字段
	Customer Customer `gorm:"foreignKey:CustomerEmail;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// JOIN operation
type Employee struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	DepartmentID uint
}

type Department struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
