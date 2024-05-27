import type { PageLoad } from '../$types';

export const load = (() => {
	return { title: 'Accueil' };
}) satisfies PageLoad;