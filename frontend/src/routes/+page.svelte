<script lang="ts">
	import type { PageServerData } from './$types';
	import AppSettings from './AppSettings.svelte';

	export let data: PageServerData;
</script>

<svelte:head>
	<title>Zeta Library | Create and save your personal collections</title>
</svelte:head>

<div class="flex items-center justify-between">
	<p class="text-gray-600">My Collections</p>

	<AppSettings settings={data.settings} />
</div>

<div class="w-11/12 mx-auto grid grid-cols-3 gap-12 mt-6">
	{#each data.collections.sort((x, y) => y.created_at - x.created_at) as collection}
		<a
			href={`/collection/${collection.id}`}
			class="py-12 px-8 flex items-center rounded-lg relative shadow-lg bg-indigo-400 hover:bg-indigo-500 text-white duration-300"
		>
			<div>
				<strong class="text-xl">{collection.name}</strong>
				<p class="line-clamp-2">
					{collection.description}
				</p>
			</div>

			<div class="text-sm">
				<span class="absolute bottom-2 right-2">{collection.type}</span>
			</div>
		</a>
	{/each}
</div>
