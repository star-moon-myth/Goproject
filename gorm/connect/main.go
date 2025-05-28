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
// 用户档案
type Profile2 struct {
	gorm.Model
	User2ID  uint   // 外键，关联 User 的 ID
	Bio     string
	Website string
}

// 语言
type Language2 struct {
	gorm.Model
	Name string `gorm:"unique"`
	// 反向关联，可以不定义，但定义了方便查询
	// Users []User `gorm:"many2many:user_languages;"`
}

	// 订单
type Order2 struct {
	gorm.Model
	User2ID      uint    // 外键，关联 User 的 ID
	OrderNumber string `gorm:"type:varchar(191);uniqueIndex"`
	Amount      float64
	// User        User // Belongs To 关联，使用指针避免循环引用
}

// 用户
type User2 struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	Profile   Profile2        // Has One 关联
	Orders    []Order2        // Has Many 关联
	Languages []Language2 `gorm:"many2many:user_languages;"`
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

	// var activeusers []User1
	// result:=db.Where("active=? and age>?",false,20).Find(&activeusers)
	// fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	// var userIn []User1
	// result=db.Where("name in ?",[]string{"Alice","Bob"}).Find(&userIn)
	// fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	// var userlike []User1
	// result=db.Where("email like ?","%example.com").Find(&userlike)
	// fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	// type NameAndEmail struct{
	// 	Name string 
	// 	Email string
	// }
	// var results []NameAndEmail
	// result:=db.Model(&User1{}).Select("name,email").Where("active=?",false).Find(&results)
	// fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	// var mapResults []map[string]interface{}
	// result=db.Model(&User1{}).Select("name","age",).Limit(10).Find(&mapResults)
	// fmt.Printf("查询到%d个用户\n",result.RowsAffected)

	// var sortedUsers []User1
	// result=db.Order("age desc,name asc").Find(&sortedUsers)
	// if result.Error!=nil{
	// 	fmt.Printf("排序失败了%v",result.Error)
	// }else{
	// 	fmt.Printf("has found %d people\n",result.RowsAffected)
	// 	for _,u:=range sortedUsers{
	// 		fmt.Printf("name:%s,age:%d,email:%s\n",u.Name,u.Age,u.Email)
	// 	}
	// }

	// var count int64 
	// result:=db.Model(&User1{}).Where("active=?",false).Count(&count)
	// if result.Error==nil{
	// 	fmt.Printf("活跃用户数量：%d\n",count)
	// }
	
	// var names []string 
	// result=db.Model(&User1{}).Where("age>?",18).Pluck("name",&names)
	// if result.Error==nil{
	// 	fmt.Println("活跃用户名称:",names)
	// }


	// rows,err:=db.Model(&User1{}).Where("age>?",18).Rows()
	// if err!=nil{
	// 	log.Printf("获取失败：%v\n",err)
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next(){
	// 	var user User1
	// 	if err:=db.ScanRows(rows,&user);err!=nil{
	// 		log.Printf("扫描行数据失败：%v\n",err)
	// 		break
	// 	}
	// 	fmt.Printf("处理用户id:%d,Name:%s\n",user.ID,user.Name)
	// }
	// if err:=rows.Err();err!=nil{
	// 	log.Printf("迭代数据是发生失败：%v\n",err)
	// }

	// result:=db.Model(&User1{Model:gorm.Model{ID:1}}).Where("active=?",false).Update("name","Aliceupdated")
	// if result.Error!=nil{ 
	// 	log.Printf("更新失败：%v\n",result.Error)
	// }
	// fmt.Printf("更新成功，更新了%d行数据\n",result.RowsAffected)
	
	// var userToupdate User1
	// db.First(&userToupdate,1)
	// userToupdate.Name="Alice_newname"
	// userToupdate.Age=88
	// result:=db.Model(&userToupdate).Updates(userToupdate)

	// result=db.Model(&User1{}).Where("id=?",1).Updates(map[string]interface{}{
	// 	"name":"Alice_Mapupdate",
	// 	"age":0,
	// 	"active":false,
	// })
	// fmt.Println("🚀开始自动迁移 User2 相关模型--------")
	// err = db.AutoMigrate(
	// 	&User2{},    // 别忘了它！
	// 	&Profile2{}, // User2 的档案
	// 	&Order2{},   // User2 的订单
	// 	&Language2{}, // User2 的语言 (如果你也要用这个表)
	// 	// 如果还有其他 User2 相关的模型，也一并加上哦！
	// )
	// if err != nil {
	// 	log.Fatalf("User2 相关模型自动迁移失败: %v", err)
	// }
	// fmt.Println("User2 相关模型自动迁移完成🎉🎉")

	// user := User2{
	// 	Name:  "Ivy",
	// 	Email: "ivy@example.com",
	// 	Profile: Profile2{ // Has One
	// 		Bio:     "Loves coding",
	// 		Website: "ivy.dev",
	// 	},
	// 	Orders: []Order2{ // Has Many
	// 		{OrderNumber: "ORD001", Amount: 100.50},
	// 		{OrderNumber: "ORD002", Amount: 55.20},
	// 	},
	// }
	// db.Create(&user)

	// var ivy User2
	// db.Where("email=?","ivy@example.com").First(&ivy)
	// order:=Order2{
	// 	OrderNumber: "ord003",
	// 	Amount: 75.0,
	// 	User2ID: ivy.ID,
	// }
	// db.Create(&order)

	// var golang,javaLang Language2
	// db.FirstOrCreate(&golang,Language2{Name: "Go"})
	// db.FirstOrCreate(&javaLang,Language2{Name: "Java"})

	// user:=User2{
	// 	Name:"Jack",
	// 	Email:"Jack@example.com",
	// 	Languages: []Language2{golang,javaLang},
	// }
	// db.Create(&user)
	
	// var userWithData []User2
	// result:=db.Preload("Orders").Preload("Profile").Find(&userWithData)
	// if result.Error==nil{
	// 	for _,u:=range userWithData{
	// 		fmt.Println("User:",u.Name)
	// 		if u.Profile.ID>0{
	// 			fmt.Printf("Profile Bio:%s\n",u.Profile.Bio)
	// 		}
	// 		if len(u.Orders)>0{
	// 			fmt.Println("orders:")
	// 			for _,o:=range u.Orders{
	// 				fmt.Printf("OrderNumber:%s,Amount:%f\n",o.OrderNumber,o.Amount)
	// 			}
	// 		}
	// 	}
	// }
	// var orderWithUser Order2
	// db.Preload("User").First(&orderWithUser,1)
	// if orderWithUser.User.ID>0{
	// 	fmt.Printf("OrderNumber:%s,Amount:%f,UserName:%s\n",orderWithUser.OrderNumber,orderWithUser.Amount,orderWithUser.User.Name)
	// }

	// var userwithlangs User2
	// db.Preload("Languages").First(&userwithlangs,1)
	// if len(userwithlangs.Languages)>0{
	// 	fmt.Println("User Languages:",userwithlangs.Name)
	// 	for _,lang :=range userwithlangs.Languages{
	// 		fmt.Printf("%s",lang.Name)
	// 	}
	// 	fmt.Println()
	// }

	var user User2
	db.Preload("Languages").First(&user,1)
	// fmt.Println(result.RowsAffected)
	// for _,u :=range user.Languages{
	// 	fmt.Println("语言名称：",u.Name)
	// }
	var phplang Language2
	db.FirstOrCreate(&phplang,Language2{Name: "PHP"})
	err=db.Model(&user).Association("Languages").Append(&phplang)
	if err!=nil{
		fmt.Println("success!")

	}

	

}