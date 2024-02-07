CREATE TABLE IF NOT EXISTS interviews (
    id SERIAL PRIMARY KEY,
    interview_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS interview_questions (
    id SERIAL PRIMARY KEY,
    text_answer VARCHAR(255) NOT NULL,
    interview_id INTEGER,
    FOREIGN KEY (interview_id) REFERENCES interviews(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS problem_questions (
    id SERIAL PRIMARY KEY,
    question_text VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS problem_answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER,
    answer_text VARCHAR(255),
    FOREIGN KEY (question_id) REFERENCES problem_questions(id) ON DELETE CASCADE
);