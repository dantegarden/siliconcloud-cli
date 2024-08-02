package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/samber/lo"
	"github.com/siliconflow/siliconcloud-cli/meta"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"unicode"
)

// Client siliconCloud client
type Client struct {
	Domain string
	ApiKey string
}

// Response the response of siliconCloud
type Response[T any] struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Status    bool   `json:"status,omitempty"`
	Data      T      `json:"data"`
	Error     string `json:"error,omitempty"`
}

// NewClient New Client
func NewClient(domain string, apiKey string) *Client {
	return &Client{
		Domain: domain,
		ApiKey: apiKey,
	}
}

func (c *Client) UserInfo() (*Response[UserInfo], error) {
	serverUrl := fmt.Sprintf("%s/%s/user/info", c.Domain, meta.APIv1)
	body, statusCode, err := c.doGet(serverUrl, nil, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, handleError(body, statusCode)
	}
	return handleResponse[UserInfo](body)
}

func (c *Client) Sign(signature string) (*Response[FilesResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/files/%s", c.Domain, meta.APIv1, signature)
	body, statusCode, err := c.doGet(serverUrl, nil, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[FilesResp](body)
}

func (c *Client) CommitFile(signature string, objectKey string) (*Response[FilesResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/files", c.Domain, meta.APIv1)
	body, statusCode, err := c.doPost(serverUrl, FileCommitReq{
		Sign:      signature,
		ObjectKey: objectKey,
	}, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[FilesResp](body)
}

func (c *Client) CommitModel(modelName string, modelType string, overwrite bool, modelFiles []*ModelFile) (*Response[ModelCommitResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/models", c.Domain, meta.APIv1)
	body, statusCode, err := c.doPost(serverUrl, ModelCommitReq{
		Name:      modelName,
		Type:      modelType,
		Overwrite: overwrite,
		Files:     modelFiles,
	}, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[ModelCommitResp](body)
}

func (c *Client) ListModel(modelType string) (*Response[ModelListResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/models", c.Domain, meta.APIv1)
	if modelType != "" {
		serverUrl = fmt.Sprintf("%s?type=%s", serverUrl, modelType)
	}
	body, statusCode, err := c.doGet(serverUrl, nil, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[ModelListResp](body)
}

func (c *Client) ListModelFiles(modelType string, modelName string, extName string) (*Response[ModelListFilesResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/models/%s/files", c.Domain, meta.APIv1, modelType)
	param := ModelListFilesReq{
		Name:    modelName,
		ExtName: extName,
	}
	body, statusCode, err := c.doGet(serverUrl, param, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[ModelListFilesResp](body)
}

func (c *Client) RemoveModel(modelType string, modelName string) (*Response[ModelDeleteResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/models", c.Domain, meta.APIv1)
	body, statusCode, err := c.doDelete(serverUrl, ModelDeleteReq{
		Name: modelName,
		Type: modelType,
	}, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[ModelDeleteResp](body)
}

func (c *Client) CheckModel(modelType string, modelName string) (*Response[CheckModelResp], error) {
	serverUrl := fmt.Sprintf("%s/x/%s/models/check", c.Domain, meta.APIv1)
	body, statusCode, err := c.doGet(serverUrl, ModelQueryReq{
		Name: modelName,
		Type: modelType,
	}, c.authHeader())
	if err != nil {
		return nil, cli.Exit(err, meta.ServerError)
	}

	if statusCode != http.StatusOK {
		return nil, cli.Exit(handleError(body, statusCode), meta.ServerError)
	}
	return handleResponse[CheckModelResp](body)
}

func (c *Client) authHeader() map[string]string {
	header := make(map[string]string)
	header[meta.HeaderAuthorization] = fmt.Sprintf("Bearer %s", c.ApiKey)
	return header
}

// doGet do get request
func (c *Client) doGet(urlStr string, queryParams interface{}, header map[string]string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, -1, err
	}

	if queryParams != nil {
		v := reflect.ValueOf(queryParams)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		query := parsedURL.Query()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Name
			fieldValue := fmt.Sprintf("%v", field.Interface())
			if fieldValue != "" && fieldValue != "0" {
				query.Add(lo.SnakeCase(fieldName), fieldValue)
			}
		}
		parsedURL.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(meta.HTTPGet, parsedURL.String(), nil)
	if err != nil {
		return nil, -1, err
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	req.Header.Set(meta.HeaderSiliconCliVersion, meta.Version)

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// doPost do post request
func (c *Client) doPost(url string, data interface{}, header map[string]string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}

	// 将数据编码为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, -1, err
	}

	req, err := http.NewRequest(meta.HTTPPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, -1, err
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	req.Header.Set(meta.HeaderSiliconCliVersion, meta.Version)
	req.Header.Set(meta.HeaderContentType, meta.JsonContentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

// doDelete do delete request
func (c *Client) doDelete(url string, data interface{}, header map[string]string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, -1, err
	}

	req, err := http.NewRequest(meta.HTTPDelete, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, -1, err
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	req.Header.Set(meta.HeaderContentType, meta.JsonContentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

func handleError(responseBody []byte, statusCode int) error {
	rawMessage := string(responseBody)
	if statusCode == http.StatusNotFound {
		return cli.Exit(fmt.Errorf("server not found, you can use \"--base_domain\" to specify the target domain"), meta.LoadError)
	}
	var parsedResponse Response[interface{}]
	err := json.Unmarshal(responseBody, &parsedResponse)
	if err != nil {
		rawMessage = strings.TrimFunc(rawMessage, func(r rune) bool {
			return unicode.Is(unicode.Quotation_Mark, r)
		})
		if rawMessage == "" {
			return cli.Exit(meta.NewErrNo("Unknown server error"), meta.ServerError)
		}
		return cli.Exit(meta.NewErrNo(rawMessage), meta.ServerError)
	}

	if errno, exists := meta.ServerErrors[parsedResponse.Code]; exists {
		return cli.Exit(errno, meta.ServerError)
	}

	return fmt.Errorf("unexcepted http status code: %d, message: %s", statusCode, rawMessage)
}

func handleResponse[T any](responseBody []byte) (*Response[T], error) {
	var parsedResponse Response[T]
	err := json.Unmarshal(responseBody, &parsedResponse)
	if err != nil {
		logs.Debugf("error: %s\n", err)
		return nil, cli.Exit(err, meta.ServerError)
	}

	if parsedResponse.Code != meta.OKCode {
		if errno, exists := meta.ServerErrors[parsedResponse.Code]; exists {
			return nil, cli.Exit(errno, meta.ServerError)
		}
		return nil, cli.Exit(fmt.Errorf("server error: %s", parsedResponse.Message), meta.ServerError)
	}
	return &parsedResponse, nil
}
