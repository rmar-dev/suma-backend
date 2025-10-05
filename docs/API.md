# 📚 Finance App API Documentation

Base URL: `http://localhost:8080/api/v1`

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## 📋 Endpoints

### Health Check

#### GET /health
Check if the API is running.

**Response:**
```json
{
  "status": "healthy",
  "service": "finance-app-api"
}
```

---

### Authentication Endpoints

#### POST /api/v1/auth/register
Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123",
  "password_confirm": "SecurePassword123",
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "+351912345678",
  "marketing_consent": false,
  "terms_accepted": true,
  "privacy_accepted": true,
  "gdpr_consent": true
}
```

**Validation Rules:**
- Email: Valid email format, unique
- Password: Minimum 8 characters
- Password Confirm: Must match password
- First Name: 2-50 characters
- Last Name: 2-50 characters
- Terms, Privacy, GDPR: Must be true

**Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "email_verified": false,
    "language": "pt",
    "currency": "EUR",
    "timezone": "Europe/Lisbon",
    "created_at": "2024-10-05T12:00:00Z"
  },
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "expires_at": "2024-10-05T13:00:00Z"
  }
}
```

---

#### POST /api/v1/auth/login
Login with email and password.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123"
}
```

**Response (200 OK):**
```json
{
  "message": "Login successful",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "last_login": "2024-10-05T12:00:00Z"
  },
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "expires_at": "2024-10-05T13:00:00Z"
  }
}
```

**Error Responses:**
- 401 Unauthorized: Invalid credentials
- 403 Forbidden: Account deactivated

---

#### POST /api/v1/auth/refresh
Refresh access token using refresh token.

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response (200 OK):**
```json
{
  "message": "Token refreshed successfully",
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "expires_at": "2024-10-05T13:00:00Z"
  }
}
```

---

#### POST /api/v1/auth/logout
Logout user (client should remove tokens).

**Response (200 OK):**
```json
{
  "message": "Logout successful. Please remove tokens from client storage."
}
```

---

### User Endpoints (Protected)

#### GET /api/v1/me
Get current authenticated user's profile.

**Headers Required:**
```
Authorization: Bearer <access-token>
```

**Response (200 OK):**
```json
{
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "+351912345678",
    "email_verified": true,
    "two_factor_enabled": false,
    "language": "pt",
    "currency": "EUR",
    "timezone": "Europe/Lisbon",
    "created_at": "2024-10-05T10:00:00Z",
    "last_login": "2024-10-05T12:00:00Z"
  }
}
```

---

## 🔄 Coming Soon

### Account Management
- `GET /api/v1/accounts` - List user's bank accounts
- `POST /api/v1/accounts` - Add a bank account
- `GET /api/v1/accounts/{id}` - Get account details
- `PUT /api/v1/accounts/{id}` - Update account
- `DELETE /api/v1/accounts/{id}` - Remove account
- `POST /api/v1/accounts/{id}/sync` - Sync transactions

### Transaction Management
- `GET /api/v1/transactions` - List transactions
- `GET /api/v1/transactions/{id}` - Get transaction details
- `PUT /api/v1/transactions/{id}` - Update transaction
- `POST /api/v1/transactions/{id}/categorize` - Categorize transaction

### Subscription Management
- `GET /api/v1/subscriptions` - List subscriptions
- `POST /api/v1/subscriptions` - Add subscription manually
- `GET /api/v1/subscriptions/{id}` - Get subscription details
- `PUT /api/v1/subscriptions/{id}` - Update subscription
- `DELETE /api/v1/subscriptions/{id}` - Cancel subscription
- `GET /api/v1/subscriptions/detect` - Auto-detect subscriptions

### Budget Management
- `GET /api/v1/budgets` - List budgets
- `POST /api/v1/budgets` - Create budget
- `GET /api/v1/budgets/{id}` - Get budget details
- `PUT /api/v1/budgets/{id}` - Update budget
- `DELETE /api/v1/budgets/{id}` - Delete budget

### Reports & Analytics
- `GET /api/v1/reports/summary` - Financial summary
- `GET /api/v1/reports/spending` - Spending analysis
- `GET /api/v1/reports/subscriptions` - Subscription costs
- `GET /api/v1/reports/export` - Export data

---

## 📝 Error Responses

All endpoints may return these error formats:

### Validation Error (400)
```json
{
  "error": "Validation failed",
  "details": {
    "field": "error message"
  }
}
```

### Unauthorized (401)
```json
{
  "error": "Authorization header missing"
}
```

### Forbidden (403)
```json
{
  "error": "Access denied"
}
```

### Not Found (404)
```json
{
  "error": "Resource not found"
}
```

### Internal Server Error (500)
```json
{
  "error": "An unexpected error occurred"
}
```

---

## 🔒 Security Headers

The API includes these security headers:
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security: max-age=31536000`

---

## 📊 Rate Limiting

- **Authentication endpoints**: 5 requests per minute
- **Other endpoints**: 100 requests per minute per user

---

## 🧪 Testing with cURL

### Register
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456","password_confirm":"Test123456","first_name":"Test","last_name":"User","terms_accepted":true,"privacy_accepted":true,"gdpr_consent":true}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test123456"}'
```

### Get Profile (with token)
```bash
curl -X GET http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

---

## 🔗 Postman Collection

Import this collection to Postman for easy testing:

1. Create a new collection named "Finance App API"
2. Add environment variables:
   - `base_url`: http://localhost:8080
   - `access_token`: (set after login)
   - `refresh_token`: (set after login)
3. Add the endpoints above with appropriate headers and bodies

---

**API Version:** 1.0.0  
**Last Updated:** October 2024