package jd

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

type Client struct {
	ServerUrl      string
	AccessToken    string
	ConnectTimeout int
	ReadTimeout    int
	AppKey         string
	AppSecret      string
	Version        string
	Format         string
	charsetUtf8    string
	jsonParamKey   string
}

func NewClient(appKey, appSecret string) *Client {
	return &Client{
		ServerUrl:    "https://api.jd.com/routerjson",
		AppKey:       appKey,
		AppSecret:    appSecret,
		Version:      "2.0",
		Format:       "json",
		charsetUtf8:  "UTF-8",
		jsonParamKey: "360buy_param_json",
	}
}

func (c *Client) generateSign(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	stringToBeSigned := c.AppSecret
	for _, k := range keys {
		v := params[k]
		if k == c.jsonParamKey {
			jsonBytes, _ := json.Marshal(v)
			stringToBeSigned += k + string(jsonBytes)
		} else {
			stringToBeSigned += k + fmt.Sprintf("%v", v)
		}
	}
	stringToBeSigned += c.AppSecret

	hash := md5.Sum([]byte(stringToBeSigned))
	return strings.ToUpper(fmt.Sprintf("%x", hash))
}

func (c *Client) curl(serverUrl string, sysParams map[string]interface{}) (string, error) {
	apiParams, _ := json.Marshal(sysParams[c.jsonParamKey])
	delete(sysParams, c.jsonParamKey)

	requestUrl := serverUrl + "?"
	for k, v := range sysParams {
		requestUrl += k + "=" + url.QueryEscape(fmt.Sprintf("%v", v)) + "&"
	}
	requestUrl = strings.TrimSuffix(requestUrl, "&")

	client := &http.Client{
		Timeout: time.Duration(c.ReadTimeout) * time.Second,
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(c.jsonParamKey+"="+url.QueryEscape(string(apiParams)))))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) execute(request Request, accessToken string) (interface{}, error) {
	sysParams := map[string]interface{}{
		"app_key":   c.AppKey,
		"v":         c.Version,
		"method":    request.GetApiMethodName(),
		"timestamp": c.getCurrentTimeFormatted(),
	}

	if accessToken != "" {
		sysParams["access_token"] = accessToken
	}

	sysParams[c.jsonParamKey] = request.GetApiParas()
	sysParams["sign"] = c.generateSign(sysParams)

	resp, err := c.curl(c.ServerUrl, sysParams)
	if err != nil {
		return map[string]interface{}{
			"code": err.Error(),
			"msg":  err.Error(),
		}, nil
	}

	var respObject interface{}
	if c.Format == "json" {
		err = json.Unmarshal([]byte(resp), &respObject)
		if err != nil {
			return map[string]interface{}{
				"code": 0,
				"msg":  "HTTP_RESPONSE_NOT_WELL_FORMED",
			}, nil
		}
	} else if c.Format == "xml" {
		// XML parsing can be added here if needed
		return map[string]interface{}{
			"code": 0,
			"msg":  "HTTP_RESPONSE_NOT_WELL_FORMED",
		}, nil
	}

	return respObject, nil
}

func (c *Client) Exec(paramsArray map[string]interface{}) (interface{}, error) {
	if _, ok := paramsArray["method"]; !ok {
		return nil, errors.New("no api name passed")
	}

	method := paramsArray["method"].(string)
	requestClassName := UnderscoreToUpperCamelCase(method[9:]) + "Request"

	req := NewRequest(requestClassName)
	reqValue := reflect.ValueOf(req)

	for paraKey, paraValue := range paramsArray {
		setterMethodName := "Set" + cases.Title(language.Und, cases.NoLower).String(strings.Replace(paraKey, "_", "", -1))

		mtd := reqValue.MethodByName(setterMethodName)
		if !mtd.IsValid() {
			continue
		}
		mtd.Call([]reflect.Value{reflect.ValueOf(paraValue)})
	}

	session, ok := paramsArray["session"].(string)
	if !ok {
		session = ""
	}

	return c.execute(req, session)
}

func (c *Client) getCurrentTimeFormatted() string {
	utcOffset := c.getStandardOffsetUTC(time.Now().Location().String())
	return fmt.Sprintf("%s.000%s", time.Now().Format("2006-01-02 15:04:05"), utcOffset)
}

func (c *Client) getStandardOffsetUTC(timezone string) string {
	if timezone == "UTC" {
		return "+0000"
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return ""
	}

	now := time.Now().In(location)
	_, offset := now.Zone()

	return fmt.Sprintf("%+03d%02d", offset/3600, abs(offset)%3600/60)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func UnderscoreToUpperCamelCase(s string) string {
	if s == "" {
		return s
	}
	s = strings.Replace(s, ".", " ", -1)
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.Replace(s, " ", "", -1)
}
