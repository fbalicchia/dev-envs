# Predictive HPA

This sub-folder contains experiments where is implemented a basic Predictive autoscale based on a linear regression model.

At the moment of writing Solution based on cloud-native components like [Keda](https://keda.sh/) to control HPA
with the support of metrics collectors like [Prometheus](https://prometheus.io/).

A custom scaler based on metrics received from Prometheus calls an inference service based on kserve and the result provided by the inference instructs HPA to scale or downscale.

[Kserve](https://kserve.github.io/website/0.10/) provides a simple  custom model based exposed via REST API build on top of [OLS](https://www.statsmodels.org/stable/generated/statsmodels.formula.api.ols.html#statsmodels.formula.api.ols)

## Subfolder layout

├── README.md
├── kind-with-istio.sh (Script the prepare infrastructure basd on k8s )
├── lab (Custom model that need to be deployed on k8s)
├── target-app (Simple application that expose metrics and need to be scaled)
├── external-scaler (Keda external scaler that link inferences service, Metrics and HPA)

-


