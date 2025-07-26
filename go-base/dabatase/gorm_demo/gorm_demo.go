package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code     string
	Price    uint
	Birthday *time.Time // A pointer to time.Time, can be null

}

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:182598@yp@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//// 迁移 schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

	// Create
	now := time.Now()
	//user := User{Name: "Jinzhu", Age: 18, Birthday: &now}
	//result := db.Create(&user) // 通过数据的指针来创建
	//if result.Error != nil {
	//	panic("failed to create user")
	//}
	//fmt.Println("user id: ", user.ID, "row affected: ", result.RowsAffected)

	users := []*User{
		{Name: "jack", Age: 18, Birthday: &now},
		{Name: "Jackson", Age: 19, Birthday: &now},
	}

	result := db.Create(users) // pass a slice to insert multiple row
	if result.Error != nil {
		panic("failed to create users")
	}
	fmt.Println("row affected: ", result.RowsAffected)
}
