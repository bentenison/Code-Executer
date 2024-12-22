-- USERS TABLE
CREATE TABLE user_metrics (
    user_id VARCHAR(100) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    level INT NOT NULL DEFAULT 1, -- Default level is 1 (Easy)
    total_score INT NOT NULL DEFAULT 0,
    accuracy FLOAT NOT NULL DEFAULT 0.0,
    speed_avg FLOAT NOT NULL DEFAULT 0.0,
    penalty_points INT NOT NULL DEFAULT 0,
    rank INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CHALLENGES TABLE
CREATE TABLE challenges (
    challenge_id VARCHAR(100) PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    difficulty INT NOT NULL, -- 1 = Easy, 2 = Medium, 3 = Hard
    score INT NOT NULL,
    execution_time INT NOT NULL, -- Time in milliseconds
    optimal_solution TEXT
);

-- SUBMISSIONS TABLE
CREATE TABLE submission_stats (
    submission_id VARCHAR(100) PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL REFERENCES user_metrics(user_id) ON DELETE CASCADE,
    challenge_id VARCHAR(100) NOT NULL REFERENCES challenges(challenge_id) ON DELETE CASCADE,
    is_correct BOOLEAN NOT NULL,
    attempts INT NOT NULL DEFAULT 1,
    time_taken INT NOT NULL, -- Time in seconds to submit
    code_quality FLOAT NOT NULL DEFAULT 0.0,
    penalty_points INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
