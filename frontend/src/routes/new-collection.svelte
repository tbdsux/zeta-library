<script lang="ts">
	import { goto } from '$app/navigation';
	import { collectionTypes } from '$lib/collection';
	import { apiUrl } from '$lib/config';
	import Modal from '$lib/modal.svelte';
	import type { PartialCollectionProps } from '$lib/types/collection';
	import {
		DialogTitle,
		Listbox,
		ListboxButton,
		ListboxOption,
		ListboxOptions,
		Transition
	} from '@rgossiaux/svelte-headlessui';
	import {
		SelectorIcon,
		CheckIcon,
		DocumentAddIcon,
		FolderAddIcon
	} from '@rgossiaux/svelte-heroicons/solid';

	let selectedType = collectionTypes[0];
	let inputName: string = '';
	let inputDescription: string = '';
	let isOpen = false;

	const createCollection = async () => {
		if (inputName == '' || selectedType.value == '') return;

		const body: PartialCollectionProps = {
			name: inputName,
			type: selectedType.value,
			description: inputDescription,
			created_at: Math.floor(new Date().getTime() / 1000)
		};

		const r = await fetch(apiUrl + '/collections/create', {
			method: 'POST',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify(body)
		});
		const data = await r.json();

		if (!r.ok) {
			// error in here
		}

		isOpen = false;
		goto(`/collection/${data.data.id}`);
	};
</script>

<Modal
	{isOpen}
	closeModal={() => {
		isOpen = false;
	}}
	className="max-w-xl"
>
	<DialogTitle class="text-xl font-extrabold text-indigo-500">Create New Collection</DialogTitle>

	<div class="p-6">
		<div class="flex flex-col my-1">
			<label for="collection-name" class="text-gray-700">Collection Name</label>
			<input
				bind:value={inputName}
				type="text"
				name="collection-name"
				id=""
				class="py-2 px-5 rounded-lg border text-gray-700"
				placeholder="Enter collection name"
			/>
		</div>

		<div class="flex flex-col my-1">
			<label for="collection-description" class="text-gray-700">Type</label>
			<Listbox value={selectedType} on:change={(e) => (selectedType = e.detail)}>
				<div class="relative mt-1">
					<ListboxButton
						class="relative w-full py-3 pl-5 pr-10 text-left bg-white rounded-lg border hover:bg-gray-100 duration-300 cursor-pointer focus:outline-none focus-visible:ring-2 focus-visible:ring-opacity-75 focus-visible:ring-white focus-visible:ring-offset-orange-300 focus-visible:ring-offset-2 focus-visible:border-indigo-500 sm:text-sm text-gray-700"
					>
						<span class="block truncate">
							{#if selectedType.value == ''}
								Select a type...
							{:else}
								{selectedType.name}
							{/if}
						</span>
						<span class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
							<SelectorIcon class="w-5 h-5 text-gray-400" aria-hidden="true" />
						</span>
					</ListboxButton>

					<Transition
						leave="transition ease-in duration-100"
						leaveFrom="opacity-100"
						leaveTo="opacity-0"
					>
						<ListboxOptions
							class="absolute w-full py-1 mt-1 overflow-auto text-base bg-white rounded-md shadow-lg max-h-40 ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
						>
							{#each collectionTypes as i, index (index)}
								<ListboxOption
									class={({ active }) =>
										`cursor-default select-none relative py-2 pl-10 pr-4 ${
											active ? 'text-gray-900 bg-gray-100' : 'text-gray-700'
										}`}
									value={i}
									let:selected
								>
									<span class={`block truncate ${selected ? 'font-medium' : 'font-normal'}`}>
										{i.name}
									</span>
									{#if selected}
										<span class="absolute inset-y-0 left-0 flex items-center pl-3 text-gray-600">
											<CheckIcon class="w-5 h-5" aria-hidden="true" />
										</span>
									{/if}
								</ListboxOption>
							{/each}
						</ListboxOptions>
					</Transition>
				</div>
			</Listbox>
		</div>

		<div class="flex flex-col my-1">
			<label for="collection-description" class="text-gray-700">Description</label>
			<textarea
				bind:value={inputDescription}
				type="text"
				name="collection-description"
				id=""
				class="py-2 px-5 rounded-lg border text-gray-700"
				placeholder="Add a short description (not required)"
			/>
		</div>

		<div class="text-right mt-6">
			<button
				on:click={createCollection}
				class="inline-flex items-center py-3 px-6 rounded-lg bg-indigo-400 hover:bg-indigo-500 duration-300 text-white text-sm"
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
	title="New Collection"
	on:click={() => (isOpen = true)}
	class="inline-flex items-center py-2 px-6 rounded-lg bg-indigo-400 hover:bg-indigo-500 duration-300 text-white"
>
	<FolderAddIcon class="h-4 w-4 mr-1" aria-hidden="true" />

	<small class="hidden sm:block">new collection</small>
</button>
