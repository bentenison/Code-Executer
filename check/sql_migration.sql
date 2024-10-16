-- Enable UUID extension (if needed, to generate UUIDs)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Updated submissions table with UUID stored as VARCHAR(36)
CREATE TABLE submissions (
    id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),  -- UUID (as string) for each submission
    user_id VARCHAR(36) NOT NULL,                           -- UUID (as string) for the user who made the submission
    language_id VARCHAR(36) NOT NULL,                       -- UUID (as string) for the programming language used
    code_snippet TEXT NOT NULL,                             -- The code that was submitted
    submission_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Timestamp of when the submission was made
    execution_status VARCHAR(20) NOT NULL,                  -- Status of the execution (e.g., 'Pending', 'Executed', 'Failed')
    result_id VARCHAR(36), 
    question_id VARCHAR(36) NOT NULL;                                 -- UUID (as string) referencing the execution result in code_execution_stats
    is_public BOOLEAN DEFAULT FALSE,                        -- Visibility of the submission (public or private)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- Timestamp of when the entry was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- Timestamp for last update
    FOREIGN KEY (user_id) REFERENCES users(id),             -- Foreign key referencing users table
    FOREIGN KEY (language_id) REFERENCES languages(id)     -- Foreign key referencing languages table
);
ALTER TABLE submissions
ADD COLUMN question_id VARCHAR(36) NOT NULL;
-- Indexes for optimization
CREATE INDEX idx_submissions_user ON submissions(user_id);
CREATE INDEX idx_submissions_language ON submissions(language_id);
CREATE INDEX idx_submissions_status ON submissions(execution_status);

-- Updated performance_metrics table with UUID as string
CREATE TABLE performance_metrics (
    id SERIAL PRIMARY KEY,                                -- Unique identifier for each metric record
    submission_id VARCHAR(36) NOT NULL,                   -- UUID (as string) referencing the submission
    execution_time INTERVAL NOT NULL,                     -- Time taken for the execution
    memory_usage INT NOT NULL,                            -- Memory used during execution (in bytes)
    status VARCHAR(20) NOT NULL,                          -- Execution status (e.g., 'Success', 'Errorc')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- Timestamp of when the metric was recorded
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- Timestamp for last update
    FOREIGN KEY (submission_id) REFERENCES submissions(id) -- Ensure referential integrity
);

-- Updated code_execution_stats table with UUID as string
CREATE TABLE code_execution_stats (
    id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for each execution record as UUID string
    user_id VARCHAR(36) NOT NULL,                           -- UUID (as string) for the user who executed the code
    language_id VARCHAR(36) NOT NULL,                       -- UUID (as string) for the programming language
    execution_time DOUBLE PRECISION NOT NULL,                         -- Time taken for the execution in milliseconds (int64)
    memory_usage BIGINT,                                   -- Memory used during execution (in bytes, int64)
    total_memory BIGINT,                                   -- Total memory available (in bytes, int64)
    cpu_usage BIGINT,                                      -- CPU usage (in milliseconds, int64)
    memory_percentage FLOAT,                                -- Memory usage percentage (float64)
    status VARCHAR(20) NOT NULL,                            -- Execution status (e.g., 'Success', 'Error', 'Timeout')
    error_message TEXT,                                     -- Error message if execution failed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- Timestamp of when the execution was recorded
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- Timestamp for last update
    code_snippet TEXT NOT NULL,                             -- Code snippet that was executed
    container_id VARCHAR(255) NOT NULL,                     -- ID of the container used for execution
    FOREIGN KEY (user_id) REFERENCES users(id),             -- Foreign key referencing users table
    FOREIGN KEY (language_id) REFERENCES languages(id)      -- Foreign key referencing languages table
);

-- Indexes for optimization
CREATE INDEX idx_code_execution_user ON code_execution_stats(user_id);
CREATE INDEX idx_code_execution_language ON code_execution_stats(language_id);
CREATE INDEX idx_code_execution_status ON code_execution_stats(status);


-- Updated users table with UUID as string
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),-- Unique identifier for each user as UUID string
    username VARCHAR(50) NOT NULL UNIQUE,                 -- Unique username for the user
    email VARCHAR(100) NOT NULL UNIQUE,                   -- Unique email address for the user
    password_hash VARCHAR(255) NOT NULL,                  -- Hashed password for security
    first_name VARCHAR(50),                               -- User's first name
    last_name VARCHAR(50),                                -- User's last name
    role VARCHAR(20) CHECK (role IN ('admin', 'creator', 'user')) DEFAULT 'user', -- User role
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- Timestamp of when the user was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP        -- Timestamp for last update
);

-- Indexes for optimization
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);

-- Updated languages table with UUID as string
CREATE TABLE languages (
    id VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4(),-- Unique identifier for the language entry as UUID string
    code VARCHAR(10) NOT NULL,                            -- Language code (e.g., 'PY', 'JAVA', 'CPP')
    name VARCHAR(50) NOT NULL,                            -- Name of the programming language (e.g., 'Python', 'Java')
    container_id VARCHAR(255) NOT NULL,                   -- Unique identifier for the Docker container
    container_name VARCHAR(255) NOT NULL,                 -- Name of the Docker container (e.g., 'python:3.9')
    version VARCHAR(20),                                  -- Version of the language (e.g., '3.9', '11', '14')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- Timestamp of when the entry was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       -- Timestamp for last update
    documentation_url VARCHAR(255),                       -- Optional URL to official documentation
    is_active BOOLEAN DEFAULT TRUE                        -- Status indicating if the language is active
);

-- Indexes for optimization
CREATE INDEX idx_languages_code ON languages(code);
CREATE INDEX idx_languages_name ON languages(name);

-- Seed data for the languages table (with UUID as strings)
INSERT INTO languages (id, code, name, container_id, container_name, version, documentation_url, is_active)
VALUES
  (uuid_generate_v4()::varchar, 'PY', 'Python', 'python_3_9', 'python:3.9', '3.9', 'https://docs.python.org/3/', TRUE),
  (uuid_generate_v4()::varchar, 'JAVA', 'Java', 'openjdk_11', 'openjdk:11', '11', 'https://docs.oracle.com/en/java/javase/11/', TRUE),
  (uuid_generate_v4()::varchar, 'CPP', 'C++', 'gcc_10', 'gcc:10', '10', 'https://gcc.gnu.org/onlinedocs/', TRUE),
  (uuid_generate_v4()::varchar, 'JS', 'JavaScript', 'node_14', 'node:14', '14', 'https://nodejs.org/en/docs/', TRUE),
  (uuid_generate_v4()::varchar, 'GO', 'Go', 'golang_1_17', 'golang:1.17', '1.17', 'https://golang.org/doc/', TRUE),
  (uuid_generate_v4()::varchar, 'RUBY', 'Ruby', 'ruby_3_0', 'ruby:3.0', '3.0', 'https://www.ruby-lang.org/en/documentation/', TRUE),
  (uuid_generate_v4()::varchar, 'PHP', 'PHP', 'php_7_4', 'php:7.4', '7.4', 'https://www.php.net/docs.php', TRUE),
  (uuid_generate_v4()::varchar, 'RUST', 'Rust', 'rust_1_55', 'rust:1.55', '1.55', 'https://doc.rust-lang.org/', TRUE),
  (uuid_generate_v4()::varchar, 'CS', 'C#', 'dotnet_5', 'mcr.microsoft.com/dotnet/sdk:5.0', '5.0', 'https://docs.microsoft.com/en-us/dotnet/csharp/', TRUE);

