CREATE TABLE IF NOT EXISTS settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    notifications JSONB DEFAULT '{
        "email": true,
        "push": true,
        "sms": false,
        "schedule": "instant"
    }',
    preferences JSONB DEFAULT '{
        "language": "en",
        "currency": "USD",
        "timezone": "UTC",
        "date_format": "YYYY-MM-DD",
        "number_format": "1,234.56"
    }',
    security JSONB DEFAULT '{
        "two_factor_enabled": false,
        "login_notifications": true,
        "suspicious_activity": true,
        "password_expires_days": 90
    }',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on user_id for faster lookups
CREATE UNIQUE INDEX idx_settings_user_id ON settings(user_id);

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_settings_updated_at
    BEFORE UPDATE ON settings
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();