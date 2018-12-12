package utils

import (
	"github.com/ngaut/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"bytes"
	"time"
	"net"
)

func GetClient(timeOut int64) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeOut))
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeOut)))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeOut),
			DisableKeepAlives:     true,
			MaxIdleConnsPerHost:   0,
		},
	}
	return client
}

func Get(url string, timeOut int64) ([]byte, error) {
	c := GetClient(timeOut)
	resp, err := c.Get(url)
	if err != nil {
		return []byte{}, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func PostJson(urls string, info []byte) ([]byte, error) {
	infoBody := bytes.NewBuffer(info)
	resp, err := http.Post(urls, "application/json;charset=utf-8", infoBody)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func Post(urls string, args map[string]string, head map[string]string, timeOut int64) ([]byte, error) {
	c := GetClient(timeOut)
	Va := url.Values{}
	for k, v := range args {
		Va.Set(k, v)
	}
	request, err := http.NewRequest("POST", urls, strings.NewReader(Va.Encode()))
	if err != nil {
		return []byte{}, err
	}
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	for k, v := range head {
		request.Header.Add(k, v)
	}
	resp, err := c.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func PostWithRetry(processId, url string, param, head map[string]string) ([]byte, error) {
	var body []byte
	var err error
	i := 0
	for i = 0; i < 3; i++ {
		body, err = Post(url, param, head, 15)
		if err != nil {
			log.Warnf("PostWithRetry Post error,processId = %v,err = %v,i = %v", processId, err, i)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if i == 3 {
		log.Warnf("PostWithRetry retry error,processId = %v", processId)
		return body, errors.New("PostWithRetry retry end")
	}
	return body, nil
}
