import { mount, flushPromises } from '@vue/test-utils'
import Dashboard from '@/views/Dashboard.vue'
import { createRouter, createWebHistory } from 'vue-router'
import axios from 'axios'

// Mock axios properly
vi.mock('axios')

const routes = [
  { path: '/dashboard', name: 'Dashboard' },
  { path: '/login', name: 'Login' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

describe('Dashboard', () => {
  let wrapper

  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
  })

  it('renders dashboard title', () => {
    wrapper = mount(Dashboard, {
      global: {
        plugins: [router]
      }
    })
    expect(wrapper.find('h2').text()).toContain('Welcome to your Dashboard')
  })

  it('fetches user data when authenticated', async () => {
    // Set up mock before mounting component
    const mockUserData = { id: 1, name: 'Test User' }
    localStorage.setItem('authToken', 'test-token')
    
    // Important: Set up the mock implementation before mounting
    axios.get.mockResolvedValueOnce({ data: mockUserData })
    
    // Mount component
    wrapper = mount(Dashboard, {
      global: {
        plugins: [router]
      }
    })
    
    // Wait for all promises to resolve
    await flushPromises()
    
    // Verify axios was called correctly
    expect(axios.get).toHaveBeenCalledWith('https://go-auth-ftrw.onrender.com/api/user', {
      headers: {
        'Authorization': 'Bearer test-token'
      }
    })
    
    // Verify data was set correctly
    expect(wrapper.vm.dashboardData).toEqual(mockUserData)
  })

  it('shows error when not authenticated', async () => {
    wrapper = mount(Dashboard, {
      global: {
        plugins: [router]
      }
    })
    
    await flushPromises()
    expect(wrapper.vm.apiError).toBe('You are not authenticated. Please log in.')
  })

  it('handles API error', async () => {
    localStorage.setItem('authToken', 'test-token')
    
    // Mock a server error response
    axios.get.mockRejectedValueOnce({
      response: { 
        status: 500,
        data: {} // Empty data object since we're not using response.data.error in the test
      }
    })
    
    wrapper = mount(Dashboard, {
      global: {
        plugins: [router]
      }
    })
    
    // Wait for promises to resolve
    await flushPromises()
    
    // This should now match the actual error message in the component
    expect(wrapper.vm.apiError).toBe('An error occurred. Please try again.')
  })
})