<template>
  <div class="font-sans text-gray-900 antialiased">
    <Navbar />
    <div class="min-h-screen bg-gray-50">
      <router-view></router-view>
    </div>
    <Footer />
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'

export default {
  components: {
    Navbar,
    Footer
  }
}
</script>