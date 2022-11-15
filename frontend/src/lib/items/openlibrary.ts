import { BOOKS_API } from './api';
import type { CollectionItemProps } from './props';

export const booksFetcher = async (query: string) => {
	const r = await fetch(BOOKS_API.replace('[query]', query));
	const data = await r.json();

	const results: CollectionItemProps[] = [];
	for (const i of data.docs) {
		results.push({
			name: i.title,
			image: i.cover_i ? `https://covers.openlibrary.org/b/id/${i.cover_i}-L.jpg` : null,
			type: 'books',
			url: `https://openlibrary.org${i.key}`,
			item_id: i.key
		});
	}

	return results;
};
