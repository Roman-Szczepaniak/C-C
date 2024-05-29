import type { PageLoad } from './$types';
import { get } from '$lib/api';

export const load = (async ({ params }) => {
	return {
		title: 'Create monster'
        /* Besoin de get card_id */
	};
}) satisfies PageLoad;