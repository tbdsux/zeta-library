export type CollectionTypes =
	| 'series'
	| 'movies'
	| 'anime'
	| 'manga'
	| 'asian_drama'
	| 'books'
	| '';

export interface CollectionTypeProps {
	name: string;
	value: CollectionTypes;
}

export interface PartialCollectionProps {
	name: string;
	type: CollectionTypes;
	description: string;
}

export interface CollectionProps extends PartialCollectionProps {
	id: string;
	key: string;
}
