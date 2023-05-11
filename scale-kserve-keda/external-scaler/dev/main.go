package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	url_pkg "net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type ReplicaHistory struct {
	Replicas int    `json:"replicas"`
	Time     string `json:"time"`
}

type Instances struct {
	LookAhead      string `json:"lookAhead"`
	CurrentTime    string `json:"currentTime"`
	ReplicaHistory []ReplicaHistory
}

type InferenceRequest struct {
	Instances []Instances
}

// Struct that permit me to define how to build inferenceRequest
type InferenceRequired struct {
	replicaHistoryLength int
	historyInterval      int
	lookAhead            int
}

func main() {

	timeout := time.Duration(10000) / time.Millisecond
	httpClient := http.Client{
		Timeout: timeout,
	}
	inferenceRequired := InferenceRequired{replicaHistoryLength: 5,
		historyInterval: 5,
		lookAhead:       2000,
	}
	inferenceReq := CreateFakeInfRequest(inferenceRequired)
	CallInferenceService(inferenceReq, httpClient, "Http://localhost")
	//fake request
	//_, err := CreateInfRequest(inferenceRequired, "count(count by (pod) (rate(http_request_duration_seconds_count[range])))", "http://localhost:9090", *httpClient)

}

type promQueryResult struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func executePromQuery(query, promAddress, time string, httpClient http.Client) (float64, error) {

	queryEscaped := url_pkg.QueryEscape(query)

	url := fmt.Sprintf("%s/api/v1/query?query=%s&time=%s", promAddress, queryEscaped, time)

	log.Info("query %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}

	r, err := httpClient.Do(req)

	if err != nil {
		return -1, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return -1, err
	}
	r.Body.Close()

	if !(r.StatusCode >= 200 && r.StatusCode <= 299) {
		return -1, fmt.Errorf("prometheus query api returned error. status: %d response: %s", r.StatusCode, string(b))
	}

	var result promQueryResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return -1, err
	}

	// allow for zero element or single element result sets
	if len(result.Data.Result) == 0 {
		return 0, nil
	} else if len(result.Data.Result) > 1 {

		_, err := retrieveValFromJson(result)
		if err != nil {
			return -1, err
		}
		return -1, fmt.Errorf("prometheus query %s returned multiple elements", query)
	}

	return retrieveValFromJson(result)
}

func retrieveValFromJson(promResult promQueryResult) (float64, error) {
	val := promResult.Data.Result[0].Value[1]
	var result float64 = -1
	var err error
	if val != nil {
		s := val.(string)
		result, err = strconv.ParseFloat(s, 64)
		if err != nil {
			log.Error(err, "error converting prometheus value", "prometheus_value", s)
			return -1, err
		}
	}
	return removeInf(result), nil
}

func removeInf(val float64) float64 {

	if math.IsInf(val, 0) {
		return 0
	} else if math.IsNaN(val) {
		return 0
	}
	return val
}

func CreateInfRequest(required InferenceRequired, query, promAddress string, httpClient http.Client) (*InferenceRequest, error) {
	result := InferenceRequest{}
	currentTime := time.Now().UTC().Format(time.RFC3339)
	replicaHistoryList := []ReplicaHistory{}
	for i := 1; i < required.replicaHistoryLength; i++ {
		iterTime := time.Now()
		iterTimeoffset := iterTime.Add(time.Hour * time.Duration(-required.replicaHistoryLength*i))

		value, err := executePromQuery(query, promAddress, iterTimeoffset.UTC().Format(time.RFC3339), httpClient)
		if err != nil {
			return nil, err
		}

		replicaHistoryItem := ReplicaHistory{
			Replicas: int(value),
			Time:     iterTimeoffset.UTC().Format(time.RFC3339),
		}
		replicaHistoryList = append(replicaHistoryList, replicaHistoryItem)

	}

	instancesItem := Instances{
		LookAhead:      strconv.Itoa(required.lookAhead),
		CurrentTime:    currentTime,
		ReplicaHistory: replicaHistoryList,
	}
	result.Instances = append(result.Instances, instancesItem)
	return &result, nil

}

func CallInferenceService(inferenceReq InferenceRequest, client http.Client, inferenceSeviceAddr string) (string, error) {

	resultRequest, err := json.Marshal(inferenceReq)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader(resultRequest)

	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, inferenceSeviceAddr, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err

	}

	return string(body), nil
}

func CreateFakeInfRequest(required InferenceRequired) InferenceRequest {
	result := InferenceRequest{}
	currentTime := time.Now().UTC().Format(time.RFC3339)
	replicaHistoryList := []ReplicaHistory{}
	for i := 1; i < required.replicaHistoryLength; i++ {
		iterTime := time.Now()
		iterTimeoffset := iterTime.Add(time.Hour * time.Duration(-required.historyInterval*i))
		randoNumber := rand.Intn(10)
		replicaHistoryItem := ReplicaHistory{
			Replicas: randoNumber,
			Time:     iterTimeoffset.UTC().Format(time.RFC3339),
		}
		replicaHistoryList = append(replicaHistoryList, replicaHistoryItem)

	}

	instancesItem := Instances{
		LookAhead:      strconv.Itoa(required.lookAhead),
		CurrentTime:    currentTime,
		ReplicaHistory: replicaHistoryList,
	}
	result.Instances = append(result.Instances, instancesItem)
	return result

}
