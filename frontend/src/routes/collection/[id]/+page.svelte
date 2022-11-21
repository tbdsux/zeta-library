<script lang="ts">
	import { ExternalLinkIcon } from '@rgossiaux/svelte-heroicons/solid';
	import AddItem from './AddItem.svelte';
	import Settings from './Settings.svelte';
	import type { PageServerData } from './$types';
	import RemoveItem from './RemoveItem.svelte';
	import Context from './Context.svelte';

	export let data: PageServerData;
</script>

<svelte:head>
	<title>{data.collection.name} - Collection</title>
</svelte:head>

<Context {data}>
	<div class="flex items-center justify-between">
		<div>
			<div class="inline-flex items-center">
				<h3 class="text-4xl font-black text-indigo-500">{data.collection.name}</h3>

				<p class="ml-2 text-sm py-0.5 px-2 rounded-lg bg-gray-400 text-white">
					{data.collection.type}
				</p>
			</div>
			<p class="text-lg text-gray-700">{data.collection.description}</p>
			<small class="mt-1 text-gray-600">{data.items.length} items</small>
		</div>

		<div class="inline-flex items-center">
			<Settings />

			<AddItem />
		</div>
	</div>

	<hr class="my-6" />

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

						<RemoveItem collectionId={data.collection.id} {item} />
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
</Context>
