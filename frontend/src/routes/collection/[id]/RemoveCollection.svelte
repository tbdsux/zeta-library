<script lang="ts">
	import { goto } from '$app/navigation';
	import Modal from '$lib/modal.svelte';
	import { DialogDescription, DialogTitle } from '@rgossiaux/svelte-headlessui';
	import { ArrowLeftIcon, TrashIcon } from '@rgossiaux/svelte-heroicons/solid';
	import toast from 'svelte-french-toast';
	import { getContext } from 'svelte';
	import type { Writable } from 'svelte/store';
	import { additemKey, type ContextProps } from './context';

	let state = getContext<Writable<ContextProps>>(additemKey);

	let isOpen = false;
	let removing = false;

	const removeCollection = async () => {
		removing = true;

		const r = await fetch(`/api/collections/${$state.collection.id}`, {
			method: 'DELETE'
		});

		const data = await r.json();

		if (!r.ok) {
			toast.error(data.message);
			return;
		}

		removing = false;

		// re-direct to index page
		goto('/');
	};
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-3xl p-6">
	<DialogTitle class="text-xl font-bold text-gray-800">Delete Collection</DialogTitle>
	<DialogDescription class="text-lg text-gray-700">
		Are you sure you want to delete this collection?
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
			on:click={removeCollection}
			disabled={removing}
			class="ml-1 inline-flex items-center py-3 px-8 rounded-lg bg-red-400 disabled:opacity-90 disabled:hover:bg-red-400 hover:bg-red-500 text-white duration-300"
		>
			<TrashIcon class="h-4 w-4 mr-1" />
			<small>
				{#if removing}
					Deleting...
				{:else}
					DELETE
				{/if}
			</small>
		</button>
	</div>
</Modal>

<button
	on:click={() => (isOpen = true)}
	title="Delete collection"
	class="p-2 rounded-lg bg-red-400 hover:bg-red-500 text-white duration-300"
>
	<TrashIcon class="h-4 w-4" aria-hidden="true" />
</button>
