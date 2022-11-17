import type { CollectionItemProps } from '$lib/items/props';
import type { SettingsProps } from '$lib/settings';
import type { CollectionProps } from '$lib/types/collection';
import { getContext } from 'svelte';

export const additemKey = Symbol();

export interface ContextProps {
	collection: CollectionProps;
	items: CollectionItemProps[];
	itemsKeys: string[];
	settings: SettingsProps;
}

export const getPageContext = () => getContext<ContextProps>(additemKey);
