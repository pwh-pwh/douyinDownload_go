package client

import (
	"crypto/tls"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	pattern = "video/(\\d+)?"
	cVUrl   = "https://m.douyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s&a_bogus="
)

var esc_re, _ = regexp.Compile(pattern)

func doGet(url string) (*http.Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal("err", err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/14.2 Chrome/87.0.4280.141 Mobile Safari/537.36")
	return client.Do(request)
}

func GetVideoId(url string) (s string, err error) {
	resp, err := doGet(url)
	if err != nil {
		fmt.Println("get error", err)
		return
	}
	defer resp.Body.Close()
	s = resp.Request.URL.String()
	s = esc_re.FindStringSubmatch(s)[1]
	return
}

func GetVideoDlUrl(id string) (s string, err error) {
	ts := fmt.Sprintf(cVUrl, id)
	//fmt.Println(ts)
	resp, err := doGet(ts)
	if err != nil {
		fmt.Println("get error", err)
		return
	}
	defer resp.Body.Close()
	var bytes = make([]byte, 128)
	var builder strings.Builder
	for {
		rLen, err := resp.Body.Read(bytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		builder.Write(bytes[:rLen])
	}

	if err != nil {
		return
	}
	//fmt.Println("===\n")
	//fmt.Println(builder.String())
	s = gjson.Get(builder.String(), "item_list.0.video.play_addr.uri").String()
	return
}

func DownloadVideo(url, fileName string) error {
	resp, err := doGet(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return err
}

func GetBody(url string) (io.ReadCloser, error, int64) {
	id, err := GetVideoId(url)
	if err != nil {
		return nil, err, 0
	}
	dlUrl, err := GetVideoDlUrl(id)
	if err != nil {
		return nil, err, 0
	}
	resultUrl := "https://www.douyin.com/aweme/v1/play/?video_id=" + dlUrl
	fmt.Println("resultUrl:", resultUrl)
	resp, err := doGet(resultUrl)
	if err != nil {
		return nil, err, 0
	}
	return resp.Body, nil, resp.ContentLength
}
