<script lang="ts">
	import { invalidate } from '$app/navigation';
	import { apiUrl } from '$lib/config';
	import { fetcher } from '$lib/items/fetcher';
	import type { CollectionItemProps } from '$lib/items/props';
	import type { CollectionProps } from '$lib/types/collection';
	import { DialogDescription, DialogTitle } from '@rgossiaux/svelte-headlessui';
	import { CheckIcon, SearchIcon } from '@rgossiaux/svelte-heroicons/solid';
	import { getAddItemsContext } from './additem';
	import AddItemPreview from './AddItemPreview.svelte';
	import AddItemSearchResult from './AddItemSearchResult.svelte';

	const { collection } = getAddItemsContext();
	export let isOpen: boolean;

	let inputQuery: string = '';
	let fetching = false;
	let error = false;

	let saving = false;

	let searchResults: CollectionItemProps[] = [];
	let selectedItems: CollectionItemProps[] = [];
	let selectedIds: string[] = [];

	const search = async () => {
		if (inputQuery == '') return;

		fetching = true;

		try {
			searchResults = await fetcher[collection.type](inputQuery);
		} catch (e) {
			error = true;
			console.error(e);
		}

		fetching = false;
	};

	const addItems = async () => {
		saving = true;

		const _body = {
			id: collection.id,
			data: selectedItems
		};

		const r = await fetch(apiUrl + '/items', {
			method: 'PATCH',
			body: JSON.stringify(_body),
			headers: {
				'content-type': 'application/json'
			}
		});

		const data = await r.json();

		if (!r.ok) {
			// handle error
			console.log(data);
		}

		// re-load fetch
		invalidate(apiUrl + `/items/${collection.id}`);

		saving = false;

		// close modal
		isOpen = false;
	};
</script>

<div class="flex items-center justify-between">
	<div>
		<DialogTitle class="text-xl font-extrabold text-gray-700">Add Items</DialogTitle>
		<DialogDescription class="text-gray-600">
			Add items to the collection
			<span class="font-medium">{collection.name}</span>
		</DialogDescription>
	</div>

	<div class="inline-flex items-center">
		<input
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

		<div class="inline-flex items-stretch ml-6 ">
			<AddItemPreview bind:selectedItems bind:selectedIds />

			<button
				on:click={addItems}
				disabled={saving}
				title="Done with selection"
				class="py-2 px-6 rounded-lg ml-1 inline-flex items-center bg-blue-400 disabled:opacity-90 disabled:hover:bg-blue-400 hover:bg-blue-500 text-white duration-300"
			>
				<CheckIcon class="h-4 w-4" aria-hidden="true" />
				<small class="ml-1" title="Done with selection">
					{#if saving}
						Saving....
					{:else}
						Done
					{/if}</small
				>
			</button>
		</div>
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
		<AddItemSearchResult {searchResults} bind:selectedItems bind:selectedIds />
	{/if}
</div>
