<template>
  <div class="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Welcome to your Dashboard
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        You're logged in!
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
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

        <!-- Content for the dashboard -->
        <div v-if="dashboardData">
          <p>Data fetched from the server:</p>
          <pre>{{ dashboardData }}</pre>
        </div>
        
        <!-- Loader spinner when data is being fetched -->
        <div v-if="loading" class="flex justify-center">
          <div class="loader">Loading...</div> <!-- You can add your loader here -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
  setup() {
    const loading = ref(false)
    const apiError = ref('')
    const dashboardData = ref(null)

    const fetchDashboardData = async () => {
      const token = localStorage.getItem('authToken')
      if (!token) {
        apiError.value = 'You are not authenticated. Please log in.'
        return
      }

      loading.value = true
      apiError.value = ''
      
      try {
        // Use the token for authentication in the API request
        const response = await axios.get('http://localhost:8080/api/user', {
          headers: {
            'Authorization': `Bearer ${token}`
          }
        })

        // Handle the data
        dashboardData.value = response.data
      } catch (error) {
        console.error('Error fetching dashboard data:', error)

        // Handle errors
        if (error.response) {
          apiError.value = error.response.data.error || 'An error occurred. Please try again.'
        } else if (error.request) {
          apiError.value = 'Cannot reach the server. Please check your connection.'
        } else {
          apiError.value = 'Failed to fetch dashboard data. Please try again.'
        }
      } finally {
        loading.value = false
      }
    }

    // Fetch the dashboard data when the component is mounted
    onMounted(() => {
      fetchDashboardData()
    })

    return {
      loading,
      apiError,
      dashboardData
    }
  }
}
</script>

<style scoped>
/* Add custom loader styles */
.loader {
  border: 4px solid #f3f3f3;
  border-top: 4px solid #3498db;
  border-radius: 50%;
  width: 50px;
  height: 50px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
