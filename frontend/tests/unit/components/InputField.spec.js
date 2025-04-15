import { mount } from '@vue/test-utils'
import InputField from '@/components/InputField.vue'

describe('InputField', () => {
  it('renders label and input correctly', () => {
    const wrapper = mount(InputField, {
      props: {
        id: 'test-input',
        label: 'Test Label',
        modelValue: '',
        type: 'text'
      }
    })

    expect(wrapper.find('label').text()).toBe('Test Label')
    expect(wrapper.find('input').attributes('id')).toBe('test-input')
    expect(wrapper.find('input').attributes('type')).toBe('text')
  })

  it('emits update event on input', async () => {
    const wrapper = mount(InputField, {
      props: {
        id: 'test-input',
        label: 'Test Label',
        modelValue: '',
        type: 'text'
      }
    })

    const input = wrapper.find('input')
    await input.setValue('test value')
    
    expect(wrapper.emitted('update:modelValue')).toBeTruthy()
    expect(wrapper.emitted('update:modelValue')[0]).toEqual(['test value'])
  })

  it('shows error message when error prop is provided', () => {
    const errorMessage = 'This field is required'
    const wrapper = mount(InputField, {
      props: {
        id: 'test-input',
        label: 'Test Label',
        modelValue: '',
        type: 'text',
        error: errorMessage
      }
    })

    expect(wrapper.find('p').text()).toBe(errorMessage)
    expect(wrapper.find('input').classes()).toContain('border-red-300')
  })
})