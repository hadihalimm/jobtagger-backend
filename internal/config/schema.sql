CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP
);

CREATE TYPE progress AS ENUM('Applied', 'Assessment', 'Interview', 'Offering', 'Accepted', 'Withdrew', 'Closed');

CREATE TABLE IF NOT EXISTS Applications (
    application_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES Users(user_id) ON DELETE CASCADE,
    job_position VARCHAR(100) NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    job_location VARCHAR(100) NOT NULL,
    job_source VARCHAR(100) NOT NULL,
    application_progress PROGRESS DEFAULT 'Applied',
    applied_date DATE,
    notes TEXT,
    created_at TIMESTAMP
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Timelines (
    timeline_id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES Applications(application_id) ON DELETE CASCADE,
    timeline_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Interviews (
    interview_id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES Applications(application_id) ON DELETE CASCADE,
    interview_title VARCHAR(100) NOT NULL,
    interview_date DATE,
    notes TEXT,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Contacts (
    contact_id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES Applications(application_id) ON DELETE CASCADE,
    contact_name VARCHAR(100) NOT NULL,
    contact_email VARCHAR(100),
    contact_phone VARCHAR(20),
    position VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMP
);