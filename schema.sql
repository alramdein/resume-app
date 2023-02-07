CREATE TABLE IF NOT EXISTS resumes (
    id BIGINT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    linkedin_url TEXT,
    portfolio_url TEXT,
    achievements TEXT[]
);

CREATE TABLE IF NOT EXISTS occupations (
    id BIGINT PRIMARY KEY NOT NULL,
    resume_id BIGINT NOT NULL,
    name TEXT,
    position TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    status TEXT,
    achievements TEXT[]
);

CREATE TABLE IF NOT EXISTS educations (
    id BIGINT PRIMARY KEY NOT NULL,
    resume_id BIGINT NOT NULL,
    name TEXT,
    degree TEXT,
    faculty TEXT,
    city TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    score INTEGER
);