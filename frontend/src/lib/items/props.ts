import type { CollectionTypes } from '$lib/types/collection';

export interface CollectionItemProps {
	name: string;
	image: string | null;
	type: CollectionTypes;
	item_id: string;
	url: string;
	date_added?: number;
	key?: string; // auto added by base
}
