<script lang="ts">
	import { page } from '$app/stores';
	import AddItem from './AddItem.svelte';
	import AddItemContext from './AddItemContext.svelte';
	import Settings from './Settings.svelte';

	const { collection, items } = $page.data;
</script>

<svelte:head>
	<title>{$page.data.collection.name} - Collection</title>
</svelte:head>

<div class="flex items-center justify-between">
	<div>
		<div class="inline-flex items-center">
			<h3 class="text-4xl font-black text-indigo-500">{collection.name}</h3>

			<p class="ml-2 text-sm py-0.5 px-2 rounded-lg bg-gray-400 text-white">{collection.type}</p>
		</div>
		<p class="text-lg text-gray-700">{collection.description}</p>
		<small class="mt-1 text-gray-600">{items.length} items</small>
	</div>

	<div class="inline-flex items-center">
		<Settings {collection} />

		<AddItemContext {collection} {items}>
			<AddItem />
		</AddItemContext>
	</div>
</div>

<hr class="my-6" />

<slot />
