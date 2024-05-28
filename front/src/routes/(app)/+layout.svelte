<script lang="ts">
    import '../../app.postcss';
    import { onMount } from 'svelte';
    import { 
			Avatar,
			DarkMode,
			Dropdown,
		  	DropdownHeader,
		   	DropdownItem,
		    Hr,
			Tooltip,
			Sidebar,
			SidebarGroup,
			SidebarItem,
			SidebarWrapper
	} from 'flowbite-svelte';
    import { page } from '$app/stores';
    import { logout } from '$lib/auth';
    import { darkMode } from '$lib/stores';

    onMount(() => {
        $darkMode = document.body.parentElement.className === 'dark';
        const observer: MutationObserver = new MutationObserver((mutations: MutationRecord[]) =>
            mutations.forEach(
                (mutation: MutationRecord) => ($darkMode = mutation.target.className === 'dark')
            )
        );
        observer.observe(document.body.parentElement, { attributeFilter: ['class'] });
        return () => observer.disconnect();
    });

    let hidden: boolean = (localStorage.getItem('crafncombat.aside') || 'false') === 'true';
</script>

<svelte:head>
    <title>Craft & Combat!</title>
</svelte:head>

<div class="relative">
    <nav class="fixed z-30 w-full bg-white border-b border-gray-200 dark:bg-gray-800 dark:border-gray-700">
        <div class="mx-auto flex flex-wrap items-center justify-between">
            <div class="w-full p-3 lg:px-5 lg:pl-3">
                <div class="flex items-center justify-between">
                    <div class="flex items-center justify-between">
                        <button on:click={() => {
                            hidden = !hidden;
                            localStorage.setItem('craftncombat.aside', hidden.toString());
                        }}
                            class="mr-3 cursor-pointer rounded p-2 text-gray-600 hover:bg-gray-100 hover:text-gray-900 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white lg:inline">
                            <span class="sr-only">Toggle sidebar</span>
                            <svg stroke="currentColor" fill="currentColor" stroke-width="0" viewBox="0 0 20 20"
                                class="h-6 w-6" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg">
                                <path fill-rule="evenodd"
                                    d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                                    clip-rule="evenodd"></path>
                            </svg>
                        </button>
                    </div>
                    <div class="flex items-center lg:gap-3">
                        <div class="flex items-center md:order-2">
                            <Avatar id="avatar-menu" />
                        </div>
                        <Dropdown placement="bottom" triggeredBy="#avatar-menu">
                            {#if $page.data.user?.lastname}
                                <DropdownHeader>
                                    {$page.data.user.firstname} {$page.data.user.lastname}
                                </DropdownHeader>
                            {/if}
                            <DropdownItem on:click={logout}>Sign out</DropdownItem>
                        </Dropdown>
                    </div>
                </div>
            </div>
        </div>
    </nav>
    <div class="flex items-start pt-16">
        <div class={!hidden ? 'lg:!block-hidden' : 'lg:!block'}>
            {#if hidden}
                <Tooltip class="z-30" placement="right" triggeredBy="[href='/monsters']">Monsters</Tooltip>
                <Tooltip class="z-30" placement="right" triggeredBy="[href='/encounters']">Encounters</Tooltip>
                <Tooltip class="z-30" placement="right" triggeredBy="[href='/user']">Profil</Tooltip>
            {/if}
            <Sidebar
                class="{!hidden ? 'w-64' : '!w-16'} flex fixed top-0 left-0 flex-col flex-shrink-0 pt-16 h-full duration-75 border-r border-gray-200 lg:flex transition-[width] dark:border-gray-700">
                <SidebarWrapper class="overflow-x-hidden h-full">
                    <SidebarGroup>
                        <SidebarItem
                            spanClass="ml-3 {hidden ? '!hidden' : '!block'}"
                            href="/monsters"
                            active={$page.url.pathname === '/monsters'}
                            label="Monsters">
                            <svelte:fragment slot="icon">
                                <!-- Add your monster icon here -->
                            </svelte:fragment>
                        </SidebarItem>
                        <SidebarItem
                            spanClass="ml-3 {hidden ? '!hidden' : '!block'}"
                            href="/encounters"
                            active={$page.url.pathname === '/encounters'}
                            label="Encounters">
                            <svelte:fragment slot="icon">
                                <!-- Add your generator icon here -->
                            </svelte:fragment>
                        </SidebarItem>
                        <SidebarItem
                            spanClass="ml-3 {hidden ? '!hidden' : '!block'}"
                            href="/user"
                            active={$page.url.pathname === '/user'}
                            label="Profil">
                            <svelte:fragment slot="icon">
                                <!-- Add your profile icon here -->
							</svelte:fragment>
                        </SidebarItem>
                        <Hr />
                        <DarkMode />
                    </SidebarGroup>
                </SidebarWrapper>
            </Sidebar>
        </div>
        <div class="overflow-y-auto relative w-full h-full mb-10 {!hidden ? 'ml-64' : 'ml-16'}">
            <slot />
        </div>
    </div>
</div>
