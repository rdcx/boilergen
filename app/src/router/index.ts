import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
    },
    {
      path: '/project',
      name: 'projects',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/SelectProjectView.vue'),
    },
    {
      name: 'login',
      path: '/login',
      component: () => import('../views/LoginView.vue'),
    },
    {
      name: 'register',
      path: '/register',
      component: () => import('../views/RegisterView.vue'),
    },
    {
      path: '/projects/:id',
      name: 'project',
      meta: {
        requiresAuth: true
      },
      component: () => import('../views/ProjectView.vue'),
    }
  ]
})

router.beforeEach((to, from, next) => {
  const loggedIn = localStorage.getItem('user')

  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    next('/login')
  }
  next()
})

export default router
