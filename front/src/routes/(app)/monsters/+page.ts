import type { PageLoad } from './$types';
import { get } from '$lib/api';

export const load = (async ({ depends }) => {
	depends('app:monsters');

	const monsters = await get('/monsters');
	return {
		title: 'Monsters',
		monsters: monsters || []  // Assurez-vous que monsters est une liste
	};
}) satisfies PageLoad;