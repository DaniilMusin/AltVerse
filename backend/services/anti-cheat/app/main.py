from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from models.inference import predict
import uvicorn

app = FastAPI(title="Anti-Cheat API")

class FeatureVector(BaseModel):
    player_id: str
    features: list[float]

@app.post("/predict")
def detect_cheat(vec: FeatureVector):
    score = predict(vec.features)
    return {"player_id": vec.player_id, "risk": score}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
