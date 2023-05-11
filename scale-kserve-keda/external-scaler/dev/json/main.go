package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

func main() {

	// infReq := InferenceRequest{

	// 	Instances: []Instances{{
	// 		LookAhead:   "20000",
	// 		CurrentTime: "2020-02-01T00:56:12Z",
	// 		ReplicaHistory: []ReplicaHistory{{
	// 			Replicas: 1,
	// 			Time:     "2020-02-01T00:55:33Z",
	// 		}},
	// 	},
	// 	},
	// }

	inferenceReq := CreateInfRequest(2000, 5, 4)

	requests := inferenceReq.Instances[0].LookAhead
	fmt.Printf("check %s", requests)

	resultRequest, err := json.Marshal(inferenceReq)
	if err != nil {
		panic(err)
	}

	fmt.Println("##########")
	fmt.Println(string(resultRequest))

	instancesJson := `{
	"instances": [{
		"lookAhead": 20000,
		"currentTime": "2020-02-01T00:56:12Z",
		"replicaHistory": [
			{
				"replicas": 1,
				"time": "2020-02-01T00:55:33Z"
			},
			{
				"replicas": 2,
				"time": "2020-02-01T00:55:43Z"
			},
			{
				"replicas": 3,
				"time": "2020-02-01T00:55:53Z"
			},
			{
				"replicas": 5,
				"time": "2020-02-01T00:56:03Z"
			}
		]
	  }
	]
  }`
	var inferenceRequest InferenceRequest
	json.Unmarshal([]byte(instancesJson), &inferenceRequest)
	fmt.Println(inferenceRequest.Instances[0].CurrentTime)

}

func CreateInfRequest(lookAhead, numberOfIteraction, step int) InferenceRequest {

	result := InferenceRequest{}

	currentTime := time.Now().UTC().Format(time.RFC3339)

	replicaHistoryList := []ReplicaHistory{}

	for i := 1; i < numberOfIteraction; i++ {
		iterTime := time.Now()
		iterTimeoffset := iterTime.Add(time.Hour * time.Duration(-step*i))

		fmt.Println("instant watch " + iterTimeoffset.UTC().Format(time.RFC3339))
		randoNumber := rand.Intn(10)
		replicaHistoryItem := ReplicaHistory{
			Replicas: randoNumber,
			Time:     iterTimeoffset.UTC().Format(time.RFC3339),
		}
		replicaHistoryList = append(replicaHistoryList, replicaHistoryItem)

	}

	instancesItem := Instances{
		LookAhead:      strconv.Itoa(lookAhead),
		CurrentTime:    currentTime,
		ReplicaHistory: replicaHistoryList,
	}
	result.Instances = append(result.Instances, instancesItem)
	return result

}
