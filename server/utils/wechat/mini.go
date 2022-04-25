package wechat

const (
	APP_ID            = "wxb9602186e96b60af"
	APP_SECRET        = "4c131d822552c3bd86906cc6a0834f30"
	MINI_AUTH_URL     = "https://api.weixin.qq.com/sns/jscode2session?appid=wxb9602186e96b60af"
	MINI_ACCESS_TOKEN = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxb9602186e96b60af"
	MINI_QRCODE       = "https://api.weixin.qq.com/wxa/getwxacode?access_token=wxb9602186e96b60af"
)

var mini struct{}

func (*mini) GetAccessToken() string {

}
