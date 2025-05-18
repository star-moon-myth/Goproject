package main
import (
	"fmt"
	"os"
	"time"
	"text/template"
)
type LeaveRequest struct {
	TeacherName string
	ClassName   string
	StudentName string
	Reason      string
	Days        int
	StartDate   string
	EndDate     string
	Date        string
}
func main(){
	// 创建模板,解析文件
temp,err:=template.ParseFiles("ch4\\4-6\\playground\\request.txt")
if err!=nil{
	fmt.Println("模型解析失败",err)
	os.Exit(1)
}
// 准备数据
data := LeaveRequest{
	TeacherName: "王老师",
	ClassName:   "三年二班",
	StudentName: "小红",
	Reason:      "感冒发烧",
	Days:        3,
	StartDate:   "2024-05-16",
	EndDate:     "2024-05-18",
	Date:        time.Now().Format("2006-01-02"), // 获取当前日期，并格式化
}
// 执行模板
err =temp.Execute(os.Stdout,data)
if err != nil {
	fmt.Println("模板执行失败:", err)
	os.Exit(1)
}
}