<template>
  <AppLayout>
    <div class="space-y-4">
      <div class="flex items-center justify-end gap-2">
        <div class="mr-auto flex rounded-lg bg-gray-100 p-1 dark:bg-dark-800">
          <button
            type="button"
            :class="tabButtonClass('subscription')"
            @click="activeTab = 'subscription'"
          >
            {{ t('payment.admin.subscriptionPlans') }}
          </button>
          <button
            type="button"
            :class="tabButtonClass('recharge')"
            @click="activeTab = 'recharge'"
          >
            {{ t('payment.admin.rechargePlans') }}
          </button>
        </div>
        <button @click="refreshActiveTab" :disabled="currentLoading" class="btn btn-secondary" :title="t('common.refresh')">
          <Icon name="refresh" size="md" :class="currentLoading ? 'animate-spin' : ''" />
        </button>
        <button v-if="activeTab === 'subscription'" @click="openPlanEdit(null)" class="btn btn-primary">{{ t('payment.admin.createPlan') }}</button>
        <button v-else @click="openRechargePlanEdit(null)" class="btn btn-primary">{{ t('payment.admin.createRechargePlan') }}</button>
      </div>

      <DataTable v-if="activeTab === 'subscription'" :columns="planColumns" :data="plans" :loading="plansLoading">
        <template #cell-name="{ value, row }">
          <span class="text-sm font-medium" :class="getPlanNameClass(row.group_id)">{{ value }}</span>
        </template>
        <template #cell-group_id="{ value }">
          <span v-if="isGroupMissing(value)" class="text-sm">
            <span class="text-gray-400">#{{ value }}</span>
            <span class="ml-1 badge badge-danger">{{ t('payment.admin.groupMissing') }}</span>
          </span>
          <GroupBadge
            v-else-if="getGroup(value)"
            :name="getGroup(value)!.name"
            :platform="getGroup(value)!.platform"
            :rate-multiplier="getGroup(value)!.rate_multiplier"
          />
          <span v-else class="text-sm text-gray-400">-</span>
        </template>
        <template #cell-price="{ value, row }">
          <div class="text-sm">
            <span class="font-medium text-gray-900 dark:text-white">${{ (value ?? 0).toFixed(2) }}</span>
            <span v-if="row.original_price" class="ml-1 text-xs text-gray-400 line-through">${{ row.original_price.toFixed(2) }}</span>
          </div>
        </template>
        <template #cell-validity_days="{ value, row }">
          <span class="text-sm">{{ value }} {{ t('payment.admin.' + (row.validity_unit || 'days')) }}</span>
        </template>
        <template #cell-for_sale="{ value, row }">
          <SaleToggle :value="value" @toggle="toggleForSale(row)" />
        </template>
        <template #cell-actions="{ row }">
          <RowActions @edit="openPlanEdit(row)" @delete="confirmDeletePlan(row)" />
        </template>
      </DataTable>

      <DataTable v-else :columns="rechargePlanColumns" :data="rechargePlans" :loading="rechargePlansLoading">
        <template #cell-name="{ value, row }">
          <div class="min-w-0">
            <span class="text-sm font-medium text-gray-900 dark:text-white">{{ value }}</span>
            <span v-if="row.badge" class="ml-2 rounded-full bg-primary-50 px-2 py-0.5 text-xs text-primary-600 dark:bg-primary-900/30 dark:text-primary-300">{{ row.badge }}</span>
          </div>
        </template>
        <template #cell-price="{ value, row }">
          <div class="text-sm">
            <span class="font-medium text-gray-900 dark:text-white">¥{{ (value ?? 0).toFixed(2) }}</span>
            <span v-if="row.original_price" class="ml-1 text-xs text-gray-400 line-through">¥{{ row.original_price.toFixed(2) }}</span>
          </div>
        </template>
        <template #cell-credit_amount="{ value }">
          <span class="text-sm font-medium text-green-600 dark:text-green-400">${{ (value ?? 0).toFixed(2) }}</span>
        </template>
        <template #cell-purchase_url="{ value }">
          <a :href="value" target="_blank" rel="noopener noreferrer" class="inline-flex max-w-xs items-center gap-1 truncate text-sm text-primary-600 hover:text-primary-700 dark:text-primary-400">
            <Icon name="externalLink" size="xs" />
            <span class="truncate">{{ value }}</span>
          </a>
        </template>
        <template #cell-for_sale="{ value, row }">
          <SaleToggle :value="value" @toggle="toggleRechargeForSale(row)" />
        </template>
        <template #cell-actions="{ row }">
          <RowActions @edit="openRechargePlanEdit(row)" @delete="confirmDeleteRechargePlan(row)" />
        </template>
      </DataTable>
    </div>

    <PlanEditDialog :show="showPlanDialog" :plan="editingPlan" :groups="groups" @close="showPlanDialog = false" @saved="loadPlans" />
    <RechargePlanEditDialog :show="showRechargePlanDialog" :plan="editingRechargePlan" @close="showRechargePlanDialog = false" @saved="loadRechargePlans" />

    <ConfirmDialog :show="showDeletePlanDialog" :title="t('payment.admin.deletePlan')" :message="t('payment.admin.deletePlanConfirm')" :confirm-text="t('common.delete')" danger @confirm="handleDeletePlan" @cancel="showDeletePlanDialog = false" />
    <ConfirmDialog :show="showDeleteRechargePlanDialog" :title="t('payment.admin.deleteRechargePlan')" :message="t('payment.admin.deleteRechargePlanConfirm')" :confirm-text="t('common.delete')" danger @confirm="handleDeleteRechargePlan" @cancel="showDeleteRechargePlanDialog = false" />
  </AppLayout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, defineComponent, h } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/app'
