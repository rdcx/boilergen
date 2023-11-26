<script lang="ts" setup>
import api from '@/api';
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAlertStore } from '@/stores/alert';
import { useUserStore } from '@/stores/user';
const { t } = useI18n()

const router = useRouter()
const username = ref('')
const email = ref('')
const password = ref('')

const error = ref('')

const register = () => {
    error.value = ''
    api.auth.register(username.value, email.value, password.value)
        .then((res) => {
            const alertStore = useAlertStore();
            alertStore.setAlert('Successfully registered', 'success');

            api.auth.login(username.value, password.value)
                .then((res) => {
                    if (res.status == 200 && res.data.token) {
                        const userStore = useUserStore();
                        const alertStore = useAlertStore();
                        alertStore.setAlert('Successfully logged in', 'info');
                        userStore.setToken(res.data.token);
                        api.auth.user()
                            .then((res) => {
                                userStore.setUser(res.data);
                                router.push({ name: 'projects' })
                            })
                            .catch((err) => {
                                alertStore.setAlert('Failed to get user', 'error');
                            })
                    }
                    return res;
                })
                .catch(err => {
                    if (err.response.status == 400) {
                        const alertStore = useAlertStore();
                        alertStore.setAlert(err.response.data.message, 'error');
                        error.value = t(err.response.data.message)
                        return err;
                    }

                    return err;
                })

        })
        .catch((err) => {
            if (err.response.status == 400) {
                error.value = t(err.response.data.message)
                return
            }
            error.value = t("register_error")
        })
}
</script>
<template>
    <div class="max-w-sm mx-auto">
        <div>
            <h1 class="text-2xl font-bold">Register</h1>
            <p class="text-gray-400">It's free, but you can pay in the form of a follow on X. <a
                    href="https://x.com/LaKitchenSink" target="_blank">@LaKitchenSink</a></p>
        </div>
        <div v-if="error" class="mt-4 text-red-500">{{ error }}</div>
        <div class="mt-8">
            <input class="input input-bordered w-full" v-model="username" placeholder="Username" />
        </div>
        <div class="mt-4">
            <input class="input input-bordered w-full" v-model="email" placeholder="Email" />
        </div>
        <div class="mt-4">
            <input class="input input-bordered w-full" type="password" @keyup.enter="register" v-model="password"
                placeholder="Password" />
        </div>

        <div class="mt-4 flex">
            <button @click="register" class="btn btn-primary">Register</button>
            <div class="grow"></div>
            <RouterLink :to="{ name: 'login' }">Already have an account?</RouterLink>
        </div>
    </div>
</template>