import gradio as gr
from graphql_client import create_story


def submit_story(title, content, tags_string, category, mood):
    tags_list = [tag.strip() for tag in tags_string.split(',')]
    
    
with gr.Blocks() as demo:
    gr.Markdown("## ✍️ Submit a Job Story")
    
    title = gr.Textbox(label="Title")
    content = gr.Textbox(label="Content", lines=5)
    tags_string = gr.Textbox(label="Tags", placeholder="e.g., ai, startup, remote work")
    category = gr.Textbox(label="Category", placeholder="e.g., bug fix, testing, database")
    mood = gr.Dropdown(label="Mood",
    choices=[
    "😍 excitement",
    "🚀 flow state euphoria", 
    "😤 pride",
    "😌 satisfaction",
    "🏆 accomplishment",
    "😮‍💨 relief",
    "🤔 curiosity",
    "💪 confidence",
    "😠 frustration",
    "😵‍💫 confusion",
    "😰 overwhelm",
    "🤡 imposter syndrome",
    "😟 anxiety",
    "🔥 burnout",
    "😭 despair",
    "🤷‍♂️ self-doubt",
    "😤 determination",
    "⏰ impatience",
    "😴 boredom",
    "🙄 procrastination",
    "🔍 perfectionism",
    "😓 stress",
    "😔 loneliness",
    "😳 embarrassment",
])
    submit = gr.Button("Submit")
    output = gr.Textbox(label="Confirmation", lines=6)
    
    submit.click(fn=submit_story, inputs=[title, content, tags_string, category, mood], outputs=output)
    
if __name__ == "__main__":
    demo.launch(server_name="0.0.0.0", server_port=4103)