package main

import (
	pb "external-scaler/externalscaler"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime/debug"

	"net/http"
	"time"

	scaler "external-scaler/pkg/scaler"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var (
	metricServerURL   string
	connectionTimeout int
	port              string
	logSeverity       string
	cacheExpriedDays  int
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {

	//enable gc please to remainder please refer to
	//https://pkg.go.dev/runtime/debug#SetGCPercent
	debug.SetGCPercent(10)

	flag.StringVar(&metricServerURL, "url", "http://prometheus-operated.metrics.svc.cluster.local:9090", "Prometheus server url")
	flag.IntVar(&connectionTimeout, "timeout", 10000, "Connection timeout")
	flag.StringVar(&port, "port", "6000", "The server port")
	flag.StringVar(&logSeverity, "logLevel", "info", "Log serverity from panic to  trace")
	flag.IntVar(&cacheExpriedDays, "expiredDays", 1, "1")

	flag.Parse()

	level, err := log.ParseLevel(logSeverity)
	if err != nil {
		fmt.Println("error parameter")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}

	grpcServer := grpc.NewServer()
	log.Info("Starting server with the following parameters")
	log.Info("Port: ", port)
	log.Info("ConnectionTimeout: ", connectionTimeout)
	log.Info("Prometheus url: ", metricServerURL)
	log.Info("logLevel: ", logSeverity)
	log.Info("cache expiredDays: ", cacheExpriedDays)
	lis, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	timeout := time.Duration(connectionTimeout) / time.Millisecond
	httpClient := &http.Client{
		Timeout: timeout,
	}

	meta := scaler.PrometheusMetadata{}

	meta.ServerAddress = metricServerURL

	pb.RegisterExternalScalerServer(grpcServer, &scaler.ExternalScaler{
		HttpClient: httpClient,
		Metadata:   &meta,
	})

	go func() {
		log.Info("Listening on:", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// Shutdown when we receive Ctrl+c (interrupt)
	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c
}
