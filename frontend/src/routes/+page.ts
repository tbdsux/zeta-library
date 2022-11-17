import { apiUrl } from '$lib/config';
import type { APIResponseProps } from '$lib/types/api';
import type { CollectionProps } from '$lib/types/collection';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, parent }) => {
	const res = await fetch(apiUrl + `/collections`);
	const data: APIResponseProps<CollectionProps[]> = await res.json();

	const { settings } = await parent();

	if (res.status != 200) {
		throw error(res.status, data.message);
	}

	return {
		collections: data.data,
		settings
	};
};
