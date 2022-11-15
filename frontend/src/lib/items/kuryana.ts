import { ASIANDRAMA_API } from './api';
import type { CollectionItemProps } from './props';

export const asiandramaFetcher = async (query: string) => {
	const r = await fetch(ASIANDRAMA_API.replace('[query]', query));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.results.dramas) {
		results.push({
			name: i.title,
			image: i.thumb.replace('s.jpg', 'f.jpg'),
			item_id: String(i.mdl_id),
			url: `https://mydramalist.com/${i.slug}`,
			type: 'asian_drama'
		});
	}

	return results;
};
