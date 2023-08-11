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

type StudentClient interface {
	StudentList(token string) ([]*model.Mahasiswa, error)
	AddStudent(token string, mahasiswa model.Mahasiswa) (respCode int, err error)
	UpdateStudent(token string, mahasiswa model.Mahasiswa) (respCode int, err error)
	DeleteStudent(token string, id int) (respCode int, err error)
}

type studentClient struct {
}

func NewStudentClient() *studentClient {
	return &studentClient{}
}

func (t *studentClient) StudentList(token string) ([]*model.Mahasiswa, error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.SetUrl("/api/v1/student/list"), nil)
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

	var student []*model.Mahasiswa
	err = json.Unmarshal(b, &student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (t *studentClient) AddStudent(token string, mahasiswa model.Mahasiswa) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		// "id" : mahasiswa.ID,
		"name" : mahasiswa.Name,
		"address" : mahasiswa.Address,
		"class_id" : mahasiswa.ClassId,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/student/add"), bytes.NewBuffer(data))
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

func (t *studentClient) UpdateStudent(token string, mahasiswa model.Mahasiswa) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	datajson := map[string]interface{}{
		"id" : mahasiswa.ID,
		"name" : mahasiswa.Name,
		"address" : mahasiswa.Address,
		"class_id" : mahasiswa.ClassId,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("PUT", config.SetUrl("/api/v1/student/update/"+strconv.Itoa(mahasiswa.ID)), bytes.NewBuffer(data))
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

func (t *studentClient) DeleteStudent(token string, id int) (respCode int, err error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("DELETE", config.SetUrl("/api/v1/student/delete/"+strconv.Itoa(id)), nil)
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
