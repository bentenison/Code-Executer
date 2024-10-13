CREATE TABLE users (
    id SERIAL PRIMARY KEY,                        -- Unique identifier for each user
    username VARCHAR(50) NOT NULL UNIQUE,        -- Unique username for the user
    email VARCHAR(100) NOT NULL UNIQUE,           -- Unique email address for the user
    password_hash VARCHAR(255) NOT NULL,         -- Hashed password for security
    first_name VARCHAR(50),                       -- User's first name
    last_name VARCHAR(50),                        -- User's last name
    role VARCHAR(20) CHECK (role IN ('admin', 'creator', 'user')) DEFAULT 'user', -- User role
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the user was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);

-- Indexes for optimization (if needed)
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
