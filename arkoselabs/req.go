package arkoselabs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Token *string `json:"token"`
}

func (a *Arkoselabs) GetLoginArkoseToken() *Result {
	if a.login_api != "" {
		r1 := a.getArkoseToken(a.login_api)
		fmt.Println("api:", a.login_api)
		if r1 != nil {
			return r1
		}
	}
	return nil
}

func (a *Arkoselabs) GetPlatformLoginArkoseToken(publicKey string) *Result {
	if a.login_api != "" {
		r1 := a.getArkoseToken(fmt.Sprintf("%s?key=%s", a.platformLoginApi, publicKey))
		fmt.Println("platformLoginApi:", a.platformLoginApi)
		if r1 != nil {
			return r1
		}
	}
	return nil
}

func (a *Arkoselabs) getArkoseToken(api string) *Result {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, api, nil)
	if err != nil {
		fmt.Println(0, err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(1, err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(2, err)
		return nil
	}
	fmt.Println(string(body))
	var r *Result
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(3, err)

		return nil
	}
	if r.Token != nil {
		return r
	}
	return nil
}
