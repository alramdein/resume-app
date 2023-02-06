CREATE TABLE IF NOT EXISTS resumes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone_number INT NOT NULL,
    linkedin_url TEXT,
    portofolio_url TEXT,
    achievements TEXT[]
);

CREATE TABLE IF NOT EXISTS occupations (
    id SERIAL PRIMARY KEY,
    name TEXT,
    position TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE
    status TEXT,
    achievements TEXT[],
);

CREATE TABLE IF NOT EXISTS educations (
    id SERIAL PRIMARY KEY,
    name TEXT,
    degree TEXT,
    faculty TEXT,
    city TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    score INTEGER
);