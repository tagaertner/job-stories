import gradio as gr
from datetime import datetime
from graphql_client import create_story, get_stories, update_story, delete_story, get_story_by_id

# TODO (JWT Integration Reminder):
# ---------------------------------------------------------
# Currently using a hardcoded userId ("demo-user-123") 
# to satisfy GraphQL schema requirements.
#
# Once authentication is implemented:
#   - Remove the `userId` field from this payload.
#   - Extract the userId dynamically from the JWT or 
#     session context (depending on your auth setup).
#   - Update the Go GraphQL resolver to pull userId 
#     from the context (not from input).
#
# Example later:
#   userId = get_user_from_jwt(request.headers["Authorization"])
#   variables["input"]["userId"] = userId
# ---------------------------------------------------------

def submit_story(title, content, tags_string, category, mood):
    tags_list = [t.strip() for t in tags_string.split(",") if t.strip()]
    variables = {
        "input": {
            "title": title,
            "content": content,
            "tags": tags_list,
            "category": category,
            "mood": mood,
            "userId": "demo-user-123"  # temporary
        }
    }

    result = create_story(variables["input"])
   
    if result is None:
        return "❌ Error: No response from GraphQL server", fetch_stories(10), title, content, tags_string, category, mood
    
    if "error" in result:
        return f"❌ Error: {result['error']}", fetch_stories(10), title, content, tags_string, category, mood
    
    if "createStory" not in result:
        return f"❌ Error: Unexpected response format: {result}", fetch_stories(10), title, content, tags_string, category, mood
    
    story = result["createStory"]
    message = f"✅ Story Created!\n\nID: {story['id']}\nTitle: {story['title']}\nCategory: {story['category']}\nMood: {story['mood']}"

    return message, fetch_stories(10), "", "", "", "", ""

def fetch_stories(limit, from_date=None, to_date=None, search_text=None):
    stories = get_stories(limit, from_date=from_date, to_date=to_date, search_text=search_text)
    
    if isinstance(stories, str):
        return [[stories, "", "", "", "", ""]]

    if not stories:
        return [["No stories found", "", "", "", "", ""]]

    table_data = []
    for s in stories:
        created_raw = s.get("createdAt", "")
        created_formatted = ""
        if created_raw:
            try:
                created_formatted = datetime.fromisoformat(created_raw.replace("Z", "+00:00")).strftime("%m-%d-%Y")
            except:
                created_formatted = created_raw
                
        content_raw = s.get("content", "") or ""
        content_preview = (content_raw[:80] + "…") if len(content_raw) > 100 else content_raw
                   
        table_data.append([
            s.get("id", ""),
            s.get("title", ""),
            content_preview,
            s.get("category", ""),
            ", ".join(s.get("tags", [])) if s.get("tags") else "",
            s.get("mood", ""),
            created_formatted
        ])
        
    return table_data

def change_story(story_id, title, content, tags_string, category, mood):
    # Validate
    if not story_id:
        return "❌ Error: Story ID is required", fetch_stories(10)
    
    # Parse tags
    tags_list = [t.strip() for t in tags_string.split(",") if t.strip()] if tags_string else []
    
    # Build input data
    input_data = {
        "id": story_id,
        "userId": "demo-user-123",
    }
   
    if title:
        input_data["title"] = title
    if content:
        input_data["content"] = content
    if tags_list:
        input_data["tags"] = tags_list
    if category:
        input_data["category"] = category 
    if mood:
        input_data["mood"] = mood
    result = update_story(input_data)
    
    # Handle response
    if "error" in result:
        return f"❌ Error: {result['error']}", fetch_stories(10)
    
    if "updateStory" not in result or result["updateStory"] is None:
        return "❌ Story not found", fetch_stories(10)
    
    story = result["updateStory"]
    return f"✅ Updated!\n\nTitle: {story['title']}\nUpdated: {story['updatedAt']}", fetch_stories(10)

def get_selected_id(event_data: gr.SelectData):
    # Check which column was clicked
    col_index = event_data.index[1]
    
    # only process if ID column 0 was clicked
    if col_index != 0:
        return "", ""
    story_id = event_data.value
    return story_id, f"📋 Selected: {story_id}\n\n👉 Now click 'Send to Update' or 'Send to Delete'"

def view_full_story(story_id):
    if not story_id:
        return "⚠️ Please click a Story ID in the table first, then try again."
    
    # Fetch just that story from GraphQL
    from graphql_client import get_story_by_id  # 👈 you’ll add this function below
    
    story = get_story_by_id(story_id)
    if not story or "error" in story:
        return f"❌ Could not load story: {story.get('error', 'Unknown error')}"
    
    content = story.get("content", "(No content)")
    return f"🧠 {story.get('title', '')}\n\n{content}"

def remove_story(story_id):
    # Validate
    if not story_id:
        return "❌ Error: Story ID is required", fetch_stories(10)
    
    # Build input data
    input_data = {
        "id": story_id,
    }
    result = delete_story(input_data)
    
    # Handle response
    if "error" in result:
        return f"❌ Error: {result['error']}", fetch_stories(10)
    
    if result.get("deleteStory") == True:
        return f"✅ Story ID:{story_id} was deleted!", fetch_stories(10)
    else:
        return f"❌ Story not found or could not be deleted", fetch_stories(10)
    
# Todo add user handlers