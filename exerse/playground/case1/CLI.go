package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type task struct {
	ID   int    `json:"id"`
	DESC string `json:"desc"`
	DONE bool   `json:"done"`
}

var tasks []task
var datafile = "task.json"

func initialdata() {
	data, err := os.ReadFile(datafile)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}
}
func savedata() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(datafile, data, 0644)
	fmt.Println("恭喜你🎉，保存成功了哦😄")
}
func addata(data string) {
	I := 1
	if len(tasks)>0{
		I=tasks[len(tasks)-1].ID+1
	}
	tasks = append(tasks, task{ID: I, DESC: data, DONE: false})
	fmt.Println("数据添加成功了")
}
func complete(id int) {
	if id < 0 || id > len(tasks) {
		fmt.Println("id不存在哦")
		return
	}
	tasks[id].DONE = true
	fmt.Println("标记为已完成")
	savedata()
}
func list() {
	if len(tasks) == 0 {
		fmt.Println("没有任务哦")
	}
	for _, data := range tasks {
		status:="⌛️"
		if data.DONE{
			status="✅"
		}
		fmt.Printf("  %s [%d] %s\n", status, data.ID, data.DESC)
	}
}
func main() {
	addcmd := flag.String("add", "", "添加这条数据")
	listcmd := flag.Bool("list", false, "列出数据")
	comcmd := flag.Int("complete", -1, "数据完成了")
	flag.Parse()
	initialdata()
	if *addcmd != "" {
		addata(*addcmd)
		savedata()
	} else if *listcmd {
		list()
	} else if *comcmd >= 0 {
		complete(*comcmd)
	} else {
		fmt.Println("输入-help提供帮助")
		list()
	}

}
