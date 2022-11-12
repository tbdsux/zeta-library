export interface CollectionTypeProps {
	name: string;
	value: string;
}

export interface PartialCollectionProps {
	name: string;
	type: string;
	description: string;
}

export interface CollectionProps extends PartialCollectionProps {
	id: string;
	key: string;
}
