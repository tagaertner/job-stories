# test_update.py
from graphql_client import update_story

# Test updating a story
result = update_story({
    "id": "7b12c3df-8ebe-4d7a-aaa3-3d06ad40b576",  # Replace with a real ID
    "userId": "demo-user-123",
    "title": "Testing Update Function"
})

print("Result:", result)