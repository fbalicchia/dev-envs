package scaler

import (
	"context"

	pb "external-scaler/externalscaler"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	url_pkg "net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type ExternalScaler struct {
	Metadata   *PrometheusMetadata
	HttpClient *http.Client
}

const (
	// Scaler Input start
	predictScale = "predictScale"
	promQuery    = "query"
	//Thrashold to stop downscale
	threshold = "threshold"
	//Refer to the HPA docs for how HPA calculates replicaCount based on metric value and target value.
	//KEDA uses the metric target type AverageValue for external metrics.
	//This will cause the metric value returned by the external scaler to be divided by the number of replicas.
	//As ugly workaround waiting issue https://github.com/kedacore/keda/issues/2030

)

type PrometheusMetadata struct {
	ServerAddress string
}

type promQueryError struct {
	promError        string
	numberOfRecords  int
	firstValueResult float64
}

func (pqe *promQueryError) Error() string {
	return pqe.promError
}

func (pqe *promQueryError) tooManyRecords() bool {
	return pqe.numberOfRecords > 1
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

func (e *ExternalScaler) IsActive(ctx context.Context, scaledObject *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	//here is necessary to call inference service to undestand the number of instance necessary
	//for now return true
	//query := scaledObject.ScalerMetadata[promQuery]
	//result := val > 0
	return &pb.IsActiveResponse{
		Result: true,
	}, nil
}

func (e *ExternalScaler) GetMetricSpec(_ context.Context, in *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	metricName := in.ScalerMetadata[predictScale]
	stopValue, err := strconv.ParseInt(in.ScalerMetadata[threshold], 10, 64)
	if err != nil {
		return nil, err
	}

	log.Debugf("called GetMetricSpec for metrics %s  return value %d ", metricName, stopValue)

	return &pb.GetMetricSpecResponse{
		MetricSpecs: []*pb.MetricSpec{{
			MetricName: normalizeString(fmt.Sprintf("%s-%s-%s", "fbalicchia", e.Metadata.ServerAddress, metricName)),
			TargetSize: stopValue,
		}},
	}, nil
}

func normalizeString(s string) string {
	s = strings.ReplaceAll(s, "/", "-")
	s = strings.ReplaceAll(s, ".", "-")
	s = strings.ReplaceAll(s, ":", "-")
	s = strings.ReplaceAll(s, "%", "-")
	return s
}

func (e *ExternalScaler) GetMetrics(_ context.Context, metricRequest *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	metricName := metricRequest.ScaledObjectRef.ScalerMetadata[predictScale]
	query := metricRequest.ScaledObjectRef.ScalerMetadata[promQuery]

	val, _ := e.ExecutePromQuery(query)

	return &pb.GetMetricsResponse{
		MetricValues: []*pb.MetricValue{{
			MetricName:  metricName,
			MetricValue: int64(math.Round(val)),
		}},
	}, nil
}

func (e *ExternalScaler) StreamIsActive(scaledObject *pb.ScaledObjectRef, epsServer pb.ExternalScaler_StreamIsActiveServer) error {
	log.Infof("stream is active is called.")
	return nil
}

func (s *ExternalScaler) IsDownScalingPhase(numberOfworkers, currentValue, threshold int64) bool {
	return false
}

// ExecutePromQuery execute query on prometheus and result value
// in case -Inf or +Inf it return 0 value
func (s *ExternalScaler) ExecutePromQuery(query string) (float64, *promQueryError) {
	t := time.Now().UTC().Format(time.RFC3339)
	queryEscaped := url_pkg.QueryEscape(query)
	url := fmt.Sprintf("%s/api/v1/query?query=%s&time=%s", s.Metadata.ServerAddress, queryEscaped, t)
	log.Debugf("query %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
	}

	r, err := s.HttpClient.Do(req)
	if err != nil {
		return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
	}
	r.Body.Close()

	if !(r.StatusCode >= 200 && r.StatusCode <= 299) {
		return -1, &promQueryError{promError: fmt.Sprintf("prometheus query api returned error. status: %d response: %s", r.StatusCode, string(b)), numberOfRecords: 0}
	}

	var result promQueryResult
	err = json.Unmarshal(b, &result)
	if err != nil {
		return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
	}

	// allow for zero element or single element result sets
	if len(result.Data.Result) == 0 {
		return 0, nil
	} else if len(result.Data.Result) > 1 {

		value, err := retrieveValFromJson(result)
		if err != nil {
			return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
		}
		return -1, &promQueryError{promError: fmt.Sprintf("prometheus query %s returned multiple elements", query), numberOfRecords: len(result.Data.Result), firstValueResult: value}
	}

	return retrieveValFromJson(result)

}

func retrieveValFromJson(promResult promQueryResult) (float64, *promQueryError) {
	val := promResult.Data.Result[0].Value[1]
	var result float64 = -1
	var err error
	if val != nil {
		s := val.(string)
		result, err = strconv.ParseFloat(s, 64)
		if err != nil {
			log.Error(err, "error converting prometheus value", "prometheus_value", s)
			return -1, &promQueryError{promError: err.Error(), numberOfRecords: 0}
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
