import { writable, derived, type Writable } from 'svelte/store';

export type NotificationType = 'success' | 'info' | 'warning' | 'error';

export interface Notification {
	id?: number;
	type?: NotificationType;
	timeout?: number;
	message: string;
}

// Store notification list
const _notifications = writable<Notification[]>([]);

// Apply timeout on each notification
export const notifications = derived<Writable<Notification[]>, Notification[]>(
	_notifications,
	($_notifications, set) => {
		set($_notifications);
		if ($_notifications.length > 0) {
			const timer = setTimeout(() => {
				_notifications.update((state) => {
					state.shift();
					return state;
				});
			}, $_notifications[0].timeout);
			return () => {
				clearTimeout(timer);
			};
		}
	}
);

// Add notification to the store
export function add(notification: Notification): void {
	// Push the notification to the top of the list of notifications
	_notifications.update((all): Notification[] => [
		...all,
		{
			id: new Date().getTime(),
			type: 'success',
			timeout: 3000,
			...notification
		}
	]);
}

export function success(message: string) {
	add({ message, type: 'success' });
}

export function info(message: string) {
	add({ message, type: 'info' });
}

export function warning(message: string) {
	add({ message, type: 'warning' });
}

export function error(message: string) {
	add({ message, type: 'error' });
}