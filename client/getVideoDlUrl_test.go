package client

import (
	"fmt"
	"testing"
)

func TestGetVideoDlUrl(t *testing.T) {
	url, _ := GetVideoDlUrl("6870423037087436046")
	fmt.Println(url)
}
