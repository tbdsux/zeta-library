<script lang="ts">
	import Modal from '$lib/modal.svelte';
	import { DialogTitle } from '@rgossiaux/svelte-headlessui';

	let isOpen = false;
</script>

<Modal
	{isOpen}
	closeModal={() => {
		isOpen = false;

		console.log(isOpen);
	}}
	className="max-w-xl"
>
	<DialogTitle class="text-xl font-extrabold text-indigo-500">Create New Collection</DialogTitle>

	<div class="p-6">
		<div class="flex flex-col">
			<label for="collection-name" class="text-gray-700">Collection Name:</label>
			<input
				type="text"
				name="collection-name"
				id=""
				class="py-2 px-5 rounded-lg border font-medium text-gray-700"
				placeholder="Enter collection name"
			/>
		</div>

		<div class="text-right mt-4">
			<button
				class="inline-flex items-center py-2 px-6 rounded-lg bg-indigo-400 hover:bg-indigo-500 duration-300 text-white text-sm"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="w-3 h-3 mr-1"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M12 10.5v6m3-3H9m4.06-7.19l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z"
					/>
				</svg> create
			</button>
		</div>
	</div>
</Modal>

<button
	on:click={() => (isOpen = true)}
	class="inline-flex items-center py-2 px-6 rounded-lg bg-indigo-400 hover:bg-indigo-500 duration-300 text-white text-sm"
>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		fill="none"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		stroke="currentColor"
		class="w-5 h-5 mr-1"
	>
		<path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M12 10.5v6m3-3H9m4.06-7.19l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z"
		/>
	</svg>

	new collection
</button>
