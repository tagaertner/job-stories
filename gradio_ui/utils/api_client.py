import requests

GATEWAY_URL = "http://gateway:4100/query"

def send_graphql_query(query: str, variables: dict = None):
    response = requests.post(
        GATEWAY_URL,
        json={"query": query, "variables": variables or {}},
        headers={"Content-Type": "application/json"}
    )
    response.raise_for_status()
    return response.json()