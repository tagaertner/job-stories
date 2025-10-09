import requests
import os

GQL_ENDPOINT = os.getenv("GRAPHQL_ENDPOINT", "http://gateway:4100/graphql")

def create_story(title, content, tags, category, mood):
    mutation = """ 
    mutation CreateStroy($input: CreateStoryInput!){
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
    
    variables = {
        "input":{
            "title": title,
            "content": content,
            "tags": tags,
            "category": category,
            "mood": mood,
        }
    }
    
    try:
        response = requests.post(GQL_ENDPOINT, json={"query": mutation, "variables": variables})
        response.raise_for_status()
        data = response.json()

        if "errors" in data:
            return f"⚠️ GraphQL Error: {data['errors'][0]['message']}"

        story = data["data"]["createStory"]
        return f"✅ Story created!\n\nTitle: {story['title']}\nCategory: {story['category']}\nMood: {story['mood']}"

    except Exception as e:
        return f"❌ Failed to submit story: {e}"
    
def get_stories(limit=10, after=None):
    query ="""
    query GetStories($limit: Int, $after: String) {
        stories(first: $limit, after: $after){
            id
            title
            content
            tags
            category
            mood
            createAt
        }
    }
    """
    
    variables = {"limit": limit}
    if after:
        variables["after"] = after
        
    try:
        response = requests.post(GQL_ENDPOINT, json={"query": query, "variables": variables})
        response.raise_for_status()
        data = response.json()

        if "errors" in data:
            return f"⚠️ GraphQL Error: {data['errors'][0]['message']}"

        stories = data.get("data", {}).get("stories", [])
        return stories

    except Exception as e:
        return f"❌ Failed to fetch stories: {e}"