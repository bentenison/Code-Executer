CREATE TABLE languages (
    id SERIAL PRIMARY KEY,          -- Unique identifier for the language entry
    code VARCHAR(10) NOT NULL,     -- Language code (e.g., 'PY', 'JAVA', 'CPP')
    name VARCHAR(50) NOT NULL,      -- Name of the programming language (e.g., 'Python', 'Java')
    container_id VARCHAR(255) NOT NULL, -- Unique identifier for the Docker container
    container_name VARCHAR(255) NOT NULL, -- Name of the Docker container (e.g., 'python:3.9')
    version VARCHAR(20),            -- Version of the language (e.g., '3.9', '11', '14')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the entry was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
    documentation_url VARCHAR(255), -- Optional URL to official documentation
    is_active BOOLEAN DEFAULT TRUE   -- Status indicating if the language is currently active
);

-- Indexes for optimization (if needed)
CREATE INDEX idx_languages_code ON languages(code);
CREATE INDEX idx_languages_name ON languages(name);

-- Seed data for the languages table
INSERT INTO languages (code, name, container_id, container_name, version, documentation_url, is_active)
VALUES
  ('PY', 'Python', 'python_3_9', 'python:3.9', '3.9', 'https://docs.python.org/3/', TRUE),
  ('JAVA', 'Java', 'openjdk_11', 'openjdk:11', '11', 'https://docs.oracle.com/en/java/javase/11/', TRUE),
  ('CPP', 'C++', 'gcc_10', 'gcc:10', '10', 'https://gcc.gnu.org/onlinedocs/', TRUE),
  ('JS', 'JavaScript', 'node_14', 'node:14', '14', 'https://nodejs.org/en/docs/', TRUE),
  ('GO', 'Go', 'golang_1_17', 'golang:1.17', '1.17', 'https://golang.org/doc/', TRUE),
  ('RUBY', 'Ruby', 'ruby_3_0', 'ruby:3.0', '3.0', 'https://www.ruby-lang.org/en/documentation/', TRUE),
  ('PHP', 'PHP', 'php_7_4', 'php:7.4', '7.4', 'https://www.php.net/docs.php', TRUE),
  ('RUST', 'Rust', 'rust_1_55', 'rust:1.55', '1.55', 'https://doc.rust-lang.org/', TRUE),
  ('CS', 'C#', 'dotnet_5', 'mcr.microsoft.com/dotnet/sdk:5.0', '5.0', 'https://docs.microsoft.com/en-us/dotnet/csharp/', TRUE);
INSERT INTO languages (code, name, container_id, container_name, version, documentation_url, is_active)
VALUES
  ('C', 'C', 'gcc_9_3', 'gcc:9.3', '9.3', 'https://gcc.gnu.org', TRUE);



-- Optional: Insert more languages as needed

