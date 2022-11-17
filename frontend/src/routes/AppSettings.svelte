<script lang="ts">
	import { invalidate } from '$app/navigation';
	import { apiUrl } from '$lib/config';
	import Modal from '$lib/modal.svelte';
	import type { SettingsProps } from '$lib/settings';
	import { DialogDescription, DialogTitle } from '@rgossiaux/svelte-headlessui';
	import { CogIcon, PencilIcon } from '@rgossiaux/svelte-heroicons/solid';

	export let settings: SettingsProps;

	let isOpen = false;
	let updating = false;

	let inputApiKey = settings.moviedbApiKey;

	const updateSettings = async () => {
		if (inputApiKey.trim() == settings.moviedbApiKey) {
			return;
		}

		updating = true;

		const r = await fetch(apiUrl + '/settings', {
			method: 'PATCH',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({ moviedbApiKey: inputApiKey })
		});
		const data = await r.json();

		if (!r.ok) {
			// handle error in here
		}

		// re-load
		invalidate(apiUrl + '/settings');

		updating = false;
	};
</script>

<Modal {isOpen} closeModal={() => (isOpen = false)} className="max-w-2xl">
	<DialogTitle class="text-xl font-bold text-gray-700">App Settings</DialogTitle>
	<DialogDescription class="text-gray-600"
		>Modify the configurations being used by this application.</DialogDescription
	>

	<div class="mt-6 w-5/6 mx-auto">
		<div class="flex flex-col my-1">
			<label for="api-key" class="text-gray-700">
				MovieDB API Key
				<small>
					<a
						title="How to get api key?"
						class="text-gray-500 hover:underline"
						href="https://www.themoviedb.org/documentation/api"
						target="_blank"
						rel="noreferrer"
					>
						( how to? )
					</a>
				</small>
			</label>
			<input
				bind:value={inputApiKey}
				type="text"
				name="api-key"
				id=""
				class="py-2 px-5 rounded-lg border text-gray-700"
				placeholder="Enter api key"
			/>
		</div>

		<div class="text-right mt-6">
			<button
				disabled={updating}
				on:click={updateSettings}
				class="inline-flex items-center py-3 px-6 rounded-lg bg-indigo-400 disabled:opacity-90 disabled:hover:bg-indigo-400 hover:bg-indigo-500 duration-300 text-white text-sm"
			>
				<PencilIcon />
				{#if updating}
					Updating...
				{:else}
					Update
				{/if}
			</button>
		</div>
	</div>
</Modal>

<button
	on:click={() => (isOpen = true)}
	title="App Settings"
	class="ml-1 py-1 px-3 rounded-lg bg-gray-400 hover:bg-gray-500 text-white duration-300 inline-flex items-center"
>
	<CogIcon class="h-4 w-4" aria-hidden="true" />
	<small class="ml-1">Config</small>
</button>
