package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"ioutil"
)

// transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

/*
	if you want to compare two Chinese characters,just rewrite the less function.
	for example:
	func (t list) Less(i, j int) bool {
		a, _ := utils.UTF82GBK(t[i].Title)
		b, _ := utils.UTF82GBK(t[j].Title)
		for idx, chr := range a {
			if idx > len(b)-1 {
				return false
			}
			if chr != b[idx] {
				return chr < b[idx]
			}
		}
		return true
	}
 */
