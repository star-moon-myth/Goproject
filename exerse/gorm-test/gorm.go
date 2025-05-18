package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	Create_At time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
}
func main() {
	dsn := "root:521334@tcp(127.0.0.1:3306)/crashcourse?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("failed to connect to database:",err)
		return 
	}
	db.AutoMigrate(&User{})
	var users []User
	db.Find(&users)
	fmt.Println(users)
	fmt.Println("----------------------------")
	for _,u:=range users{
		fmt.Printf("user:id=%d,name=%s,CreateAt=%s\n",u.ID,u.Name,u.Create_At)
	}
	db.Model(&User{}).Where("id=?",1).Update("name","Frank")
	db.Delete(&User{},1)
}