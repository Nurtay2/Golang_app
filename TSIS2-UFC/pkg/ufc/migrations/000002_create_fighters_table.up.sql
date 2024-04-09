
CREATE TABLE fighters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight_class VARCHAR(50),
    reach DECIMAL(5, 2),
    wins INT,
    losses INT
);

CREATE TABLE fighting_styles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    duration INTERVAL,
    winner_fighter_id INT
);

CREATE TABLE gyms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255)
);

CREATE TABLE promotions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    founder_id INT
);