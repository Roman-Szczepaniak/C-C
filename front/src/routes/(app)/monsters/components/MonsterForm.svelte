<script lang="ts">
	import { Card, Heading, Input, Label } from 'flowbite-svelte';
	import { goto, invalidate } from '$app/navigation';
	import { del, formDataToJson, post, put } from '$lib/api';
	import type { Monster } from '$lib/interfaces';
	import * as notification from '$lib/notifications';
	import SubmitFormButtons from '$lib/components/SubmitFormButtons.svelte';

	export let monster: Monster = {} as Monster;

	async function deleteMonster() {
		try {
			const response = await del(`/monsters/${monster.id}`);
			if (response?.detail) {
				notification.error('Impossible to delete');
				return;
			}
			notification.success('Monster deleted');
			await goto(`/monsters`);
		} catch (error) {
			notification.error('An error occurred while deleting the monster');
		}
	}

	let loading = false;
	let savePromise = Promise.resolve();
	$: savePromise, savePromise.then(() => (loading = false));
	async function handleSubmit(event: SubmitEvent) {
		loading = true;
		const formData = new FormData(event.target as HTMLFormElement);

		const data = formDataToJson(formData);
		if (monster?.id) {
			data.id = monster.id;
			const updateMonster = await put(`/monsters/${monster.id}`, data);
			if (updateMonster?.detail) {
				notification.error('Something wrong!');
				return;
			}
			monster = updateMonster;
			await invalidate('app:monster');
		} else {
			delete data.id;
			const createMonster = await post(`/monsters`, data);
			if (createMonster?.detail) {
				notification.error('Something wrong!');
				return;
			}
			notification.success('Monster created successfully');
			await goto(`/monsters/${createMonster?.id}`);
		}
	}
</script>

<Card size="rounded-lg">
	<form on:submit|preventDefault={(event) => (savePromise = handleSubmit(event))}>
		{#if monster.id}
			<Heading tag="h3" class="mb-5">Monster "{monster.name}"</Heading>
		{/if}

		<div class="grid gap-4 mb-6">
			<Label>
				name
				<Input id="name" name="name" type="text" bind:value={monster.name} required />
			</Label>
			<Label>
				alignment
				<Input id="alignment" name="alignment" type="text" bind:value={monster.alignment} required />
			</Label>
			<Label>
				size
				<Input id="size" name="size" type="text" bind:value={monster.size} required />
			</Label>
			<Label>
				type
				<Input id="type" name="type" type="text" bind:value={monster.type} required />
			</Label>
			<Label>
				environment
				<Input id="environment" name="environment" type="text" bind:value={monster.environment} required />
			</Label>
			<Label>
				status
				<Input id="status" name="status" type="text" bind:value={monster.status} required />
			</Label>
			<Label>
				cr
				<Input id="cr" name="cr" type="text" bind:value={monster.cr} required />
			</Label>
		</div>
		<SubmitFormButtons bind:loading canDelete={monster.id && monster.card_id === null} deleteCallback={deleteMonster} />
	</form>
</Card>
