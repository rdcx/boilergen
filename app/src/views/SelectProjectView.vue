<script lang="ts" setup>
import api from '@/api';
import { onMounted } from 'vue';
import type { Project } from '@/types';
import { ref } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()
const projects = ref<Project[]>([])
const name = ref('')

const createProject = () => {
    api.projects.create(name.value)
        .then((res) => {
            projects.value.push(res.data)
        })
        .catch((err) => {
            console.log(err)
        })
}

const openProject = (id: number) => {
    router.push({ name: 'project', params: { id } })
}

const deleteProject = (id: number) => {
    if (!confirm('Are you sure you want to delete this project?')) return

    api.projects.delete(id)
        .then((res) => {
            const index = projects.value.findIndex((p) => p.id == id)
            projects.value.splice(index, 1)
        })
        .catch((err) => {
            console.log(err)
        })
}

onMounted(() => {
    api.projects.all()
        .then((res) => {
            projects.value = res.data
        })
        .catch((err) => {
            console.log(err)
        })
})

</script>
<template>
    <div class="max-w-sm mx-auto mt-8">
        <div>
            <h1 class="text-2xl font-bold">Projects</h1>
            <p class="text-gray-400">Your projects</p>

            <div class="mt-4 flex">
                <input class="input input-bordered focus:outline-none rounded-r-none w-full" v-model="name"
                    placeholder="Name" />
                <div class="grow"></div>
                <button class="btn btn-primary rounded-l-none" @click="createProject">Create project</button>
            </div>
        </div>
        <div v-for="project in projects" class="card bordered my-4 shadow-lg cursor-pointer">
            <div class="card-body">
                <h2 class="card-title">{{ project.name }}</h2>

                <div class="flex">
                    <div class="w-1/2 mr-4">
                        <button class="btn btn-primary mt-4 w-full" @click="openProject(project.id)">Open</button>
                    </div>
                    <div class="w-1/2 ml-4">
                        <button class="btn btn-error mt-4 w-full" @click="deleteProject(project.id)">Delete</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>