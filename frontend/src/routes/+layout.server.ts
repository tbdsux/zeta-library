import { fetchHeaders, serverApiUrl } from '$lib/server/config';
import type { SettingsProps } from '$lib/settings';
import type { APIResponseProps } from '$lib/types/api';
import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ fetch, depends }) => {
	const r = await fetch(serverApiUrl + '/settings', {
		headers: fetchHeaders
	});
	depends('load:settings');

	const data: APIResponseProps<SettingsProps> = await r.json();

	if (!r.ok) {
		throw error(r.status, data.message);
	}

	return {
		settings: data.data
	};
};
