import type { CollectionItemProps } from '$lib/items/props';
import type { SettingsProps } from '$lib/settings';
import type { CollectionProps } from '$lib/types/collection';

export const additemKey = Symbol();

export interface ContextProps {
	collection: CollectionProps;
	items: CollectionItemProps[];
	itemsKeys: string[];
	settings: SettingsProps;
}
