<script lang="ts">
	import { collectionTypes, filterCollection } from '$lib/collection';
	import Modal from '$lib/modal.svelte';
	import type { CollectionProps } from '$lib/types/collection';
	import {
		DialogDescription,
		DialogTitle,
		Listbox,
		ListboxButton,
		ListboxOption,
		ListboxOptions,
		Transition
	} from '@rgossiaux/svelte-headlessui';
	import {
		CheckIcon,
		CogIcon,
		PencilIcon,
		SelectorIcon,
		TrashIcon
	} from '@rgossiaux/svelte-heroicons/solid';

	export let collection: CollectionProps;

	let isOpen = false;
	let inputNewName = collection.name;
	let newSelectedType = filterCollection(collection.type) ?? collectionTypes[0];
	let inputNewDescription = collection.description;

	const updateCollection = () => {};
	const removeCollection = () => {};
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-2xl">
	<div class="flex items-center justify-between">
		<div>
			<DialogTitle class="text-lg font-bold text-gray-700">Collection Settings</DialogTitle>
			<DialogDescription class="text-gray-600">
				Remove / update the collection's details
			</DialogDescription>
		</div>

		<button
			on:click={removeCollection}
			title="Delete collection"
			class="p-2 rounded-lg bg-red-400 hover:bg-red-500 text-white duration-300"
		>
			<TrashIcon class="h-4 w-4" aria-hidden="true" />
		</button>
	</div>

	<div class="mt-6 w-5/6 mx-auto">
		<div class="flex flex-col my-1">
			<label for="collection-id" class="text-gray-700">ID / Key</label>
			<input
				value={collection.id}
				type="text"
				name="collection-id"
				id=""
				disabled
				class="py-2 px-5 rounded-lg border bg-gray-200 text-gray-700"
			/>
		</div>

		<div class="flex flex-col my-1">
			<label for="collection-name" class="text-gray-700">Name</label>
			<input
				bind:value={inputNewName}
				type="text"
				name="collection-name"
				id=""
				class="py-2 px-5 rounded-lg border text-gray-700"
				placeholder="Enter collection name"
			/>
		</div>

		<div class="flex flex-col my-1">
			<label for="collection-description" class="text-gray-700">Type</label>
			<Listbox value={newSelectedType} on:change={(e) => (newSelectedType = e.detail)}>
				<div class="relative mt-1">
					<ListboxButton
						class="relative w-full py-3 pl-5 pr-10 text-left bg-white rounded-lg border hover:bg-gray-100 duration-300 focus:outline-none focus-visible:ring-2 focus-visible:ring-opacity-75 focus-visible:ring-white focus-visible:ring-offset-orange-300 focus-visible:ring-offset-2 focus-visible:border-indigo-500 sm:text-sm text-gray-700"
					>
						<span class="block truncate">
							{#if newSelectedType.value == ''}
								Select a type...
							{:else}
								{newSelectedType.name}
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
				bind:value={inputNewDescription}
				type="text"
				name="collection-description"
				id=""
				class="py-2 px-5 rounded-lg border text-gray-700"
				placeholder="Add a short description (not required)"
			/>
		</div>

		<div class="text-right mt-6">
			<button
				on:click={updateCollection}
				class="inline-flex items-center py-3 px-6 rounded-lg bg-indigo-400 hover:bg-indigo-500 duration-300 text-white text-sm"
			>
				<PencilIcon /> Update
			</button>
		</div>
	</div>
</Modal>

<button
	on:click={() => (isOpen = true)}
	title="Settings"
	class="p-2 rounded-lg m-1 inline-flex items-center bg-gray-400 hover:bg-gray-500 text-white duration-300"
>
	<CogIcon class="h-4 w-4" aria-hidden="true" />
	<small class="ml-1">Settings</small>
</button>
