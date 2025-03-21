package jd

import (
	"net/url"
	"testing"
)

// 生成授权地址，浏览器访问后会重定向到redirect_uri，获取access_token
func Test_ToLoginGetAccessToken(t *testing.T) {
	// https://open-oauth.jd.com/oauth2/to_login?app_key=XXXXX&response_type=code&redirect_uri=XXXXX&state=20180416&scope=snsapi_base

	u, _ := url.Parse("https://open-oauth.jd.com/oauth2/to_login")
	q := u.Query()
	q.Add("app_key", "BB5039D6E7D11C7A481E092BF6350CB3")
	q.Add("response_type", "code")
	q.Add("redirect_uri", "https://api.vitodiagnosis.com/mini/jd/token")
	q.Add("state", "20250318")
	q.Add("scope", "snsapi_base")
	u.RawQuery = q.Encode()
	t.Log(u.String())
}
