-- Drop the trigger first
DROP TRIGGER IF EXISTS update_settings_updated_at ON settings;

-- Drop the index
DROP INDEX IF EXISTS idx_settings_user_id;

-- Drop the table
DROP TABLE IF EXISTS settings;