import { adminPaymentAPI } from '@/api/admin/payment'
import { extractI18nErrorMessage } from '@/utils/apiError'
import adminAPI from '@/api/admin'
import type { RechargePlan, SubscriptionPlan } from '@/types/payment'
import type { AdminGroup } from '@/types'
import type { Column } from '@/components/common/types'
import AppLayout from '@/components/layout/AppLayout.vue'
import DataTable from '@/components/common/DataTable.vue'
import ConfirmDialog from '@/components/common/ConfirmDialog.vue'
import Icon from '@/components/icons/Icon.vue'
import GroupBadge from '@/components/common/GroupBadge.vue'
import PlanEditDialog from './PlanEditDialog.vue'
import RechargePlanEditDialog from './RechargePlanEditDialog.vue'
import { platformTextClass } from '@/utils/platformColors'

const { t } = useI18n()
const appStore = useAppStore()
const activeTab = ref<'subscription' | 'recharge'>('subscription')

const SaleToggle = defineComponent({
  props: { value: { type: Boolean, required: true } },
  emits: ['toggle'],
  setup(props, { emit }) {
    return () => h('button', {
      type: 'button',
      class: [
        'relative inline-flex h-5 w-9 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2',
        props.value ? 'bg-primary-500' : 'bg-gray-300 dark:bg-dark-600',
      ],
      onClick: () => emit('toggle'),
    }, [
      h('span', {
        class: [
          'pointer-events-none inline-block h-4 w-4 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out',
          props.value ? 'translate-x-4' : 'translate-x-0',
        ],
      }),
    ])
  },
})

const RowActions = defineComponent({
  emits: ['edit', 'delete'],
  setup(_, { emit }) {
    return () => h('div', { class: 'flex items-center gap-2' }, [
      h('button', {
        class: 'flex flex-col items-center gap-0.5 rounded-lg p-1.5 text-gray-500 transition-colors hover:bg-blue-50 hover:text-blue-600 dark:hover:bg-blue-900/20 dark:hover:text-blue-400',
        onClick: () => emit('edit'),
      }, [h(Icon, { name: 'edit', size: 'sm' }), h('span', { class: 'text-xs' }, t('common.edit'))]),
      h('button', {
        class: 'flex flex-col items-center gap-0.5 rounded-lg p-1.5 text-gray-500 transition-colors hover:bg-red-50 hover:text-red-600 dark:hover:bg-red-900/20 dark:hover:text-red-400',
        onClick: () => emit('delete'),
      }, [h(Icon, { name: 'trash', size: 'sm' }), h('span', { class: 'text-xs' }, t('common.delete'))]),
    ])
  },
})

const groups = ref<AdminGroup[]>([])

async function loadGroups() {
  try {
    groups.value = await adminAPI.groups.getAll()
  } catch { /* ignore */ }
}

function getGroup(id: number): AdminGroup | undefined {
  return groups.value.find(g => g.id === id)
}

function isGroupMissing(id: number): boolean {
  return id > 0 && !groups.value.find(g => g.id === id)
}

function getPlanNameClass(groupId: number): string {
  const group = getGroup(groupId)
  return group ? platformTextClass(group.platform) : 'text-gray-900 dark:text-white'
}

const plansLoading = ref(false)
const plans = ref<SubscriptionPlan[]>([])
const showPlanDialog = ref(false)
const showDeletePlanDialog = ref(false)
const editingPlan = ref<SubscriptionPlan | null>(null)
const deletingPlanId = ref<number | null>(null)

