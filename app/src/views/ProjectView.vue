<script lang="ts" setup>
import api from '@/api';
import { onMounted } from 'vue';
import type { List, Project } from '@/types';
import { ref } from 'vue';
import { useRoute } from 'vue-router'
import { computed } from 'vue';

const route = useRoute()
const project = ref<Project>()

const loading = ref(false)

const createForms = ref<Array<{
    name: '',
    description: '',
    input: HTMLInputElement
}>>([])

const createListName = ref('')

const createList = () => {
    loading.value = true
    api.lists.create(createListName.value, route.params.id as unknown as number)
        .then((res) => {
            loading.value = false
            loadLists()
            createListName.value = ''
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const deleteList = (list: List) => {

    if (!confirm('Are you sure you want to delete this list?')) return
    loading.value = true
    api.lists.delete(list.id)
        .then((res) => {
            loading.value = false
            const index = project.value?.lists.findIndex((l: List) => l.id == list.id)
            project.value?.lists.splice(index as number, 1)
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const addTask = (listId: number) => {
    loading.value = true
    api.tasks.create(createForms.value[listId].name, createForms.value[listId].description, listId)
        .then((res) => {
            loading.value = false
            project.value?.lists.forEach((list: List) => {
                if (list.id == listId) {
                    list.tasks.push(res.data)
                }
            })

            createForms.value[listId].name = ''
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const deleteTask = (task: any) => {

    if (!confirm('Are you sure you want to delete this task?')) return
    loading.value = true
    api.tasks.delete(task.id)
        .then((res) => {
            loading.value = false
            project.value?.lists.forEach((list: List) => {
                if (list.id == task.list_id) {
                    const index = list.tasks.findIndex((t: any) => t.id == task.id)
                    list.tasks.splice(index, 1)
                }
            })
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const moveListLeft = (list: List) => {

    // Find the list to the left
    const leftList = project.value?.lists.find((l: List) => l.position == list.position - 1)
    if (!leftList) return
    loading.value = true
    // Update the list
    api.lists.update(list.id, list.name, leftList?.position)
        .then((res) => {
            loading.value = false
            loadLists()
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const moveListRight = (list: List) => {

    // Find the list to the left
    const rightList = project.value?.lists.find((l: List) => l.position == list.position + 1)

    if (!rightList) return
    loading.value = true
    // Update the list
    api.lists.update(list.id, list.name, rightList?.position)
        .then((res) => {
            loading.value = false
            loadLists()
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const moveLeft = (task: any) => {

    // Find the list
    const list = project.value?.lists.find((list: List) => list.id == task.list_id)

    if (!list) return

    // Find the list to the left
    const leftList = project.value?.lists.find((l: List) => l.position == list?.position - 1)

    if (!leftList) return

    // Find the index of the task
    const index = list?.tasks.findIndex((t: any) => t.id == task.id)

    // Remove the task from the list
    list?.tasks.splice(index as number, 1)
    loading.value = true
    // Update the task
    api.tasks.update(task.id, task.name, task.description, leftList?.id)
        .then((res) => {
            loading.value = false
            project.value?.lists.forEach((list: List) => {
                if (list.id == res.data.list_id) {
                    list.tasks.push(res.data)
                }
            })
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

const moveRight = (task: any) => {

    // Find the list
    const list = project.value?.lists.find((list: List) => list.id == task.list_id)

    if (!list) return

    // Find the list to the left
    const rightList = project.value?.lists.find((l: List) => l.position == list.position + 1)

    if (!rightList) return

    // Find the index of the task
    const index = list?.tasks.findIndex((t: any) => t.id == task.id)

    // Remove the task from the list
    list?.tasks.splice(index as number, 1)
    loading.value = true
    // Update the task
    api.tasks.update(task.id, task.name, task.description, rightList?.id)
        .then((res) => {
            loading.value = false
            project.value?.lists.forEach((list: List) => {
                if (list.id == res.data.list_id) {
                    list.tasks.push(res.data)
                }
            })
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

function timeSince(date: any) {

    var seconds = Math.floor((new Date() as any - date) / 1000);

    var interval = seconds / 31536000;

    if (interval > 1) {
        return Math.floor(interval) + " years";
    }
    interval = seconds / 2592000;
    if (interval > 1) {
        return Math.floor(interval) + " months";
    }
    interval = seconds / 86400;
    if (interval > 1) {
        return Math.floor(interval) + " days";
    }
    interval = seconds / 3600;
    if (interval > 1) {
        return Math.floor(interval) + " hours";
    }
    interval = seconds / 60;
    if (interval > 1) {
        return Math.floor(interval) + " minutes";
    }
    return Math.floor(seconds) + " seconds";
}

const loadLists = () => {
    loading.value = true
    api.projects.get(route.params.id as unknown as number)
        .then((res) => {
            loading.value = false
            project.value = res.data

            // Create forms
            project.value?.lists.forEach((list: List) => {
                createForms.value[list.id] = {
                    name: '',
                    description: '',
                } as any
            })
        })
        .catch((err) => {
            loading.value = false
            console.log(err)
        })
}

onMounted(() => {
    loadLists()
})
</script>
<template>
    <div v-if="loading" class="fixed left-[50%] top-3 loading loading-infinity loading-lg"></div>
    <div class="md:flex md:max-w-[120rem] md:mx-auto">
        <div v-for="list in project?.lists.sort((a, b) => a.position - b.position)"
            class="w-full bordered shadow-lg md:m-8 relative my-8 md:my-0">
            <div class="absolute -top-3 -right-3">
                <button class="mb-2" @click="deleteList(list)"><svg xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                </button>
            </div>
            <div class="flex p-8">
                <div>
                    <h2 class="card-title">{{ list.name }}</h2>
                </div>
                <div class="grow"></div>
                <div>
                    <button @click="moveListLeft(list)" class="text-gray-400 hover:text-gray-500"><svg
                            xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                            <path fill-rule="evenodd"
                                d="M20.25 12a.75.75 0 01-.75.75H6.31l5.47 5.47a.75.75 0 11-1.06 1.06l-6.75-6.75a.75.75 0 010-1.06l6.75-6.75a.75.75 0 111.06 1.06l-5.47 5.47H19.5a.75.75 0 01.75.75z"
                                clip-rule="evenodd" />
                        </svg>
                    </button>

                    <button @click="moveListRight(list)" class="text-gray-400 hover:text-gray-500"><svg
                            xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                            <path fill-rule="evenodd"
                                d="M3.75 12a.75.75 0 01.75-.75h13.19l-5.47-5.47a.75.75 0 011.06-1.06l6.75 6.75a.75.75 0 010 1.06l-6.75 6.75a.75.75 0 11-1.06-1.06l5.47-5.47H4.5a.75.75 0 01-.75-.75z"
                                clip-rule="evenodd" />
                        </svg>
                    </button>
                </div>

            </div>


            <div class="card-body">

                <input class="input input-bordered w-full" v-model="createForms[list.id].name"
                    ref="createForms[list.id].input" placeholder="Task name" />

                <button class="btn btn-primary" @click="addTask(list.id)">Add task</button>
            </div>

            <div class="px-8">
                <div v-for="task in list.tasks" class="relative my-2 p-2 rounded-xl">
                    <div class="text-sm text-secondary font-bold">{{ task.name }}</div>
                    <div class="text-sm text-secondary" v-if="task.description">{{ task.description || 'No description' }}
                    </div>
                    <div class="text-xs">{{ timeSince(new Date(task.created_at)) }} ago</div>

                    <div class="flex mt-2">
                        <button @click="moveLeft(task)" class="text-gray-400 hover:text-gray-500"><svg
                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                                <path fill-rule="evenodd"
                                    d="M20.25 12a.75.75 0 01-.75.75H6.31l5.47 5.47a.75.75 0 11-1.06 1.06l-6.75-6.75a.75.75 0 010-1.06l6.75-6.75a.75.75 0 111.06 1.06l-5.47 5.47H19.5a.75.75 0 01.75.75z"
                                    clip-rule="evenodd" />
                            </svg>
                        </button>

                        <button @click="moveRight(task)" class="text-gray-400 hover:text-gray-500"><svg
                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                                <path fill-rule="evenodd"
                                    d="M3.75 12a.75.75 0 01.75-.75h13.19l-5.47-5.47a.75.75 0 011.06-1.06l6.75 6.75a.75.75 0 010 1.06l-6.75 6.75a.75.75 0 11-1.06-1.06l5.47-5.47H4.5a.75.75 0 01-.75-.75z"
                                    clip-rule="evenodd" />
                            </svg>
                        </button>

                        <div class="grow"></div>
                        <button class="mb-2" @click="deleteTask(task)"><svg xmlns="http://www.w3.org/2000/svg" fill="none"
                                viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round"
                                    d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <div class="w-full bordered shadow-lg md:m-8 my-8 md:my-0">
            <div class="card-body">
                <h2 class="card-title">Add a list</h2>
            </div>

            <div class="card-body">
                <input class="input input-bordered w-full" v-model="createListName" placeholder="List name" />

                <button class="btn btn-primary" @click="createList">Create list</button>
            </div>
        </div>
    </div>
</template>