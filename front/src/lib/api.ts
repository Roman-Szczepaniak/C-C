import { error, redirect } from '@sveltejs/kit';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { PUBLIC_API_URL } from '$env/static/public';
import jwt_decode, { type JwtPayload } from 'jwt-decode';
import * as notifications from '$lib/notifications';

export function formDataToJson(formData: FormData): Record<string, FormDataEntryValue | any> {
	return Object.fromEntries(formData.entries());
}

export interface APIRequestOptions {
	asResponse?: boolean;
	headers?: HeadersInit;
	timeout?: number;
}

async function send({
	method,
	path,
	data,
	options
}: {
	method: string;
	path: string | URL;
	data?: any;
	options?: APIRequestOptions;
}): Promise<any> {
	let opts: RequestInit & { headers: Headers } = {
		method,
		headers: new Headers(options?.headers)
	};

	if (data) {
		if (data instanceof URLSearchParams) {
			opts.headers.set('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
			opts.body = data;
		} else if (data instanceof FormData) {
			opts.headers.set('Content-Type', 'application/json');
			opts.body = JSON.stringify(formDataToJson(data));
		} else {
			opts.headers.set('Content-Type', 'application/json');
			opts.body = JSON.stringify(data);
		}
	}

	if (localStorage.getItem('access_token')) {
		opts.headers.set('Authorization', `Bearer ${localStorage.getItem('access_token')}`);
	}

	try {
		let timeoutId;
		if (options?.timeout) {
			const controller = new AbortController();
			timeoutId = setTimeout(() => controller.abort(), options.timeout);
			opts.signal = controller.signal;
		}
		const res = await fetch(new URL(path, PUBLIC_API_URL), opts);
		if (timeoutId) clearTimeout(timeoutId);
		if (res.status === 401) {
			const access_token = localStorage.getItem('access_token');
			if (!access_token) throw redirect(307, '/login');

			const decoded_token: JwtPayload = jwt_decode(access_token);
			if (!decoded_token.exp) throw redirect(307, '/login');
			const expiresAt = new Date(decoded_token.exp * 1000);
			if (expiresAt < new Date()) {
				const refresh_token = localStorage['refresh_token'];
				const jwt_refresh_response = await fetch(new URL('/users/token/refresh', PUBLIC_API_URL), {
					method: 'POST',
					headers: { Authorization: `Bearer ${refresh_token}` }
				});

				if (jwt_refresh_response.ok) {
					const jwt = await jwt_refresh_response.json();
					// notifications.success('Refresh token');
					localStorage.setItem('access_token', jwt['access_token']);
					localStorage.setItem('refresh_token', jwt['refresh_token']);

					return await send({ method, path, data, options });
				}
			}
			notifications.error('Expired token');
			localStorage.removeItem('access_token');
			localStorage.removeItem('refresh_token');
			throw redirect(307, '/login');
		} else if (options?.asResponse) {
			return res;
		} else if (res.ok || res.status === 422) {
			const text = await res.text();
			return text ? JSON.parse(text) : {};
		}

		throw error(res.status);
	} catch (exception: any) {
		// Properly forward redirection
		if ('status' in exception && 'location' in exception && exception.status === 307) {
			// redirect didn't work client side, use goto
			if (browser) {
				return goto(exception.location);
			}
			throw exception;
		}

		console.error(exception);
		if (options?.asResponse) {
			throw exception;
		}
		throw error(500);
	}
}

export function get(path: string | URL, options?: APIRequestOptions) {
	return send({ method: 'GET', path, options });
}

export function del(path: string | URL, options?: APIRequestOptions) {
	return send({ method: 'DELETE', path, options });
}

export function post(path: string | URL, data: any, options?: APIRequestOptions) {
	return send({ method: 'POST', path, data, options });
}

export function put(path: string | URL, data: any, options?: APIRequestOptions) {
	return send({ method: 'PUT', path, data, options });
}

export function patch(path: string | URL, data: any, options?: APIRequestOptions) {
	return send({ method: 'PATCH', path, data, options });
}