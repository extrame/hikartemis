package hikartemis

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/extrame/hikartemis/camera"
	"github.com/extrame/hikartemis/region"
	"github.com/gofrs/uuid"
)

// HKConfig 海康OpenAPI配置参数
type HKConfig struct {
	Ip      string //平台ip
	Port    int    //平台端口
	AppKey  string //平台APPKey
	Secret  string //平台APPSecret
	IsHttps bool   //是否使用HTTPS协议
	client  *http.Client
}

// 返回结果
type BaseResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回值data
type ListData struct {
	Total    int         `json:"total"`
	PageSize int         `json:"pageSize"`
	PageNo   int         `json:"pageNo"`
	List     interface{} `json:"list"`
}

func (d *ListData) UnmarshalJSON(data []byte) error {
	type Alias ListData
	aux := &struct {
		Total    int             `json:"total"`
		PageSize int             `json:"pageSize"`
		PageNo   int             `json:"pageNo"`
		List     json.RawMessage `json:"list"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if d.List == nil {
		return nil
	}
	fmt.Printf("%T, %v", d.List, d.List)
	if err := json.Unmarshal(aux.List, d.List); err != nil {
		return errors.Wrap(err, "failed to unmarshal list"+string(aux.List))
	}
	return nil
}

func (hk *HKConfig) Init(timeout int) {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client.Timeout = time.Duration(timeout) * time.Second
	if hk.IsHttps {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}
	hk.client = client
}

// @title		HTTP Post请求
// @url			HTTP接口Url		string				 HTTP接口Url，不带协议和端口，如/artemis/api/resource/v1/org/advance/orgList
// @body		请求参数			map[string]string
// @return		请求结果			参数类型
func (hk *HKConfig) HttpPost(url string, body interface{}, resp interface{}) (result *BaseResult, err error) {
	var header = make(map[string]string)
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = hk.initRequest(header, url, string(bodyJson), true)
	if err != nil {
		return nil, err
	}
	var sb []string
	if hk.IsHttps {
		sb = append(sb, "https://")
	} else {
		sb = append(sb, "http://")
	}
	sb = append(sb, fmt.Sprintf("%s:%d", hk.Ip, hk.Port))
	sb = append(sb, url)

	req, err := http.NewRequest("POST", strings.Join(sb, ""), bytes.NewReader(bodyJson))
	if err != nil {
		return
	}

	req.Header.Set("Accept", header["Accept"])
	req.Header.Set("Content-Type", header["Content-Type"])
	for k, v := range header {
		if strings.Contains(k, "x-ca-") {
			req.Header.Set(k, v)
		}
	}
	httpresp, err := hk.client.Do(req)
	if err != nil {
		return
	}
	defer httpresp.Body.Close()
	if httpresp.StatusCode == http.StatusOK {
		var resBody []byte
		resBody, err = io.ReadAll(httpresp.Body)
		if err != nil {
			return
		}
		result = &BaseResult{
			Data: resp,
		}
		err = json.Unmarshal(resBody, result)
	} else if httpresp.StatusCode == http.StatusFound || httpresp.StatusCode == http.StatusMovedPermanently {
		reqUrl := httpresp.Header.Get("Location")
		err = fmt.Errorf("HttpPost Response StatusCode：%d，Location：%s", httpresp.StatusCode, reqUrl)
	} else {
		err = fmt.Errorf("HttpPost Response StatusCode：%d", httpresp.StatusCode)
	}
	return
}

// @title		HTTP Post请求
// @url			HTTP接口Url		string				 HTTP接口Url，不带协议和端口，如/artemis/api/resource/v1/org/advance/orgList
// @body		请求参数			map[string]string
// @return		请求结果			参数类型
func (hk *HKConfig) RawHttpPost(url string, body map[string]string) (result BaseResult, err error) {
	var header = make(map[string]string)
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = hk.initRequest(header, url, string(bodyJson), true)
	if err != nil {
		return BaseResult{}, err
	}
	var sb []string
	if hk.IsHttps {
		sb = append(sb, "https://")
	} else {
		sb = append(sb, "http://")
	}
	sb = append(sb, fmt.Sprintf("%s:%d", hk.Ip, hk.Port))
	sb = append(sb, url)

	req, err := http.NewRequest("POST", strings.Join(sb, ""), bytes.NewReader(bodyJson))
	if err != nil {
		return
	}

	req.Header.Set("Accept", header["Accept"])
	req.Header.Set("Content-Type", header["Content-Type"])
	for k, v := range header {
		if strings.Contains(k, "x-ca-") {
			req.Header.Set(k, v)
		}
	}
	resp, err := hk.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var resBody []byte
		resBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(resBody, &result)
	} else if resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusMovedPermanently {
		reqUrl := resp.Header.Get("Location")
		err = fmt.Errorf("HttpPost Response StatusCode：%d，Location：%s", resp.StatusCode, reqUrl)
	} else {
		err = fmt.Errorf("HttpPost Response StatusCode：%d", resp.StatusCode)
	}
	return
}

// initRequest 初始化请求头
func (hk HKConfig) initRequest(header map[string]string, url, body string, isPost bool) error {
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	if isPost {
		var err error
		header["content-md5"], err = computeContentMd5(body)
		if err != nil {
			return err
		}
	}
	header["x-ca-timestamp"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	header["x-ca-nonce"] = uid.String()
	header["x-ca-key"] = hk.AppKey

	var strToSign string
	if isPost {
		strToSign = buildSignString(header, url, "POST")
	} else {
		strToSign = buildSignString(header, url, "GET")
	}
	signedStr, err := computeForHMACSHA256(strToSign, hk.Secret)
	if err != nil {
		return err
	}
	header["x-ca-signature"] = signedStr
	return nil
}

// computeContentMd5 计算content-md5
func computeContentMd5(body string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return "", err
	}
	md5Str := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(md5Str)), nil
}

// computeForHMACSHA256 计算HMACSHA265
func computeForHMACSHA256(str, secret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// buildSignString 计算签名字符串
func buildSignString(header map[string]string, url, method string) string {
	var sb []string
	sb = append(sb, strings.ToUpper(method))
	sb = append(sb, "\n")

	if header != nil {
		if _, ok := header["Accept"]; ok {
			sb = append(sb, header["Accept"])
			sb = append(sb, "\n")
		}
		if _, ok := header["Content-MD5"]; ok {
			sb = append(sb, header["Content-MD5"])
			sb = append(sb, "\n")
		}
		if _, ok := header["Content-Type"]; ok {
			sb = append(sb, header["Content-Type"])
			sb = append(sb, "\n")
		}
		if _, ok := header["Date"]; ok {
			sb = append(sb, header["Date"])
			sb = append(sb, "\n")
		}
	}
	sb = append(sb, buildSignHeader(header))
	sb = append(sb, url)
	return strings.Join(sb, "")
}

// buildSignHeader 计算签名头
func buildSignHeader(header map[string]string) string {
	var sortedDicHeader map[string]string
	sortedDicHeader = header

	var sslice []string
	for key, _ := range sortedDicHeader {
		sslice = append(sslice, key)
	}
	sort.Strings(sslice)

	var sbSignHeader []string
	var sb []string
	//在将key输出
	for _, k := range sslice {
		if strings.Contains(strings.ReplaceAll(k, " ", ""), "x-ca-") {
			sb = append(sb, k+":")
			if sortedDicHeader[k] != "" {
				sb = append(sb, sortedDicHeader[k])
			}
			sb = append(sb, "\n")
			if len(sbSignHeader) > 0 {
				sbSignHeader = append(sbSignHeader, ",")
			}
			sbSignHeader = append(sbSignHeader, k)
		}
	}

	header["x-ca-signature-headers"] = strings.Join(sbSignHeader, "")
	return strings.Join(sb, "")
}

func (hk HKConfig) GetCameraList(regions ...string) (camera.CameraList, error) {
	body := map[string]interface{}{
		"pageNo":           "1",
		"pageSize":         "100",
		"regionIndexCodes": regions,
		"isSubRegion":      true,
	}
	var resq = make(camera.CameraList, 0)
	result, err := hk.HttpPost("/artemis/api/resource/v2/camera/search", body, &ListData{
		List: &resq,
	})
	if err != nil {
		return nil, err
	}
	var list = result.Data.(*ListData)
	var data = list.List.(*camera.CameraList)
	return *data, nil
}

func (hk HKConfig) GetCameraUrl(cam *camera.Camera) (*camera.Url, error) {
	body := map[string]interface{}{
		"cameraIndexCode": cam.IndexCode,
		"streamType":      0,
		"protocol":        "wss",
		"streamform":      "ps",
	}
	var resq camera.Url
	result, err := hk.HttpPost("/artemis/api/video/v2/cameras/previewURLs", body, &resq)
	if err != nil {
		return nil, err
	}
	var data = result.Data.(*camera.Url)
	return data, nil
}

func (hk HKConfig) GetResourceList() (camera.CameraList, error) {
	body := map[string]string{
		"pageNo":   "1",
		"pageSize": "100",
	}
	var resp = make(camera.CameraList, 0)
	result, err := hk.HttpPost("/artemis/api/irds/v2/resource/resourcesByParams", body, &ListData{
		List: &resp,
	})
	if err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, errors.New(result.Msg)
	}
	var list = result.Data.(*ListData)
	var data = list.List.(*camera.CameraList)
	return *data, nil
}

func (hk HKConfig) GetRootRegion() (*region.Region, error) {
	body := map[string]string{
		"treeCode": "0",
	}
	var resp region.Region
	result, err := hk.HttpPost("/artemis/api/resource/v1/regions/root", body, &resp)
	if err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, errors.New(result.Msg)
	}
	var data = result.Data.(*region.Region)
	return data, nil
}

func (hk HKConfig) GetSubRegion(parentIndexCode string) (region.RegionList, error) {
	body := map[string]string{
		"parentIndexCode": parentIndexCode,
		"pageNo":          "1",
		"pageSize":        "100",
	}
	var resp region.RegionList
	result, err := hk.HttpPost("/artemis/api/resource/v1/regions/subRegions", body, &resp)
	if err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, errors.New(result.Msg)
	}
	var data = result.Data.(*region.RegionList)
	return *data, nil
}
