CREATE TABLE performance_metrics (
    id SERIAL PRIMARY KEY,                             -- Unique identifier for each metric record
    submission_id INT NOT NULL,                        -- Foreign key referencing the submission
    execution_time INTERVAL NOT NULL,                  -- Time taken for the execution
    memory_usage INT NOT NULL,                         -- Memory used during execution (in bytes)
    status VARCHAR(20) NOT NULL,                       -- Execution status (e.g., 'Success', 'Errorc')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Timestamp of when the metric was recorded
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
    FOREIGN KEY (submission_id) REFERENCES submissions(id) -- Ensure referential integrity
);
-- Seed data for the performance_metrics table
INSERT INTO performance_metrics (submission_id, execution_time, memory_usage, status)
VALUES
  (1, INTERVAL '0:00:01.5', 102400, 'Success'),      -- 1.5 seconds, 100 KB, successful execution
  (2, INTERVAL '0:00:02.3', 204800, 'Errorc'),        -- 2.3 seconds, 200 KB, execution with an error
  (3, INTERVAL '0:00:00.8', 51200, 'Success'),       -- 0.8 seconds, 50 KB, successful execution
  (4, INTERVAL '0:00:04.0', 307200, 'Success'),      -- 4 seconds, 300 KB, successful execution
  (5, INTERVAL '0:00:01.2', 153600, 'Errorc'),        -- 1.2 seconds, 150 KB, execution with an error
  (6, INTERVAL '0:00:05.5', 409600, 'Success');      -- 5.5 seconds, 400 KB, successful execution
