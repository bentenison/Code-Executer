CREATE TABLE languages (
    id SERIAL PRIMARY KEY,          -- Unique identifier for the language entry
    code VARCHAR(10) NOT NULL,     -- Language code (e.g., 'PY', 'JAVA', 'CPP')
    name VARCHAR(50) NOT NULL,      -- Name of the programming language (e.g., 'Python', 'Java')
    container_id VARCHAR(255) NOT NULL, -- Unique identifier for the Docker container
    container_name VARCHAR(255) NOT NULL, -- Name of the Docker container (e.g., 'python:3.9')
    version VARCHAR(20),            -- Version of the language (e.g., '3.9', '11', '14')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the entry was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
    documentation_url VARCHAR(255), -- Optional URL to official documentation
    is_active BOOLEAN DEFAULT TRUE   -- Status indicating if the language is currently active
);

-- Indexes for optimization (if needed)
CREATE INDEX idx_languages_code ON languages(code);
CREATE INDEX idx_languages_name ON languages(name);
