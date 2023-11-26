import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { User } from '@/types'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')!) : null)
  const token = ref<string>(localStorage.getItem('token') || '')
  const setUser = (u: User) => {
    user.value = u
    localStorage.setItem('user', JSON.stringify(u))
  }

  const setToken = (t: string) => {
    token.value = t
    localStorage.setItem('token', t)
  }

  const logout = () => {
    user.value = null
    token.value = ''
    localStorage.removeItem('user')
    localStorage.removeItem('token')
  }

  return { user, setUser, token, setToken, logout }
})
