import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { jwtDecode } from "jwt-decode";

export const load = (({ route, url }) => {
	if (!localStorage.getItem('access_token')) {
		if (route.id !== '/(notconnected)/login' && route.id !== '/(notconnected)/signup') {
			// Enregistre l'URL actuelle dans le stockage local
			localStorage.setItem('redirect_after_login', url.pathname);
			throw redirect(307, '/login');
		}
	} else {
		return {
			user: jwtDecode(localStorage.getItem('access_token'))
		};
	}
}) satisfies LayoutServerLoad;

// Pas de rendu côté serveur
export const ssr = false;