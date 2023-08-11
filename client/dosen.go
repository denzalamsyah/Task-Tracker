package client

import (
	"a21hc3NpZ25tZW50/config"
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DosenClient interface {
	DosenList(token string) ([]*model.Dosen, error)
	AddDosen(token string, dosen model.Dosen) (respCode int, err error)
	UpdateDosen(token string, dosen model.Dosen) (respCode int, err error)
	DeleteDosen(token string, id int) (respCode int, err error)
}

type dosenClient struct {
}

func NewDosenClient() *dosenClient {
	return &dosenClient{}
}

func (t *dosenClient) DosenList(token string) ([]*model.Dosen, error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.SetUrl("/api/v1/dosen/list"), nil)
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

	var dosen []*model.Dosen
	err = json.Unmarshal(b, &dosen)
	if err != nil {
		return nil, err
	}

	return dosen, nil
}

func (t *dosenClient) AddDosen(token string, dosen model.Dosen) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		"name": dosen.Name,
		"address" : dosen.Address,
		"matkul_id" : dosen.MatkulId,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/dosen/add"), bytes.NewBuffer(data))
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

func (t *dosenClient) UpdateDosen(token string, dosen model.Dosen) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		"name": dosen.Name,
		"address" : dosen.Address,
		"matkul_id" : dosen.MatkulId,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("PUT", config.SetUrl("/api/v1/dosen/update/"+strconv.Itoa(dosen.ID)), bytes.NewBuffer(data))
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

func (t *dosenClient) DeleteDosen(token string, id int) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("DELETE", config.SetUrl("/api/v1/dosen/delete/"+strconv.Itoa(id)), nil)
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
