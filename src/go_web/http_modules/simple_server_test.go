package http_modules

import (
	"fmt"
	"net/http"
	"testing"
)

func welcome(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(res, "Learn Go website development")
}
func Test_simpleServer(t *testing.T){
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8090", nil)
	fmt.Println("服务已启动")
}