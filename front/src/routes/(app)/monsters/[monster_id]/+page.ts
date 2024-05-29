import type { PageLoad } from './$types';
import { get } from '$lib/api';

export const load = (async ({ params }) => {
	return {
		title: 'Monster',
        monster: get(`/monsters/${params.monster_id}`)
        /* Besoin de get card_id */
	};
}) satisfies PageLoad;