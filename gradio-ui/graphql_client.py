import requests
import os

GQL_ENDPOINT = os.getenv("GRAPHQL_ENDPOINT", "http://gateway:4100/graphql")

def create_story(input_data):
    mutation = """ 
    mutation CreateStory($input: CreateStoryInput!){
        createStory(input: $input){
            id
            title
            content
            tags
            category
            mood
            createdAt
        }
    }
    """
    try:
        response = requests.post(
            GQL_ENDPOINT,
            json={"query":mutation, "variables": {"input": input_data}},
            headers={"Content-Type": "application/json"}
        )
        response.raise_for_status()
        data = response.json()
        
        if "errors" in data:
            return {"error": data["errors"][0]["message"]}
        return data["data"]
    
    except Exception as e:
        return {"error": f"❌ Failed to submit story: {e}"}
   
    
def get_stories(limit=10, offset=0):
    query = """
    query GetStories($limit: Int, $offset: Int) {
        stories(limit: $limit, offset: $offset) {
            id
            title
            category
            mood
            tags
            createdAt
        }
    }
    """
    
    variables = {"limit": limit, "offset": offset}
        
    try:
        response = requests.post(
            GQL_ENDPOINT, 
            json={"query": query, 
                  "variables": variables},
            headers={"Content-Type": "application/json"},
            )
        response.raise_for_status()
        data = response.json()

        if "errors" in data:
            return f"⚠️ GraphQL Error: {data['errors'][0]['message']}"

        stories = data.get("data", {}).get("stories", [])
        return stories

    except Exception as e:
        return f"❌ Failed to fetch stories: {e}"