<script lang="ts">
	import type { CollectionItemProps } from '$lib/items/props';
	import Modal from '$lib/modal.svelte';
	import { DialogTitle } from '@rgossiaux/svelte-headlessui';
	import {
		CheckIcon,
		ClipboardCheckIcon,
		ExternalLinkIcon,
		PlusCircleIcon,
		XIcon
	} from '@rgossiaux/svelte-heroicons/solid';

	export let selectedItems: CollectionItemProps[];
	export let selectedIds: string[];

	let isOpen = false;

	const unselect = (item: CollectionItemProps) => {
		selectedItems = selectedItems.filter((i) => i.item_id != item.item_id);
		selectedIds = selectedIds.filter((i) => i != item.item_id);
	};
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-4xl p-6">
	<DialogTitle class="font-medium text-gray-700">Preview selected items</DialogTitle>

	<div class="grid grid-cols-4 gap-6 min-h-[24rem] max-h-screen overflow-auto mt-6">
		{#each selectedItems as item}
			<div class="group relative w-full h-full">
				{#if selectedIds.includes(item.item_id)}
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
							href={item.url}
							target="_blank"
							rel="noreferrer"
							class="p-1 bg-white text-gray-700 hover:bg-gray-100 duration-300 rounded-lg absolute top-2 right-2"
							title="Visit original"
						>
							<ExternalLinkIcon class="h-4 w-4" aria-hidden="true" />
						</a>

						<strong class="text-white line-clamp-3">{item.name}</strong>

						<button
							on:click={() => unselect(item)}
							class="text-white bg-red-400 hover:bg-red-500 duration-300 py-2 px-3 rounded-lg inline-flex items-center text-sm mt-6"
						>
							<XIcon class="h-4 w-4 ml-1" aria-hidden="true" />
							Unselect Item
						</button>
					</div>
				</div>

				{#if item.image}
					<div class="">
						<img src={item.image} class="h-72 w-full object-cover rounded-t-lg" alt={item.name} />
					</div>
				{:else}
					<div class="h-72 w-full bg-gray-200 border rounded-t-lg flex items-center justify-center">
						No image found
					</div>
				{/if}

				<div class="py-4 px-2 border rounded-b-lg text-center bg-gray-50">
					<strong title={item.name} class="line-clamp-1 text-gray-800">{item.name}</strong>
				</div>
			</div>
		{/each}
	</div>
</Modal>

<button
	on:click={() => (isOpen = true)}
	title="Preview selected items"
	class="py-2 px-6 rounded-lg inline-flex items-center bg-gray-400 hover:bg-gray-500 text-white duration-300"
>
	<small class="mr-1">({selectedIds.length})</small>
	<ClipboardCheckIcon class="h-4 w-4" aria-hidden="true" />
</button>
