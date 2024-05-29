<script lang="ts">
	import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
	import MonsterForm from '../components/MonsterForm.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="p-4">
	<Breadcrumb class="mb-5">
		<BreadcrumbItem href="/" home>Dashboard</BreadcrumbItem>
		<BreadcrumbItem href="/monsters/">Monsters</BreadcrumbItem>
	</Breadcrumb>

	<div class="grid gap-6 mb-6">
		<MonsterForm bind:monster={data.monster} />
	</div>
</div>