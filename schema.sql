CREATE TABLE assets (
    id INT PRIMARY KEY,
    title TEXT,
    description TEXT,
    creator TEXT,
    version TEXT,
    repository_url TEXT,
    stars INT,
    first_commit TIMESTAMP,
    latest_commit TIMESTAMP
)