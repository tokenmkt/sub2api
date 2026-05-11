<template>
  <div v-if="homeContent" class="min-h-screen">
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <div v-else v-html="homeContent"></div>
  </div>

  <div
    v-else
    class="relative flex min-h-screen flex-col overflow-hidden bg-[radial-gradient(circle_at_top,rgba(45,212,191,0.18),transparent_34%),linear-gradient(180deg,#f4fbfa_0%,#eff7f5_42%,#f8fafc_100%)] dark:bg-[radial-gradient(circle_at_top,rgba(20,184,166,0.24),transparent_28%),linear-gradient(180deg,#051816_0%,#091d1a_40%,#020617_100%)]"
  >
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="absolute -left-20 top-16 h-72 w-72 rounded-full bg-emerald-300/25 blur-3xl"></div>
      <div class="absolute right-0 top-0 h-96 w-96 rounded-full bg-cyan-300/20 blur-3xl"></div>
      <div class="absolute bottom-0 left-1/3 h-80 w-80 rounded-full bg-teal-400/15 blur-3xl"></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(15,118,110,0.07)_1px,transparent_1px),linear-gradient(90deg,rgba(15,118,110,0.07)_1px,transparent_1px)] bg-[size:72px_72px] [mask-image:linear-gradient(to_bottom,rgba(0,0,0,0.25),transparent_80%)]"
      ></div>
    </div>

    <header class="relative z-20 px-6 py-5">
      <nav class="mx-auto flex max-w-6xl items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="h-11 w-11 overflow-hidden rounded-2xl bg-white/90 p-1 shadow-[0_16px_40px_rgba(15,118,110,0.18)] ring-1 ring-emerald-900/5 dark:bg-slate-900/80 dark:ring-white/10">
            <img :src="siteLogo || '/logo.png'" alt="Logo" class="h-full w-full object-contain" />
          </div>
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.28em] text-teal-700 dark:text-teal-300">
              {{ siteName }}
            </p>
            <p class="text-sm text-slate-500 dark:text-slate-300">
              tokenMKT official API gateway
            </p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <LocaleSwitcher />

          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="rounded-xl border border-white/60 bg-white/70 p-2.5 text-slate-600 shadow-sm backdrop-blur transition hover:border-teal-200 hover:text-slate-900 dark:border-white/10 dark:bg-slate-900/60 dark:text-slate-300 dark:hover:border-teal-400/40 dark:hover:text-white"
            :title="t('home.viewDocs')"
          >
            <Icon name="book" size="md" />
          </a>

          <button
            @click="toggleTheme"
            class="rounded-xl border border-white/60 bg-white/70 p-2.5 text-slate-600 shadow-sm backdrop-blur transition hover:border-teal-200 hover:text-slate-900 dark:border-white/10 dark:bg-slate-900/60 dark:text-slate-300 dark:hover:border-teal-400/40 dark:hover:text-white"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          >
            <Icon v-if="isDark" name="sun" size="md" />
            <Icon v-else name="moon" size="md" />
          </button>

          <router-link
            v-if="isAuthenticated"
            :to="dashboardPath"
            class="inline-flex items-center gap-2 rounded-full bg-slate-950 px-4 py-2 text-sm font-medium text-white shadow-[0_16px_30px_rgba(15,23,42,0.2)] transition hover:bg-slate-800 dark:bg-white dark:text-slate-900 dark:hover:bg-slate-100"
          >
            <span
              class="flex h-6 w-6 items-center justify-center rounded-full bg-gradient-to-br from-teal-400 to-emerald-500 text-[11px] font-semibold text-white"
            >
              {{ userInitial }}
            </span>
            <span>{{ t('home.dashboard') }}</span>
          </router-link>
          <router-link
            v-else
            to="/login"
            class="inline-flex items-center rounded-full bg-slate-950 px-4 py-2 text-sm font-medium text-white shadow-[0_16px_30px_rgba(15,23,42,0.2)] transition hover:bg-slate-800 dark:bg-white dark:text-slate-900 dark:hover:bg-slate-100"
          >
            {{ t('home.login') }}
          </router-link>
        </div>
      </nav>
    </header>

    <main class="relative z-10 flex-1 px-6 pb-10 pt-6">
      <div class="mx-auto flex max-w-6xl flex-col gap-16">
        <section class="grid items-center gap-8 lg:grid-cols-[0.84fr_1.16fr] lg:gap-10">
          <div class="max-w-2xl">
            <div
              class="inline-flex items-center gap-2 rounded-full border border-teal-200/80 bg-white/80 px-4 py-1.5 text-[11px] font-semibold uppercase tracking-[0.2em] text-teal-700 shadow-sm backdrop-blur dark:border-teal-400/20 dark:bg-slate-900/70 dark:text-teal-300"
            >
              <span class="h-2 w-2 rounded-full bg-teal-500"></span>
              {{ t('home.badge') }}
            </div>

            <p class="mt-5 text-xs font-semibold uppercase tracking-[0.28em] text-slate-500 dark:text-slate-400">
              {{ siteName }}
            </p>
            <h1
              class="mt-3 max-w-3xl text-[38px] font-semibold leading-[1.02] text-slate-950 dark:text-white md:text-[46px] lg:text-[58px]"
            >
              {{ t('home.heroTitle') }}
            </h1>
            <p class="mt-4 max-w-2xl text-[15px] leading-7 text-slate-600 dark:text-slate-300 md:text-[17px]">
              {{ t('home.heroSubtitle') }}
            </p>
            <p class="mt-4 max-w-2xl text-sm leading-7 text-slate-500 dark:text-slate-400 md:text-[15px]">
              {{ t('home.heroDescription') }}
            </p>

            <div class="mt-7 flex flex-col gap-3 sm:flex-row">
              <router-link
                :to="isAuthenticated ? dashboardPath : '/login'"
                class="inline-flex items-center justify-center rounded-full bg-teal-600 px-6 py-2.5 text-sm font-semibold text-white shadow-[0_18px_40px_rgba(13,148,136,0.28)] transition hover:bg-teal-500"
              >
                {{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}
                <Icon name="arrowRight" size="md" class="ml-2" :stroke-width="2" />
              </router-link>
              <router-link
                to="/install-guide"
                class="inline-flex items-center justify-center rounded-full border border-slate-200 bg-white/80 px-6 py-2.5 text-sm font-semibold text-slate-700 backdrop-blur transition hover:border-teal-200 hover:text-slate-950 dark:border-white/10 dark:bg-slate-900/70 dark:text-slate-200 dark:hover:border-teal-400/40 dark:hover:text-white"
              >
                {{ t('home.installGuide') }}
              </router-link>
              <a
                v-if="docUrl"
                :href="docUrl"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center justify-center rounded-full border border-slate-200 bg-white/80 px-6 py-2.5 text-sm font-semibold text-slate-700 backdrop-blur transition hover:border-teal-200 hover:text-slate-950 dark:border-white/10 dark:bg-slate-900/70 dark:text-slate-200 dark:hover:border-teal-400/40 dark:hover:text-white"
              >
                {{ t('home.heroSecondaryCta') }}
              </a>
            </div>

            <div class="mt-7 grid gap-3 sm:grid-cols-3">
              <div
                v-for="metric in heroMetrics"
                :key="metric.label"
                class="rounded-[24px] border border-white/70 bg-white/75 p-4 shadow-[0_20px_45px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/68"
              >
                <p class="text-[11px] font-semibold uppercase tracking-[0.22em] text-slate-500 dark:text-slate-400">
                  {{ metric.label }}
                </p>
                <p class="mt-2.5 text-[30px] font-semibold leading-none text-slate-950 dark:text-white">
                  {{ metric.value }}
                </p>
              </div>
            </div>
          </div>

          <div class="relative lg:-mr-3">
            <div
              class="absolute -right-5 top-8 h-32 w-32 rounded-full bg-cyan-300/25 blur-3xl dark:bg-cyan-400/20"
            ></div>
            <div
              class="relative mx-auto w-[94%] overflow-hidden rounded-[26px] bg-white/72 px-0 pb-0 pt-5 text-slate-900 shadow-[0_22px_52px_rgba(15,118,110,0.10)] backdrop-blur-2xl dark:bg-slate-950/66 dark:text-slate-100"
            >
              <div class="absolute inset-x-0 top-0 h-20 bg-[linear-gradient(180deg,rgba(45,212,191,0.16),rgba(45,212,191,0))] dark:bg-[linear-gradient(180deg,rgba(45,212,191,0.1),rgba(45,212,191,0))]"></div>
              <div class="relative flex items-center justify-between px-6">
                <div>
                  <p class="text-[10px] font-semibold uppercase tracking-[0.22em] text-teal-700 dark:text-teal-300">
                    {{ t('home.quickstart.label') }}
                  </p>
                </div>
                <div class="rounded-full border border-emerald-400/35 bg-emerald-50 px-2.5 py-1 text-[10px] font-semibold text-emerald-700 shadow-sm dark:bg-emerald-400/12 dark:text-emerald-200">
                  运行正常
                </div>
              </div>

              <div class="mt-3.5 overflow-hidden">
                <div class="relative w-full overflow-hidden rounded-t-[30px] bg-[radial-gradient(circle_at_top,rgba(59,130,246,0.18),transparent_38%),linear-gradient(180deg,#0f172a_0%,#111827_45%,#0b1120_100%)] px-6 pb-6 pt-5 shadow-[0_24px_60px_rgba(15,23,42,0.28)]">
                  <div class="mb-3.5 flex items-center gap-3 text-slate-400">
                    <span class="h-3 w-3 rounded-full bg-rose-400"></span>
                    <span class="h-3 w-3 rounded-full bg-amber-400"></span>
                    <span class="h-3 w-3 rounded-full bg-emerald-400"></span>
                    <span class="ml-3 text-[12px] font-medium tracking-[0.08em]">{{ t('home.quickstart.filename') }}</span>
                  </div>

                  <div class="space-y-1.5 font-mono text-[13px] leading-6.5 md:text-[14px]">
                    <div class="text-slate-300">
                      <span class="text-violet-400">from</span>
                      <span class="mx-2 text-white">openai</span>
                      <span class="text-violet-400">import</span>
                      <span class="ml-2 text-sky-300">OpenAI</span>
                    </div>

                    <div class="pt-2 text-slate-500">{{ t('home.quickstart.comment') }}</div>

                    <div class="text-sky-300">client <span class="text-white">=</span> OpenAI<span class="text-white">(</span></div>
                    <div class="pl-7 text-slate-100">
                      api_key<span class="text-white">=</span><span class="text-emerald-400">"sk-tokenmkt-..."</span><span class="text-white">,</span>
                    </div>
                    <div class="pl-7 text-slate-100">
                      base_url<span class="text-white">=</span><span class="text-emerald-400">"https://api.tokenmkt.cc/v1"</span>
                    </div>
                    <div class="text-white">)</div>

                    <div class="pt-3 text-sky-300">
                      response <span class="text-white">= client.chat.completions.</span><span class="text-amber-300">create</span><span class="text-white">(</span>
                    </div>
                    <div class="pl-7 text-slate-100">
                      model<span class="text-white">=</span><span class="text-emerald-400">"gpt-4o"</span><span class="text-white">,</span>
                    </div>
                    <div class="pl-7 text-slate-100">
                      messages<span class="text-white">=[{"{"}</span><span class="text-emerald-400">"role"</span><span class="text-white">: </span><span class="text-emerald-400">"user"</span><span class="text-white">, </span><span class="text-emerald-400">"content"</span><span class="text-white">: </span><span class="text-emerald-400">"你好！"</span><span class="text-white">{"}"}</span><span class="text-white">]</span>
                    </div>
                    <div class="text-white">)</div>
                  </div>

                  <div class="mt-4 rounded-2xl border border-white/10 bg-white/5 px-4 py-2.5 font-mono text-[12px] text-slate-300 md:text-[13px]">
                    <span class="text-slate-500">&gt;</span>
                    <span class="ml-2">{{ t('home.quickstart.response') }}</span>
                  </div>
                </div>
              </div>

            </div>
          </div>
        </section>

        <section>
          <div class="mx-auto max-w-3xl text-center">
            <p class="text-xs font-semibold uppercase tracking-[0.28em] text-teal-700 dark:text-teal-300">
              Built for production traffic
            </p>
            <h2 class="mt-4 text-3xl font-semibold text-slate-950 dark:text-white md:text-4xl">
              {{ t('home.trust.title') }}
            </h2>
            <p class="mt-4 text-base leading-7 text-slate-600 dark:text-slate-300">
              {{ t('home.trust.subtitle') }}
            </p>
          </div>

          <div class="mt-10 grid gap-5 lg:grid-cols-3">
            <article
              v-for="card in featureCards"
              :key="card.title"
              class="group rounded-[28px] border border-white/70 bg-white/80 p-7 shadow-[0_24px_60px_rgba(15,118,110,0.08)] backdrop-blur transition duration-300 hover:-translate-y-1 hover:shadow-[0_30px_70px_rgba(15,118,110,0.14)] dark:border-white/10 dark:bg-slate-900/72 dark:hover:shadow-[0_26px_70px_rgba(13,148,136,0.16)]"
            >
              <div
                class="inline-flex h-12 w-12 items-center justify-center rounded-2xl text-sm font-semibold text-white shadow-lg"
                :class="card.accent"
              >
                {{ card.tag }}
              </div>
              <h3 class="mt-6 text-xl font-semibold text-slate-950 dark:text-white">
                {{ card.title }}
              </h3>
              <p class="mt-4 text-sm leading-7 text-slate-600 dark:text-slate-300">
                {{ card.desc }}
              </p>
            </article>
          </div>
        </section>

        <section class="rounded-[32px] border border-white/70 bg-white/72 p-8 shadow-[0_24px_70px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72 lg:px-10">
          <div class="flex flex-col gap-8 lg:flex-row lg:items-center lg:justify-between lg:gap-10">
            <div class="mx-auto max-w-xl text-center lg:mx-0 lg:text-left">
              <h2 class="text-2xl font-semibold text-slate-950 dark:text-white md:text-3xl">
                {{ t('home.providers.title') }}
              </h2>
              <p class="mt-3 text-sm leading-7 text-slate-600 dark:text-slate-300">
                {{ t('home.providers.description') }}
              </p>
            </div>

            <div class="flex flex-wrap items-center justify-center gap-4 lg:max-w-[760px] lg:justify-end">
            <div
              v-for="provider in providers"
              :key="provider.name"
              class="flex min-w-[200px] items-center gap-3 rounded-2xl border px-5 py-3 backdrop-blur"
              :class="provider.available
                ? 'border-teal-200 bg-teal-50/70 text-slate-800 dark:border-teal-400/20 dark:bg-teal-400/10 dark:text-slate-100'
                : 'border-slate-200 bg-slate-100/70 text-slate-500 dark:border-white/10 dark:bg-white/5 dark:text-slate-400'"
            >
              <div
                class="flex h-10 w-10 items-center justify-center overflow-hidden rounded-2xl"
                :class="
                  provider.logo
                    ? 'bg-white ring-1 ring-slate-200/80 dark:bg-slate-900 dark:ring-white/10'
                    : provider.badgeClass
                "
              >
                <img
                  v-if="provider.logo"
                  :src="provider.logo"
                  :alt="provider.name"
                  class="h-full w-full object-cover"
                />
                <span v-else class="text-sm font-bold text-white">{{ provider.short }}</span>
              </div>
              <div>
                <p class="text-sm font-semibold">{{ provider.name }}</p>
                <p class="text-xs uppercase tracking-[0.18em]">
                  {{ provider.available ? t('home.providers.supported') : t('home.providers.soon') }}
                </p>
              </div>
            </div>
          </div>
          </div>
        </section>

        <section
          class="overflow-hidden rounded-[36px] border border-slate-900/5 bg-[radial-gradient(circle_at_top,rgba(59,130,246,0.18),transparent_38%),linear-gradient(180deg,#0f172a_0%,#111827_45%,#0b1120_100%)] px-8 py-10 text-center text-white shadow-[0_32px_90px_rgba(2,6,23,0.35)] dark:border-white/10"
        >
          <div class="mx-auto max-w-3xl">
            <p class="text-xs font-semibold uppercase tracking-[0.28em] text-teal-300">
              tokenMKT
            </p>
            <h2 class="mt-4 text-3xl font-semibold md:text-4xl">
              {{ t('home.cta.title') }}
            </h2>
            <p class="mt-4 text-base leading-7 text-slate-300">
              {{ t('home.cta.description') }}
            </p>
            <div class="mt-8 flex justify-center">
              <router-link
                :to="isAuthenticated ? dashboardPath : '/login'"
                class="inline-flex items-center rounded-full bg-white px-6 py-3 text-sm font-semibold text-slate-950 transition hover:bg-teal-50"
              >
                {{ t('home.cta.button') }}
              </router-link>
            </div>
          </div>
        </section>
      </div>
    </main>

    <footer class="relative z-10 px-6 py-8">
      <div
        class="mx-auto flex max-w-6xl flex-col items-center gap-4 border-t border-slate-200/70 pt-6 text-center dark:border-white/10"
      >
        <p class="text-sm text-slate-500 dark:text-slate-400">
          &copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}
        </p>
        <div class="flex items-center gap-4">
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-sm text-slate-500 transition hover:text-slate-900 dark:text-slate-400 dark:hover:text-white"
          >
            {{ t('home.docs') }}
          </a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()

