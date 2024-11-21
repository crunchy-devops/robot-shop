from flask import Flask, jsonify
import redis
import os

app = Flask(__name__)

# Redis connection settings from environment variables
REDIS_HOST = os.getenv("REDIS_HOST", "localhost")
REDIS_PORT = int(os.getenv("REDIS_PORT", 6379))

# Connect to Redis
redis_client = redis.StrictRedis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)

@app.route('/')
def index():
    # Test Redis connection
    redis_client.set("message", "Hello from Redis!")
    message = redis_client.get("message")
    return jsonify({"message": message})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
