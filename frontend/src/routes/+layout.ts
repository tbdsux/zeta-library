import { apiUrl } from '$lib/config';
import type { SettingsProps } from '$lib/settings';
import type { APIResponseProps } from '$lib/types/api';
import { error } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ fetch }) => {
	const r = await fetch(apiUrl + '/settings');
	const data: APIResponseProps<SettingsProps> = await r.json();

	if (!r.ok) {
		throw error(r.status, data.message);
	}

	return {
		settings: data.data
	};
};
