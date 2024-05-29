import type { PageLoad } from './$types';
import { get } from '$lib/api';

export const load = (async ({ }) => {
	const profile = await get(`/profile`);
	return {
		title: 'Profile',
		profile: profile
	};
}) satisfies PageLoad;