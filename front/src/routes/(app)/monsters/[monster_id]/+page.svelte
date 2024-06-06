<script lang="ts">
	import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
	import MonsterForm from '../components/MonsterForm.svelte';
	import CardForm from '../components/CardForm.svelte';
	import Box from '$lib/components/Box.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="p-4">
	<Breadcrumb class="mb-5">
		<BreadcrumbItem href="/" home>Dashboard</BreadcrumbItem>
		<BreadcrumbItem href="/monsters/">Monsters</BreadcrumbItem>
	</Breadcrumb>

	<div class="flex flex-row gap-4">
		<div class="grid grid-cols-1 gap-4 flex-grow">
			<Box>
				<MonsterForm bind:monster={data.monster} />
			</Box>
			{#if data.monster.card}
				<Box>
					<CardForm bind:card={data.monster.card} monster_id={data.monster.id}/>
				</Box>
			{:else}
				<Box>
					<p>No card associated with this monster. Create one!</p>
					<CardForm monster_id={data.monster.id}/>
				</Box>
			{/if}
		</div>
	</div>
</div>