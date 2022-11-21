<script lang="ts">
	import { invalidate } from '$app/navigation';
	import { apiUrl } from '$lib/config';
	import type { CollectionItemProps } from '$lib/items/props';
	import Modal from '$lib/modal.svelte';
	import { DialogDescription, DialogTitle } from '@rgossiaux/svelte-headlessui';
	import { ArrowLeftIcon, TrashIcon, XIcon } from '@rgossiaux/svelte-heroicons/solid';

	export let item: CollectionItemProps;
	export let collectionId: string;

	let isOpen = false;
	let removing = false;

	const removeItem = async () => {
		removing = true;
		const key = item.key;

		const r = await fetch(apiUrl + `/items/${collectionId}/${key}`, {
			method: 'DELETE'
		});

		const data = await r.json();

		if (!r.ok) {
			// handle error
			console.log(data);
		}

		// re-load fetch
		invalidate('load:items');

		removing = false;

		// close modal
		isOpen = false;
	};
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-3xl p-6">
	<DialogTitle class="text-xl font-bold text-gray-800">Remove Item</DialogTitle>
	<DialogDescription class="text-lg text-gray-700">
		Are you sure you want to remove this item <span class="font-medium">{item.name}</span> from the collection?
	</DialogDescription>

	<div class="text-right mt-4">
		<button
			on:click={() => (isOpen = false)}
			class="inline-flex items-center py-3 px-8 rounded-lg bg-gray-400 hover:bg-gray-500 text-white duration-300"
		>
			<ArrowLeftIcon class="h-4 w-4 mr-1" />
			<small>CANCEL</small>
		</button>
		<button
			on:click={removeItem}
			disabled={removing}
			class="ml-1 inline-flex items-center py-3 px-8 rounded-lg bg-red-400 disabled:opacity-90 disabled:hover:bg-red-400 hover:bg-red-500 text-white duration-300"
		>
			<TrashIcon class="h-4 w-4 mr-1" />
			<small>
				{#if removing}
					Removing...
				{:else}
					REMOVE
				{/if}
			</small>
		</button>
	</div>
</Modal>

<button
	on:click={() => {
		isOpen = true;
	}}
	class="text-white bg-red-400 hover:bg-red-500 duration-300 py-2 px-6 rounded-lg inline-flex items-center text-sm mt-6"
>
	<XIcon class="h-4 w-4 ml-1" aria-hidden="true" />
	Remove Item
</button>
