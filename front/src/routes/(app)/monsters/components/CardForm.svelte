<script lang="ts">
	import { Heading, Input, Label, Textarea } from 'flowbite-svelte';
	import { goto } from '$app/navigation';
	import { del, formDataToJson, post, put } from '$lib/api';
	import type { Card } from '$lib/interfaces';
	import * as notification from '$lib/notifications';
	import SubmitFormButtons from '$lib/components/SubmitFormButtons.svelte';

	export let card: Card = {} as Card;
    export let monster_id: number;

	async function deleteCard() {
		try {
			const response = await del(`/cards/${card.id}`);
			if (response?.detail) {
				notification.error('Impossible to delete');
				return;
			}
			notification.success('Card deleted');
			await goto(`/monsters`);
		} catch (error) {
			notification.error('An error occurred while deleting the card');
		}
	}

	let loading = false;
	let savePromise = Promise.resolve();
	$: savePromise, savePromise.then(() => (loading = false));
	async function handleSubmit(event: SubmitEvent) {
		loading = true;
		const formData = new FormData(event.target as HTMLFormElement);
		const data = formDataToJson(formData);
        data.ca = Number(data.ca);
		data.pv = Number(data.pv);
		
		if (card.id) {
			data.id = card.id;
			const updateCard = await put(`/cards/${card.id}`, data);
			if (updateCard?.detail) {
				notification.error('Something wrong!');
				return;
			}
			card = updateCard;
			notification.success('Card updated successfully');
		} else {
			delete data.id;
			const createCard = await post(`/cards`, data);
			if (createCard?.detail) {
				notification.error('Something wrong!');
				return;
			}
			card = createCard;
            const updateMonster = await put(`/monsters/${monster_id}`, { card_id: card.id });
			if (updateMonster?.detail) {
				notification.error('Something wrong when updating monster!');
				return;
			}
			notification.success('Card created and associated with monster successfully');
            await goto(`/monsters`);
		}
	}
</script>

<form on:submit|preventDefault={(event) => (savePromise = handleSubmit(event))}>
	{#if card.id}
		<Heading tag="h3" class="mb-5">Card</Heading>
	{/if}

	<div class="grid gap-4 mb-6">
		<Label>
			Description
			<Textarea id="description" name="description" rows="3" bind:value={card.description} required />
		</Label>
		<Label>
			CA
			<Input id="ca" name="ca" type="number" bind:value={card.ca} required />
		</Label>
		<Label>
			PV
			<Input id="pv" name="pv" type="number" bind:value={card.pv} required />
		</Label>
		<Label>
			Speed
			<Input id="speed" name="speed" type="text" bind:value={card.speed} required />
		</Label>
		<div class="grid grid-cols-6 gap-4">
			<Label>
				Strength
				<Input id="strength" name="strength" type="text" bind:value={card.strength} required />
			</Label>
			<Label>
				Dexterity
				<Input id="dexterity" name="dexterity" type="text" bind:value={card.dexterity} required />
			</Label>
			<Label>
				Constitution
				<Input id="constitution" name="constitution" type="text" bind:value={card.constitution} required />
			</Label>
			<Label>
				Intelligence
				<Input id="intelligence" name="intelligence" type="text" bind:value={card.intelligence} required />
			</Label>
			<Label>
				Wisdom
				<Input id="wisdom" name="wisdom" type="text" bind:value={card.wisdom} required />
			</Label>
			<Label>
				Charisma
				<Input id="charisma" name="charisma" type="text" bind:value={card.charisma} required />
			</Label>
		</div>
		<Label>
			Save Throws
			<Input id="save_throws" name="save_throws" type="text" bind:value={card.save_throws} />
		</Label>
		<Label>
			Damage Resist
			<Input id="damage_resist" name="damage_resist" type="text" bind:value={card.damage_resist} />
		</Label>
		<Label>
			Damage Immune
			<Input id="damage_immune" name="damage_immune" type="text" bind:value={card.damage_immune} />
		</Label>
		<Label>
			Condition Immune
			<Input id="condition_immune" name="condition_immune" type="text" bind:value={card.condition_immune} />
		</Label>
		<Label>
			Senses
			<Input id="senses" name="senses" type="text" bind:value={card.senses} />
		</Label>
		<Label>
			Languages
			<Input id="languages" name="languages" type="text" bind:value={card.languages} />
		</Label>
	</div>
	<SubmitFormButtons bind:loading canDelete={card.id} deleteCallback={deleteCard} />
</form>
