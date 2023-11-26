<script lang="ts" setup>
import { useAlertStore } from '@/stores/alert';
import type { Alert } from '@/types';
import { computed } from 'vue';
const alertStore = useAlertStore()
const alerts = computed(() => alertStore.alerts);

const dismiss = (a: Alert) => {
    alertStore.dismiss(a)
}
</script>

<template>
    <div v-if="alerts.length > 0" class="fixed top-10 right-10 space-y-4">
        <div v-for="alert in alerts" class="alert" :class="'alert-' + alert.type">
            <svg v-if="alert.type == 'success'" xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6"
                fill="none" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-if="alert.type == 'error'" xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6"
                fill="none" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>

            <svg v-if="alert.type == 'info'" xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6"
                fill="none" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 6v2m0 4h.01m-6.938 5.657a9 9 0 1113.458 0A9 9 0 015.062 15.657z" />
            </svg>

            <svg v-if="alert.type == 'warning'" xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6"
                fill="none" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m0 3v2m-9-5h18" />
            </svg>
            <span>{{ alert.message }}</span>

            <button class="btn btn-sm" @click="dismiss(alert)">Dismiss</button>
        </div>
    </div>
</template>