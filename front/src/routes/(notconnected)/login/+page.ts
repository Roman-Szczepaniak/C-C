import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load = (() => {
	if (localStorage.getItem('access_token')) {
		throw redirect(307, '/');
	}

	return {
		title: 'Login'
	};
}) satisfies LayoutServerLoad;