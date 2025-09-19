
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

-- Seed users
INSERT INTO users (id, name, email, role, active) VALUES 
('1', 'John Doe', 'john@example.com', 'writer', true),
('2', 'Jane Smith', 'jane@example.com', 'editor', true),
('3', 'Bob Wilson', 'bob@example.com', 'writer', true),
('4', 'Alice Johnson', 'alice@example.com', 'editor', true),
('5', 'Mike Chen', 'mike@example.com', 'writer', true),
('6', 'Sarah Wilson', 'sarah@example.com', 'reader', true),
('7', 'David Brown', 'david@example.com', 'reader', false),
('8', 'Emma Davis', 'emma@example.com', 'writer', true),
('9', 'James Miller', 'james@example.com', 'editor', true),
('10', 'Lisa Garcia', 'lisa@example.com', 'writer', true)
ON CONFLICT (id) DO NOTHING;

-- Wait for job_stories table
DO $$
BEGIN
    WHILE NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'job_stories') LOOP
        PERFORM pg_sleep(1);
    END LOOP;
END $$;
-- Insert job stories for pagination testing
INSERT INTO job_stories (id, user_id, title, content, tags, category, mood, created_at, updated_at) VALUES
  ('1', '1', 'Fix login bug', 'Fixed OAuth issue during login.', ARRAY['auth', 'bugfix'], 'bug fix', '😤 pride', NOW(), NOW()),
  ('2', '1', 'Refactor database layer', 'Simplified GORM repository logic.', ARRAY['refactor', 'gorm'], 'refactor', '😌 satisfaction', NOW(), NOW()),
  ('3', '1', 'Wrote tests for payment', 'Added unit tests for payment gateway.', ARRAY['testing', 'payments'], 'testing', '💪 confidence', NOW(), NOW()),
  ('4', '1', 'Added search filtering', 'Implemented tags/category filters.', ARRAY['search', 'filters'], 'feature', '🤔 curiosity', NOW(), NOW()),
  ('5', '1', 'Story pagination', 'Paginated storiesByUser query.', ARRAY['pagination', 'graphql'], 'backend', '🚀 flow state euphoria', NOW(), NOW()),
  ('6', '1', 'Improve Dockerfile', 'Optimized caching and layer ordering.', ARRAY['docker', 'devops'], 'infrastructure', '😮‍💨 relief', NOW(), NOW()),
  ('7', '1', 'Updated README', 'Clarified setup and env configs.', ARRAY['docs', 'setup'], 'documentation', '🏆 accomplishment', NOW(), NOW()),
  ('8', '1', 'Mock data support', 'Enabled mock mode for dev testing.', ARRAY['mock', 'dev'], 'tooling', '😴 boredom', NOW(), NOW()),
  ('9', '1', 'CI pipeline added', 'Added Jenkins pipeline for Go builds.', ARRAY['ci', 'jenkins'], 'devops', '😤 determination', NOW(), NOW()),
  ('10', '1', 'Error handling refactor', 'Improved structured logging.', ARRAY['logging', 'errors'], 'refactor', '😓 stress', NOW(), NOW()),
  ('11', '1', 'GraphQL @key directive', 'Resolved subgraph federation bug.', ARRAY['graphql', 'federation'], 'bug fix', '😭 despair', NOW(), NOW()),
  ('12', '1', 'Hooked up Gradio UI', 'Gradio is working locally.', ARRAY['ui', 'gradio'], 'frontend', '🤷‍♂️ self-doubt', NOW(), NOW()),
  ('13', '1', 'Wrote entity resolver', 'Returned user from story.', ARRAY['graphql', 'resolvers'], 'backend', '😵‍💫 confusion', NOW(), NOW()),
  ('14', '1', 'Postgres schema tweak', 'Updated constraints + GORM models.', ARRAY['postgres', 'gorm'], 'db', '😳 embarrassment', NOW(), NOW()),
  ('15', '1', 'Added search indexing', 'Improved lookup speed on tags.', ARRAY['search', 'index'], 'performance', '⏰ impatience', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;

