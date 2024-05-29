<script lang="ts">
    import type { PageData } from './$types';
    import {
        Breadcrumb,
        BreadcrumbItem,
        Card,
        Heading,
        Input,
        Label,
    } from 'flowbite-svelte';

    export let data: PageData;
</script>

<div class="p-4">
    <Breadcrumb class="mb-5">
        <BreadcrumbItem href="/" home>Dashboard</BreadcrumbItem>
        <BreadcrumbItem>Profile</BreadcrumbItem>
    </Breadcrumb>

    <Heading tag="h3" class="mb-5">Profile</Heading>

    <Card>
        <h1 class="text-2xl font-bold">Vos données</h1>
        <div>
            <Label class="space-y-2">
                <span>Email</span>
                <Input type="email" name="email" placeholder="email" bind:value={data.profile.email}/>
            </Label>
        </div>
        <div>
            <Label class="space-y-2">
                <span>Firstname</span>
                <Input type="text" name="firstname" placeholder="firstname" bind:value={data.profile.firstname}/>
            </Label>
        </div>
        <div>
            <Label class="space-y-2">
                <span>Lastname</span>
                <Input type="text" name="lastname" placeholder="lastname" bind:value={data.profile.lastname}/>
            </Label>
        </div>
    </Card>
</div>

