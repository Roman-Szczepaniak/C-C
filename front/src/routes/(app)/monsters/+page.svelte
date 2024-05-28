<script lang="ts">
	import type { PageData } from '../monster/$types';
	import {
		Accordion,
		AccordionItem,
		Alert,
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Input,
		Modal,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import Key from 'svelte-heros-v2/Key.svelte';
	import ArrowDownOnSquare from 'svelte-heros-v2/ArrowDownOnSquare.svelte';
	import PencilSquare from 'svelte-heros-v2/PencilSquare.svelte';
	import Pencil from 'svelte-heros-v2/Pencil.svelte';
	import Trash from 'svelte-heros-v2/Trash.svelte';
	import { goto, invalidate } from '$app/navigation';
	import { get, del, post, put } from '$lib/api';
	import type { Monster } from '$lib/interfaces.js';
	import * as notification from '$lib/notifications';

	 export let data: PageData;

	/** async function deleteMonster(monster) {
		const response = await del(`monster/${monster.id}/`);
		if (response?.detail) {
			notification.error('The attribute group contains attributes');
			return;
		}
		await invalidate('app:monsters');
		await goto(`/monsters`);
	}

	async function saveMonster(monster) {
		try {
			let response;
			if (monster.id === null) {
				response = await post('monster/', { code: monster.code });
			} else {
				response = await put(`monster/${monster.id}/`, { id: monster.id, code: monster.code });
			}

			if (response?.id) {
				invalidate('app:monsters');
			} else if (response) {
				invalidate('app:monsters');
			} else {
				notification.error('Failed to save the group');
			}
		} catch (error) {
			notification.error('An error occurred while saving the group');
		}
	}

	async function fetchMonsters(searchParams: URLSearchParams, requestCount: boolean) {
		const monstersResponse = await get('/monsters', { asResponse: true });
		if (monstersResponse.status !== 200) {
			notification.error('Something wrong!');
			throw new Error();
		}
		return monstersResponse;
	}
	let modalDeleteMonster: Monster; **/
</script>

<div class="p-4">
    <Breadcrumb class="mb-5">
        <BreadcrumbItem href="/" home>Dashboard</BreadcrumbItem>
        <BreadcrumbItem>Monsters</BreadcrumbItem>
    </Breadcrumb>

    <Table slot="data" let:items hoverable>
        <TableHead>
            <TableHeadCell>Name</TableHeadCell>
            <TableHeadCell>Alignment</TableHeadCell>
            <TableHeadCell>Size</TableHeadCell>
            <TableHeadCell>Type</TableHeadCell>
            <TableHeadCell>CR</TableHeadCell>
            <TableHeadCell />
        </TableHead>

        <TableBody>
            {#if Array.isArray(data.monsters)}
                {#each data.monsters as monster}
                    <TableBodyRow>
                        <TableBodyCell>{monster.name}</TableBodyCell>
                        <TableBodyCell>{monster.alignment}</TableBodyCell>
                        <TableBodyCell>{monster.size}</TableBodyCell>
                        <TableBodyCell>{monster.type}</TableBodyCell>
                        <TableBodyCell>{monster.cr}</TableBodyCell>
                        <TableBodyCell>
                            <div class="flex items-center justify-end gap-x-3">
                                <Button size="sm" href={`/monster/${monster.id}`}>
                                    <PencilSquare size="15" />
                                </Button>
                            </div>
                        </TableBodyCell>
                    </TableBodyRow>
                {/each}
            {:else}
                <TableBodyRow>
                    <TableBodyCell colspan="6" class="text-center">
                        No monsters available.
                    </TableBodyCell>
                </TableBodyRow>
            {/if}
        </TableBody>
    </Table>
</div>
		