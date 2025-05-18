// package memo
// type Memo struct {
// 	requests chan request // 请求 channel
// }

// type request struct {
// 	key      string
// 	response chan<- result // 响应 channel
// }

// func New(f Func) *Memo {
// 	memo := &Memo{requests: make(chan request)}
// 	go memo.server(f) // 启动一个 server goroutine 处理请求
// 	return memo
// }

// func (memo *Memo) Get(key string) (interface{}, error) {
// 	response := make(chan result) // 创建一个响应 channel
// 	memo.requests <- request{key, response} // 发送请求到 server goroutine
// 	res := <-response // 等待 server goroutine 返回结果
// 	return res.value, res.err
// }

// func (memo *Memo) Close() { close(memo.requests) } // 关闭请求 channel

// func (memo *Memo) server(f Func) {
// 	cache := make(map[string]*entry) //  server goroutine 内部维护一个 cache
// 	for req := range memo.requests { //  循环处理请求 channel 中的请求
// 		e := cache[req.key]
// 		if e == nil { //  缓存未命中
// 			// This is the first request for this key.
// 			e = &entry{ready: make(chan struct{})}
// 			cache[req.key] = e
// 			go e.call(f, req.key) //  启动一个 worker goroutine 异步计算结果
// 		}
// 		go e.deliver(req.response) //  启动一个 goroutine 异步发送结果给 client
// 	}
// }

// func (e *entry) call(f Func, key string) {
// 	// Evaluate the function.
// 	e.res.value, e.res.err = f(key) //  计算结果
// 	close(e.ready) //  关闭 ready channel，广播结果已准备好 📢
// }

// func (e *entry) deliver(response chan<- result) {
// 	// Wait for the ready condition.
// 	<-e.ready //  等待 ready channel 关闭 ⏳
// 	// Send the result to the client.
// 	response <- e.res //  将结果发送回 client (通过 response channel)
// }
package memo