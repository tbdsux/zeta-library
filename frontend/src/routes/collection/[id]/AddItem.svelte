<script lang="ts">
	import Modal from '$lib/modal.svelte';
	import { PlusIcon } from '@rgossiaux/svelte-heroicons/solid';
	import AddItemContainer from './AddItemContainer.svelte';

	let isOpen = false;
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-6xl 2xl:max-w-7xl p-6">
	<AddItemContainer bind:isOpen />
</Modal>

<button
	on:click={() => (isOpen = true)}
	title="Add Item"
	class="py-2 px-6 rounded-lg m-1 inline-flex items-center bg-blue-400 hover:bg-blue-500 text-white duration-300"
>
	<PlusIcon class="h-4 w-4" aria-hidden="true" />
	<small class="ml-1 uppercase font-medium">Add Item</small>
</button>
