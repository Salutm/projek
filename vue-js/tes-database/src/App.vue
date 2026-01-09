<script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const items = ref([])
  const inputNama = ref('')
  const fechData = async () => {
    try {
      const response = await axios.get('http://localhost:3000/items')
      items.value = response.data      
    } catch (err) {
      consosle.error("gagal ambil data wak: ", err);
    }
  }

  const tambahData = async () => {
    if(!inputNama.value) return
    try {
      await axios.post('http://localhost:3000/items', {
        name: inputNama.value,
        description: 'deskripsi njir'
      })
      inputNama.value = ''
      fechData()
    } catch (err) {
      console.error("gagal tambah data wak:", err)
    }
  }


  onMounted(fechData)
</script>

<template>
  <h1>halo cuy</h1>
  <div style="padding: 20px;">
    <h2>CRUD Gabut: Vue + MySQL</h2>
    
    <input v-model="inputNama" placeholder="Ketik sesuatu..." @keyup.enter="tambahData">
    <button @click="tambahData">Simpan</button>

    <ul style="margin-top: 20px;">
      <li v-for="item in items" :key="item.id">
        {{ item.name }}
      </li>
    </ul>
  </div>
</template>

<style scoped>
</style>
