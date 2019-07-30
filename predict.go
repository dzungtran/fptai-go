package fptai_go

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	PredictIntent string = "intent"
	PredictEntity string = "entity"
	PredictAll    string = "all"
)

type NPLPredictStatus struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Module  string `json:"module"`
	ApiCode string `json:"api_code"`
	ErrCode string `json:"err_code"`
}

type NPLPredictEntity struct {
	Start       int64       `json:"start"`
	End         int64       `json:"end"`
	Value       string      `json:"value"`
	RealValue   string      `json:"real_value"`
	Entity      string      `json:"entity"`
	Subentities interface{} `json:"subentities"`
}

type NPLPredictIntent struct {
	Label      string  `json:"label"`
	Confidence float64 `json:"confidence"`
}

type NPLPredict struct {
	Entities  []NPLPredictEntity `json:"entities"`
	Intents   []NPLPredictIntent `json:"intents"`
	HistoryId int64              `json:"history_id"`
}

type NPLPredictRequest struct {
	Content     string `json:"content"`
	SaveHistory bool   `json:"save_history"`
}

type NPLPredictResponse struct {
	Status NPLPredictStatus `json:"status"`
	Data   NPLPredict       `json:"data"`
}

// GetNPLPredict - https://docs.fpt.ai/#nlp-predict
func (c *Client) GetNPLPredict(content string, saveHistory bool, contain string) (*NPLPredictResponse, error) {
	copyContain := contain

	for _, c := range []string{PredictIntent, PredictEntity} {
		if c == copyContain {
			contain = "/" + copyContain
			break
		}
	}

	if copyContain == contain {
		contain = ""
	}

	predictRequest := &NPLPredictRequest{
		Content:     content,
		SaveHistory: saveHistory,
	}
	jsonRequest, err := json.Marshal(predictRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/predict"+contain, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var predict *NPLPredictResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&predict)
	return predict, err
}

// GetNPLPredictEntities -
func (c *Client) GetNPLPredictEntities(content string, saveHistory bool) (*NPLPredictResponse, error) {
	return c.GetNPLPredict(content, saveHistory, PredictEntity)
}

// GetNPLPredictIntents -
func (c *Client) GetNPLPredictIntents(content string, saveHistory bool) (*NPLPredictResponse, error) {
	return c.GetNPLPredict(content, saveHistory, PredictIntent)
}
