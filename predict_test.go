package fptai_go

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestClient_GetNPLPredict(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"status":{"code":200,"message":"Predict All successful","module":"","api_code":0,"err_code":0},"data":{"intents":[{"label":"ask_product","confidence":0.92},{"label":"ask_general_information","confidence":0.04},{"label":"ask_inventory","confidence":0.03}],"entities":[{"start":9,"end":15,"value":"iPhone","real_value":"apple-iphone","entity":"filter_brand","subentities":null}]},"history_id":0}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient("UnitTestToken")
	c.ApiBaseUrl = testServer.URL
	predictResp, _ := c.GetPredict("Test content", true, ContainAll)

	if predictResp.Status.Code != 200 {
		t.Errorf("expected status code 200, got: %v", predictResp.Status.Code)
	}

	if len(predictResp.Data.Intents) != 3 {
		t.Errorf("expected 3 intents, got: %v", predictResp.Data.Intents)
	}

	if len(predictResp.Data.Entities) != 1 {
		t.Errorf("expected 1 entities, got: %v", predictResp.Data.Entities)
	}
}

func TestClient_GetPredictIntents(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"status":{"code":200,"message":"Predict Intents successful","module":"","api_code":0,"err_code":0},"data":{"intents":[{"label":"product_info","confidence":0.99},{"label":"purchase","confidence":0.005}]}}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient("UnitTestToken")
	c.ApiBaseUrl = testServer.URL
	predictResp, _ := c.GetPredictIntents("Test content", false)

	if predictResp.Status.Code != 200 {
		t.Errorf("expected status code 200, got: %v", predictResp.Status.Code)
	}

	if len(predictResp.Data.Intents) != 2 {
		t.Errorf("expected 2 intents, got: %v", predictResp.Data.Intents)
	}
}

func TestClient_GetPredictEntities(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"status":{"code":200,"message":"Recognize Entities successful","module":"","api_code":0,"err_code":0},"data":{"entities":[{"start":32,"end":41,"value":"vegetable","real_value":"vegetable","entity":"product_name","subentities":[]}]}}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient("UnitTestToken")
	c.ApiBaseUrl = testServer.URL
	predictResp, _ := c.GetPredictEntities("Test content", true)

	if predictResp.Status.Code != 200 {
		t.Errorf("expected status code 200, got: %v", predictResp.Status.Code)
	}

	if len(predictResp.Data.Entities) != 1 {
		t.Errorf("expected 1 entities, got: %v", predictResp.Data.Entities)
	}
}

func TestClient_GetNPLPredictLive(t *testing.T) {
	client := NewClient(os.Getenv("FPTAI_TEST_TOKEN"))
	predictResp, err := client.GetPredict("ngày 16 tháng sau vào thứ mấy?", false, "")
	t.Logf("predictResp: %#v", predictResp)
	t.Logf("err: %v", err)
}