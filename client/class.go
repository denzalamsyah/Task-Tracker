package client

import (
	"a21hc3NpZ25tZW50/config"
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ClassClient interface {
	ClassList(token string) ([]*model.Class, error)
	AddClass(token string, class model.Class) (respCode int, err error)
	UpdateClass(token, id string, class model.Class) (respCode int, err error)
	DeleteClass(token, id string) (respCode int, err error)
}

type classClient struct {
}

func NewClassClient() *classClient {
	return &classClient{}
}

func (c *classClient) ClassList(token string) ([]*model.Class, error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.SetUrl("/api/v1/class/list"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status code not 200")
	}

	var Class []*model.Class
	err = json.Unmarshal(b, &Class)
	if err != nil {
		return nil, err
	}

	return Class, nil
}

func (c *classClient) AddClass(token string, class model.Class) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		"name": class.Name,
		"professor" : class.Professor,
		"room_number" :class.RoomNumber,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/class/add"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1, errors.New("status code not 200")
	}

	return resp.StatusCode, nil
}

func (c *classClient) UpdateClass(token, id string, class model.Class) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		"name": class.Name,
		"professor" : class.Professor,
		"room_number" :class.RoomNumber,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("PUT", config.SetUrl("/api/v1/class/update/"+id), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1, errors.New("status code not 200")
	}

	return resp.StatusCode, nil
}

func (c *classClient) DeleteClass(token, id string) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("DELETE", config.SetUrl("/api/v1/class/delete/"+id), nil)
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1, errors.New("status code not 200")
	}

	return resp.StatusCode, nil
}
