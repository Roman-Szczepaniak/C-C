import { goto } from '$app/navigation';
import { post } from '$lib/api';

export const signup = async (firstname: string, lastname: string, email: string, hashpassword: string) => {
	await post('/users/signup', {
		FirstName: firstname,
		LastName: lastname,
		HashPassword: hashpassword,
		Email: email
	});

	goto('/login');
}

export const login = async (email: string, hashpassword: string) => {
	const response = await post('/users/login', {
		email: email,
		hashpassword: hashpassword
	});

	localStorage.setItem('access_token', response['token']);
	localStorage.setItem('refresh_token', response['refresh_token']);

	goto(localStorage.getItem('redirect_after_login') ?? '/');
};

export const logout = () => {
	localStorage.removeItem('access_token');
	localStorage.removeItem('refresh_token');

	goto('/login');
};