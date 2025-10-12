-- Drop the trigger first
DROP TRIGGER IF EXISTS update_accounts_updated_at ON accounts;

-- Drop the indexes
DROP INDEX IF EXISTS idx_accounts_user_id;
DROP INDEX IF EXISTS idx_accounts_status;
DROP INDEX IF EXISTS idx_accounts_type;

-- Drop the table
DROP TABLE IF EXISTS accounts;

-- Drop the custom types
DROP TYPE IF EXISTS account_type;
DROP TYPE IF EXISTS account_status;