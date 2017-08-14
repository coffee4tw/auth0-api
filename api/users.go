package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Context) GetUserInfo(u *url.URL, h http.Header) (int, http.Header, interface{}, error) {

	req, err := http.NewRequest("GET", "https://samples.auth0.com/userinfo", nil)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("could not create new request")
	}

	token := u.Query().Get("access_token")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("could not get user info")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("could not read response body")
	}
	var data = map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, fmt.Errorf("could not marshal json")
	}

	return http.StatusOK, nil, data, nil
}
