from flask import Flask, request

app = Flask(__name__)

# Text - GET
@app.get("/text")
def text_get():
    return "hello world", 200

# JSON - GET
@app.get("/json")
def json_get():
    return { "hello": "world" }, 200

# JSON - POST
@app.post("/json")
def json_post():
    hello = request.json["hello"]
    return { "hello": hello }, 200