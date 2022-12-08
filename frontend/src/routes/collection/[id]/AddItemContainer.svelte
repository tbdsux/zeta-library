<script lang="ts">
	import { fetcher } from '$lib/items/fetcher';
	import type { CollectionItemProps } from '$lib/items/props';
	import { DialogDescription, DialogTitle } from '@rgossiaux/svelte-headlessui';
	import { CheckIcon, SearchIcon } from '@rgossiaux/svelte-heroicons/solid';
	import { getContext } from 'svelte';
	import type { Writable } from 'svelte/store';
	import AddItemSearchResult from './AddItemSearchResult.svelte';
	import { additemKey, type ContextProps } from './context';

	let state = getContext<Writable<ContextProps>>(additemKey);

	export let isOpen: boolean;

	let inputQuery: string = '';
	let fetching = false;
	let error = false;

	let searchResults: CollectionItemProps[] = [];

	const search = async () => {
		if (inputQuery == '') return;

		// do not search if missing api key in app settings
		if ($state.collection.type == 'movies' || $state.collection.type == 'series') {
			if ($state.settings.moviedbApiKey == '') return;
		}

		fetching = true;

		try {
			searchResults = await fetcher[$state.collection.type](
				inputQuery,
				$state.settings.moviedbApiKey
			);
		} catch (e) {
			error = true;
			console.error(e);
		}

		fetching = false;
	};
</script>

<div class="flex flex-wrap items-center justify-between">
	<div class="m-2">
		<DialogTitle class="text-xl font-extrabold text-gray-700">Add Items</DialogTitle>
		<DialogDescription class="text-gray-600">
			Add items to the collection
			<span class="font-medium">{$state.collection.name}</span>
		</DialogDescription>

		{#if $state.collection.type == 'movies' || $state.collection.type == 'series'}
			{#if $state.settings.moviedbApiKey == ''}
				<p class="text-sm text-yellow-500">TheMovieDB API Key not set in app config settings.</p>
			{/if}
		{/if}
	</div>

	<div class="inline-flex flex-wrap justify-center items-stretch m-2">
		<input
			on:keydown={async (e) => {
				if (e.key === 'Enter') {
					await search();
				}
			}}
			bind:value={inputQuery}
			type="text"
			name="search"
			id="search"
			class="py-2 px-4 rounded-lg border"
			placeholder="Search for item..."
		/>
		<button
			on:click={search}
			class="p-3 bg-indigo-400 hover:bg-indigo-500 text-white rounded-lg ml-1 duration-300"
		>
			<SearchIcon class="h-4 w-4" aria-hidden="true" />
		</button>

		<button
			on:click={() => (isOpen = false)}
			title="Done with selection"
			class="py-2 px-6 rounded-lg ml-2 inline-flex items-center bg-blue-400 disabled:opacity-90 disabled:hover:bg-blue-400 hover:bg-blue-500 text-white duration-300"
		>
			<CheckIcon class="h-4 w-4" aria-hidden="true" />
			<small class="ml-1" title="Done with selection">Done</small>
		</button>
	</div>
</div>

<hr class="my-4" />

<div class="min-h-[24rem] max-h-screen overflow-auto">
	{#if fetching}
		<p>Fetching...</p>
	{:else if error}
		<p>
			There was a problem trying to search for the query. If problem persists, please report to the
			developer.
		</p>
	{:else}
		<AddItemSearchResult {searchResults} />
	{/if}
</div>
