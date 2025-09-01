-- Создаем таблицу
CREATE SCHEMA movies;

CREATE TABLE movies.movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    is_series BOOLEAN NOT NULL DEFAULT FALSE,
    rating NUMERIC(3, 1) CHECK (
        rating >= 0
        AND rating <= 10
    ),
    imdb_rating NUMERIC(3, 1) CHECK (
        imdb_rating >= 0
        AND imdb_rating <= 10
    ),
    kinopoisk_rating NUMERIC(3, 1) CHECK (
        kinopoisk_rating >= 0
        AND kinopoisk_rating <= 10
    ),
    imdb_id VARCHAR(20),
    kinopoisk_id INTEGER,
    image_url VARCHAR(512),
    is_watched BOOLEAN NOT NULL DEFAULT FALSE,
    release_year INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создаем индекс для быстрого поиска по названию
CREATE INDEX idx_movies_title ON movies.movies(title);

-- Создаем индекс для фильтрации по году выпуска
CREATE INDEX idx_movies_release_year ON movies.movies(release_year);

-- Создаем индекс для фильтрации по просмотренным/непросмотренным
CREATE INDEX idx_movies_is_watched ON movies.movies(is_watched);

-- Индексы для быстрого поиска по рейтингам
CREATE INDEX idx_movies_imdb_rating ON movies.movies(imdb_rating);

CREATE INDEX idx_movies_kinopoisk_rating ON movies.movies(kinopoisk_rating);

-- Таблица жанров
CREATE TABLE movies.genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- Таблица связи жанров с фильмами
CREATE TABLE movies.movie_genres (
    movie_id INTEGER REFERENCES movies.movies(id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES movies.genres(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, genre_id)
);
