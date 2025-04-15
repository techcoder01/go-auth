// src/main.js
import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import './index.css'
import axios from 'axios'

// Set up axios defaults
axios.defaults.baseURL = 'http://localhost:8080'

// Import components
import Home from './views/Home.vue'
import Login from './views/Login.vue'
import Register from './views/Register.vue'

// Create router
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('./views/Dashboard.vue'),
      meta: { requiresAuth: true }
    }
    
  ]
})

// Navigation guard for protected routes
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('authToken')
  
  if (to.matched.some(record => record.meta.requiresAuth) && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

// Create Vue app
const app = createApp(App)
app.use(router)
app.mount('#app')

// Set up token in axios if it exists
const token = localStorage.getItem('authToken')
if (token) {
  axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
}
