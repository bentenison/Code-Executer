CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,                       -- Unique identifier for each submission
    user_id INT NOT NULL,                        -- Identifier for the user who made the submission
    language_id INT NOT NULL,                    -- Foreign key referencing the languages table
    code_snippet TEXT NOT NULL,                  -- The code that was submitted
    submission_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of when the submission was made
    execution_status VARCHAR(20) NOT NULL,       -- Status of the execution (e.g., 'Pending', 'Executed', 'Failed')
    result_id INT,                               -- Foreign key referencing the execution result in the code_execution_stats table
    is_public BOOLEAN DEFAULT FALSE,             -- Visibility of the submission (public or private)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the entry was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Timestamp for last update
);

-- Indexes for optimization (if needed)
CREATE INDEX idx_submissions_user ON submissions(user_id);
CREATE INDEX idx_submissions_language ON submissions(language_id);
CREATE INDEX idx_submissions_status ON submissions(execution_status);
