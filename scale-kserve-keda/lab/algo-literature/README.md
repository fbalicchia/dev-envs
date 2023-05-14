#Predictive-scale-algo

This is an experimental that was created taking inspiration from [link](https://predictive-horizontal-pod-autoscaler.readthedocs.io/en/latest/)

The idea respecting predictive-horizontal-pod project is to understand if is possible to
encapsulate [Time Series Analysis model](www.statsmodels.org) in external gRPC service scaler
that can be used from [KEDA](https://keda.sh/)

The first step was brings algo as is with tests and run is without problems 