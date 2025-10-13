# gradio-ui/app.py
import gradio as gr
from graphql_client import create_story, get_stories, update_story, delete_story

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

def fetch_stories(limit):
    print(f"🔍 DEBUG: fetch_stories called with limit={limit}")
    stories = get_stories(limit)
    print(f"🔍 DEBUG: get_stories returned {len(stories) if isinstance(stories, list) else 'error'}")
    
    if isinstance(stories, str):
        return [[stories, "", "", "", "", ""]]

    if not stories:
        return [["No stories found", "", "", "", "", ""]]

    table_data = []
    for s in stories:
        table_data.append([
            s.get("id", ""),
            s.get("title", ""),
            s.get("category", ""),
            ", ".join(s.get("tags", [])) if s.get("tags") else "",
            s.get("mood", ""),
            s.get("createdAt", "")
        ])
    print(f"🔍 DEBUG: Returning {len(table_data)} rows")
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

# ✅ Gradio Interface
with gr.Blocks() as demo:
    gr.Markdown("# 🧠 Job Stories Portal")

    limit_state = gr.State(value=10)
    selected_id_state = gr.State(value="")
    
    # --- Submit Story Tab ---
    with gr.Tab("📝 Submit Story"):
        title = gr.Textbox(label="Title")
        content = gr.Textbox(label="Content", lines=5)
        tags_string = gr.Textbox(label="Tags", placeholder="e.g., ai, startup, remote work")
        category = gr.Textbox(label="Category", placeholder="e.g., bug fix, testing, database")

        mood = gr.Dropdown(
            label="Mood",
            choices=[
                "😍 excitement", "🚀 flow state euphoria", "😤 pride", "😌 satisfaction",
                "🏆 accomplishment", "😮‍💨 relief", "🤔 curiosity", "💪 confidence",
                "😠 frustration", "😵‍💫 confusion", "😰 overwhelm", "🤡 imposter syndrome",
                "😟 anxiety", "🔥 burnout", "😭 despair", "🤷‍♂️ self-doubt", "😤 determination",
                "⏰ impatience", "😴 boredom", "🙄 procrastination", "🔍 perfectionism",
                "😓 stress", "😔 loneliness", "😳 embarrassment",
            ],
        )

        submit = gr.Button("Submit")
        output = gr.Textbox(label="Confirmation", lines=6)

    # --- View Stories Tab ---
    with gr.Tab("👀 View Stories"):
        gr.Markdown("### Recently Created Stories")
        gr.Markdown("💡 **Tip:** Click on a Story ID to auto-fill the Update/Delete forms")
        
        gr.Markdown("### Step 1: Click an ID in the table above")
        gr.Markdown("### Step 2: Choose where to send it")
        with gr.Row():
            send_to_update_btn = gr.Button("📝 Send to Update", scale=1)
            send_to_delete_btn = gr.Button("🗑️ Send to Delete", scale=1)
        selection_status = gr.Textbox(label="Status", interactive=False)
        
        limit = gr.Number(
            value=10,
            minimum=1,
            maximum=50,
            step=1,
            label="Number of Stories"
        )
        
        fetch_btn = gr.Button("Fetch Stories", scale=1)
        table = gr.Dataframe(
            headers=["ID", "Title", "Category", "Tags", "Mood", "Created At"],
            interactive=False,
            wrap=True,
            value=fetch_stories(10),  # initial display
        )

    # --- Update Story Tab ---
    with gr.Tab("📝 Update Story"):
        gr.Markdown("### Update Existing Story")
        gr.Markdown("💡 *Copy a Story ID from the View Stories tab*")
        
        update_id = gr.Text(label="Story ID *", placeholder="e.g., 7b12c3df-8ebe-4d7a-aaa3-3d06ad40b576")
        update_title = gr.Text(label="New Title (optional)")
        update_content = gr.Textbox(label="New Content (optional)", lines=5)
        update_tags = gr.Textbox(label="New Tags (optional)", placeholder="e.g., ai, startup")
        update_category = gr.Textbox(label="New Category (optional)")
        update_mood = gr.Dropdown(
            label="New Mood (optional)",
            allow_custom_value=True,
            choices=[
                "",
                "😍 excitement", "🚀 flow state euphoria", "😤 pride", "😌 satisfaction",
                "🏆 accomplishment", "😮‍💨 relief", "🤔 curiosity", "💪 confidence",
                "😠 frustration", "😵‍💫 confusion", "😰 overwhelm", "🤡 imposter syndrome",
                "😟 anxiety", "🔥 burnout", "😭 despair", "🤷‍♂️ self-doubt",
                "⏰ impatience", "😴 boredom", "🙄 procrastination", "🔍 perfectionism",
                "😓 stress", "😔 loneliness", "😳 embarrassment",
            ],
        )
        
        update_btn = gr.Button("Update Story")
        update_output = gr.Textbox(label="Confirmation", lines=6)
    
    # --- Delete Story Tab ---
    with gr.Tab("🗑️ Delete Story"):
        gr.Markdown("### Delete Existing Story")
        gr.Markdown("⚠️ *This action cannot be undone!*")
        
        delete_id = gr.Text(label="Story ID *", placeholder="e.g., 7b12c3df-8ebe-4d7a-aaa3-3d06ad40b576")
        delete_btn = gr.Button("Delete Story", variant="stop")
        delete_output = gr.Textbox(label="Confirmation", lines=6)
    
    # === ALL EVENT HANDLERS ===
    submit.click(
        fn=submit_story,
        inputs=[title, content, tags_string, category, mood],
        outputs=[output, table, title, content, tags_string, category, mood]
    )
    
    update_btn.click(
        fn=change_story,
        inputs=[update_id, update_title, update_content, update_tags, update_category, update_mood],
        outputs=[update_output, table]
    )
    
    delete_btn.click(
        fn=remove_story,
        inputs=[delete_id],
        outputs=[delete_output, table]
    )
    
    fetch_btn.click(
        fn=fetch_stories,
        inputs=[limit],
        outputs=[table]
    )
    
    table.select(
        fn=get_selected_id,
        outputs=[selected_id_state, selection_status]
    )
    
    send_to_update_btn.click(
        fn=lambda id: (id, f"✅ Sent to Update form: {id}"),
        inputs=[selected_id_state],
        outputs=[update_id, selection_status]
    )
    
    send_to_delete_btn.click(
        fn=lambda id: (id, f"✅ Sent to Delete form: {id}"),
        inputs=[selected_id_state],
        outputs=[delete_id, selection_status]
    )
    

if __name__ == "__main__":
    demo.launch(server_name="0.0.0.0", server_port=4103)