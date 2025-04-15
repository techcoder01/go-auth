<template>
  <div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Sign in to your account
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Or
        <router-link to="/register" class="font-medium text-indigo-600 hover:text-indigo-500">
          create a new account
        </router-link>
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <!-- Display API error message if exists -->
        <div v-if="apiError" class="mb-4 bg-red-50 border-l-4 border-red-400 p-4">
          <div class="flex">
            <div class="flex-shrink-0">
              <svg class="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-red-700">{{ apiError }}</p>
            </div>
          </div>
        </div>
        
        <form class="space-y-6" @submit.prevent="handleSubmit">
          <InputField
            id="email"
            label="Email address"
            type="email"
            v-model="email"
            :error="errors.email"
          />

          <InputField
            id="password"
            label="Password"
            type="password"
            v-model="password"
            :error="errors.password"
          />

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input id="remember-me" name="remember-me" type="checkbox" v-model="rememberMe" class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
              <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                Remember me
              </label>
            </div>

            <div class="text-sm">
              <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                Forgot your password?
              </a>
            </div>
          </div>

          <div>
            <FormButton type="submit" :loading="loading">
              Sign in
            </FormButton>
          </div>
        </form>

        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white text-gray-500">
                Or continue with
              </span>
            </div>
          </div>

          <div class="mt-6 grid grid-cols-3 gap-3">
            <!-- Social Login Icons here -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import InputField from '../components/InputField.vue'
import FormButton from '../components/FormButton.vue'
import axios from 'axios'

export default {
  components: {
    InputField,
    FormButton
  },
  setup() {
    const router = useRouter()

    const email = ref('')
    const password = ref('')
    const rememberMe = ref(false)
    const loading = ref(false)
    const apiError = ref('')
    const errors = reactive({
      email: '',
      password: ''
    })

    const validateForm = () => {
      let isValid = true

      // Reset errors
      errors.email = ''
      errors.password = ''
      apiError.value = ''

      // Validate email
      if (!email.value) {
        errors.email = 'Email is required'
        isValid = false
      } else if (!/\S+@\S+\.\S+/.test(email.value)) {
        errors.email = 'Email is invalid'
        isValid = false
      }

      // Validate password
      if (!password.value) {
        errors.password = 'Password is required'
        isValid = false
      } else if (password.value.length < 6) {
        errors.password = 'Password must be at least 6 characters'
        isValid = false
      }

      return isValid
    }

    const handleSubmit = async () => {
      if (!validateForm()) return

      loading.value = true
      apiError.value = ''

      try {
        // Make the API request directly here
        const response = await axios.post('http://localhost:8080/api/login', {
          email: email.value,
          password: password.value,
        })
        
        // Handle successful login
        if (response.data && response.data.token) {
  localStorage.setItem('authToken', response.data.token)
  router.push('/dashboard')
}


      } catch (error) {
        console.error('Login error:', error)

        // Handle specific error responses from API
        if (error.response) {
          switch (error.response.status) {
            case 401:
              apiError.value = 'Invalid email or password'
              break
            case 400:
              apiError.value = error.response.data.error || 'Invalid input'
              break
            case 500:
              apiError.value = 'Server error. Please try again later.'
              break
            default:
              apiError.value = 'An error occurred. Please try again.'
          }
        } else if (error.request) {
          // Request made but no response received
          apiError.value = 'Cannot reach the server. Please check your connection.'
        } else {
          // Other errors
          apiError.value = 'Login failed. Please try again.'
        }
      } finally {
        loading.value = false
      }
    }

    return {
      email,
      password,
      rememberMe,
      loading,
      errors,
      apiError,
      handleSubmit
    }
  }
}
</script>
