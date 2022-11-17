<script lang="ts">
	import '../app.css';

	import '@fontsource/work-sans/300.css';
	import '@fontsource/work-sans/400.css';
	import '@fontsource/work-sans/500.css';
	import '@fontsource/work-sans/700.css';
	import '@fontsource/work-sans/800.css';
	import '@fontsource/work-sans/900.css';
	import NewCollection from './new-collection.svelte';
</script>

<header class="py-4">
	<div class="flex items-center justify-between w-5/6 mx-auto">
		<a href="/">
			<h1 class="tracking-wide text-2xl font-black text-indigo-500">zeta</h1>
			<p class="text-sm text-gray-700">Your personal library collections</p>
		</a>

		<div class="inline-flex items-stretch">
			<NewCollection />
		</div>
	</div>
</header>

<hr />

<main class="w-5/6 mx-auto my-12">
	<slot />
</main>
