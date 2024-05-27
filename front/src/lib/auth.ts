import { goto } from '$app/navigation';
import { post } from '$lib/api';

export const login = async (username: string, password: string) => {
	const response = await post('/users/login', new URLSearchParams({ username, password }));

	localStorage.setItem('access_token', response['access_token']);
	localStorage.setItem('refresh_token', response['refresh_token']);

	goto(localStorage.getItem('redirect_after_login') ?? '/');
};

export const logout = () => {
	localStorage.removeItem('access_token');
	localStorage.removeItem('refresh_token');

	goto('/login');
};