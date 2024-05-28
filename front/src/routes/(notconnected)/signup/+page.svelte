<script lang="ts">
	import { signup } from '$lib/auth';
	import { Button, Card, FloatingLabelInput, Spinner } from 'flowbite-svelte';

	let firstname: string, lastname: string, email: string, hashpassword: string;
	let loading = false;
	let color: any = 'base';

	async function submit() {
		loading = true;
		try {
			await signup(firstname, lastname, email, hashpassword);
		} catch (error) {
			color = 'red';
		}
		loading = false;
	}
</script>

<div class="grid h-screen place-items-center">
	<Card>
		<form on:submit|preventDefault={submit} method="POST" class="flex flex-col space-y-6">
			<h1 class="text-2xl font-bold">Welcome to Craft&Combat!</h1>
            <div>
				<FloatingLabelInput
					{color}
					id="firstname"
					name="firstname"
					type="text"
					label="Firstname"
					required
					bind:value={firstname}
				/>
			</div>
            <div>
				<FloatingLabelInput
					{color}
					id="lastname"
					name="lastname"
					type="text"
					label="Lastname"
					required
					bind:value={lastname}
				/>
			</div>
			<div>
				<FloatingLabelInput
					{color}
					id="email"
					name="email"
					type="text"
					label="Email"
					required
					bind:value={email}
				/>
			</div>
			<div>
				<FloatingLabelInput
					{color}
					id="hashpassword"
					name="hashpassword"
					type="password"
					label="Hashpassword"
					required
					bind:value={hashpassword}
				/>
			</div>

			<Button type="submit" class="w-full" disabled={loading}>
				{#if loading}<Spinner class="mr-3" size="4" color="white" />{/if}
				Sign Up
			</Button>
		</form>
	</Card>
</div>