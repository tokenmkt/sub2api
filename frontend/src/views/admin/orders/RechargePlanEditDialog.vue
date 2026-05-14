<template>
  <BaseDialog :show="show" :title="plan ? t('payment.admin.editRechargePlan') : t('payment.admin.createRechargePlan')" width="wide" @close="emit('close')">
    <form id="recharge-plan-form" class="space-y-4" @submit.prevent="handleSave">
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="input-label">{{ t('payment.admin.planName') }} <span class="text-red-500">*</span></label>
          <input v-model="form.name" type="text" class="input" required />
        </div>
        <div>
          <label class="input-label">{{ t('payment.admin.badge') }}</label>
          <input v-model="form.badge" type="text" class="input" />
        </div>
      </div>

      <div>
        <label class="input-label">{{ t('payment.admin.planDescription') }}</label>
        <textarea v-model="form.description" rows="2" class="input"></textarea>
      </div>

      <div class="grid grid-cols-3 gap-4">
        <div>
          <label class="input-label">{{ t('payment.admin.price') }} <span class="text-red-500">*</span></label>
          <input v-model.number="form.price" type="number" step="0.01" min="0.01" class="input" required />
        </div>
        <div>
          <label class="input-label">{{ t('payment.admin.creditAmount') }} <span class="text-red-500">*</span></label>
          <input v-model.number="form.credit_amount" type="number" step="0.01" min="0.01" class="input" required />
        </div>
        <div>
          <label class="input-label">{{ t('payment.admin.originalPrice') }}</label>
          <input v-model.number="form.original_price" type="number" step="0.01" min="0" class="input" />
        </div>
      </div>

      <div>
        <label class="input-label">{{ t('payment.admin.purchaseUrl') }} <span class="text-red-500">*</span></label>
        <input v-model="form.purchase_url" type="url" class="input" placeholder="https://example.com/buy" required />
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="input-label">{{ t('payment.admin.sortOrder') }}</label>
          <input v-model.number="form.sort_order" type="number" min="0" class="input" />
        </div>
      </div>

      <div>
        <label class="input-label">{{ t('payment.admin.features') }}</label>
        <textarea v-model="featuresText" rows="3" class="input" :placeholder="t('payment.admin.featuresPlaceholder')"></textarea>
        <p class="mt-1 text-xs text-gray-500 dark:text-gray-400">{{ t('payment.admin.featuresHint') }}</p>
      </div>

      <div class="flex items-center gap-3">
        <label class="text-sm text-gray-700 dark:text-gray-300">{{ t('payment.admin.forSale') }}</label>
        <button
          type="button"
          :class="[
            'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2',
            form.for_sale ? 'bg-primary-500' : 'bg-gray-300 dark:bg-dark-600'
          ]"
          @click="form.for_sale = !form.for_sale"
        >
          <span :class="[
            'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
            form.for_sale ? 'translate-x-5' : 'translate-x-0'
          ]" />
        </button>
      </div>
    </form>
    <template #footer>
      <div class="flex justify-end gap-3">
        <button type="button" class="btn btn-secondary" @click="emit('close')">{{ t('common.cancel') }}</button>
        <button type="submit" form="recharge-plan-form" :disabled="saving" class="btn btn-primary">{{ saving ? t('common.saving') : t('common.save') }}</button>
      </div>
    </template>
  </BaseDialog>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/app'
import { adminPaymentAPI } from '@/api/admin/payment'
import { extractApiErrorMessage } from '@/utils/apiError'
import type { RechargePlan } from '@/types/payment'
import BaseDialog from '@/components/common/BaseDialog.vue'

const props = defineProps<{
  show: boolean
  plan: RechargePlan | null
}>()

const emit = defineEmits<{
  close: []
  saved: []
}>()

const { t } = useI18n()
const appStore = useAppStore()
const saving = ref(false)
const featuresText = ref('')
const form = reactive({
  name: '',
  description: '',
  price: 0,
  credit_amount: 0,
  original_price: 0,
  purchase_url: '',
  badge: '',
  sort_order: 0,
  for_sale: true,
})

watch(() => props.show, (visible) => {
  if (!visible) return
  if (props.plan) {
    Object.assign(form, {
      name: props.plan.name,
      description: props.plan.description,
      price: props.plan.price,
      credit_amount: props.plan.credit_amount,
      original_price: props.plan.original_price || 0,
      purchase_url: props.plan.purchase_url,
      badge: props.plan.badge || '',
      sort_order: props.plan.sort_order || 0,
      for_sale: props.plan.for_sale,
    })
    featuresText.value = (props.plan.features || []).join('\n')
  } else {
    Object.assign(form, {
      name: '',
      description: '',
      price: 0,
      credit_amount: 0,
      original_price: 0,
      purchase_url: '',
      badge: '',
      sort_order: 0,
      for_sale: true,
    })
    featuresText.value = ''
  }
})

function isHttpURL(value: string): boolean {
  try {
    const parsed = new URL(value)
    return parsed.protocol === 'http:' || parsed.protocol === 'https:'
  } catch {
    return false
  }
}

function buildPayload() {
  return {
    name: form.name,
    description: form.description,
    price: form.price,
    credit_amount: form.credit_amount,
    original_price: form.original_price || 0,
    features: featuresText.value.split('\n').map(f => f.trim()).filter(Boolean).join('\n'),
    purchase_url: form.purchase_url,
    badge: form.badge,
    sort_order: form.sort_order,
    for_sale: form.for_sale,
  }
}

async function handleSave() {
  if (!form.name.trim()) {
    appStore.showError(t('payment.admin.rechargePlanNameRequired'))
    return
  }
  if (!form.price || form.price <= 0 || !form.credit_amount || form.credit_amount <= 0) {
    appStore.showError(t('payment.admin.rechargePlanAmountRequired'))
    return
  }
  if (!isHttpURL(form.purchase_url)) {
    appStore.showError(t('payment.admin.purchaseUrlInvalid'))
    return
  }
  saving.value = true
  try {
    const payload = buildPayload()
    if (props.plan) await adminPaymentAPI.updateRechargePlan(props.plan.id, payload)
    else await adminPaymentAPI.createRechargePlan(payload)
    appStore.showSuccess(t('common.saved'))
    emit('close')
    emit('saved')
  } catch (err: unknown) {
    appStore.showError(extractApiErrorMessage(err, t('common.error')))
  } finally {
    saving.value = false
  }
}
</script>
