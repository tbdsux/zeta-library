<script lang="ts">
	import type { CollectionItemProps } from '$lib/items/props';
	import { CheckIcon, ExternalLinkIcon, PlusCircleIcon } from '@rgossiaux/svelte-heroicons/solid';

	import toast from 'svelte-french-toast';
	import { invalidate } from '$app/navigation';
	import { getContext } from 'svelte';
	import type { Writable } from 'svelte/store';
	import { additemKey, type ContextProps } from './context';

	let state = getContext<Writable<ContextProps>>(additemKey);

	export let result: CollectionItemProps;
	let adding = false;

	const addItem = async () => {
		const finalItem = {
			...result,
			date_added: Math.floor(new Date().getTime() / 1000)
		};

		const body = {
			id: $state.collection.id,
			data: [finalItem]
		};

		adding = true;

		const r = await fetch('/api/items', {
			method: 'PATCH',
			body: JSON.stringify(body),
			headers: {
				'content-type': 'application/json'
			}
		});

		const data = await r.json();
		adding = false;
		if (!r.ok) {
			toast.error(data.message);
			return;
		}

		toast.success(`Successfully added item '${result.name}'`);

		invalidate('load:items');
	};
</script>

<div class="group relative shadow rounded-lg">
	{#if $state.itemsKeys.includes(result.item_id)}
		<span
			title="Item already exists in collection"
			class="inline-flex items-center bg-indigo-500 text-white absolute z-[60] top-1 left-1 py-1 px-2 rounded-lg"
		>
			<CheckIcon class="h-4 w-4 mr-1" />
			<small>added</small>
		</span>
	{/if}

	<div
		class="absolute rounded-lg w-full h-full group-hover:bg-black/70 duration-300 flex items-center justify-center"
	>
		<div class="hidden group-hover:block p-3 text-center">
			<a
				href={result.url}
				target="_blank"
				rel="noreferrer"
				class="p-1 bg-white text-gray-700 hover:bg-gray-100 duration-300 rounded-lg absolute top-2 right-2"
				title="Visit original"
			>
				<ExternalLinkIcon class="h-4 w-4" aria-hidden="true" />
			</a>

			<strong class="text-white line-clamp-3">{result.name}</strong>

			{#if !$state.itemsKeys.includes(result.item_id)}
				<button
					on:click={() => addItem()}
					class="text-white bg-blue-400 hover:bg-blue-500 duration-300 py-2 px-6 rounded-lg inline-flex items-center text-sm mt-6"
				>
					<PlusCircleIcon class="h-4 w-4 ml-1" aria-hidden="true" />
					{#if adding}
						Adding...
					{:else}
						Add Item
					{/if}
				</button>
			{/if}
		</div>
	</div>

	{#if result.image}
		<div class="">
			<img src={result.image} class="h-72 w-full object-cover rounded-t-lg" alt={result.name} />
		</div>
	{:else}
		<div class="h-72 w-full bg-gray-200 border rounded-t-lg flex items-center justify-center">
			No image found
		</div>
	{/if}

	<div class="py-4 px-2 border rounded-b-lg text-center bg-gray-50">
		<strong title={result.name} class="line-clamp-1 text-gray-800">{result.name}</strong>
	</div>
</div>
