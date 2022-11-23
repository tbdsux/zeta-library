<script lang="ts">
	import type { CollectionItemProps } from '$lib/items/props';
	import {
		CheckIcon,
		ExternalLinkIcon,
		PlusCircleIcon,
		XIcon
	} from '@rgossiaux/svelte-heroicons/solid';
	import { getPageContext } from './context';

	const { itemsKeys } = getPageContext();

	export let searchResults: CollectionItemProps[];
	export let selectedItems: CollectionItemProps[];
	export let selectedIds: string[];

	const select = (item: CollectionItemProps) => {
		selectedItems = [
			{ ...item, date_added: Math.floor(new Date().getTime() / 1000) },
			...selectedItems
		];
		selectedIds = [item.item_id, ...selectedIds];
	};

	const unselect = (item: CollectionItemProps) => {
		selectedItems = selectedItems.filter((i) => i.item_id != item.item_id);
		selectedIds = selectedIds.filter((i) => i != item.item_id);
	};
</script>

<div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 lg:gap-8">
	{#each searchResults as result}
		<div class="group relative shadow rounded-lg">
			{#if itemsKeys.includes(result.item_id)}
				<span
					title="Item already exists in collection"
					class="inline-flex items-center bg-indigo-500 text-white absolute z-[60] top-1 left-1 py-1 px-2 rounded-lg"
				>
					<CheckIcon class="h-4 w-4 mr-1" />
					<small>added</small>
				</span>
			{/if}

			{#if selectedIds.includes(result.item_id)}
				<button
					class="absolute top-1 left-1 inline-flex items-center text-white bg-blue-500 rounded-lg py-0.5 px-2"
				>
					<CheckIcon class="h-4 w-4" />
					<small class="ml-1">selected</small>
				</button>
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

					{#if !itemsKeys.includes(result.item_id)}
						{#if selectedIds.includes(result.item_id)}
							<button
								on:click={() => unselect(result)}
								class="text-white bg-red-400 hover:bg-red-500 duration-300 py-2 px-6 rounded-lg inline-flex items-center text-sm mt-6"
							>
								<XIcon class="h-4 w-4 ml-1" aria-hidden="true" />
								Unselect Item
							</button>
						{:else}
							<button
								on:click={() => select(result)}
								class="text-white bg-blue-400 hover:bg-blue-500 duration-300 py-2 px-6 rounded-lg inline-flex items-center text-sm mt-6"
							>
								<PlusCircleIcon class="h-4 w-4 ml-1" aria-hidden="true" />
								Select Item
							</button>
						{/if}
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
	{/each}
</div>
