import { apiUrl } from '$lib/config';
import type { CollectionItemProps } from '$lib/items/props';
import type { APIResponseProps } from '$lib/types/api';
import type { CollectionProps } from '$lib/types/collection';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params, parent }) => {
	const colRes = await fetch(apiUrl + `/collections/get/${params.id}`);
	const itemsRes = await fetch(apiUrl + `/items/${params.id}`);
	const { settings } = await parent();

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
