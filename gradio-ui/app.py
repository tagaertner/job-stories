# gradio-ui/app.py
import gradio as gr
from graphql_client import create_story, get_stories, update_story

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
        return "❌ Error: No response from GraphQL server", fetch_stories(10)
    
    if "error" in result:
        return f"❌ Error: {result['error']}",fetch_stories(10), title, content, tags_string, category, mood  
    
    
    if "createStory" not in result:
        return f"❌ Error: Unexpected response format: {result}",fetch_stories(10), title, content, tags_string, category, mood  # ✅ Now 7 values
    
    
    story = result["createStory"]
    message = f"✅ Story Created!\n\nID: {story['id']}\nTitle: {story['title']}\nCategory: {story['category']}\nMood: {story['mood']}"

    return message, fetch_stories(10), "", "", "", "", ""

def fetch_stories(limit):
    stories = get_stories(limit)

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
    return table_data

# todo update story
def change_story(story_id,title, content, tags_string, category, mood):
    # Validate
    if not story_id:
        return "❌ Error: Story ID is required", fetch_stories(10)
    
    # Parce tags
    tags_list = [t.strip() for t in tags_string.split(",") if t.strip()] if tags_string else[]
    
    # Build input data
    input_data ={
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
    
# todo delete story

# ✅ 
with gr.Blocks() as demo:
    gr.Markdown("# 🧠 Job Stories Portal")

    limit_state = gr.State(value=10)
    
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
        output = gr.Textbox(label="Response", lines=6)
        # submit.click(fn=submit_story, inputs=[title, content, tags_string, category, mood], outputs=output)

    # --- View Stories Tab ---
    with gr.Tab("👀 View Stories"):
        gr.Markdown("### Recently Created Stories")
        limit = gr.Slider(1, 50, value=10, step=1, label="Number of Stories")
        fetch_btn = gr.Button("Fetch Stories")
        table = gr.Dataframe(
            headers=["ID", "Title", "Category", "Tags", "Mood", "Created At"],
            datatype=["str", "str", "str", "str", "str", "str"],
            interactive=False,
            wrap=True,
            value=fetch_stories(10),
        )
        fetch_btn.click(fetch_stories, inputs=[limit], outputs=[table])
        limit.change(lambda x: x, inputs=[limit], outputs=[limit_state])
    
        # submit = gr.Button("Submit")
        # output = gr.Textbox(label="Response", lines=6)     
    
    #--- Update Story Tab ---
    with gr.Tab("📝 Update Story"):
        gr.Markdown("### Update Existing Story")
        gr.Markdown("💡 *Copy a Story ID form the View Stories tab*")
        
        update_id = gr.Text(label="Story ID *", placeholder="e.g., 7b12c3df-8ebe-4d7a-aaa3-3d06ad40b576")
        update_title = gr.Text(label="New Title(optiona)")
        update_content = gr.Textbox(label="New Content (optional)", lines=5)
        update_tags = gr.Textbox(label="New Tags (optional)", placeholder="e.g., ai, startup")
        update_category = gr.Textbox(label="New Category (optional)")
        update_mood = gr.Dropdown(
            label="New Mood (optional)",
            choices=[
                "",  # Empty = no change
                "😍 excitement", "🚀 flow state euphoria", "😤 pride", "😌 satisfaction",
                "🏆 accomplishment", "😮‍💨 relief", "🤔 curiosity", "💪 confidence",
                "😠 frustration", "😵‍💫 confusion", "😰 overwhelm", "🤡 imposter syndrome",
                "😟 anxiety", "🔥 burnout", "😭 despair", "🤷‍♂️ self-doubt",
                "⏰ impatience", "😴 boredom", "🙄 procrastination", "🔍 perfectionism",
                "😓 stress", "😔 loneliness", "😳 embarrassment",
            ],
        )
        
        update_btn = gr.Button("Update Story")
        update_output = gr.Textbox(label="Response", lines=6)
    
    submit.click(
    fn=submit_story,
    inputs=[title, content, tags_string, category, mood],
    outputs=[
        output,        # status message
        table,         # updated stories list
        title,         # clear title
        content,       # clear content
        tags_string,   # clear tags
        category,      # clear category
        mood           # clear mood
    ]
)
    
      
    update_btn.click(
        fn=change_story,
        inputs=[update_id, update_title, update_content, update_tags, update_category, update_mood],
        outputs=[update_output, table]  
    )
        
if __name__ == "__main__":
    demo.launch(server_name="0.0.0.0", server_port=4103)