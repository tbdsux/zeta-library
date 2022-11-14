<script lang="ts">
	import { page } from '$app/stores';
	import { PlusIcon } from '@rgossiaux/svelte-heroicons/solid';
	import Settings from './settings.svelte';

	const { collection } = $page.data;
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
		<small class="mt-1 text-gray-600">12 items</small>
	</div>

	<div class="inline-flex items-center">
		<Settings {collection} />

		<button
			title="Add Item"
			class="py-2 px-6 rounded-lg m-1 inline-flex items-center bg-blue-400 hover:bg-blue-500 text-white duration-300"
		>
			<PlusIcon class="h-4 w-4" aria-hidden="true" />
			<small class="ml-1 uppercase font-medium">Add Item</small>
		</button>
	</div>
</div>

<hr class="my-6" />

<slot />
