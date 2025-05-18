package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接字符串
	dsn := "root:521334@tcp(127.0.0.1:3306)/crashcourse?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	fmt.Println("Successfully connected to MySQL!")
	create(db)
	// insert(db)
	// selectrow(db)
	// selectrows(db)
	// update(db)
	// delete(db)
	transaction(db)
}
func create(db *sql.DB){
	_,err:=db.Exec(`
	create table if not exists users(
	id int auto_increment primary key,
	name varchar(255) not null,
	created_at timestamp default current_timestamp
	)
	`)
	if err!=nil{
		fmt.Println("failed to create table:",err)
		return 
	}
	fmt.Println("Table crrated successfully!")
}

// func insert(db *sql.DB){
// 	name:="Alice"
// 	result,err:=db.Exec("insert into users (name) values(?)",name)
// 	if err!=nil{
// 		fmt.Println("Failed to insert into table:",err)
// 		return
// 	}
// 	id,_:=result.LastInsertId()
// 	fmt.Println("Inserted user with id:",id)
// }

// func selectrow(db *sql.DB){
// 	var id int 
// 	var name string 
// 	var created_at time.Time
// 	err:=db.QueryRow("select id,name,created_at from users where id=?",1).Scan(&id,&name,&created_at)
// 	if err!=nil{
// 		fmt.Println("Failed to query from users:",err)
// 	}
// 	fmt.Printf("-----------\n🔍SelectOneUser:\n🆔id %d,🏠name %s,⌛️create_at %s\n-------------\n",id,name,created_at)
// }

// func selectrows(db *sql.DB){
// 	rows,err:=db.Query("select id,name,created_at from users where name=?","Alice")
// 	if err!=nil{
// 		fmt.Println("failed to quert from users:",err)
// 		return
// 	}
// 	defer rows.Close()
// 	fmt.Println("\n------------")
// 	for rows.Next(){
// 		var id int
// 		var name string
// 		var created_at string 
// 		err:=rows.Scan(&id,&name,&created_at)
// 		if err!=nil{
// 			fmt.Println("failed to scan row:",err)
// 			return 
// 		}
// 		fmt.Printf("🔍SelectMoreUser:\n🆔id %d,🏠name %s,⌛️create_at %s\n",id,name,created_at)
// 	}
// 	fmt.Println("------------")
// 	if err=rows.Err();err!=nil{
// 		fmt.Println("error during iteration:",err)
// 		return
// 	}
// }
// func update(db *sql.DB){
// 	result,err:=db.Exec("update users set name=? where id =?","Bob",1)
// 	if err!=nil{
// 		fmt.Println("failed to update data:",err)
// 		return 
// 	}
// 	rowsaffected,_:=result.RowsAffected()
// 	fmt.Printf("update %d rows\n",rowsaffected)
// }

// func delete(db *sql.DB){
// 	result,err:=db.Exec("delete from users where id =?",6)
// 	if err!=nil{
// 		fmt.Println("failed to delete data:",err)
// 		return 
// 	}
// 	rowsaffected,_:=result.RowsAffected()
// 	fmt.Printf("delete %d rows\n",rowsaffected)
// }

func transaction(db *sql.DB){
	tx,err:=db.Begin()
	if err!=nil{
		fmt.Println("failed to begin transaction:",err)
		return
	}
	defer tx.Rollback()

	_,err=tx.Exec("insert into users (name) values(?)","Nina")
	if err!=nil{
		fmt.Println("failed to insert data:",err)
		return 
	}

	_,err=tx.Exec("update users set name=? where id =?","Mike",2)
	if err!=nil{
		fmt.Println("failed to update data:",err)
		return
	}

	err=tx.Commit()
	if err!=nil{
		fmt.Println("failed to commit transaction:",err)
		return
	}
	fmt.Println("Transaction completed successfully!🎉")
}
