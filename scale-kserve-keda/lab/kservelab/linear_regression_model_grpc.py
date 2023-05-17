import kserve
from kserve import InferRequest, InferResponse, InferOutput, Model, ModelServer
from typing import Dict
import logging
import time
import numpy as np
from kserve.utils.utils import generate_uuid
logging.basicConfig(level=kserve.constants.KSERVE_LOGLEVEL)



class OLSModel(Model):
    def __init__(self, name: str):
        super().__init__(name)
        self.name = name
        self.ready = False

    def load(self):
        logging.info("Simulate loading model..")
        time.sleep(1)
        self.ready = True
        logging.info("ready!")

    def predict(self, payload: InferRequest, headers: Dict[str, str] = None) -> InferResponse:
        req = payload.inputs[0]
        logging.info("request predict %s", req)
        values = np.array([1, 2, 3, 4])
        values.shape
        result = values.flatten().tolist()
        response_id = generate_uuid()
        infer_output = InferOutput(name="output-0", shape=list(values.shape), datatype="FP32", data=result)
        infer_response = InferResponse(model_name=self.name, infer_outputs=[infer_output], response_id=response_id)
        return infer_response


if __name__ == "__main__":
    model = OLSModel("linear_regression_model_grpc")
    model.load()
    ModelServer().start([model])
