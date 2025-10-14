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
   
    
def get_stories(limit=10, offset=0, from_date=None, to_date=None):
    query = """
    query GetStories($filter: StoryFilter, $limit: Int, $offset: Int) {
        stories(filter: $filter, limit: $limit, offset: $offset) {
            id
            title
            content
            category
            mood
            tags
            createdAt
        }
    }
    """
    
    variables = {
        "filter":{},
        "limit": limit, 
        "offset": offset
        }
    
    
    if from_date:
        variables["filter"]["dateFrom"] = from_date
    if to_date:
        variables["filter"]["dateTo"] = to_date
        
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
    
def get_story_by_id(story_id):
    query = """
    query GetStory($id: ID!) {
        story(id: $id) {
            id
            title
            content
            category
            mood
            tags
            createdAt
        }
    }
    """
    try:
        response = requests.post(
            GQL_ENDPOINT,
            json={"query": query, "variables": {"id": story_id}},
            headers={"Content-Type": "application/json"}
        )
        response.raise_for_status()
        data = response.json()

        if "errors" in data:
            return {"error": data["errors"][0]["message"]}
        
        return data.get("data", {}).get("story", {})
    except Exception as e:
        return {"error": str(e)}
    
    
# todo get_stories_by_id related to load storyies
def update_story(input_data):
    mutation = """ 
    mutation UpdateStory($input: UpdateStoryInput!){
        updateStory(input: $input){
            id
            userId
            title
            content
            tags
            category
            mood
            createdAt
            updatedAt
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
        return {"error": f"❌ Failed to update story: {e}"}

# todo delete
def delete_story(input_data):
    mutation = """ 
    mutation DeleteStory($input: DeleteStoryInput!){
        deleteStory(input: $input)
    
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
        return {"error": f"❌ Failed to delete story: {e}"}