package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//GetAllauxin returns list of auxin
func (c Client) GetAllauxin() ([]auxin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auxin/getAllauxin", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var auxin []auxin
	err = json.Unmarshal(body, &auxin)
	if err != nil {
		return nil, err
	}

	return auxin, nil
}

//Createauxin will create an auxin
func (c *Client) Createauxin(auxin auxin) (*auxin, error) {
	avg, err := json.Marshal(auxin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auxin/createNewauxin", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var insertedID InsertedResult
	err = json.Unmarshal(body, &insertedID)
	if err != nil {
		return nil, err
	}
	auxin.ID = insertedID.InsertedID
	return &auxin, nil
}

//UpdateauxinByName will update an auxin
func (c *Client) UpdateauxinByName(auxin auxin) (*UpdateResult, error) {
	avg, err := json.Marshal(auxin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/auxin/updateauxinByName", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var updateResult UpdateResult
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		return nil, err
	}

	return &updateResult, nil
}

//DeleteauxinByName will delete an auxin
func (c *Client) DeleteauxinByName(auxinName string) (*DeleteResult, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/auxin/deleteauxinByName", c.HostURL), http.NoBody)
	req.URL.Query().Add("name", auxinName)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var deleteResult DeleteResult
	err = json.Unmarshal(body, &deleteResult)
	if err != nil {
		return nil, err
	}

	return &deleteResult, nil
}
