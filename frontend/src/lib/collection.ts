import type { CollectionTypeProps } from './types/collection';

const collectionTypes: CollectionTypeProps[] = [
	{
		name: '(none)',
		value: ''
	},
	{
		name: 'Series',
		value: 'series'
	},
	{
		name: 'Movies',
		value: 'movies'
	},
	{
		name: 'Anime',
		value: 'anime'
	},
	{
		name: 'Manga',
		value: 'manga'
	},
	{
		name: 'Asian Drama',
		value: 'asian_drama'
	},
	{
		name: 'Books',
		value: 'books'
	}
];

export const filterCollection = (value: string) => {
	return collectionTypes.filter((i) => i.value === value)[0];
};

export { collectionTypes };
