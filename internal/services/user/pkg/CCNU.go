package CCNU

import (
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
	errno "user/pkg/errno"
)

const loginUrl = "https://account.ccnu.edu.cn/cas/login"
const avatarUrl = "http://yichen.api.z7zz.cn/api/sjtx2.php"
const defaultAvatat = ""

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetAvatar() string {
	resp, err := http.Get(avatarUrl)
	if err != nil {
		return defaultAvatat
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultAvatat
	}
	
	var m = make(map[string]any)
	if json.Unmarshal(body, &m) != nil {
		return defaultAvatat
	}
	avatar := m["text"].(string)
	return avatar
}

func Login(sid, psd string) error {
	return request(sid, psd)
}

func request(sid, psd string) error {
	data, err := prepare()
	if err != nil {
		return err
	}
	parseUrl, err := url.Parse(loginUrl)
	if err != nil {
		return err
	}
	
	jar, _ := cookiejar.New(&cookiejar.Options{})
	client := &http.Client{
		Timeout: 2 * time.Second,
		Jar:     jar,
	}
	
	targetUrl := "https://" + parseUrl.Host + data.Get("action")
	data.Del("action")
	data.Set("username", sid)
	data.Set("password", psd)
	
	req, err := http.NewRequest(http.MethodPost, targetUrl, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(info))
	
	if yes, err := regexp.Match("success", info); err != nil {
		return err
	} else if !yes {
		return errno.LoginWrongInfoError
	}
	return nil
}

func prepare() (url.Values, error) {
	resp, err := http.Get(loginUrl)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	
	data := url.Values{}
	data.Add("action", doc.Find("#fm1").AttrOr("action", ""))
	doc.Find("#fm1 > section.row.btn-row").Children().Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("type", "") == "hidden" {
			data.Set(s.AttrOr("name", ""), s.AttrOr("value", ""))
		}
	})
	return data, nil
}
