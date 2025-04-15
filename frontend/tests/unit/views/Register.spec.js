import { mount } from '@vue/test-utils'
import Register from '@/views/Register.vue'
import { createRouter, createWebHistory } from 'vue-router'
import axios from 'axios'

// Mock axios globally
vi.mock('axios')

const routes = [
  { path: '/register', name: 'Register' },
  { path: '/dashboard', name: 'Dashboard' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

describe('Register', () => {
  let wrapper

  beforeEach(() => {
    // Clear all mocks and localStorage before each test
    vi.clearAllMocks()
    localStorage.clear()
    
    wrapper = mount(Register, {
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

  it('renders registration form', () => {
    expect(wrapper.find('form').exists()).toBe(true)
    expect(wrapper.find('h2').text()).toContain('Create a new account')
  })

  it('validates all fields', async () => {
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.vm.errors.email).toBe('Email is required')
    expect(wrapper.vm.errors.password).toBe('Password is required')
    expect(wrapper.vm.errors.confirmPassword).toBe('Please confirm your password')
    expect(wrapper.vm.errors.terms).toBe('You must agree to the terms and conditions')
  })
  
  it('shows error when passwords dont match', async () => {
    // Modify data directly through wrapper.vm
    wrapper.vm.password = 'password123'
    wrapper.vm.confirmPassword = 'different'
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.vm.errors.confirmPassword).toBe('Passwords do not match')
  })
  
  it('submits form with valid data', async () => {
    const mockResponse = { 
      status: 201,
      data: { 
        user: { id: 1, email: 'test@example.com' },
        message: 'Registration successful'
      }
    }
    axios.post.mockResolvedValue(mockResponse)
  
    // Modify data directly through wrapper.vm
    wrapper.vm.email = 'test@example.com'
    wrapper.vm.password = 'password123'
    wrapper.vm.confirmPassword = 'password123'
    wrapper.vm.agreeToTerms = true
    
    await wrapper.find('form').trigger('submit')
  
    expect(axios.post).toHaveBeenCalledWith('http://localhost:8080/api/register', {
      email: 'test@example.com',
      password: 'password123'
    })
  })
  
  it('handles registration error', async () => {
    const error = { 
      response: { 
        status: 400, 
        data: { error: 'Email already in use' } 
      } 
    }
    axios.post.mockRejectedValue(error)
  
    // Modify data directly through wrapper.vm
    wrapper.vm.email = 'test@example.com'
    wrapper.vm.password = 'password123'
    wrapper.vm.confirmPassword = 'password123'
    wrapper.vm.agreeToTerms = true
    
    await wrapper.find('form').trigger('submit')
  
    expect(wrapper.vm.apiError).toBe('Email already in use')
  })
  
  it('redirects to dashboard on successful registration', async () => {
    const mockResponse = { 
      status: 201,
      data: { 
        user: { id: 1, email: 'test@example.com' },
        message: 'Registration successful'
      }
    }
    axios.post.mockResolvedValue(mockResponse)
    const push = vi.spyOn(router, 'push')
  
    // Modify data directly through wrapper.vm
    wrapper.vm.email = 'test@example.com'
    wrapper.vm.password = 'password123'
    wrapper.vm.confirmPassword = 'password123'
    wrapper.vm.agreeToTerms = true
    
    await wrapper.find('form').trigger('submit')
  
    expect(push).toHaveBeenCalledWith('/dashboard')
  })  
})