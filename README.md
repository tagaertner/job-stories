### 🧠 StackTrack: A lightweight developer journal to log wins, blockers, and growth over time.

## Project Structure

```
stack-track/
├── docker-compose.yml
├── go.mod
├── go.sum
├── gqlgen.yml # (if using Go gqlgen)
├── .env # local environment variables
├── .gitignore

├── gateway/ # Apollo Federation gateway (Node.js)
│ ├── gateway.js
│ ├── dockerfile
│ └── package.json

├── models/ # Shared Go structs (optional)
│ └── job_story.go

├── pkg/ # Optional: common federation helpers, utilities
│ └── federation.go

├── generated/ # gqlgen output (Go)
│ ├── generated.go
│ └── resolver.go

├── database/
│ └── init/ # Seed files for Postgres
│ ├── 01-schema.sql
│ └── 02-seed-job-stories.sql

├── services/
│ ├── job-stories/ # Go microservice for job story CRUD
│ │ ├── main.go
│ │ ├── schema.graphql
│ │ ├── resolvers/
│ │ ├── models/
│ │ ├── services/
│ │ ├── database/
│ │ └── dockerfile
│
│ ├── ai-summary/ # Python microservice for Vertex AI calls
│ │ ├── main.py
│ │ ├── vertex_ai_client.py
│ │ ├── summarizer.py
│ │ ├── Dockerfile
│ │ └── requirements.txt
│
│ └── gradio-ui/ # Gradio app for user-facing UI
│ ├── app.py
│ ├── components/
│ │ ├── story_input.py
│ │ └── interview_practice.py
│ ├── utils/
│ │ └── formatters.py
│ ├── Dockerfile
│ └── requirements.txt

└── README.md
```
