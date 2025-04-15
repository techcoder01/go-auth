import { mount } from '@vue/test-utils'
import Home from '@/views/Home.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/login', name: 'Login' },
  { path: '/register', name: 'Register' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

describe('Home', () => {
  beforeEach(() => {
    router.push('/') // Ensure the router is pushed to the homepage route
  })

  it('renders home page content', async () => {
    const wrapper = mount(Home, {
      global: {
        plugins: [router], // Include the router plugin
      }
    })

    await router.isReady() // Wait for the router to be ready before running the assertions

    expect(wrapper.find('h1').text()).toContain('Secure and simple')
    expect(wrapper.findAll('router-link-stub').length).toBe(2)
  })
})
