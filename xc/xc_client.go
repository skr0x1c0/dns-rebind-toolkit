package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	ErrorInvalidSession = fmt.Errorf("session invalid/expired")
)

type XCConfig struct {
	Root     string
	Username string
	Password string
}

type XCClient struct {
	config XCConfig
	client *http.Client
	store  XCStore
}

func NewXCClient(config XCConfig) *XCClient {
	jar, err := cookiejar.New(nil)
	AssertOk(err)
	return &XCClient{
		config: config,
		client: &http.Client{Jar: jar},
		store:  NewInMemoryXCStore(),
	}
}

func (x *XCClient) TriggerNegativeCache(name string) error {
	reqUrl := "http://" + name + "/image.png"
	err := x.doDocAddFile(reqUrl)

	if err == ErrorInvalidSession {
		if err := x.doLogin(); err != nil {
			return err
		}
		err = x.doDocAddFile(reqUrl)
	}

	if err == nil {
		return fmt.Errorf("negative cache trigger expected error, got nothing")
	}

	if strings.Contains(err.Error(), "GENERAL_ARGUMENTS_ERROR") {
		return nil
	}

	return err
}

func (x *XCClient) doLogin() error {
	fmt.Println("logging in to xc")
	form := url.Values{}
	form.Set("action", "login")
	form.Set("name", x.config.Username)
	form.Set("password", x.config.Password)
	form.Set("staySignedIn", "true")

	resp, err := x.client.Post(x.config.Root+"/appsuite/api/login",
		"application/x-www-form-urlencoded", bytes.NewBuffer([]byte(form.Encode())))
	if err != nil {
		return fmt.Errorf("cannot load cookies, error %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read response body, error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot login, status %s, error %s", resp.Status, string(body))
	}

	data := XCLoginResult{}
	if err := json.Unmarshal(body, &data); err != nil || data.Session == "" {
		return fmt.Errorf("login failed, error %s", string(body))
	}
	x.store.SetSession(data)
	return nil
}

type docAddFileRequest struct {
	AddImageUrl string `json:"add_imageurl"`
	AddExt      string `json:"add_ext"`
}

func (x *XCClient) doDocAddFile(imageUrl string) error {
	req := docAddFileRequest{
		AddImageUrl: imageUrl,
		AddExt:      "png",
	}
	reqData, err := json.Marshal(&req)
	if err != nil {
		return fmt.Errorf("cannot marshal addfile request, error %v", err)
	}

	form := url.Values{}
	form.Set("action", "addfile")
	form.Set("requestdata", string(reqData))
	form.Set("version", "1")
	form.Set("app", "text")

	session, err := x.store.GetSession()
	if err != nil {
		return err
	}

	resp, err := x.client.Post(
		x.config.Root+"/appsuite/api/oxodocumentfilter?session="+session.Session,
		"application/x-www-form-urlencoded", bytes.NewBuffer([]byte(form.Encode())))
	if err != nil {
		return fmt.Errorf("cannot execute addfile request, error %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read response body, error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("addfile request failed, status: %s, error: %s", resp.Status, string(body))
	}

	if strings.Contains(string(body), "added_filename") {
		return nil
	}
	if strings.Contains(string(body), "Your session expired") {
		return ErrorInvalidSession
	}
	return fmt.Errorf("add failed, error %s", string(body))
}
