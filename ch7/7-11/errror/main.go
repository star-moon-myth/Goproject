package main

import (
	"errors" // 引入 errors 包，用来创建自定义错误
	"fmt"   // 引入 fmt 包，用来输出
	"os"    // 引入 os 包，文件操作相关的都在这
	"syscall" // 引入 syscall 包，系统调用相关的错误码
)

// 定义一个自定义的 “文件不存在” 错误，和官方的 ErrNotExist 区分开
var ErrNotExistCustom = errors.New("file not found by custom")

func main() {
	filePath := "/no/such/file" //  故意写一个不存在的文件路径

	// 1. 尝试打开文件，肯定会报错 💥
	_, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件出错了:", err) // 打印原始错误信息

		// 2. 使用 os.IsNotExist 判断是否是 “文件不存在” 错误 🤔
		if os.IsNotExist(err) {
			fmt.Println("👉 错误类型判断: 文件不存在!")
		}

		// 3. 类型断言，获取 PathError 类型的详细信息 🕵️‍♀️
		pathErr, ok := err.(*os.PathError) // 类型断言 *os.PathError
		if ok {
			fmt.Println("👉 类型断言 PathError:")
			fmt.Println("   - 操作类型:", pathErr.Op)   //  比如 "open"
			fmt.Println("   - 出错路径:", pathErr.Path) // 就是我们定义的 filePath
			fmt.Println("   - 底层错误:", pathErr.Err)  //  更底层的错误，比如 syscall.ENOENT
		}

		// 4.  更底层的错误码 syscall.ENOENT 判断 (文件不存在的系统错误码) 🔍
		if errors.Is(err, syscall.ENOENT) { // 使用 errors.Is 判断，更通用，推荐！
			fmt.Println("👉 errors.Is(syscall.ENOENT): 底层错误码是 syscall.ENOENT, 也是文件不存在!")
		}

		// 5. 自定义的 ErrNotExistCustom 错误判断 (假设你的代码里有自定义的文件不存在错误)  💡
		if errors.Is(err, ErrNotExistCustom) { //  errors.Is 可以判断自定义错误
			fmt.Println("👉 errors.Is(ErrNotExistCustom): 这是我们自定义的文件不存在错误!")
		} else {
			fmt.Println("👉 errors.Is(ErrNotExistCustom): 不是我们自定义的文件不存在错误!")
		}

	} else {
		fmt.Println("文件打开成功了！🎉  (但这个例子里，这行代码不会执行)")
	}
}

