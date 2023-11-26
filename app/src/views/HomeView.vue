<script lang="ts" setup>
import { RouterLink, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const router = useRouter()

const goToX = () => {
    // open x.com in a new tab
    window.open('https://x.com/LaKitchenSink', '_blank')
}

const goToGithubRepo = () => {
    // open github repo in a new tab
    window.open('https://github.com/rdcx/boilergen', '_blank')
}
</script>
<template>
    <div class="max-w-7xl mx-auto mt-12">

        <div class="md:grid grid-cols-3">
            <div>
                <h1 class="text-2xl font-bold">Fast track your new project.</h1>
                <p class="text-gray-400">Configure and install in 2 minutes.</p>

                <button class="btn btn-primary mt-4" @click="router.push({ name: 'projects' })">Get started</button>
            </div>
            <div class="mt-8">
                <div>Need help? Or have feedback.</div>

                <button @click="goToX()" class="btn btn-ghost rounded-xl p-1 w-12"><img src="/x.png" class="w-12"></button>
            </div>
            <div class="mt-8">
                <div>View the source code?</div>

                <button @click="goToGithubRepo()" class="btn btn-ghost rounded-xl p-1 w-12"><img src="/github.png"
                        class="w-12"></button>
            </div>
        </div>

    </div>
</template>