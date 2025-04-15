import { mount } from '@vue/test-utils'
import Navbar from '@/components/Navbar.vue'
import { createRouter, createMemoryHistory } from 'vue-router'

// Define mock routes
const routes = [
  { path: '/', name: 'Home', component: { template: '<div>Home</div>' } },
  { path: '/login', name: 'Login', component: { template: '<div>Login</div>' } },
  { path: '/register', name: 'Register', component: { template: '<div>Register</div>' } },
  { path: '/dashboard', name: 'Dashboard', component: { template: '<div>Dashboard</div>' } }
]

// Mock localStorage
const localStorageMock = (() => {
  let store = {}
  return {
    getItem(key) {
      return store[key] || null
    },
    setItem(key, value) {
      store[key] = value.toString()
    },
    removeItem(key) {
      delete store[key]
    },
    clear() {
      store = {}
    }
  }
})()

Object.defineProperty(window, 'localStorage', {
  value: localStorageMock
})

// Helper to mount Navbar
const mountNavbar = async () => {
  const router = createRouter({
    history: createMemoryHistory(),
    routes
  })

  await router.push('/')
  await router.isReady()

  return mount(Navbar, {
    global: {
      plugins: [router]
    }
  })
}

describe('Navbar.vue', () => {
  beforeEach(() => {
    localStorage.clear()
  })

  it('renders without crashing', async () => {
    const wrapper = await mountNavbar()
    expect(wrapper.exists()).toBe(true)
  })

  it('shows Login/Register when not authenticated', async () => {
    const wrapper = await mountNavbar()
    expect(wrapper.find('[data-test="login"]').exists()).toBe(true)
    expect(wrapper.find('[data-test="register"]').exists()).toBe(true)
    expect(wrapper.find('[data-test="logout"]').exists()).toBe(false)
    expect(wrapper.find('[data-test="dashboard"]').exists()).toBe(false)
  })

  it('shows Dashboard/Logout when authenticated', async () => {
    localStorage.setItem('authToken', 'mock-token')
    localStorage.setItem('user', JSON.stringify({ id: 1, email: 'test@example.com' }))

    const wrapper = await mountNavbar()
    await wrapper.vm.$nextTick()

    expect(wrapper.find('[data-test="dashboard"]').exists()).toBe(true)
    expect(wrapper.find('[data-test="logout"]').exists()).toBe(true)
    expect(wrapper.find('[data-test="login"]').exists()).toBe(false)
  })

  it('logs out when Logout button is clicked', async () => {
    // Setup authentication state
    localStorage.setItem('authToken', 'mock-token')
    localStorage.setItem('user', JSON.stringify({ id: 1, email: 'test@example.com' }))
    
    // Mock axios.post to prevent actual API calls
    vi.mock('axios', () => ({
      default: {
        post: vi.fn().mockResolvedValue({ data: { success: true } })
      }
    }))
    
    // Mock router.push to prevent navigation
    const mockRouterPush = vi.fn()
    const routerMock = {
      push: mockRouterPush
    }
    
    const wrapper = await mountNavbar()
    
    // Replace the router instance with our mock
    wrapper.vm.$options.setup = () => {
      const originalSetup = wrapper.vm.$options.setup
      const setupResult = originalSetup()
      setupResult.router = routerMock
      return setupResult
    }
    
    // Ensure the next tick
    await wrapper.vm.$nextTick()
    
    // Find and click logout button
    const logoutButton = wrapper.find('[data-test="logout"]')
    expect(logoutButton.exists()).toBe(true)
    await logoutButton.trigger('click')
    
    // Give time for async operation to complete
    await new Promise(resolve => setTimeout(resolve, 0))
    
    // Check localStorage was cleared
    expect(localStorage.getItem('authToken')).toBeNull()
    expect(localStorage.getItem('user')).toBeNull()
  })
  
})
