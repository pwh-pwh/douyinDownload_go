package client

import (
	"fmt"
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	err := DownloadVideo("https://aweme.snssdk.com/aweme/v1/playwm/?video_id=v0200f230000btcaac52m1gham4830p0&ratio=720p&line=0",
		"aaa.mp4")
	fmt.Println(err)
}
