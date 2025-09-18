
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

-- Seed job_stories
INSERT INTO job_stories (id, user_id, title, content, tags, category, mood, created_at, updated_at) VALUES
('101', '1', 'Got promoted!', 'I finally got promoted after 2 years of effort.', ARRAY['promotion', 'career'], 'win', 'happy', NOW(), NOW()),
('102', '2', 'Failed deployment', 'Had a major bug in prod, learned a lot fixing it.', ARRAY['deployment', 'bug'], 'fail', 'frustrated', NOW(), NOW()),
('103', '3', 'First PR merged', 'My first pull request got accepted!', ARRAY['github', 'code'], 'win', 'proud', NOW(), NOW()),
('104', '4', 'Interview gone wrong', 'Forgot to share screen in an interview.', ARRAY['interview', 'oops'], 'fail', 'embarrassed', NOW(), NOW()),
('105', '5', 'Refactored legacy code', 'Cleaned up a huge mess in our codebase.', ARRAY['refactor', 'legacy'], 'win', 'satisfied', NOW(), NOW()),
('106', '6', 'Got rejected again', 'Another job rejection, but I will not give up.', ARRAY['jobhunt', 'rejection'], 'fail', 'resilient', NOW(), NOW()),
('107', '7', 'Built a dashboard', 'Created a useful dashboard for team metrics.', ARRAY['dashboard', 'metrics'], 'win', 'accomplished', NOW(), NOW()),
('108', '8', 'Mentored new dev', 'Helped onboard a junior engineer.', ARRAY['mentorship', 'team'], 'win', 'supportive', NOW(), NOW()),
('109', '9', 'Bad sprint planning', 'Estimated 5 tasks, finished only 2.', ARRAY['agile', 'planning'], 'fail', 'reflective', NOW(), NOW()),
('110', '10', 'Late night deploy', 'Had to hotfix at 2 AM 😴', ARRAY['hotfix', 'prod'], 'fail', 'tired', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;