import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';

const config: UserConfig = {
	plugins: [sveltekit()],
	ssr: { noExternal: ['@rgossiaux/svelte-headlessui'] }
};

export default config;
