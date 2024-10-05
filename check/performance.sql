CREATE TABLE performance_metrics (
    id SERIAL PRIMARY KEY,                             -- Unique identifier for each metric record
    submission_id INT NOT NULL,                        -- Foreign key referencing the submission
    execution_time INTERVAL NOT NULL,                  -- Time taken for the execution
    memory_usage INT NOT NULL,                         -- Memory used during execution (in bytes)
    status VARCHAR(20) NOT NULL,                       -- Execution status (e.g., 'Success', 'Error')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Timestamp of when the metric was recorded
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
    FOREIGN KEY (submission_id) REFERENCES submissions(id) -- Ensure referential integrity
);
