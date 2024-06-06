<script lang="ts">
	import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
	import ActionForm from '../components/ActionForm.svelte';
	import CardForm from '../components/CardForm.svelte';
	import MonsterForm from '../components/MonsterForm.svelte';
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
			<MonsterForm bind:monster={data.monster} />
			{#if data.monster.card}
				<div class="flex justify-center">
					<CardForm bind:card={data.monster.card} monster_id={data.monster.id}/>
				</div>
				<div class="flex justify-center wrap">
					<!-- <ActionForm bind:action={$newAction} card_id={data.monster.card.id} isNew={true} /> à revoir-->
					{#each data.monster.card.actions as action (action.id)}
						<ActionForm bind:action={action} />
					{/each}
				</div>
			{:else}
				<div class="flex justify-center">
					<CardForm monster_id={data.monster.id}/>
				</div>
			{/if}
		</div>
	</div>
</div>
