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
	created_at: number;
}

export interface CollectionProps extends PartialCollectionProps {
	id: string;
	key: string;
}
