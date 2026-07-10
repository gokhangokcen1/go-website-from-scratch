<script setup>
import { ref } from 'vue'
import IPScannerService from '../services/IPScannerService'

const ip = ref('')
const cidr = ref(24)
const sonuclar = ref([])
const hataMesaji = ref('')
const taraniyor = ref(false)
const yalakaModu = ref(false)



async function tara() {
  hataMesaji.value = ''
  sonuclar.value = []
  taraniyor.value = true

  try {
    const response = await IPScannerService.tara(ip.value, Number(cidr.value))
    sonuclar.value = response.data
  } catch (error) {
    hataMesaji.value = error.response?.data?.error || 'Tarama basarisiz'
  } finally {
    taraniyor.value = false
  }
}

function ayaktaSayisi() {
  return sonuclar.value.filter(s => s.ayakta).length
}

function toplamIP() { return Math.pow(2, 32 - Number(cidr.value)) - 2 }
</script>

<template>
  <div>
    <h1>IP Scanner</h1>

    <form @submit.prevent="tara">
      <input v-model="ip" placeholder="Örnek: 192.168.30.110" />
      <input v-model="cidr" type="number" min="1" max="30" placeholder="CIDR (ornek: 24)" />
      <button type="submit" :disabled="taraniyor" @click="yalakaModu = false">
        {{ taraniyor ? 'Taraniyor...' : 'Tara' }}
      </button>
      <!-- GEREKİRSE BURAYI GİZLE -->
       
      <!-- <button type="submit" :disabled="taraniyor" @click="yalakaModu = true"
      style="min-width: 240px; flex: 0 0 240px;">
        {{ taraniyor ? 'Taraniyor...' : 'Tara (ama yalaka şekilde)' }}
      </button> -->

    </form>

    <p v-if="hataMesaji" class="hata">{{ hataMesaji }}</p>

    <p v-if="sonuclar.length > 0">
      Toplam {{ toplamIP() }} IP tarandi, {{ ayaktaSayisi() }} tanesi ayakta.
    </p>

    <table v-if="sonuclar.length > 0">
      <thead>
        <tr>
          <th>IP</th>
          <th>Durum</th>
          <th>MAC Adresi</th>
          <th>Uretici</th>
          <!-- <td v-if="yalakaModu">
            <div style="min-width: 130px;">
              
            </div>
          </td>    SİL -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="s in sonuclar" :key="s.ip">
          <td>{{ s.ip }}</td>
          <td :class="s.ayakta ? 'ayakta' : 'kapali'">
            {{ s.ayakta ? 'Ayakta' : 'Kapali' }}
          </td>
          <td>{{ s.mac || '-' }}</td>
          <td>{{ s.uretici || '-' }}</td>
          <!-- <td v-if="yalakaModu">{{ '' }}</td> -->
          
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
form {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}
input, button {
  padding: 8px;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  border: 1px solid #ccc;
  padding: 8px;
  text-align: left;
}
.ayakta {
  color: green;
  font-weight: bold;
}
.kapali {
  color: #999;
}
.hata {
  color: red;
}
</style>