module.exports = {
  root: true,
  env: {
    node: true,
    "vitest-globals/env": true
  },
  extends: [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    "plugin:vitest-globals/recommended",
  ],
  parserOptions: {
    parser: '@babel/eslint-parser',
    requireConfigFile: false,
    ecmaVersion: 2022,
  },
  plugins: ['vitest'],
  rules: {
    'vue/multi-word-component-names': 'off',
    'no-unused-vars': 'warn',
  },
};
