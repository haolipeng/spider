package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
	"strconv"
)

//正则表达式每次编译需要时间，这里我们进行预先编译
//短声明变量无法在全局声明
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)岁</div>`)
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)座</div>`)

var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)

var workLocationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

var xingzuoInfo = []string{
	"白羊座",
	"金牛座",
	"双子座",
	"巨蟹座",
	"狮子座",
	"处女座",
	"天秤座",
	"天蝎座",
	"射手座",
	"摩羯座",
	"水瓶座",
	"双鱼座",
}
var marriageInfo = []string{
	"已婚",
	"未婚",
	"离异",
}

//用户信息解析器
//id:用户的唯一标识
//url：用户的url
//name:用户昵称
func ParserProfile(contents []byte, id string, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	//解析顺序按照web css div顺序来
	//1.婚姻状况  未婚 离异 已婚
	marriage := extractString(contents, marriageRe)
	for _, v := range marriageInfo {
		if marriage == v {
			profile.Marriage = marriage
		}
	}

	//2.年龄
	strAge := extractString(contents, ageRe)
	age, err := strconv.Atoi(strAge)
	if err != nil {
		fmt.Println("parser profile age failed")
	}

	profile.Age = age

	//3.星座
	xingzuo := extractString(contents, xingzuoRe)
	for _, v := range xingzuoInfo {
		if xingzuo == v {
			profile.Xingzuo = xingzuo
			fmt.Println("xingzuo matched", profile.Marriage)
		}
	}

	//4.身高
	StrHeight := extractString(contents, heightRe)
	height, err := strconv.Atoi(StrHeight)
	if err != nil {
		fmt.Println("parser profile height failed")
	}
	profile.Height = height

	//5.体重
	StrWeight := extractString(contents, weightRe)
	weight, err := strconv.Atoi(StrWeight)
	if err != nil {
		fmt.Println("parser profile weight failed")
	}
	profile.Weight = weight

	//6.工作点
	workLocation := extractString(contents, workLocationRe)
	profile.Hukou = workLocation

	//7.收入
	profile.Income = extractString(contents, incomeRe)

	//8.解析结果
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				ID:      id,
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
