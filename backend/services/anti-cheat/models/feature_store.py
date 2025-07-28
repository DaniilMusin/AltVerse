import redis
import json

r = redis.Redis.from_url("redis://redis:6379/1")

def save_feature_vector(pid: str, vec: list[float]) -> None:
    r.setex(f"fv:{pid}", 3600, json.dumps(vec))
