import type { CollectionTypes } from '$lib/collection';
import { animeFetcher, mangaFetcher } from './jikan';
import { asiandramaFetcher } from './kuryana';
import { booksFetcher } from './openlibrary';
import type { CollectionItemProps } from './props';
import { movieFetcher, seriesFetcher } from './tmdb';

export const fetcher: Record<CollectionTypes, (query: string) => Promise<CollectionItemProps[]>> = {
	anime: animeFetcher,
	asian_drama: asiandramaFetcher,
	movies: movieFetcher,
	series: seriesFetcher,
	books: booksFetcher,
	manga: mangaFetcher
};
