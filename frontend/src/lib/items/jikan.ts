import { ANIME_API, MANGA_API } from './api';
import type { CollectionItemProps } from './props';

export const animeFetcher = async (query: string) => {
	const r = await fetch(ANIME_API.replace('[query]', query));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.data) {
		results.push({
			name: i.title,
			type: 'anime',
			image: i.images.jpg.large_image_url,
			item_id: String(i.mal_id),
			url: i.url
		});
	}

	return results;
};

export const mangaFetcher = async (query: string) => {
	const r = await fetch(MANGA_API.replace('[query]', query));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.data) {
		results.push({
			name: i.title,
			type: 'manga',
			image: i.images.jpg.large_image_url,
			item_id: String(i.mal_id),
			url: i.url
		});
	}

	return results;
};
