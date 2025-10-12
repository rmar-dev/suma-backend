CREATE TYPE account_type AS ENUM ('checking', 'savings', 'credit', 'investment', 'loan', 'other');
CREATE TYPE account_status AS ENUM ('active', 'inactive', 'pending', 'closed');

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_name VARCHAR(255) NOT NULL,
    account_type account_type NOT NULL,
    account_number VARCHAR(255),
    bank_name VARCHAR(255),
    iban VARCHAR(34),
    balance DECIMAL(19,4) NOT NULL DEFAULT 0,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    is_primary BOOLEAN DEFAULT FALSE,
    status account_status DEFAULT 'active',
    metadata JSONB,
    last_sync TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_accounts_user_id ON accounts(user_id);
CREATE INDEX idx_accounts_status ON accounts(status);
CREATE INDEX idx_accounts_type ON accounts(account_type);

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_accounts_updated_at
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();