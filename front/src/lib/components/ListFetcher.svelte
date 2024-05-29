<script lang="ts">
	import { createEventDispatcher, onMount, onDestroy } from 'svelte';
	import type { Navigation } from '@sveltejs/kit';
	import { goto } from '$app/navigation';
	import { page, navigating } from '$app/stores';
	import {
		Button,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	import * as notification from '$lib/notifications';
	import ListFetcherPagination from './ListFetcherPagination.svelte';

	const dispatch = createEventDispatcher();

	export let callback: (searchParams: URLSearchParams, requestCount?: boolean) => Promise<Response>;
		
	export let formClass = 'my-4';
	export let limit: number = 20;

	export let paginationTop = false;

	let items: Array<any> = [];

	let hasMore: boolean;
	let previousOffset: number;
	let nextOffset: number;

	let formRef: HTMLFormElement;

	let loading = false;

	function getFilterId(search = $page.url.searchParams): string {
		const searchParams = new URLSearchParams(search);
		searchParams.delete('limit');
		searchParams.delete('offset');
		return searchParams.toString();
	}

	async function updateData(searchParams = $page.url.searchParams) {
		try {
			loading = true;

			const response = await callback(searchParams);
			hasMore = response.headers.get('x-has-more') === 'True';
			previousOffset = parseInt(response.headers.get('x-previous-offset') || '0');
			nextOffset = parseInt(response.headers.get('x-next-offset') || '0');
			items = await response.json();

			if (response.headers.has('x-total-count')) {
				dispatch('count', Number(response.headers.get('x-total-count')));
			}
		} catch (e) {
			notification.error('Failed to fetch list');
		} finally {
			loading = false;
		}
	}

	async function hydrateFilter(searchParams = $page.url.searchParams) {
		if (!searchParams.has('limit') || !searchParams.has('offset')) {
			searchParams.set('limit', limit.toString());
			searchParams.set('offset', '0');
			await goto(`?${searchParams.toString()}`);
		}

		for (const el of formRef.querySelectorAll('[name]') as NodeListOf<HTMLInputElement>) {
			if (searchParams.has(el.name)) {
				if (el.type === 'checkbox') {
					el.checked = true;
				} else {
					el.value = searchParams.get(el.name) as string;
				}
			}
		}
	}

	const navigatingUnsubscribe = navigating.subscribe((nav: Navigation | null) => {
		if (!nav) return;
		if (nav.to?.url.pathname === nav.from?.url.pathname) {
			updateData(nav.to?.url.searchParams);
			hydrateFilter(nav.to?.url.searchParams);
		}
	});

	onDestroy(() => {
		navigatingUnsubscribe();
	});

	async function previous() {
		$page.url.searchParams.set('offset', previousOffset.toString());
		await goto(`?${$page.url.searchParams.toString()}`);
	}

	async function next() {
		$page.url.searchParams.set('offset', nextOffset.toString());
		await goto(`?${$page.url.searchParams.toString()}`);
	}

	async function applyFilters(event: any) {
		const searchParams = new URLSearchParams({
			limit: $page.url.searchParams.get('limit') || limit.toString(),
			offset: '0' // Reset offset
		});

		// Get form data and merge them in searchParams
		const formData = new FormData(event.target);
		for (const pair of formData.entries()) {
			if (pair[1] != '') {
				searchParams.set(pair[0], pair[1].toString());
			}
		}

		// If filter changed, reset the count
		if (getFilterId() !== getFilterId(searchParams)) {
			dispatch('count', null);
		}

		await goto(`?${searchParams.toString()}`);
	}

	export async function refresh() {
		await updateData();
	}

	export async function count(): Promise<number> {
		const response = await callback($page.url.searchParams, true);
		return parseInt(response.headers.get('x-total-count') || '0');
	}

	onMount(async () => {
		hydrateFilter();
		await updateData();
	});
</script>

<form bind:this={formRef} class={formClass} on:submit|preventDefault={applyFilters}>
	<div class="flex justify-between">
		<div class="flex gap-2">
			<slot name="filters" />
		</div>
		{#if $$slots.filters}
			<Button slot="right" size="sm" type="submit" disabled={loading}>Search</Button>
		{/if}
	</div>
</form>

{#if paginationTop}
	<ListFetcherPagination
		bind:hasMore
		showPrevious={nextOffset?.toString() !== $page.url.searchParams.get('limit')}
		on:previous={previous}
		on:next={next}
	/>
{/if}

{#if loading}
	<slot name="loading">
		<Table>
			<TableHead>
				<TableHeadCell>
					<div
						class="animate-pulse h-2.5 bg-gray-200 rounded-full dark:bg-gray-800"
						style="max-width: {Math.floor(Math.random() * 300 + 200)}px"
					/>
				</TableHeadCell>
			</TableHead>
			<TableBody>
				{#each { length: limit } as _}
					<TableBodyRow>
						<TableBodyCell>
							<div
								class="animate-pulse h-2 bg-gray-200 rounded-full dark:bg-gray-700 my-2"
								style="max-width: {Math.floor(Math.random() * 300 + 200)}px"
							/>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</slot>
{:else}
	<slot name="data" {items} />
{/if}

<ListFetcherPagination
	bind:hasMore
	showPrevious={Boolean(nextOffset) &&
		nextOffset.toString() !== $page.url.searchParams.get('limit')}
	disabled={loading}
	on:previous={previous}
	on:next={next}
/>