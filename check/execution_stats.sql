CREATE TABLE code_execution_stats (
    id SERIAL PRIMARY KEY,                     -- Unique identifier for each execution record
    user_id INT NOT NULL,                      -- Identifier for the user who executed the code
    language_id INT NOT NULL,                  -- Foreign key referencing the languages table
    execution_time INTERVAL NOT NULL,          -- Time taken for the execution (e.g., '5 seconds', '1 minute')
    memory_usage INT,                          -- Memory used during the execution (in bytes)
    status VARCHAR(20) NOT NULL,               -- Execution status (e.g., 'Success', 'Error', 'Timeout')
    error_message TEXT,                        -- Error message if the execution failed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the execution was recorded
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
    code_snippet TEXT NOT NULL,                -- The code snippet that was executed
    container_id VARCHAR(255) NOT NULL         -- ID of the container used for execution
);

-- Indexes for optimization (if needed)
CREATE INDEX idx_code_execution_user ON code_execution_stats(user_id);
CREATE INDEX idx_code_execution_language ON code_execution_stats(language_id);
CREATE INDEX idx_code_execution_status ON code_execution_stats(status);
