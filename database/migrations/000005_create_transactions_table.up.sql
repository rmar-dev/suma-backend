CREATE TYPE transaction_type AS ENUM ('income', 'expense', 'transfer');
CREATE TYPE transaction_status AS ENUM ('pending', 'completed', 'failed', 'cancelled');

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    type transaction_type NOT NULL,
    amount DECIMAL(19,4) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    description TEXT NOT NULL,
    merchant VARCHAR(255),
    category VARCHAR(100),
    status transaction_status DEFAULT 'completed',
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    metadata JSONB,
    tags TEXT[],
    notes TEXT,
    is_recurring BOOLEAN DEFAULT FALSE,
    parent_transaction_id UUID REFERENCES transactions(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_type ON transactions(type);
CREATE INDEX idx_transactions_date ON transactions(date);
CREATE INDEX idx_transactions_category ON transactions(category);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_merchant ON transactions(merchant);
CREATE INDEX idx_transactions_is_recurring ON transactions(is_recurring);

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_transactions_updated_at
    BEFORE UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Create a trigger to update account balance when a transaction is inserted
CREATE OR REPLACE FUNCTION update_account_balance()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        IF NEW.type = 'income' THEN
            UPDATE accounts SET balance = balance + NEW.amount WHERE id = NEW.account_id;
        ELSIF NEW.type = 'expense' THEN
            UPDATE accounts SET balance = balance - NEW.amount WHERE id = NEW.account_id;
        END IF;
    ELSIF TG_OP = 'DELETE' THEN
        IF OLD.type = 'income' THEN
            UPDATE accounts SET balance = balance - OLD.amount WHERE id = OLD.account_id;
        ELSIF OLD.type = 'expense' THEN
            UPDATE accounts SET balance = balance + OLD.amount WHERE id = OLD.account_id;
        END IF;
    ELSIF TG_OP = 'UPDATE' THEN
        -- Revert old transaction
        IF OLD.type = 'income' THEN
            UPDATE accounts SET balance = balance - OLD.amount WHERE id = OLD.account_id;
        ELSIF OLD.type = 'expense' THEN
            UPDATE accounts SET balance = balance + OLD.amount WHERE id = OLD.account_id;
        END IF;
        
        -- Apply new transaction
        IF NEW.type = 'income' THEN
            UPDATE accounts SET balance = balance + NEW.amount WHERE id = NEW.account_id;
        ELSIF NEW.type = 'expense' THEN
            UPDATE accounts SET balance = balance - NEW.amount WHERE id = NEW.account_id;
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_account_balance
    AFTER INSERT OR UPDATE OR DELETE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_account_balance();

-- Create a function to handle transfers between accounts
CREATE OR REPLACE FUNCTION create_transfer(
    p_from_account_id UUID,
    p_to_account_id UUID,
    p_amount DECIMAL,
    p_currency VARCHAR(3),
    p_description TEXT,
    p_date TIMESTAMP WITH TIME ZONE
) RETURNS UUID AS $$
DECLARE
    v_user_id UUID;
    v_transfer_id UUID;
BEGIN
    -- Get the user_id from the from_account
    SELECT user_id INTO v_user_id FROM accounts WHERE id = p_from_account_id;
    
    -- Create the transfer group ID
    v_transfer_id := gen_random_uuid();
    
    -- Create the outgoing transaction
    INSERT INTO transactions (
        user_id,
        account_id,
        type,
        amount,
        currency,
        description,
        date,
        parent_transaction_id
    ) VALUES (
        v_user_id,
        p_from_account_id,
        'expense',
        p_amount,
        p_currency,
        p_description || ' (Transfer Out)',
        p_date,
        v_transfer_id
    );
    
    -- Create the incoming transaction
    INSERT INTO transactions (
        user_id,
        account_id,
        type,
        amount,
        currency,
        description,
        date,
        parent_transaction_id
    ) VALUES (
        v_user_id,
        p_to_account_id,
        'income',
        p_amount,
        p_currency,
        p_description || ' (Transfer In)',
        p_date,
        v_transfer_id
    );
    
    RETURN v_transfer_id;
END;
$$ LANGUAGE plpgsql;