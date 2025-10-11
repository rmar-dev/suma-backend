-- Drop the function for creating transfers
DROP FUNCTION IF EXISTS create_transfer(UUID, UUID, DECIMAL, VARCHAR(3), TEXT, TIMESTAMP WITH TIME ZONE);

-- Drop the trigger and function for updating account balance
DROP TRIGGER IF EXISTS update_account_balance ON transactions;
DROP FUNCTION IF EXISTS update_account_balance();

-- Drop the trigger for updated_at
DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;

-- Drop all indexes
DROP INDEX IF EXISTS idx_transactions_user_id;
DROP INDEX IF EXISTS idx_transactions_account_id;
DROP INDEX IF EXISTS idx_transactions_type;
DROP INDEX IF EXISTS idx_transactions_date;
DROP INDEX IF EXISTS idx_transactions_category;
DROP INDEX IF EXISTS idx_transactions_status;
DROP INDEX IF EXISTS idx_transactions_merchant;
DROP INDEX IF EXISTS idx_transactions_is_recurring;

-- Drop the table
DROP TABLE IF EXISTS transactions;

-- Drop the custom types
DROP TYPE IF EXISTS transaction_type;
DROP TYPE IF EXISTS transaction_status;