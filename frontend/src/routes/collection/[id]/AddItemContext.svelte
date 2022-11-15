<script lang="ts">
	import type { CollectionItemProps } from '$lib/items/props';
	import type { CollectionProps } from '$lib/types/collection';
	import { setContext } from 'svelte';
	import { additemKey, type AddItemContextProps } from './additem';

	export let collection: CollectionProps;
	export let items: CollectionItemProps[];

	const itemsKeys: string[] = [];
	if (items) {
		for (const i of items) {
			itemsKeys.push(i.item_id);
		}
	}

	setContext<AddItemContextProps>(additemKey, { collection, items, itemsKeys });
</script>

<div>
	<slot />
</div>
