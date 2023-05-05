import kserve
from kserve import ModelServer
import math
from typing import Dict
import json
import logging
import time
from datetime import datetime, timedelta
import statsmodels.api as sm
from dataclasses import dataclass
from dataclasses_json import dataclass_json, LetterCase
from typing import List, Optional
logging.basicConfig(level=kserve.constants.KSERVE_LOGLEVEL)


def custom_predict(instance):
    json_instance = json.dumps(instance, indent = 4)
    algorithm_input = AlgorithmInput.from_json(json_instance)
    if algorithm_input.current_time is not None:
        try:
            current_time = datetime.strptime(algorithm_input.current_time, "%Y-%m-%dT%H:%M:%SZ")
        except ValueError as ex:
            logging.info("ready!")
            logging.error(f"Invalid datetime format: {str(ex)}")
            return -1
    search_time = datetime.timestamp(current_time + timedelta(milliseconds=int(algorithm_input.look_ahead)))
    x = []
    y = []
    # Build up data for linear model, in order to not deal with huge values and get rounding errors, use the difference
    # between the time being searched for and the metric recorded time in seconds
    for i, timestamped_replica in enumerate(algorithm_input.replica_history):
        try:
            created = datetime.strptime(timestamped_replica.time, "%Y-%m-%dT%H:%M:%SZ")
        except ValueError as ex:
            logging.error(f"Invalid datetime format: {str(ex)}")
            return -1

        x.append(search_time - datetime.timestamp(created))
        y.append(timestamped_replica.replicas)
    # Add constant for OLS, constant is 1.0
    x = sm.add_constant(x)
    model = sm.OLS(y, x).fit()
    return math.ceil(model.predict([[1, 0]])[0])


# for REST predictor the preprocess handler converts to input dict to the v1 REST protocol dict
class OLSModel(kserve.Model):
    def __init__(self, name: str):
        super().__init__(name)
        self.name = name
        self.ready = False

    def load(self):
        logging.info("Simulate loading model..")
        time.sleep(1)
        self.ready = True
        logging.info("ready!")

    def predict(self, payload: Dict, headers: Dict[str, str] = None) -> Dict:
        return {"predictions": [custom_predict(instance) for instance in payload['instances']]}


@dataclass_json(letter_case=LetterCase.CAMEL)
@dataclass
class TimestampedReplica:
    """
    JSON data representation of a timestamped evaluation
    """
    time: str
    replicas: int

@dataclass_json(letter_case=LetterCase.CAMEL)
@dataclass
class AlgorithmInput:
    """
    JSON data representation of the data this algorithm requires to be provided to it.
    """
    look_ahead: int
    replica_history: List[TimestampedReplica]
    current_time: Optional[str] = None



if __name__ == "__main__":
    model = OLSModel("linear-regression-model")
    model.load()
    ModelServer(workers=1).start(models=[model])
