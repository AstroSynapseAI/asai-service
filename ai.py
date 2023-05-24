import sys
import os
import configparser
import json


from langchain import OpenAI, SerpAPIWrapper
from langchain.agents import initialize_agent, Tool
from flask import Flask, request, jsonify

app = Flask(__name__)

config = configparser.ConfigParser()
config.read("/app/config.ini")

os.environ["OPENAI_API_KEY"] = config.get("openai", "key")
os.environ["SERPAPI_API_KEY"] = config.get("serp_api", "key")

llm = OpenAI(temperature=0.6)
search = SerpAPIWrapper()
tools = [
    Tool(
        name="Intermediate Answer",
        description="A tool to search for intermediate answers",
        func=search.run
    )
]

@app.route('/test', methods=['POST'])

def test():
    output = {"content": "Test ping is coming back!"}
    return jsonify(output)

@app.route('/prompt', methods=['POST'])

def prompt():
    data = request.get_json()
    user_prompt = data['user_prompt']
    history = data['history']

    query = user_prompt
    self_ask_with_search = initialize_agent(tools, llm, agent="self-ask-with-search", verbose=False)
    response = self_ask_with_search.run(query)

    output = {"content": response}
    return jsonify(output)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

