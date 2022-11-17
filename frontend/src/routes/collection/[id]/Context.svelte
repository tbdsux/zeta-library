<script lang="ts">
	import { setContext } from 'svelte';
	import { additemKey, type ContextProps } from './context';
	import type { PageData } from './$types';

	export let data: PageData;

	const itemsKeys: string[] = [];
	if (data.items) {
		for (const i of data.items) {
			itemsKeys.push(i.item_id);
		}
	}

	setContext<ContextProps>(additemKey, { ...data, itemsKeys });
</script>

<div>
	<slot />
</div>
