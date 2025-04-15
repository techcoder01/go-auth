<template>
    <div class="mb-4">
      <label :for="id" class="block text-sm font-medium text-gray-700">{{ label }}</label>
      <div class="mt-1 relative rounded-md shadow-sm">
        <input 
          :id="id" 
          :type="type" 
          :value="modelValue" 
          @input="$emit('update:modelValue', $event.target.value)"
          class="appearance-none block w-full px-3 py-2 border rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          :class="{ 'border-red-300 text-red-900': error, 'border-gray-300': !error }"
        />
      </div>
      <p v-if="error" class="mt-2 text-sm text-red-600">{{ error }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      id: {
        type: String,
        required: true
      },
      label: {
        type: String,
        required: true
      },
      modelValue: {
        type: String,
        default: ''
      },
      type: {
        type: String,
        default: 'text'
      },
      error: {
        type: String,
        default: ''
      }
    },
    emits: ['update:modelValue']
  }
  </script>