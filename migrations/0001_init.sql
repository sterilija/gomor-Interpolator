-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Role VARCHAR(255) NOT NULL,
    Active BOOLEAN NOT NULL,
    Shadow_Rating INT NOT NULL
);
CREATE TABLE IF NOT EXISTS actions (
    Id SERIAL PRIMARY KEY,
    Action_Type VARCHAR(255) NOT NULL,
    User_Id INT NOT NULL,
    Approved BOOLEAN NOT NULL
);
CREATE TABLE IF NOT EXISTS victims (
    Id SERIAL PRIMARY KEY,
    Victim_Type VARCHAR(255) NOT NULL,
    User_Id INT NOT NULL,
    Victim_Name VARCHAR(255) NOT NULL,
Victim_avg_Rating FLOAT NOT NULL,
    Victim_median_Rating FLOAT NOT NULL,
    Victim_Tags_Id INT NOT NULL,
    Victim_Legend_Id INT NOT NULL
);
CREATE TABLE IF NOT EXISTS victim_tags (
    Id SERIAL PRIMARY KEY,
    Tag VARCHAR(255) NOT NULL,
    User_Id INT NOT NULL
);
CREATE TABLE IF NOT EXISTS victim_legends (
    Id SERIAL PRIMARY KEY,
    Legend VARCHAR(255) NOT NULL,
    User_Id INT NOT NULL,
    Victim_Corps VARCHAR(255) NOT NULL,
    Victim_Teams VARCHAR(255) NOT NULL,
    Proofs VARCHAR(255) NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS actions;
DROP TABLE IF EXISTS victims;
DROP TABLE IF EXISTS victim_tags;
DROP TABLE IF EXISTS victim_legends;
