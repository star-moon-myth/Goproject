package main

import (
	_"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
type Product1 struct{
	gorm.Model
	Code string `gorm:"unique"`
	Price uint 
}

type User1 struct {
	gorm.Model // 内嵌 gorm.Model 结构体

	Name     string `gorm:"comment:用户名"` // 使用 tag 添加数据库注释
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"` // 设置类型、唯一索引、非空
	Age      uint8  `gorm:"default:18"`   // 设置默认值
	Birthday *time.Time                 // 使用指针类型，允许 NULL 值
	Active   bool                       // bool 类型
	// 如果字段名和数据库列名不一致，可以使用 `column` tag
	// Password string `gorm:"column:user_password"`
}

type Order1 struct {
	gorm.Model
	UserID    uint // 外键，对应 User 模型
	ProductID uint // 外键，对应 Product 模型
	Amount    float64
	// ... 其他字段
}

func main() {
	dsn := "root:521334@tcp(127.0.0.1:3306)/crashcourse?charset=utf8mb4&parseTime=True&loc=Local"
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db,err:=gorm.Open(mysql.Open(dsn),gormConfig)
	if err!=nil{
		log.Fatalf("数据库连接失败：%v",err)
	}

	sqlDB,err:=db.DB()
	if err!=nil{
		log.Fatalf("获取*sql.db失败:%v",err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("连接配置完成😄")
	// fmt.Println("🚀开始自动迁移--------")
	// err=db.AutoMigrate(&User1{},&Order1{},&Product1{})
    // if err!=nil{
	// 	log.Fatalf("自动迁移失败")
	// }
	// fmt.Println("自动迁移完成🎉🎉")
	// user:=User1{Name: "Alice",Email: "alice@qq.com",Age: 26}
	// result:=db.Create(&user)
	// if result.Error!=nil{
	// 	log.Printf("error for creating:%v\n",result.Error)
	// }else{
	// 	fmt.Printf("\nsuccess for creating!---id:%d,---rowsaffected:%d\n",user.ID,result.RowsAffected)
	// }

	// users:=[]User1{
	// 	{Name: "Bob",Email: "bob@qq.com",Age: 22},
	// 	{Name: "Charlie", Email: "charlie@example.com"},
	// 	{Name: "David", Email: "david@example.com"},
	// }
	// result:=db.Create(&users)
	// if result.Error!=nil{
	// 	log.Printf("error for more creating:%v\n",result.Error)
	// }else{
	// 	fmt.Printf("success for creating %d people\n",result.RowsAffected)
	// 	for _,u:=range users{
	// 		fmt.Printf("user:id=%d,name=%s,email=%s,age=%d\n",u.ID,u.Name,u.Email,u.Age)
	// 	}
	// }
	// var firstuser User1
	// result:=db.First(&firstuser)
	// if result.Error!=nil{
	// 	if errors.Is(result.Error,gorm.ErrRecordNotFound){
	// 		fmt.Println("未找到用户记录")
	// 	}else{
	// 		log.Printf("查询第一条用户失败:%v",result.Error)
	// 	}
	// }else{
	// 	fmt.Printf("查到第一条用户,用户id:%d,用户名:%s\n",firstuser.ID,firstuser.Name)
	// }

	// var userbyid User1
	// db.First(&userbyid,4)
	// fmt.Printf("查到用户id:4,用户名:%s\n",userbyid.Name)

	// var activeuser User1
	// db.Where("name=? and active=?","David",false).First(&activeuser)
	// fmt.Printf("查到用户,用户id:%d,用户名:%s,邮箱:%s\n",activeuser.ID,activeuser.Name,activeuser.Email)
	// var allusers []User1
	// result:=db.Find(&allusers)
	// if result.Error!=nil{
	// 	log.Printf("查询所有用户失败%v\n",result.Error)		
	// }else{
	// 	fmt.Printf("查询到%d个用户:\n",result.RowsAffected)
	// 	for _,u:=range allusers{
	// 		fmt.Printf("ID=%d,Name=%s,Email=%s\n",u.ID,u.Name,u.Email)
	// 	}
	// }

	var activeusers []User1
	result:=db.Where("active=? and age>?",false,20).Find(&activeusers)
	fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	var userIn []User1
	result=db.Where("name in ?",[]string{"Alice","Bob"}).Find(&userIn)
	fmt.Printf("查询到%d个用户\n",result.RowsAffected)


}