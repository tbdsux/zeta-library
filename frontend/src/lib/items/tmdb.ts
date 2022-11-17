import { MOVIE_API, SERIES_API, TMDB_IMG } from './api';
import type { CollectionItemProps } from './props';

export const seriesFetcher = async (query: string, apiKey: string) => {
	const r = await fetch(SERIES_API.replace('[query]', query).replace('[apikey]', apiKey));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.results) {
		results.push({
			name: i.original_name,
			image: i.poster_path ? TMDB_IMG + i.poster_path : null,
			item_id: String(i.id),
			type: 'series',
			url: `https://www.themoviedb.org/tv/${i.id}`
		});
	}

	return results;
};

export const movieFetcher = async (query: string, apiKey: string) => {
	const r = await fetch(MOVIE_API.replace('[query]', query).replace('[apikey]', apiKey));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.results) {
		results.push({
			name: i.original_title,
			image: i.poster_path ? TMDB_IMG + i.poster_path : null,
			item_id: String(i.id),
			type: 'movies',
			url: `https://www.themoviedb.org/movie/${i.id}`
		});
	}

	return results;
};
