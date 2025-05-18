package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	dsn := "root:521334@tcp(127.0.0.1:3306)/crashcourse?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql",dsn)
	if err!=nil{
		log.Fatalf("failed to open mysql:%v",err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute*3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err=db.Ping()
	if err!=nil{
		log.Fatalf("failed to connect to mysql:%v",err)
	}
	fmt.Println("A success for connecting mysql🎉")
	// id,err:=insertuser(db,"Nina","888888@qq.com")
	// if err!=nil{
	// 	fmt.Println("failed")
	// }
	// fmt.Println(id,err)

	// _,err=queryuserbyid(db,1)
	// if err!=nil{
	// 	fmt.Println(err)
	// }

	// allusers,err:=quertyAllusers(db)
	// if err!=nil{
	// 	fmt.Println("sorry:",err)
	// }else{
	// 	for _,user:=range allusers{
	// 		fmt.Println(user)
	// 	}
	// }
	// updateusers(db,1,"88888@gmail.com")
	usersToInsert := []User{
			{Name: "Bob", Email:sql.NullString{String:"bob@example.com", Valid: true}},
			{Name: "Charlie", Email:sql.NullString{String:"charlie@example.com", Valid: true}},
		}
	err=batchInsertUsers(db,usersToInsert)
	if err!=nil{
		fmt.Println("❌",err)
	}
}

// func insertuser(db *sql.DB,name string,email string)(int64,error){
// 	sqlstr:="insert into users2 (name,email) values (?,?)"
// 	result,err:=db.Exec(sqlstr,name,email)
// 	if err!=nil{
// 		fmt.Println("failed to insert into users2:",err)
// 		return 0,err
// 	}
// 	thelastid,err:=result.LastInsertId()
// 	if err!=nil{
// 		fmt.Println("failed to get the last id:",err)
// 		return 0,err
// 	}
// 	fmt.Println("succeed to insert into users2🎉")
// 	return thelastid,nil
// }


// func queryuserbyid(db *sql.DB,id int64)(User,error){
// 	querystr:="select id,name,email,create_at from users2 where id=?"
// 	var u User
// 	row:=db.QueryRow(querystr,id)
// 	err:=row.Scan(&u.ID,&u.Name,&u.Email,&u.CreatedAt)
// 	if err!=nil{
// 		if err==sql.ErrNoRows{
// 			fmt.Println("cant find user with id ",id)
// 			return User{},fmt.Errorf("cant find id %d",id)
// 		}
// 		fmt.Println("scan err",err)
// 		return User{},err
// 	}
// 	if u.Email.Valid{
// 		fmt.Printf("查询成功！用户: ID=%d, Name=%s, Email=%s, CreatedAt=%s\n", u.ID, u.Name, u.Email.String, u.CreatedAt.Format("2006-01-02 15:04:05"))
// 	}else{
// 		fmt.Printf("查询成功！用户: ID=%d, Name=%s, Email=NULL, CreatedAt=%s\n", u.ID, u.Name, u.CreatedAt.Format("2006-01-02 15:04:05"))
// 	}
// 	return u,nil 
// }

// func quertyAllusers(db *sql.DB)([]User,error){
// 	sql:="select id,name,email,create_at from users2 where id>?"
// 	var users []User
// 	rows,err:=db.Query(sql,0)
// 	if err!=nil{
// 		fmt.Println("failed to select something",err)
// 		return nil,err
// 	}
// 	defer rows.Close()
// 	for rows.Next(){
// 		var u User
// 		err:=rows.Scan(&u.ID,&u.Name,&u.Email,&u.CreatedAt)
// 		if err!=nil{
// 			fmt.Println("failed to select one row for the table",err)
// 			return users,err
// 		}
// 		users=append(users, u)
// 	}
// 	err=rows.Err()
// 	if err!=nil{
// 		fmt.Println("have err when scanning",err)
// 		return users,err
// 	}
// 	fmt.Printf("select %d users\n",len(users))
// 	return users,nil
// }

// func updateusers(db *sql.DB,id int64,newEmail string)(int64,error){
// 	sqlstr:="update users2 set email=? where id=?"
// 	result,err:=db.Exec(sqlstr,newEmail,id)
// 	if err!=nil{
// 		fmt.Println("failed to update:",err)
// 		return 0,err
// 	}
// 	rowsAff,err:=result.RowsAffected()
// 	if err!=nil{
// 		fmt.Println("failed to get rows affected:",err)
// 		return 0,err
// 	}
// 	fmt.Printf("succeed to update %d lines:",rowsAff)
// 	return rowsAff,nil
// }
type User struct{
	ID int64
	Name string
	Email sql.NullString
	CreatedAt time.Time
}
func batchInsertUsers(db *sql.DB,users []User)error{
	querystr:="insert into users2 (name,email) values(?,?)"
	stmt,err:=db.Prepare(querystr)
	if err!=nil{
		fmt.Println("预处理错误❌：",err)
		return err
	}
	defer stmt.Close()
	for _,u:=range users {
		_,err:=stmt.Exec(u.Name,u.Email.String)
		if err!=nil{
			fmt.Println("执行错误：",err)
			return err
		}		
	}
	fmt.Println("Success!🎉")
	return nil
}