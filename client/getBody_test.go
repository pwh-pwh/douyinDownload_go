package client

import (
	"fmt"
	"testing"
)

func TestGetBody(t *testing.T) {
	body, err, _ := GetBody("https://v.douyin.com/JyCk5gy/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
}
