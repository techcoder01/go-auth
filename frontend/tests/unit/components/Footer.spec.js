import { mount } from '@vue/test-utils'
import Footer from '@/components/Footer.vue'

describe('Footer', () => {
  it('renders correctly', () => {
    const wrapper = mount(Footer)
    expect(wrapper.find('footer').exists()).toBe(true)
    expect(wrapper.text()).toContain('VueAuth')
    expect(wrapper.text()).toContain(new Date().getFullYear().toString())
  })

  it('contains social media links', () => {
    const wrapper = mount(Footer)
    expect(wrapper.find('a').exists()).toBe(true)
    expect(wrapper.find('svg').exists()).toBe(true)
  })
})