module.exports = {
    root: true,
    vitest: true,  // If using Vitest
    env: {
      node: true
    },
    extends: [
      'plugin:vue/vue3-essential',
      'eslint:recommended'
    ],
    parserOptions: {
      parser: '@babel/eslint-parser',
      requireConfigFile: false
    },
    rules: {
      'vue/multi-word-component-names': 'off',
      'no-unused-vars': 'warn'
    }
  }