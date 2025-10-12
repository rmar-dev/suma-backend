-- Drop the trigger first
DROP TRIGGER IF EXISTS update_sessions_updated_at ON sessions;

-- Drop the indexes
DROP INDEX IF EXISTS idx_sessions_user_id;
DROP INDEX IF EXISTS idx_sessions_access_token;
DROP INDEX IF EXISTS idx_sessions_refresh_token;

-- Drop the table
DROP TABLE IF EXISTS sessions;