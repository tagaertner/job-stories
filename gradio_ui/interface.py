import gradio as gr
from .handlers import (
    submit_story, fetch_stories, change_story, 
    get_selected_id, view_full_story, remove_story
)


# === ✅ Gradio Interface ===# 
# todo Track usage history (e.g., how many logs created)
def build_interface():
    
    with gr.Blocks() as demo:
        gr.Markdown("# 🧠 Job Stories Portal")

        limit_state = gr.State(value=10)
        selected_id_state = gr.State(value="")

        # --- 📝 Submit Story Tab ---
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
            
        # === 👀 View Stories Tab ===# 
        with gr.Tab("👀 View Stories"):
            gr.Markdown("### Recently Created Stories")
            gr.Markdown("💡 **Tip:** Click on a Story ID to auto-fill the Update/Delete forms")

            gr.Markdown("### Step 1: Click an ID in the table below")
            gr.Markdown("### Step 2: Choose where to send it")
            gr.Markdown("### Step 3: View Full Story")
            gr.Markdown("👉 Click a Story ID in the table below first, then click the button below to view full story.")

            with gr.Row():
                send_to_update_btn = gr.Button("📝 Send to Update", scale=1)
                send_to_delete_btn = gr.Button("🗑️ Send to Delete", scale=1)
                selection_status = gr.Textbox(label="Status", interactive=False)

            # Date range filters
            with gr.Row():
                from_date = gr.Textbox(label="From (MM-DD-YYYY)", placeholder="10-01-2025")
                to_date = gr.Textbox(label="To (MM-DD-YYYY)", placeholder="10-13-2025")
                
            search_box = gr.Textbox(label="Search Stories", placeholder="Search in title or content (e.g., backend, debugging)")

            # Pagination controls
            limit = gr.Number(
                value=10,
                minimum=1,
                maximum=50,
                step=1,
                label="Number of Stories"
            )

            fetch_btn = gr.Button("Fetch Stories", scale=1)

            # Stories table
            table = gr.Dataframe(
                headers=["ID", "Title", "Category", "Tags", "Mood", "Created At"],
                interactive=False,
                wrap=True,
                value=[],
            )

            # 👁️ View full story section
            gr.Markdown("### Step 3: View Full Story")
            gr.Markdown("👉 Click a Story ID in the table above, then click the button below to view its full content.")

            with gr.Row():
                view_full_btn = gr.Button("👁️ View Full Story", scale=1)
            full_story_box = gr.Textbox(label="Full Story Content", lines=10, interactive=False)

        # === 📝 Update Story Tab ===# 
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

        # --- 🗑️ Delete Story Tab ---
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
            inputs=[limit,from_date, to_date,search_box],
            outputs=[table]
        )
        
        
        limit.change(
            fn=fetch_stories,
            inputs=[limit, from_date, to_date, search_box],
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
        
        view_full_btn.click(
            fn=view_full_story, 
            inputs=[selected_id_state],
            outputs=[full_story_box]
        )

        demo.load(
            fn=lambda:fetch_stories(10),
            outputs=[table]
        )

    return demo 