package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
    "sync" // 需要 map 的并发安全
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items = make(map[int]Item) // 简易内存数据库
	mu    sync.Mutex        // 保护 items map
	nextID = 1               // 简单 ID 生成器
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // 在处理请求前锁定，结束后解锁
	defer mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id") // 尝试获取 ID query 参数
		if idStr == "" {                // 如果没有 ID，则列出所有
			var itemList []Item
			for _, item := range items {
				itemList = append(itemList, item)
			}
			json.NewEncoder(w).Encode(itemList)
		} else { // 如果有 ID，则查找特定 item
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID format", http.StatusBadRequest)
				return
			}
			item, ok := items[id]
			if !ok {
				http.NotFound(w, r) // 找不到则 404
				return
			}
			json.NewEncoder(w).Encode(item) // 找到了，返回 JSON
		}

	case http.MethodPost:
		var newItem Item
        // 解析请求体 JSON
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}
        if newItem.Name == "" { // 简单校验
            http.Error(w, "Item name cannot be empty", http.StatusBadRequest)
			return
        }

		newItem.ID = nextID
		items[newItem.ID] = newItem
		nextID++
		w.WriteHeader(http.StatusCreated) // 设置状态码 201 Created
		json.NewEncoder(w).Encode(newItem) // 返回创建的 item
		log.Printf("✅ Item created: %+v\n", newItem)

	// case http.MethodPut: // 实现更新逻辑...
	// case http.MethodDelete: // 实现删除逻辑...

	default:
		// 不支持的方法
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	log.Println("🚀 Server starting on :8080, supporting GET/POST on /items...")
	log.Fatal(http.ListenAndServe(":8080", nil))
    // 测试:
    // GET /items -> 列出所有
    // GET /items?id=1 -> 获取 ID 为 1 的项目
    // POST /items (Body: {"name": "New Gadget"}) -> 创建新项目
}
