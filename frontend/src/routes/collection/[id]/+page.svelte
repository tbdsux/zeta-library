<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>{data.collection.name}</title>
</svelte:head>

<p>{JSON.stringify(data)}</p>
