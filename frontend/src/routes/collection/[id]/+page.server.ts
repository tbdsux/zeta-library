import { fetchHeaders, serverApiUrl } from '$lib/server/config';
import type { CollectionItemProps } from '$lib/items/props';
import type { APIResponseProps } from '$lib/types/api';
import type { CollectionProps } from '$lib/types/collection';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params, parent, depends }) => {
	const colRes = await fetch(serverApiUrl + `/collections/get/${params.id}`, {
		headers: fetchHeaders
	});
	const itemsRes = await fetch(serverApiUrl + `/items/${params.id}`, {
		headers: fetchHeaders
	});
	const { settings } = await parent();

	depends('load:items');

	const colData: APIResponseProps<CollectionProps> = await colRes.json();
	const itemsData: APIResponseProps<CollectionItemProps[]> = await itemsRes.json();

	if (colRes.status != 200) {
		throw error(colRes.status, colData.message);
	}

	if (itemsRes.status != 200) {
		throw error(itemsRes.status, itemsData.message);
	}

	return {
		collection: colData.data,
		items: itemsData.data,
		settings
	};
};
