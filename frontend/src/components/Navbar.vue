<template>
  <nav class="bg-white shadow-md">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0 flex items-center">
            <router-link to="/" class="text-indigo-600 font-bold text-xl">VueAuth</router-link>
          </div>
          <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
            <router-link to="/" class="inline-flex items-center px-1 pt-1 border-b-2" :class="[$route.path === '/' ? 'border-indigo-500 text-gray-900' : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700']">
              Home
            </router-link>
            <router-link v-if="isAuthenticated" to="/dashboard" data-test="dashboard" class="inline-flex items-center px-1 pt-1 border-b-2" :class="[$route.path === '/dashboard' ? 'border-indigo-500 text-gray-900' : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700']">
              Dashboard
            </router-link>
          </div>
        </div>
        <div class="hidden sm:ml-6 sm:flex sm:items-center">
          <div v-if="isAuthenticated" class="flex items-center space-x-4">
            <div class="text-sm text-gray-700">{{ user?.email }}</div>
            <button @click="logout" data-test="logout" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Logout
            </button>
          </div>
          <div v-else class="flex items-center space-x-2">
            <router-link to="/login" data-test="login" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-indigo-600 bg-white hover:bg-indigo-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Login
            </router-link>
            <router-link to="/register" data-test="register" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Register
            </router-link>
          </div>
        </div>

        <!-- Mobile menu button -->
        <div class="flex items-center sm:hidden">
          <button @click="toggleMobileMenu" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500">
            <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile menu -->
    <div v-if="mobileMenuOpen" class="sm:hidden">
      <div class="pt-2 pb-3 space-y-1">
        <router-link to="/"  class="block pl-3 pr-4 py-2 border-l-4" :class="[$route.path === '/' ? 'border-indigo-500 text-indigo-700 bg-indigo-50' : 'border-transparent text-gray-600 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-800']" @click="closeMobileMenu">
          Home
        </router-link>
        <router-link v-if="isAuthenticated" to="/dashboard" data-test="dashboard" class="block pl-3 pr-4 py-2 border-l-4" :class="[$route.path === '/dashboard' ? 'border-indigo-500 text-indigo-700 bg-indigo-50' : 'border-transparent text-gray-600 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-800']" @click="closeMobileMenu">
          Dashboard
        </router-link>
      </div>
      <div class="pt-4 pb-3 border-t border-gray-200">
        <div v-if="isAuthenticated" class="flex items-center px-4">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 rounded-full bg-indigo-100 flex items-center justify-center">
              <span class="text-indigo-800 font-medium">{{ user?.email?.charAt(0).toUpperCase() }}</span>
            </div>
          </div>
          <div class="ml-3">
            <div class="text-base font-medium text-gray-800">{{ user?.id ? `User ID: ${user.id}` : '' }}</div>
            <div class="text-sm font-medium text-gray-500">{{ user?.email }}</div>
          </div>
        </div>
        <div class="mt-3 space-y-1">
          <div v-if="isAuthenticated">
            <button @click="logout" class="block w-full text-left px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100">
              Logout
            </button>
          </div>
          <div v-else class="space-y-1">
            <router-link to="/login" data-test="login" class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100" @click="closeMobileMenu">
              Login
            </router-link>
            <router-link to="/register" data-test="register" class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100" @click="closeMobileMenu">
              Register
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import axios from 'axios'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  setup() {
    const router = useRouter()
    const mobileMenuOpen = ref(false)

    // Computed properties for authentication and user
    const isAuthenticated = computed(() => localStorage.getItem('authToken') !== null)
    const user = computed(() => JSON.parse(localStorage.getItem('user')))

    // Mobile menu toggle
    const toggleMobileMenu = () => {
      mobileMenuOpen.value = !mobileMenuOpen.value
    }

    // Close mobile menu after selection
    const closeMobileMenu = () => {
      mobileMenuOpen.value = false
    }

    // Logout function using Axios
    const logout = async () => {
      try {
        await axios.post('http://localhost:8080/api/logout')  // Assuming your API has a logout route
        localStorage.removeItem('authToken')
        localStorage.removeItem('user')
        router.push('/login')
      } catch (error) {
        console.error('Logout failed:', error)
        localStorage.removeItem('authToken')
        localStorage.removeItem('user')
        router.push('/login')
      }
    }

    return {
      mobileMenuOpen,
      toggleMobileMenu,
      closeMobileMenu,
      isAuthenticated,
      user,
      logout
    }
  }
}
</script>
