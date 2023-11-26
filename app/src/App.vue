<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useUserStore } from './stores/user';
import { useRoute } from 'vue-router'
import Alerts from './components/Alerts.vue';
const route = useRoute()
const userStore = useUserStore()
const logout = () => {
    userStore.logout()
    window.location.href = '/'
}

</script>

<template>
    <div class="flex min-h-full flex-col">
        <Alerts />
        <header class="shrink-0 shadow-md px-8">
            <div class="mx-auto flex h-16 max-w-7xl items-center justify-between lg:px-0">
                <RouterLink to="/" class="font-bold text-2xl">Boilergen <span
                        class="text-xs font-light text-accent">v0.0.1</span></RouterLink>

                <div class="flex items-center gap-x-8">
                    <div class="flex-none">
                        <div v-if="userStore.user" class="dropdown dropdown-end">
                            <label tabindex="0" class="btn btn-ghost btn-circle avatar">
                                <div class="w-10 rounded-full">
                                    <img src="/avatar.png" />
                                </div>
                            </label>
                            <ul tabindex="0"
                                class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
                                <li @click="logout"><a>Logout</a></li>
                            </ul>
                        </div>

                        <div v-else>
                            <RouterLink to="/login" class="btn btn-sm"
                                :class="route.name == 'login' ? 'btn-primary' : 'btn-ghost'">Login</RouterLink>
                            <RouterLink to="/register" class="ml-2 btn btn-sm"
                                :class="route.name == 'register' ? 'btn-primary' : 'btn-ghost'">Register</RouterLink>
                        </div>
                    </div>
                </div>
            </div>
        </header>

        <div class="px-8 mt-8">
            <RouterView />
        </div>
    </div>
</template>