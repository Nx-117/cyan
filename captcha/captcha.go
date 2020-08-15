package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

/*
如有需要可以自行更改以下变量
*/

/**
验证码图片高度
*/
var Height = 60

/**
验证码图片长度
*/
var Width = 120

/**
验证码图片噪点
*/
var NoiseCount = 0

/**
验证码图片最小行数
*/
var ShowLineOptionsMin = 2

/**
验证码图片最大行数
*/
var ShowLineOptionsMax = 4

/**
验证码图片字符长度
*/
var Length = 4

/**
验证码字典
*/
var Source = "1234567890qwertyuioplkjhgfdsazxcvbnm"

/**
RGBA为颜色编号，可以自行设置也可以使用默认值
*/
var R uint8 = 254
var G uint8 = 254
var B uint8 = 254
var A uint8 = 254

/**
验证码字体
*/
var FontType = "chromohv.ttf"

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// 生成图形化验证码
func GenerateCaptcha() (*CaptchaResult, error) {

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          Height,
		Width:           Width,
		NoiseCount:      NoiseCount,
		ShowLineOptions: ShowLineOptionsMin | ShowLineOptionsMax,
		Length:          Length,
		Source:          Source,
		BgColor: &color.RGBA{
			R: R,
			G: G,
			B: B,
			A: A,
		},
		Fonts: []string{FontType},
	}
	captcha := base64Captcha.NewCaptcha(captchaConfig.ConvertFonts(), store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	captchaResult := CaptchaResult{
		Id:          id,
		Base64Blob:  b64s,
		VerifyValue: "200",
	}

	return &captchaResult, nil
}

/**
验证码校验
key 验证ID
verifyValue 验证码value
clear 是否删除
*/
func VerfiyCaptcha(key, verifyValue string, del bool) bool {
	return store.Verify(key, verifyValue, del)
}
