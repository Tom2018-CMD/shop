package models

import (
	pb "shop/proto/captcha"
)

//var store base64Captcha.Store

// 获取验证码
func MakeCaptcha(height int, width int, length int) (id, b64s, answer string, err error) {
	//var driver base64Captcha.Driver
	//
	//driveString := base64Captcha.DriverString{
	//	Height:          height,
	//	Width:           width,
	//	NoiseCount:      0,
	//	ShowLineOptions: 2 | 4,
	//	Length:          length,
	//	Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
	//	BgColor: &color.RGBA{
	//		R: 102,
	//		G: 102,
	//		B: 214,
	//		A: 125,
	//	},
	//	Fonts: []string{"wqy-microhei.ttc"},
	//}
	//driver = driveString.ConvertFonts()
	//
	////if redisEnable {
	////	store = redisStore
	////} else {
	////	store = defaultStore
	////}
	//c := base64Captcha.NewCaptcha(driver, store)
	//id, b64s, answer, err = c.Generate()
	//fmt.Println(reflect.TypeOf(store))

	captchaSrv := pb.NewCaptchaService("captcha", MicroCaptchaClient)
	res, err := captchaSrv.MakeCaptcha(ctx, &pb.MakeCaptchaRequest{
		Height: int32(height),
		Width:  int32(width),
		Length: int32(length),
	})
	if err != nil {
		return "", "", "", err
	}
	return res.Id, res.B64S, res.Answer, err
}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	//fmt.Println(1321, id, VerifyValue)
	//fmt.Println(store.Verify(id, VerifyValue, true))
	//if store.Verify(id, VerifyValue, true) {
	//	return true
	//} else {
	//	return false
	//}
	captchaSrv := pb.NewCaptchaService("captcha", MicroCaptchaClient)
	res, err := captchaSrv.VerifyCaptcha(ctx, &pb.VerifyCaptchaRequest{
		Id:          id,
		VerifyValue: VerifyValue,
	})

	if err != nil {
		return false
	}
	return res.Success
}
