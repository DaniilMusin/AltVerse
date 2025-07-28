import joblib
from pathlib import Path
import numpy as np

_model = joblib.load(Path(__file__).with_suffix('.pkl'))

def predict(vec) -> float:
    x = np.asarray(vec).reshape(1, -1)
    score = _model.predict_proba(x)[0, 1]
    return float(score)
