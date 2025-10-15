# gradio-ui/app.py
from gradio_ui.interface import build_interface

if __name__ == "__main__":
    demo = build_interface()
    demo.launch(server_name="0.0.0.0", server_port=4103)