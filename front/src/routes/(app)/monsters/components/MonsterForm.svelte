<script lang="ts">
	import { Heading, Input, Label } from 'flowbite-svelte';
	import { goto, invalidate } from '$app/navigation';
	import { del, formDataToJson, post, put } from '$lib/api';
	import type { Monster } from '$lib/interfaces';
	import * as notification from '$lib/notifications';
    import SubmitFormButtons from '$lib/components/SubmitFormButtons.svelte';

	export let monster: Monster = {} as Monster;

	async function deleteMonster() {
		try {
			const response = await del(`/monsters/${monster.id}/`);
			if (response?.detail) {
				notification.error('Impossible to delete');
				return;
			}
			notification.success('Monster deleted');
			await goto(`/monsters`);
		} catch (error) {
			notification.error('An error occurred while deleting the attribute');
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
			const updateAttribute = await put(`/monsters/${monster.id}`, data);
			if (updateAttribute?.detail) {
				notification.error('Something wrong!');
				return;
			}
			monster = updateAttribute;
			await invalidate('app:attribute');
		} else {
			delete data.id;
			const createMonster = await post(`/monsters`, data);
			if (createMonster?.detail) {
				notification.error('Something wrong!');
				return;
			}
			notification.success('Monster created successfully');
			await goto(`/monsters`);
		}
	}
</script>

<form on:submit|preventDefault={(event) => (savePromise = handleSubmit(event))}>
	{#if monster.id}
		<Heading tag="h3" class="mb-5">Monster "{monster.name}"</Heading>
	{/if}

	<div class="grid gap-4 mb-6">
        <Label>
            name
            <Input id="name" name="name" type="text" value={monster.name} required />
        </Label>
        <Label>
            alignment
            <Input id="alignment" name="alignment" type="text" value={monster.alignment} required />
        </Label>
        <Label>
            size
            <Input id="size" name="size" type="text" value={monster.size} required />
        </Label>
        <Label>
            type
            <Input id="type" name="type" type="text" value={monster.type} required />
        </Label>
        <Label>
            environment
            <Input id="environment" name="environment" type="text" value={monster.environment} required />
        </Label>
        <Label>
            status
            <Input id="status" name="status" type="text" value={monster.status} required />
        </Label>
        <Label>
            cr
            <Input id="cr" name="cr" type="text" value={monster.cr} required />
        </Label>
	</div>
	<SubmitFormButtons
		bind:loading
		deleteCallback={deleteMonster}
	/>
</form>