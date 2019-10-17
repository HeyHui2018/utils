package utils

import (
	"bytes"
	"errors"
	"github.com/ngaut/log"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
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
			// DisableKeepAlives为false时，在新建connection之前，会去缓存map中寻找同类型连接，若没有则新建并存入
			// 缓存map供重复使用，key为请求方法和请求host，此方法适用于短时间大量访问相同host的情况；当短时间大量访问不同的host时，需将DisableKeepAlives设为true，这样底层会在收到请求后直接关闭连接
			MaxIdleConnsPerHost: 0,
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
	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
