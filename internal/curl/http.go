package curl

import (
	"io"
	"log"
	"net/http"
	"os"
)

//get方法
func Get(url string) error{
	resp,err := http.Get(url)

	if err != nil {
		log.Fatalf("get 请求错误: %v",err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout,resp.Body)

	return nil
}

func Post(url string) error  {
	return nil
}
