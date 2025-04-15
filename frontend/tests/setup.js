import { config } from '@vue/test-utils'

// Mock global components
config.global.stubs = {
  'router-link': true,
  'router-view': true
}