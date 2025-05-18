package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 假设这是我们的员工数据库 (一个 Employee 类型的切片)
var employees = []Employee{
	{ID: 1, Name: "Dilbert", Position: "Junior Programmer", ManagerID: 5},
	{ID: 5, Name: "Pointy-haired Boss", Position: "Pointy-haired boss", ManagerID: 0}, // 假设老板没有上级
	{ID: 2, Name: "Wally", Position: "Senior Colleague", ManagerID: 5},
	// ... 可以添加更多员工
}

// EmployeeByID 函数：根据员工 ID 查找员工
func EmployeeByID(id int) *Employee {
	for i := range employees { // 遍	历员工切片
		if employees[i].ID == id { // 如果找到 ID 匹配的员工
			return &employees[i] // 返回指向该员工的指针
		}
	}
	return nil // 如果遍历完所有员工都没有找到匹配的 ID，返回 nil (空指针)
}

func main() {
	// 测试 EmployeeByID 函数
	fmt.Println(EmployeeByID(1).Name)                               // 打印 Dilbert 的名字
	fmt.Println(EmployeeByID(5).Position)                           // 打印 Pointy-haired Boss 的职位
	fmt.Println(EmployeeByID(10)) //尝试一个不存在的id

	id := 1
	EmployeeByID(id).Salary = 0 // 解雇 Dilbert (把 Dilbert 的薪水设置为 0)
	fmt.Println(EmployeeByID(1).Salary)

}

