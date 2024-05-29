<script lang="ts">
	import {
		Breadcrumb,
		BreadcrumbItem,
		Button,
		Select,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import ListFetcher from '$lib/components/ListFetcher.svelte';
	import PencilSquare from 'svelte-heros-v2/PencilSquare.svelte';
	import { get } from '$lib/api';
	import { goto } from '$app/navigation';
	import * as notification from '$lib/notifications';

	async function fetchMonsters(searchParams: URLSearchParams, requestCount: boolean = false) {
		const monstersResponse = await get(`/monsters?${searchParams.toString()}`, {
			asResponse: true,
			headers: { 'x-request-count': requestCount.toString() }
		});
		if (monstersResponse.status !== 200) {
			notification.error('Something wrong!');
			throw new Error();
		}
		return monstersResponse;
	}
</script>

<div class="p-4">
    <Breadcrumb class="mb-5">
        <BreadcrumbItem href="/" home>Dashboard</BreadcrumbItem>
        <BreadcrumbItem>Monsters</BreadcrumbItem>
    </Breadcrumb>
	<Button size="sm" on:click={() => goto(`/monsters/create`)}>
		Create
	</Button>

	<ListFetcher callback={fetchMonsters} paginationTop>
		<svelte:fragment slot="filters">
			<Select
					size="sm"
					value=""
					items={[
						{ name: 'Tiny', value: 'Tiny' },
						{ name: 'Small', value: 'Small' },
						{ name: 'Medium', value: 'Medium' },
						{ name: 'Large', value: 'Large' },
						{ name: 'Huge', value: 'Huge' }
					]}
					name="size"
					placeholder="Size"
				/>
			<Select
					size="sm"
					value=""
					items={[
						{ name: 'Aberration', value: 'Aberration' },
						{ name: 'Beast', value: 'Beast' },
						{ name: 'Celestial', value: 'Celestial' },
						{ name: 'Dragon', value: 'Dragon' },
						{ name: 'Undead', value: 'Undead' }
					]}
					name="type"
					placeholder="Type"
				/>
			<Select
					size="sm"
					value=""
					items={[
						{ name: 'Forest', value: 'Forest' },
						{ name: 'Mountain', value: 'Mountain' },
						{ name: 'Aquatic', value: 'Aquatic' },
						{ name: 'Swamp', value: 'Swamp' },
						{ name: 'Desert', value: 'Desert' }
					]}
					name="environment"
					placeholder="Environment"
				/>
			<Select
					size="sm"
					value=""
					items={[
						{ name: 'Lawful', value: 'Lawful' },
						{ name: 'Neutral', value: 'Neutral' },
						{ name: 'Chaotic', value: 'Chaotic' },
						{ name: 'Unaligned', value: 'Unaligned' }
					]}
					name="alignment"
					placeholder="Alignment"
				/>
			<Select
					size="sm"
					value=""
					items={[
						{ name: 'Ordinary', value: 'Ordinary' },
						{ name: 'Legendary', value: 'Legendary' }
					]}
					name="status"
					placeholder="Status"
				/>
		</svelte:fragment>

		<Table id="monsters" slot="data" let:items hoverable>
			<TableHead>
				<TableHeadCell>Name</TableHeadCell>
				<TableHeadCell>Alignment</TableHeadCell>
				<TableHeadCell>Size</TableHeadCell>
				<TableHeadCell>Type</TableHeadCell>
				<TableHeadCell>Environment</TableHeadCell>
				<TableHeadCell>Status</TableHeadCell>
				<TableHeadCell>CR</TableHeadCell>
				<TableHeadCell />
			</TableHead>

			<TableBody>
				{#each items as monster (monster.id)}
					<TableBodyRow>
						<TableBodyCell>{monster.name}</TableBodyCell>
						<TableBodyCell>{monster.alignment}</TableBodyCell>
						<TableBodyCell>{monster.size}</TableBodyCell>
						<TableBodyCell>{monster.type}</TableBodyCell>
						<TableBodyCell>{monster.environment}</TableBodyCell>
						<TableBodyCell>{monster.status}</TableBodyCell>
						<TableBodyCell>{monster.cr}</TableBodyCell>
						<TableBodyCell>
							<div class="flex items-center justify-end gap-x-3">
								<Button size="sm" href={`/monsters/${monster.id}`}>
									<PencilSquare size="15" />
								</Button>
							</div>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</ListFetcher>
</div>
		