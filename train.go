package fptai_go

import (
	"encoding/json"
	"net/http"
)

type TrainingStatusState int64

const (
	WaitingToTrain TrainingStatusState = 0
	Training       TrainingStatusState = 1
	TrainSuccess   TrainingStatusState = 2
	TrainFailed    TrainingStatusState = 3
)

type TrainingStatus struct {
	LastSuccessTime string              `json:"last_success_time"`
	Status          TrainingStatusState `json:"status"`
	Trainable       bool                `json:"trainable"`
}

type TrainingModelStatus struct {
	Code         string  `json:"code"`
	AppCode      string  `json:"app_code"`
	Type         int64   `json:"type"`
	State        int64   `json:"state"`
	Message      string  `json:"message"`
	ModelFile    string  `json:"model_file"`
	Accuracy     float64 `json:"accuracy"`
	Precision    float64 `json:"precision"`
	Recall       int64   `json:"recall"`
	F1           int64   `json:"f1"`
	CreatedTime  string  `json:"created_time"`
	FinishedTime string  `json:"finished_time"`
	StartedTime  string  `json:"started_time"`
}

type TrainingModelStatusResponse struct {
	Message string                `json:"message"`
	Status  []TrainingModelStatus `json:"status"`
}

// GetCurrentTrainingStatus - https://docs.fpt.ai/#training-status
func (c *Client) GetCurrentTrainingStatus(model string) (*TrainingStatus, error) {
	copyModel := model

	for _, c := range []string{"intent", "entity"} {
		if c == copyModel {
			model = "/" + copyModel
			break
		}
	}

	if copyModel == model {
		model = ""
	}

	resp, err := c.request(http.MethodPost, "/train/status", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var trainingStatus *TrainingStatus
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&trainingStatus)
	return trainingStatus, err
}

// GetTrainingModelStatus - https://docs.fpt.ai/#nlp-train
func (c *Client) GetTrainingModelStatus(model string) (*TrainingModelStatusResponse, error) {
	copyModel := model

	for _, c := range []string{"intent", "entity"} {
		if c == copyModel {
			model = "/" + copyModel
			break
		}
	}

	if copyModel == model {
		model = ""
	}

	resp, err := c.request(http.MethodPost, "/train"+model, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var trainingStatus *TrainingModelStatusResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&trainingStatus)
	return trainingStatus, err
}

// GetTrainingIntentStatus -
func (c *Client) GetTrainingIntentStatus() (*TrainingModelStatusResponse, error) {
	return c.GetTrainingModelStatus("intent")
}

// GetTrainingEntityStatus -
func (c *Client) GetTrainingEntityStatus() (*TrainingModelStatusResponse, error) {
	return c.GetTrainingModelStatus("entity")
}
