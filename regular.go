package utils

import (
	"regexp"
)

func HaveLetter(s string) (bool, error) {
	ok, err := regexp.MatchString(`.*[A-Za-z]+.*`, s)
	return ok, err
}

/*
数字："^[0-9]+$"
中文："^[\u4E00-\u9FA5]*$"
非空："^[\\s\\S]+$"
用户名："^[\u4E00-\u9FA5a-zA-Z0-9_.]{0,40}$"
十进制："^\\d+\\.[0-9]+$"
手机号："^1[0-9]{10}$"
座机号："^[0-9{8}$]"
*/