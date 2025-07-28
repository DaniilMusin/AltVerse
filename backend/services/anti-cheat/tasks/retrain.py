import joblib, pandas as pd
from pathlib import Path
from catboost import CatBoostClassifier

DATA_PATH = Path("/data/cheat_events.parquet")
MODEL_OUT = Path(__file__).with_suffix(".pkl")


def load_dataset():
    df = pd.read_parquet(DATA_PATH)
    X = df.drop(columns=["is_cheat"])
    y = df["is_cheat"]
    return X, y


def retrain():
    X, y = load_dataset()
    model = CatBoostClassifier(
        iterations=300, depth=6, loss_function="Logloss", verbose=False
    )
    model.fit(X, y)
    joblib.dump(model, MODEL_OUT)


if __name__ == "__main__":
    retrain()
