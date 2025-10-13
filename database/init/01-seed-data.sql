-- Wait for users and job_stories tables to exist
DO $$ 
BEGIN
    WHILE NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') LOOP
        PERFORM pg_sleep(1);
    END LOOP;

    WHILE NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'job_stories') LOOP
        PERFORM pg_sleep(1);
    END LOOP;
END $$;

-- Seed users with proper UUIDs
INSERT INTO users (id, name, email, role, active) VALUES 
('00000000-0000-0000-0000-000000000001', 'John Doe', 'john@example.com', 'writer', true),
('00000000-0000-0000-0000-000000000002', 'Jane Smith', 'jane@example.com', 'editor', true),
('00000000-0000-0000-0000-000000000003', 'Bob Wilson', 'bob@example.com', 'writer', true),
('00000000-0000-0000-0000-000000000004', 'Alice Johnson', 'alice@example.com', 'editor', true),
('00000000-0000-0000-0000-000000000005', 'Mike Chen', 'mike@example.com', 'writer', true),
('00000000-0000-0000-0000-000000000006', 'Sarah Wilson', 'sarah@example.com', 'reader', true),
('00000000-0000-0000-0000-000000000007', 'David Brown', 'david@example.com', 'reader', false),
('00000000-0000-0000-0000-000000000008', 'Emma Davis', 'emma@example.com', 'writer', true),
('00000000-0000-0000-0000-000000000009', 'James Miller', 'james@example.com', 'editor', true),
('00000000-0000-0000-0000-000000000010', 'Lisa Garcia', 'lisa@example.com', 'writer', true)
ON CONFLICT (id) DO NOTHING;

-- Wait for job_stories table
DO $$
BEGIN
    WHILE NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'job_stories') LOOP
        PERFORM pg_sleep(1);
    END LOOP;
END $$;

-- Insert job stories with proper UUIDs
INSERT INTO job_stories (id, user_id, title, content, tags, category, mood, created_at, updated_at) VALUES
  ('10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', 'Fix login bug', 'Fixed OAuth issue during login.', ARRAY['auth', 'bugfix'], 'bug fix', '😤 pride', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'Refactor database layer', 'Simplified GORM repository logic.', ARRAY['refactor', 'gorm'], 'refactor', '😌 satisfaction', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000001', 'Wrote tests for payment', 'Added unit tests for payment gateway.', ARRAY['testing', 'payments'], 'testing', '💪 confidence', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000001', 'Added search filtering', 'Implemented tags/category filters.', ARRAY['search', 'filters'], 'feature', '🤔 curiosity', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000005', '00000000-0000-0000-0000-000000000001', 'Story pagination', 'Paginated storiesByUser query.', ARRAY['pagination', 'graphql'], 'backend', '🚀 flow state euphoria', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000006', '00000000-0000-0000-0000-000000000001', 'Improve Dockerfile', 'Optimized caching and layer ordering.', ARRAY['docker', 'devops'], 'infrastructure', '😮‍💨 relief', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000007', '00000000-0000-0000-0000-000000000001', 'Updated README', 'Clarified setup and env configs.', ARRAY['docs', 'setup'], 'documentation', '🏆 accomplishment', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000008', '00000000-0000-0000-0000-000000000001', 'Mock data support', 'Enabled mock mode for dev testing.', ARRAY['mock', 'dev'], 'tooling', '😴 boredom', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000009', '00000000-0000-0000-0000-000000000001', 'CI pipeline added', 'Added Jenkins pipeline for Go builds.', ARRAY['ci', 'jenkins'], 'devops', '😤 determination', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000010', '00000000-0000-0000-0000-000000000001', 'Error handling refactor', 'Improved structured logging.', ARRAY['logging', 'errors'], 'refactor', '😓 stress', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000001', 'GraphQL @key directive', 'Resolved subgraph federation bug.', ARRAY['graphql', 'federation'], 'bug fix', '😭 despair', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000012', '00000000-0000-0000-0000-000000000001', 'Hooked up Gradio UI', 'Gradio is working locally.', ARRAY['ui', 'gradio'], 'frontend', '🤷‍♂️ self-doubt', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000001', 'Wrote entity resolver', 'Returned user from story.', ARRAY['graphql', 'resolvers'], 'backend', '😵‍💫 confusion', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000014', '00000000-0000-0000-0000-000000000001', 'Postgres schema tweak', 'Updated constraints + GORM models.', ARRAY['postgres', 'gorm'], 'db', '😳 embarrassment', NOW(), NOW()),
  ('10000000-0000-0000-0000-000000000015', '00000000-0000-0000-0000-000000000001', 'Added search indexing', 'Improved lookup speed on tags.', ARRAY['search', 'index'], 'performance', '⏰ impatience', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;