import { mount } from '@vue/test-utils'
import FormButton from '@/components/FormButton.vue'

describe('FormButton', () => {
  it('renders button with default slot content', () => {
    const wrapper = mount(FormButton, {
      slots: {
        default: 'Submit'
      }
    })

    expect(wrapper.find('button').text()).toBe('Submit')
    expect(wrapper.find('button').classes()).toContain('bg-indigo-600')
  })

  it('shows loading spinner when loading prop is true', () => {
    const wrapper = mount(FormButton, {
      props: {
        loading: true
      },
      slots: {
        default: 'Submit'
      }
    })

    expect(wrapper.find('svg').exists()).toBe(true)
    expect(wrapper.find('button').classes()).toContain('opacity-75')
  })

  it('emits click event when clicked', async () => {
    const wrapper = mount(FormButton, {
      slots: {
        default: 'Submit'
      }
    })

    await wrapper.find('button').trigger('click')
    expect(wrapper.emitted('click')).toBeTruthy()
  })
})