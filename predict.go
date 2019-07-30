package fptai_go

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	ContainIntent string = "intent"
	ContainEntity string = "entity"
	ContainAll    string = "all"
)

type PredictStatus struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Module  string `json:"module"`
	ApiCode int64  `json:"api_code"`
	ErrCode int64  `json:"err_code"`
	Detail  string `json:"detail"`
	AppCode string `json:"app_code"`
}

type PredictEntity struct {
	Start       int64       `json:"start"`
	End         int64       `json:"end"`
	Value       string      `json:"value"`
	RealValue   interface{} `json:"real_value"`
	Entity      string      `json:"entity"`
	Subentities interface{} `json:"subentities"`
}

type PredictIntent struct {
	Label      string  `json:"label"`
	Confidence float64 `json:"confidence"`
}

type Predict struct {
	Entities  []PredictEntity `json:"entities"`
	Intents   []PredictIntent `json:"intents"`
	HistoryId int64              `json:"history_id"`
}

type PredictRequest struct {
	Content     string `json:"content"`
	SaveHistory bool   `json:"save_history"`
}

type PredictResponse struct {
	Status PredictStatus `json:"status"`
	Data   Predict       `json:"data"`
}

// GetNPLPredict - https://docs.fpt.ai/#nlp-predict
func (c *Client) GetPredict(content string, saveHistory bool, contain string) (*PredictResponse, error) {
	copyContain := contain

	for _, c := range []string{ContainIntent, ContainEntity} {
		if c == copyContain {
			contain = "/" + copyContain
			break
		}
	}

	if copyContain == contain {
		contain = ""
	}

	predictRequest := &PredictRequest{
		Content:     content,
		SaveHistory: saveHistory,
	}
	jsonRequest, err := json.Marshal(predictRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "predict"+contain, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var predict *PredictResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&predict)
	return predict, err
}

// GetNPLPredictEntities -
func (c *Client) GetPredictEntities(content string, saveHistory bool) (*PredictResponse, error) {
	return c.GetPredict(content, saveHistory, ContainEntity)
}

// GetNPLPredictIntents -
func (c *Client) GetPredictIntents(content string, saveHistory bool) (*PredictResponse, error) {
	return c.GetPredict(content, saveHistory, ContainIntent)
}
