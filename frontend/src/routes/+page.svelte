<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Zeta Library | Create and save your personal collections</title>
</svelte:head>

<div class="">
	<p class="text-gray-600">My Collections</p>
</div>

<div class="w-11/12 mx-auto grid grid-cols-3 gap-12 mt-6">
	{#each data.collections as collection}
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
				<span class="absolute bottom-2 left-2">( 12 items )</span>
				<span class="absolute bottom-2 right-2">{collection.type}</span>
			</div>
		</a>
	{/each}
</div>
