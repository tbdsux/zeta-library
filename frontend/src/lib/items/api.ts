import { moviedbKey } from '$lib/config';

const MOVIE_API = `https://api.themoviedb.org/3/search/movie?api_key=${moviedbKey}&language=en-US&query=[query]&page=1&include_adult=true`;
const SERIES_API = `https://api.themoviedb.org/3/search/tv?api_key=${moviedbKey}&language=en-US&query=[query]&page=1&include_adult=true`;
const ANIME_API = 'https://api.jikan.moe/v4/anime?q=[query]';
const MANGA_API = 'https://api.jikan.moe/v4/manga?q=[query]';
const BOOKS_API = 'https://openlibrary.org/search.json?title=[query]&page=1&limit=20';
const ASIANDRAMA_API = 'https://kuryana.vercel.app/search/q/[query]';

const TMDB_IMG = 'https://image.tmdb.org/t/p/w600_and_h900_bestv2';

export { MOVIE_API, SERIES_API, ANIME_API, MANGA_API, BOOKS_API, ASIANDRAMA_API, TMDB_IMG };
