<template>
  <div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Create a new account
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Or
        <router-link to="/login" class="font-medium text-indigo-600 hover:text-indigo-500">
          sign in to your existing account
        </router-link>
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <!-- API Error Alert -->
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

          <InputField
            id="confirm-password"
            label="Confirm Password"
            type="password"
            v-model="confirmPassword"
            :error="errors.confirmPassword"
          />

          <div class="flex items-center">
            <input id="terms" name="terms" type="checkbox" v-model="agreeToTerms" class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
            <label for="terms" class="ml-2 block text-sm text-gray-900">
              I agree to the
              <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">Terms and Conditions</a>
            </label>
          </div>
          <p v-if="errors.terms" class="mt-1 text-sm text-red-600">{{ errors.terms }}</p>

          <div>
            <FormButton type="submit" :loading="loading">
              Register
            </FormButton>
          </div>
        </form>
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
    const confirmPassword = ref('')
    const agreeToTerms = ref(false)
    const loading = ref(false)
    const apiError = ref('')
    const errors = reactive({
      email: '',
      password: '',
      confirmPassword: '',
      terms: ''
    })

    const validateForm = () => {
      let isValid = true

      // Reset errors
      errors.email = ''
      errors.password = ''
      errors.confirmPassword = ''
      errors.terms = ''
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

      // Validate confirm password
      if (!confirmPassword.value) {
        errors.confirmPassword = 'Please confirm your password'
        isValid = false
      } else if (confirmPassword.value !== password.value) {
        errors.confirmPassword = 'Passwords do not match'
        isValid = false
      }

      // Validate terms
      if (!agreeToTerms.value) {
        errors.terms = 'You must agree to the terms and conditions'
        isValid = false
      }

      return isValid
    }

    const handleSubmit = async () => {
  if (!validateForm()) return

  loading.value = true

  try {
    // Make the API call for registration
    const response = await axios.post('http://localhost:8080/api/register', {
      email: email.value,
      password: password.value
    })

    // If successful, process the response
    if (response.status === 201) {
      // Optionally store the user data or token
      localStorage.setItem('userId', response.data.user.id)
      localStorage.setItem('userEmail', response.data.user.email)
      apiError.value = response.data.message

      // Redirect to the dashboard
      router.push('/dashboard')
    }
  } catch (error) {
    console.error('Registration error:', error)

    // Handle API errors from your backend
    if (error.response) {
      switch (error.response.status) {
        case 400:
          apiError.value = error.response.data.error || 'Email may already be in use'
          break
        case 500:
          apiError.value = 'Server error. Please try again later.'
          break
        default:
          apiError.value = 'Registration failed. Please try again.'
      }
    } else if (error.request) {
      // Request made but no response received
      apiError.value = 'Cannot reach the server. Please check your connection.'
    } else {
      // Other errors
      apiError.value = 'Registration failed. Please try again.'
    }
  } finally {
    loading.value = false
  }
}


    return {
      email,
      password,
      confirmPassword,
      agreeToTerms,
      loading,
      errors,
      apiError,
      handleSubmit
    }
  }
}
</script>
