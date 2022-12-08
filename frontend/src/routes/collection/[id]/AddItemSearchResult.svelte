<script lang="ts">
	import type { CollectionItemProps } from '$lib/items/props';
	import SearchResultItem from './SearchResultItem.svelte';

	export let searchResults: CollectionItemProps[];
</script>

<div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 lg:gap-8">
	{#each searchResults as result}
		<SearchResultItem {result} />
	{/each}
</div>