const rechargePlansLoading = ref(false)
const rechargePlans = ref<RechargePlan[]>([])
const showRechargePlanDialog = ref(false)
const showDeleteRechargePlanDialog = ref(false)
const editingRechargePlan = ref<RechargePlan | null>(null)
const deletingRechargePlanId = ref<number | null>(null)

const currentLoading = computed(() => activeTab.value === 'subscription' ? plansLoading.value : rechargePlansLoading.value)

const planColumns = computed((): Column[] => [
  { key: 'id', label: 'ID' },
  { key: 'name', label: t('payment.admin.planName') },
  { key: 'group_id', label: t('payment.admin.group') },
  { key: 'price', label: t('payment.admin.price') },
  { key: 'validity_days', label: t('payment.admin.validityDays') },
  { key: 'for_sale', label: t('payment.admin.forSale') },
  { key: 'sort_order', label: t('payment.admin.sortOrder') },
  { key: 'actions', label: t('common.actions') },
])

const rechargePlanColumns = computed((): Column[] => [
  { key: 'id', label: 'ID' },
  { key: 'name', label: t('payment.admin.planName') },
  { key: 'price', label: t('payment.admin.price') },
  { key: 'credit_amount', label: t('payment.admin.creditAmount') },
  { key: 'purchase_url', label: t('payment.admin.purchaseUrl') },
  { key: 'for_sale', label: t('payment.admin.forSale') },
  { key: 'sort_order', label: t('payment.admin.sortOrder') },
  { key: 'actions', label: t('common.actions') },
])

function tabButtonClass(tab: 'subscription' | 'recharge') {
  return [
    'rounded-md px-3 py-1.5 text-sm font-medium transition',
    activeTab.value === tab
      ? 'bg-white text-gray-900 shadow dark:bg-dark-700 dark:text-white'
      : 'text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300',
  ]
}

function normalizeFeatures<T extends { features: string | string[] }>(item: T): Omit<T, 'features'> & { features: string[] } {
  return {
    ...item,
    features: typeof item.features === 'string'
      ? item.features.split('\n').map((f: string) => f.trim()).filter(Boolean)
      : (item.features || []),
  }
}

async function loadPlans() {
  plansLoading.value = true
  try {
    const res = await adminPaymentAPI.getPlans()
    plans.value = (res.data || []).map(normalizeFeatures) as SubscriptionPlan[]
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  } finally {
    plansLoading.value = false
  }
}

async function loadRechargePlans() {
  rechargePlansLoading.value = true
  try {
    const res = await adminPaymentAPI.getRechargePlans()
    rechargePlans.value = (res.data || []).map(normalizeFeatures) as RechargePlan[]
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  } finally {
    rechargePlansLoading.value = false
  }
}

function refreshActiveTab() {
  if (activeTab.value === 'subscription') loadPlans()
  else loadRechargePlans()
}

function openPlanEdit(plan: SubscriptionPlan | null) {
  editingPlan.value = plan
  showPlanDialog.value = true
}

function openRechargePlanEdit(plan: RechargePlan | null) {
  editingRechargePlan.value = plan
  showRechargePlanDialog.value = true
}

async function toggleForSale(plan: SubscriptionPlan) {
  try {
    await adminPaymentAPI.updatePlan(plan.id, { for_sale: !plan.for_sale })
    plan.for_sale = !plan.for_sale
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  }
}

async function toggleRechargeForSale(plan: RechargePlan) {
  try {
    await adminPaymentAPI.updateRechargePlan(plan.id, { for_sale: !plan.for_sale })
    plan.for_sale = !plan.for_sale
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  }
}

function confirmDeletePlan(plan: SubscriptionPlan) {
  deletingPlanId.value = plan.id
  showDeletePlanDialog.value = true
}

function confirmDeleteRechargePlan(plan: RechargePlan) {
  deletingRechargePlanId.value = plan.id
  showDeleteRechargePlanDialog.value = true
}

async function handleDeletePlan() {
  if (!deletingPlanId.value) return
  try {
    await adminPaymentAPI.deletePlan(deletingPlanId.value)
    appStore.showSuccess(t('common.deleted'))
    showDeletePlanDialog.value = false
    loadPlans()
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  }
}

async function handleDeleteRechargePlan() {
  if (!deletingRechargePlanId.value) return
  try {
    await adminPaymentAPI.deleteRechargePlan(deletingRechargePlanId.value)
    appStore.showSuccess(t('common.deleted'))
    showDeleteRechargePlanDialog.value = false
    loadRechargePlans()
  } catch (err: unknown) {
    appStore.showError(extractI18nErrorMessage(err, t, 'payment.errors', t('common.error')))
  }
}

onMounted(() => {
  loadGroups()
  loadPlans()
  loadRechargePlans()
})
</script>
