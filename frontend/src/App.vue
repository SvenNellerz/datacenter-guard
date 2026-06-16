<template>
  <div class="min-h-screen bg-gray-950 text-gray-100">
    <Navbar />
    <div class="flex">
      <Sidebar />
      <main class="flex-1 ml-64 pt-16 min-h-screen">
        <div class="p-6">
          <RouterView />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import Navbar from '@/components/Navbar.vue'
import Sidebar from '@/components/Sidebar.vue'
</script>