const authStore = useAuthStore()
const appStore = useAppStore()

const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'tokenMKT')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')

const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const isDark = ref(document.documentElement.classList.contains('dark'))

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => isAdmin.value ? '/admin/dashboard' : '/dashboard')
const userInitial = computed(() => {
  const user = authStore.user
  if (!user || !user.email) return ''
  return user.email.charAt(0).toUpperCase()
})

const currentYear = computed(() => new Date().getFullYear())
const chatgptLogo = '/ChatGPT.png'
const claudeLogo = '/Claude.png'
const geminiLogo = '/Gemini.jpg'

const heroMetrics = computed(() => [
  { label: t('home.metrics.official.label'), value: t('home.metrics.official.value') },
  { label: t('home.metrics.stability.label'), value: t('home.metrics.stability.value') },
  { label: t('home.metrics.purity.label'), value: t('home.metrics.purity.value') }
])

const featureCards = computed(() => [
  {
    tag: 'API',
    accent: 'bg-gradient-to-br from-sky-500 to-cyan-500',
    title: t('home.features.official.title'),
    desc: t('home.features.official.desc')
  },
  {
    tag: 'SLA',
    accent: 'bg-gradient-to-br from-teal-500 to-emerald-500',
    title: t('home.features.stability.title'),
    desc: t('home.features.stability.desc')
  },
  {
    tag: 'RAW',
    accent: 'bg-gradient-to-br from-slate-700 to-slate-900',
    title: t('home.features.purity.title'),
    desc: t('home.features.purity.desc')
  }
])

const providers = computed(() => [
  {
    name: 'GPT',
    short: 'G',
    logo: chatgptLogo,
    available: true,
    badgeClass: 'bg-gradient-to-br from-emerald-500 to-green-600'
  },
  {
    name: t('home.providers.claude'),
    short: 'C',
    logo: claudeLogo,
    available: false,
    badgeClass: 'bg-gradient-to-br from-orange-400 to-orange-500'
  },
  {
    name: t('home.providers.gemini'),
    short: 'G',
    logo: geminiLogo,
    available: false,
    badgeClass: 'bg-gradient-to-br from-sky-500 to-blue-600'
  }
])

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (
    savedTheme === 'dark' ||
    (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)
  ) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

onMounted(() => {
  initTheme()
  authStore.checkAuth()

  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>

<style scoped>
.router-link-active {
  text-decoration: none;
}
</style>
