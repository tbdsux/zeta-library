<script lang="ts">
	import { ExternalLinkIcon, XIcon } from '@rgossiaux/svelte-heroicons/solid';
	import type { PageData } from './$types';

	export let data: PageData;

	const removeItem = (id: string) => {};
</script>

<div class="grid grid-cols-5 gap-6">
	{#each data.items.sort((x, y) => Number(y.date_added) - Number(x.date_added)) as item}
		<div class="group relative shadow rounded-lg">
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
						on:click={() => {}}
						class="text-white bg-red-400 hover:bg-red-500 duration-300 py-2 px-6 rounded-lg inline-flex items-center text-sm mt-6"
					>
						<XIcon class="h-4 w-4 ml-1" aria-hidden="true" />
						Remove Item
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
