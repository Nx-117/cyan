package httpClient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const UPDATE = "UPDATE"
const DELETE = "DELETE"
const X_WWW_FORM = "application/x-www-form-urlencoded"
const FORM = "form"
const JSON = "application/json"

/**
get请求方式
*/
func Get(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

/**
get请求并携带参数
*/
func GetByParams(url string, params map[string]string) (string, error) {
	return Get(concatUrl(url, params))
}

/**
get请求并携带参数和请求头
*/
func GetByParamsAndHeads(url string, params, heads map[string]string) (string, error) {
	req, err := http.NewRequest(GET, concatUrl(url, params), nil)

	if 0 != len(heads) {
		for v, _ := range heads {
			req.Header.Set(v, heads[v])
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), err
}

/**
拼接url
*/
func concatUrl(url string, params map[string]string) string {
	length := len(params)
	if 0 != length {
		param := "?"
		for v, _ := range params {
			param = param + v + "=" + params[v] + "&"
		}
		url = url + param[:length-1]
	}
	return url
}

/**
post请求form表单方式
*/
func PostForm(urlStr string, params map[string]string) (string, error) {
	values := url.Values{}

	if 0 != len(params) {
		for v, _ := range params {
			values.Add(v, params[v])
		}
	}

	resp, err := http.PostForm(urlStr, values)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

/**
post 请求 form方式
并设置head
*/
func PostFormHeads(url string, params map[string]string, heads map[string]string) (string, error) {
	return postHeads(url, FORM, params, heads)
}

/**
post 请求 x-www-form方式
并设置head
*/
//func Post_X_WWW_FORM_Heads(url string, params map[string]string, heads map[string]string) (string, error) {
//	return postHeads(url, X_WWW_FORM, params, heads)
//}

func postHeads(url, clientType string, params map[string]string, heads map[string]string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if 0 != len(params) {
		for v, _ := range params {
			writer.WriteField(v, params[v])
		}
	}

	err := writer.Close()

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(POST, url, payload)

	if err != nil {
		return "", err
	}

	if 0 != len(heads) {
		for v, _ := range heads {
			req.Header.Add(v, heads[v])
		}
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

/**
post 请求 json方式
*/
func PostJson(url string, params map[string]interface{}) (string, error) {
	client := &http.Client{}
	bytesData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(POST, url, bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", JSON)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/**
post请求json方式
并设置head
*/
func PostJsonHead(url string, params map[string]interface{}, heads map[string]string) (string, error) {
	bytesData, err := json.Marshal(params)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(POST, url, bytes.NewReader(bytesData))

	if err != nil {
		return "", err
	}

	if 0 != len(heads) {
		for v, _ := range heads {
			req.Header.Add(v, heads[v])
		}
	}

	req.Header.Add("Content-Type", JSON)
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

/**
post请求 发送文件 并携带head
可传多个文件和参数，head
*/
func PostSendFileAndHead(url string, params, files, heads map[string]string) (string, error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if 0 != len(files) {
		for v, _ := range files {
			file, err := os.Open(files[v])
			defer file.Close()
			part,
				err := writer.CreateFormFile(v, filepath.Base(files[v]))
			_, err = io.Copy(part, file)
			if err != nil {
				return "", err
			}
		}
	}

	if 0 != len(params) {
		for v, _ := range params {
			_ = writer.WriteField(v, params[v])
		}
	}

	err := writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(POST, url, payload)

	if err != nil {
		return "", err
	}
	if 0 != len(heads) {
		for v, _ := range heads {
			req.Header.Add(v, heads[v])
		}
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
