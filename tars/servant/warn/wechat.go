package warn

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	MsgTypeText     = "text"
	MsgTypeMarkDown = "markdown"
	MsgTypeNews     = "news"
	MsgTypeImage    = "image"

	contentType = "Content-Type"
	jsonType    = "application/json;charset=UTF-8"
)

var wechatUrl = ""

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type PostSend struct {
	MsgType  string    `json:"msgtype"`
	MarkDown *MarkDown `json:"markdown"`
	News     *News     `json:"news"`
	Image    *Image    `json:"image"`
	Text     *Text     `json:"text"`
}

/*
{
    "msgtype": "text",
    "text": {
        "content": "广州今日天气：29度，大部分多云，降雨概率：60%",
        "mentioned_list":["wangqing","@all"],
        "mentioned_mobile_list":["13800001111","@all"]
    }
}
*/

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type MarkDown struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type News struct {
	Articles []*Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

func InitWechat(url string) error {
	wechatUrl = url
	return nil
}

func Send(url string, send *PostSend) error {
	bytesData, err := json.Marshal(send)
	body := bytes.NewReader(bytesData)
	if err != nil {
		return fmt.Errorf("json marshal: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("new request error: %v", err)
	}
	req.Header.Set(contentType, jsonType)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code %d", resp.StatusCode)
	}
	return nil
}

func SendMarkDown(url, content string) error {
	markdown := &MarkDown{
		Content: content,
	}
	send := &PostSend{
		MsgType:  MsgTypeMarkDown,
		MarkDown: markdown,
	}
	return Send(url, send)
}

func SendNews(url, title, desc, newsUrl, picUrl string) error {
	article := &Article{
		Title:       title,
		Description: desc,
		Url:         newsUrl,
		PicUrl:      picUrl,
	}
	articles := []*Article{article}
	news := &News{
		Articles: articles,
	}
	send := &PostSend{
		MsgType: MsgTypeNews,
		News:    news,
	}
	return Send(url, send)
}

func SendImage(url string, data []byte) error {
	content := base64.StdEncoding.EncodeToString(data)
	h := md5.New()
	h.Write(data)
	hash := hex.EncodeToString(h.Sum(nil))
	image := &Image{
		Base64: content,
		Md5:    hash,
	}
	send := &PostSend{
		MsgType: MsgTypeImage,
		Image:   image,
	}
	return Send(url, send)
}

var alarmChan = make(chan *alarm, 30)

type alarm struct {
	url      string
	postSend *PostSend
}

func (a *alarm) send() error {
	return Send(a.url, a.postSend)
}

func init() {
	go func() {
		for {
			select {
			case a := <-alarmChan:
				_ = a.send()
			}
		}
	}()
}

func ServerAlarm(content string) error {
	if wechatUrl == "" {
		return nil
	}
	url := wechatUrl
	return Alarm(url, content)
}

func appendAlarm(url string, send *PostSend) error {
	a := &alarm{
		url:      url,
		postSend: send,
	}
	select {
	case alarmChan <- a:
		return nil
	default:
		return errors.New("full")
	}
}

func Alarm(url, content string) error {
	markdown := &MarkDown{
		Content: content,
	}
	send := &PostSend{
		MsgType:  MsgTypeMarkDown,
		MarkDown: markdown,
	}
	return appendAlarm(url, send)
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return fmt.Sprintf("get ip fail: %v", err)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if strings.HasPrefix(ipnet.IP.String(), "192.168") || strings.HasPrefix(ipnet.IP.String(), "172.25") {
					return ipnet.IP.String()
				}
			}

		}
	}
	return "unknown ip"
}
