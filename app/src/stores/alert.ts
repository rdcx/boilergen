import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { Alert, AlertType } from '../types'

export const useAlertStore = defineStore('alert', () => {
    const alerts = ref<Array<Alert>>([])

    const setAlert = (message: string, type: AlertType) => {
        const alert: Alert = {
            message,
            type,
        }
        alerts.value.push(alert)
        setTimeout(() => {
            alerts.value = alerts.value.filter((a) => a.message !== alert.message)
        }, 5000)
    }

    const dismiss = (alert: Alert) => {
        console.log(alert)
        alerts.value = alerts.value.filter((a) => a.message !== alert.message)
    }

    return { alerts, setAlert, dismiss }
})