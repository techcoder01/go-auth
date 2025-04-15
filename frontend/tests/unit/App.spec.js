import { mount } from '@vue/test-utils'
import App from '@/App.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', component: { template: '<div>Home</div>' } }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

describe('App', () => {
  it('renders navbar, router view and footer', () => {
    const wrapper = mount(App, {
      global: {
        plugins: [router],
        stubs: ['Navbar', 'Footer', 'RouterView']
      }
    })

    expect(wrapper.findComponent({ name: 'Navbar' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'Footer' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'RouterView' }).exists()).toBe(true)
  })
})