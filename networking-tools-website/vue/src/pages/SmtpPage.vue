<script setup>
import { ref } from 'vue'
import SmtpService from '../services/SmtpService'

const port = ref('25')
const host = ref('')
const from = ref('')
const to = ref('')
const subject = ref('')
const body = ref('')
const isSending = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const trace = ref([])

async function sendEmail() {
  errorMessage.value = ''
  successMessage.value = ''
  trace.value = []
  isSending.value = true
  try {
    const { data } = await SmtpService.send({ host: host.value, port: port.value, from: from.value, to: to.value, subject: subject.value, body: body.value })
    successMessage.value = data.message
    trace.value = data.trace || []
  } catch (error) {
    errorMessage.value = error.response?.data?.message || 'SMTP sunucusuna bağlanırken hata oluştu.'
    trace.value = error.response?.data?.trace || []
  } finally {
    isSending.value = false
  }
}
</script>

<template>
  <div class="smtp-page">
    <div class="intro">
      <p class="eyebrow">SMTP E-POSTA GÖNDERİMİ</p>
      <h1>E-posta gönder</h1>
      <p>Belirttiğiniz SMTP relay sunucusu üzerinden e-posta teslimi yapın.</p>
    </div>

    <form class="mail-form" @submit.prevent="sendEmail">
      <div class="relay-info">
        <label>SMTP sunucusu
          <input v-model.trim="host" type="text" placeholder="mail01.kurum.local veya 172.16.0.6" required />
        </label>
        <label>Port
          <select v-model="port"><option value="25">25</option><option value="587">587</option></select>
        </label>
      </div>
      <label>Gönderen e-posta adresi<input v-model.trim="from" type="email" placeholder="gonderen@ornek.com" required /></label>
      <label>Alıcı e-posta adresi<input v-model.trim="to" type="email" placeholder="alici@ornek.com" required /></label>
      <label>Konu<input v-model="subject" type="text" maxlength="200" placeholder="Mesaj konusu" required /></label>
      <label>Mesaj<textarea v-model="body" rows="9" maxlength="10000" placeholder="Mesajınızı yazın..." required></textarea></label>
      <button type="submit" :disabled="isSending">{{ isSending ? 'Gönderiliyor...' : 'E-posta gönder' }}</button>
    </form>

    <p v-if="successMessage" class="success">{{ successMessage }}</p>
    <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    <section v-if="trace.length" class="trace">
      <h2>SMTP işlem kaydı</h2>
      <ol><li v-for="(step, index) in trace" :key="index">{{ step }}</li></ol>
    </section>
  </div>
</template>

<style scoped>
.smtp-page { max-width: 760px; }.intro { margin-bottom: 28px; }.intro h1 { margin: 4px 0 8px; font-size: clamp(1.8rem, 4vw, 2.5rem); }.intro p:last-child { margin: 0; color: #aabed0; }.intro code { color: #8debd7; }.eyebrow { margin: 0; color: #65e5cc; font-size: .76rem; font-weight: 800; letter-spacing: .13em; }
.mail-form { display: grid; gap: 18px; }.mail-form > label { display: grid; gap: 8px; }.mail-form input, .mail-form textarea { width: 100%; }.relay-info { display: flex; flex-wrap: wrap; align-items: center; gap: 12px; padding: 14px; border: 1px solid rgba(76,224,193,.25); border-radius: 12px; background: rgba(76,224,193,.07); }.relay-info label { display: flex; align-items: center; gap: 8px; }.relay-info label:first-child { flex: 1 1 360px; }.relay-info input { min-height: 36px; }.relay-info select { min-height: 36px; padding: 4px 8px; }.success { margin-top: 20px; padding: 12px 14px; border: 1px solid rgba(76,224,193,.35); border-radius: 10px; background: rgba(76,224,193,.1); color: #a6f1dc; }.trace { margin-top: 24px; padding: 20px; border: 1px solid rgba(132,204,239,.2); border-radius: 14px; background: rgba(5,15,29,.45); }.trace h2 { margin: 0 0 12px; font-size: 1.05rem; }.trace ol { margin: 0; padding-left: 24px; color: #c7d8e8; font-family: ui-monospace, SFMono-Regular, Menlo, monospace; font-size: .84rem; line-height: 1.65; }.trace li { padding-left: 4px; }
</style>
