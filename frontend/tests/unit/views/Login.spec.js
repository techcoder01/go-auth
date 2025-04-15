import { mount } from '@vue/test-utils'
import Login from '@/views/Login.vue'
import { createRouter, createWebHistory } from 'vue-router'
import axios from 'axios'

vi.mock('axios')

const routes = [
  { path: '/', name: 'Home' },
  { path: '/login', name: 'Login' },
  { path: '/dashboard', name: 'Dashboard' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

describe('Login', () => {
  let wrapper

  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
    wrapper = mount(Login, {
      global: {
        plugins: [router],
        stubs: {
          InputField: true,
          FormButton: true,
          RouterLink: true
        }
      }
    })
  })

  it('renders login form', () => {
    expect(wrapper.find('form').exists()).toBe(true)
    expect(wrapper.find('h2').text()).toContain('Sign in to your account')
  })

  it('validates email and password', async () => {
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.vm.errors.email).toBe('Email is required')
    expect(wrapper.vm.errors.password).toBe('Password is required')
  })

  it('shows error for invalid email format', async () => {
    wrapper.vm.email = 'invalid-email' // Set the data directly
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.vm.errors.email).toBe('Email is invalid')
  })

  it('submits form with valid data', async () => {
    const mockResponse = { data: { token: 'test-token' } }
    axios.post.mockResolvedValue(mockResponse)

    wrapper.vm.email = 'test@example.com' // Set the data directly
    wrapper.vm.password = 'password123' // Set the data directly
    await wrapper.find('form').trigger('submit')

    expect(axios.post).toHaveBeenCalledWith('https://go-auth-ftrw.onrender.com/api/login', {
      email: 'test@example.com',
      password: 'password123'
    })
  })

  it('stores token and redirects on successful login', async () => {
    const mockResponse = { data: { token: 'test-token' } }
    axios.post.mockResolvedValue(mockResponse)
    const push = vi.spyOn(router, 'push')

    wrapper.vm.email = 'test@example.com' // Set the data directly
    wrapper.vm.password = 'password123' // Set the data directly
    await wrapper.find('form').trigger('submit')

    expect(localStorage.getItem('authToken')).toBe('test-token')
    expect(push).toHaveBeenCalledWith('/dashboard')
  })

  it('handles login error', async () => {
    const error = { response: { status: 401 } }
    axios.post.mockRejectedValue(error)

    wrapper.vm.email = 'test@example.com' // Set the data directly
    wrapper.vm.password = 'password123' // Set the data directly
    await wrapper.find('form').trigger('submit')

    expect(wrapper.vm.apiError).toBe('Invalid email or password')
  })
})
