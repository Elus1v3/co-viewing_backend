CREATE TABLE "user" (
    id_pk SERIAL PRIMARY KEY,
    nickname VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE user_friends (
    user_id_fk INT REFERENCES "user"(id_pk) ON DELETE CASCADE,
    friend_id_fk INT REFERENCES "user"(id_pk) ON DELETE CASCADE,    
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    PRIMARY KEY (user_id_fk, friend_id_fk)
);

CREATE TABLE watched_movies (
    user_id_fk INT REFERENCES "user"(id_pk) ON DELETE CASCADE,
    movie_id VARCHAR(20) NOT NULL,
    PRIMARY KEY (user_id_fk, movie_id)
);