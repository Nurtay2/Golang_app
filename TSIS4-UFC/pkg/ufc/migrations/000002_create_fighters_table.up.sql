
CREATE TABLE IF NOT EXISTS fighters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    weight_class VARCHAR(50),
    reach DECIMAL(5, 2),
    wins INT,
    losses INT
);


CREATE TABLE IF NOT EXISTS fighting_styles (
    id SERIAL PRIMARY KEY,
    FOREIGN KEY fighter_id  REFERENCES fighters(id),
    fighting_style VARCHAR(100),
    nickname VARCHAR(100),
    popular_phrase TEXT
);

CREATE TABLE IF NOT EXISTS promotion (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    fighting_style VARCHAR(100),
    nickname VARCHAR(100),
    popular_phrase TEXT
);
