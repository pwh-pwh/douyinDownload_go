package client

import (
	"fmt"
	"testing"
)

func TestGetVideoId(t *testing.T) {
	id, err := GetVideoId("https://v.douyin.com/JyCk5gy/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id, "--")
	}

}
