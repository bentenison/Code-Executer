-- Insert sample data for Python

CREATE TABLE language_specifications (
    id SERIAL PRIMARY KEY,           -- Auto-incremented ID for each language entry
    language_name VARCHAR(100) NOT NULL,  -- Language name (e.g., Python, Java, etc.)
    file_extension VARCHAR(10) NOT NULL, -- File extension (e.g., .py, .java, etc.)
    docker_image VARCHAR(255) NOT NULL,  -- Docker image (e.g., python:latest, openjdk:latest, etc.)
    command TEXT[] NOT NULL,           -- Command (can store multiple commands, e.g., ["go", "run"])
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the record is created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Timestamp when the record is updated
);

-- Index for faster filtering by language name
CREATE INDEX idx_language_name ON language_specifications(language_name);

INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('Python', '.py', 'python:latest', ARRAY['python']);

-- Insert sample data for Java
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('Java', '.java', 'openjdk:latest', ARRAY['javac', 'java']);

-- Insert sample data for Go
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('Go', '.go', 'golang:latest', ARRAY['go', 'run']);

-- Insert sample data for C++
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('C++', '.cpp', 'gcc:latest', ARRAY['g++', '-o', 'program', 'program.cpp', '&&', './program']);

-- Insert sample data for JavaScript (Node.js)
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('JavaScript', '.js', 'node:latest', ARRAY['node']);

-- Insert sample data for Ruby
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('Ruby', '.rb', 'ruby:latest', ARRAY['ruby']);

-- Insert sample data for PHP
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('PHP', '.php', 'php:latest', ARRAY['php']);

-- Insert sample data for Rust
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('Rust', '.rs', 'rust:latest', ARRAY['rustc', 'program.rs', '&&', './program']);

-- Insert sample data for C#
INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('C#', '.cs', 'mcr.microsoft.com/dotnet/core/sdk:latest', ARRAY['csc', 'program.cs', '&&', 'mono', 'program.exe']);

INSERT INTO language_specifications (language_name, file_extension, docker_image, command)
VALUES 
    ('C', '.c', 'gcc:latest', ARRAY['gcc', 'program.c', '-o', 'program', '&&', './program']);
