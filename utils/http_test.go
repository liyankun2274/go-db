package utils

import (
	"fmt"
	"testing"
)

func Test_Get(t *testing.T) {
	h := NewHttpSend(GetUrlBuild("http://api.map.baidu.com/telematics/v3/weather", map[string]string{"location": "嘉兴", "output": "json", "ak": "5slgyqGDENN7Sy7pw29IUvrZ"}))
	rs, err := h.Get()
	if err != nil {
		fmt.Println("请求错误!")
	} else {
		fmt.Println(rs)
	}
}

//func Test_Post(t *testing.T) {
//	h := NewHttpSend("liyankun://127.0.0.1/test.php")
//	h.SetBody(map[string]string{"name": "xiaochuan"})
//	_, err := h.Post()
//	if err != nil {
//		t.Error("请求错误:", err)
//	} else {
//		t.Log("正常返回")
//	}
//}
//
//func Test_Json(t *testing.T) {
//	h := NewHttpSend("liyankun://127.0.0.1/test.php")
//	h.SetSendType("JSON")
//	h.SetBody(map[string]string{"name": "xiaochuan"})
//	_, err := h.Post()
//	if err != nil {
//		t.Error("请求错误:", err)
//	} else {
//		t.Log("正常返回")
//	}
//}
//
//func Benchmark_GET(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		h := NewHttpSend(GetUrlBuild("liyankun://127.0.0.1/test.php", map[string]string{"name": "xiaochuan"}))
//		_, err := h.Get()
//		if err != nil {
//			b.Error("请求错误:", err)
//		} else {
//			b.Log("正常返回")
//		}
//	}
//}
