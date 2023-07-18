package oneDayOneLibray

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ContextTimeOutTest  http://127.0.0.1:8081/
func ContextTimeOutTest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建一个带有超时的 context
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		// 在 goroutine 中执行耗时操作
		go func() {
			// 模拟一个耗时操作
			time.Sleep(100 * time.Second)

			// 检查 context 是否已经超时
			if ctx.Err() != nil {
				fmt.Println("context timeout")
				return
			}

			// 输出处理结果
			fmt.Fprintln(w, "hello world")
		}()

		// 等待处理结果或超时
		select {
		case <-ctx.Done():
			fmt.Println("request canceled or timed out")
		}
	})

	http.ListenAndServe(":8081", nil)
}
