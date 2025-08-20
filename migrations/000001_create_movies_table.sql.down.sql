-- Удаляем таблицу связей жанров с фильмами (сначала, так как она зависит от других таблиц)
DROP TABLE IF EXISTS movies.movie_genres;

-- Удаляем таблицу жанров
DROP TABLE IF EXISTS movies.genres;

-- Удаляем все индексы таблицы movies
DROP INDEX IF EXISTS idx_movies_title;

DROP INDEX IF EXISTS idx_movies_release_year;

DROP INDEX IF EXISTS idx_movies_is_watched;

DROP INDEX IF EXISTS idx_movies_imdb_rating;

DROP INDEX IF EXISTS idx_movies_kinopoisk_rating;

-- Наконец, удаляем саму таблицу movies
DROP TABLE IF EXISTS movies.movies;

DROP SCHEMA IF EXISTS movies;
