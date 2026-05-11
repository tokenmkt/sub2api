<template>
  <div
    class="relative flex min-h-screen flex-col overflow-hidden bg-[radial-gradient(circle_at_top,rgba(45,212,191,0.16),transparent_34%),linear-gradient(180deg,#f4fbfa_0%,#eff7f5_42%,#f8fafc_100%)] dark:bg-[radial-gradient(circle_at_top,rgba(20,184,166,0.24),transparent_28%),linear-gradient(180deg,#051816_0%,#091d1a_40%,#020617_100%)]"
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
        <router-link to="/home" class="flex items-center gap-3">
          <div class="h-11 w-11 overflow-hidden rounded-2xl bg-white/90 p-1 shadow-[0_16px_40px_rgba(15,118,110,0.18)] ring-1 ring-emerald-900/5 dark:bg-slate-900/80 dark:ring-white/10">
            <img :src="siteLogo || '/logo.png'" alt="Logo" class="h-full w-full object-contain" />
          </div>
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.28em] text-teal-700 dark:text-teal-300">
              {{ siteName }}
            </p>
            <p class="text-sm text-slate-500 dark:text-slate-300">
              tokenMKT install guide
            </p>
          </div>
        </router-link>

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
            to="/home"
            class="inline-flex items-center rounded-full bg-slate-950 px-4 py-2 text-sm font-medium text-white shadow-[0_16px_30px_rgba(15,23,42,0.2)] transition hover:bg-slate-800 dark:bg-white dark:text-slate-900 dark:hover:bg-slate-100"
          >
            {{ t('installGuide.backHome') }}
          </router-link>
        </div>
      </nav>
    </header>

    <main class="relative z-10 flex-1 px-6 pb-12 pt-6">
      <div class="mx-auto max-w-6xl">
        <section class="text-center">
          <div
            class="inline-flex items-center gap-2 rounded-full border border-teal-200/80 bg-white/80 px-4 py-1.5 text-[11px] font-semibold uppercase tracking-[0.2em] text-teal-700 shadow-sm backdrop-blur dark:border-teal-400/20 dark:bg-slate-900/70 dark:text-teal-300"
          >
            <span class="h-2 w-2 rounded-full bg-teal-500"></span>
            {{ t('installGuide.badge') }}
          </div>
          <h1 class="mx-auto mt-5 max-w-4xl text-[34px] font-semibold leading-[1.05] text-slate-950 dark:text-white md:text-[44px] lg:text-[56px]">
            {{ t('installGuide.title') }}
          </h1>
          <p class="mx-auto mt-4 max-w-3xl text-[15px] leading-7 text-slate-600 dark:text-slate-300 md:text-[17px]">
            {{ t('installGuide.subtitle') }}
          </p>
        </section>

        <section class="mt-10 rounded-[32px] border border-white/70 bg-white/72 p-4 shadow-[0_24px_70px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72 sm:p-6">
          <div class="flex flex-wrap items-center justify-center gap-3">
            <button
              v-for="tool in tools"
              :key="tool.id"
              @click="activeTool = tool.id"
              class="inline-flex items-center gap-3 rounded-2xl border px-4 py-3 text-sm font-semibold transition"
              :class="activeTool === tool.id
                ? 'border-teal-200 bg-teal-50/80 text-slate-900 shadow-sm dark:border-teal-400/20 dark:bg-teal-400/10 dark:text-white'
                : 'border-slate-200 bg-white/80 text-slate-500 hover:border-teal-200 hover:text-slate-900 dark:border-white/10 dark:bg-slate-950/40 dark:text-slate-300 dark:hover:border-teal-400/30 dark:hover:text-white'"
            >
              <img :src="tool.logo" :alt="tool.name" class="h-6 w-6 rounded-md object-contain bg-white" />
              <span>{{ tool.name }}</span>
            </button>
          </div>
        </section>

        <section class="mt-8 space-y-5">
          <article class="rounded-[28px] border border-white/70 bg-white/78 p-6 shadow-[0_22px_50px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72">
            <p class="text-xs font-semibold uppercase tracking-[0.24em] text-teal-700 dark:text-teal-300">
              {{ currentGuide.summaryTitle }}
            </p>
            <h2 class="mt-3 text-2xl font-semibold text-slate-950 dark:text-white">
              {{ t('installGuide.overviewTitle') }}
            </h2>
            <p class="mt-4 text-sm leading-7 text-slate-600 dark:text-slate-300">
              {{ currentGuide.summary }}
            </p>
            <div class="mt-5 grid gap-3 md:grid-cols-3">
              <div
                v-for="point in currentGuide.highlights"
                :key="point.label"
                class="rounded-2xl border border-teal-100 bg-teal-50/70 p-4 dark:border-teal-400/15 dark:bg-teal-400/8"
              >
                <p class="text-[11px] font-semibold uppercase tracking-[0.2em] text-teal-700 dark:text-teal-300">
                  {{ point.label }}
                </p>
                <p class="mt-2 text-sm leading-6 text-slate-700 dark:text-slate-200">
                  {{ point.value }}
                </p>
              </div>
            </div>
          </article>

          <article class="rounded-[28px] border border-white/70 bg-white/78 p-6 shadow-[0_22px_50px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72">
            <div class="flex items-start gap-4">
              <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-2xl bg-teal-500 text-base font-semibold text-white shadow-sm">
                1
              </div>
              <div class="min-w-0 flex-1">
                <p class="text-xs font-semibold uppercase tracking-[0.24em] text-teal-700 dark:text-teal-300">
                  {{ t('installGuide.stepPrefix') }} 01
                </p>
                <h2 class="mt-2 text-2xl font-semibold text-slate-950 dark:text-white">
                  {{ t('installGuide.checklistTitle') }}
                </h2>
                <ul class="mt-5 space-y-3">
                  <li
                    v-for="item in currentGuide.checklist"
                    :key="item"
                    class="flex items-start gap-3 rounded-2xl border border-slate-200/80 bg-slate-50/70 px-4 py-3 dark:border-white/10 dark:bg-slate-950/35"
                  >
                    <span class="mt-1 inline-flex h-5 w-5 flex-shrink-0 items-center justify-center rounded-full bg-teal-500/12 text-teal-600 dark:text-teal-300">
                      <Icon name="check" size="sm" />
                    </span>
                    <span class="text-sm leading-6 text-slate-700 dark:text-slate-200">{{ item }}</span>
                  </li>
                </ul>
              </div>
            </div>
          </article>

          <article
            v-for="(block, index) in currentGuide.primaryBlocks"
            :key="block.title"
            class="overflow-hidden rounded-[28px] border border-white/70 bg-white/78 shadow-[0_22px_50px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72"
          >
            <div class="border-b border-slate-200/70 px-6 py-5 dark:border-white/10">
              <div class="flex items-start gap-4">
                <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-2xl bg-teal-500 text-base font-semibold text-white shadow-sm">
                  {{ index + 2 }}
                </div>
                <div class="min-w-0 flex-1">
                  <div class="flex items-start justify-between gap-4">
                    <div>
                      <p class="text-xs font-semibold uppercase tracking-[0.24em] text-teal-700 dark:text-teal-300">
                        {{ t('installGuide.stepPrefix') }} {{ String(index + 2).padStart(2, '0') }}
                      </p>
                      <h3 class="mt-2 text-xl font-semibold text-slate-950 dark:text-white">
                        {{ block.title }}
                      </h3>
                      <p v-if="block.description" class="mt-2 text-sm leading-6 text-slate-500 dark:text-slate-400">
                        {{ block.description }}
                      </p>
                    </div>
                    <button
                      @click="copyToClipboard(block.code, t('installGuide.copySuccess'))"
                      class="inline-flex items-center gap-2 rounded-full border border-slate-200 bg-white/80 px-3 py-2 text-xs font-semibold text-slate-700 transition hover:border-teal-200 hover:text-slate-950 dark:border-white/10 dark:bg-slate-950/50 dark:text-slate-200 dark:hover:border-teal-400/30 dark:hover:text-white"
                    >
                      <Icon name="copy" size="sm" />
                      {{ t('installGuide.copy') }}
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div class="bg-[radial-gradient(circle_at_top,rgba(59,130,246,0.18),transparent_38%),linear-gradient(180deg,#0f172a_0%,#111827_45%,#0b1120_100%)] px-6 pb-6 pt-5">
              <div class="mb-4 flex items-center gap-3 text-slate-400">
                <span class="h-3 w-3 rounded-full bg-rose-400"></span>
                <span class="h-3 w-3 rounded-full bg-amber-400"></span>
                <span class="h-3 w-3 rounded-full bg-emerald-400"></span>
                <span class="ml-3 text-[12px] font-medium tracking-[0.08em] text-slate-300">{{ block.filename }}</span>
              </div>
              <pre class="overflow-x-auto whitespace-pre-wrap break-words font-mono text-[13px] leading-6 text-slate-100 md:text-[14px]"><code>{{ block.code }}</code></pre>
            </div>
          </article>

          <article
            v-if="currentGuide.optionalBlocks.length"
            class="rounded-[28px] border border-white/70 bg-white/78 p-6 shadow-[0_22px_50px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72"
          >
            <p class="text-xs font-semibold uppercase tracking-[0.24em] text-teal-700 dark:text-teal-300">
              {{ t('installGuide.optionalEyebrow') }}
            </p>
            <h2 class="mt-2 text-2xl font-semibold text-slate-950 dark:text-white">
              {{ t('installGuide.optionalTitle') }}
            </h2>
            <div class="mt-5 space-y-5">
              <div
                v-for="block in currentGuide.optionalBlocks"
                :key="block.title"
                class="overflow-hidden rounded-[24px] border border-slate-200/80 bg-slate-50/70 dark:border-white/10 dark:bg-slate-950/35"
              >
                <div class="border-b border-slate-200/70 px-5 py-4 dark:border-white/10">
                  <div class="flex items-start justify-between gap-4">
                    <div>
                      <h3 class="text-lg font-semibold text-slate-950 dark:text-white">
                        {{ block.title }}
                      </h3>
                      <p v-if="block.description" class="mt-2 text-sm leading-6 text-slate-500 dark:text-slate-400">
                        {{ block.description }}
                      </p>
                    </div>
                    <button
                      @click="copyToClipboard(block.code, t('installGuide.copySuccess'))"
                      class="inline-flex items-center gap-2 rounded-full border border-slate-200 bg-white/80 px-3 py-2 text-xs font-semibold text-slate-700 transition hover:border-teal-200 hover:text-slate-950 dark:border-white/10 dark:bg-slate-950/50 dark:text-slate-200 dark:hover:border-teal-400/30 dark:hover:text-white"
                    >
                      <Icon name="copy" size="sm" />
                      {{ t('installGuide.copy') }}
                    </button>
                  </div>
                </div>
                <div class="bg-[radial-gradient(circle_at_top,rgba(59,130,246,0.18),transparent_38%),linear-gradient(180deg,#0f172a_0%,#111827_45%,#0b1120_100%)] px-5 pb-5 pt-4">
                  <div class="mb-4 flex items-center gap-3 text-slate-400">
                    <span class="h-3 w-3 rounded-full bg-rose-400"></span>
                    <span class="h-3 w-3 rounded-full bg-amber-400"></span>
                    <span class="h-3 w-3 rounded-full bg-emerald-400"></span>
                    <span class="ml-3 text-[12px] font-medium tracking-[0.08em] text-slate-300">{{ block.filename }}</span>
                  </div>
                  <pre class="overflow-x-auto whitespace-pre-wrap break-words font-mono text-[13px] leading-6 text-slate-100 md:text-[14px]"><code>{{ block.code }}</code></pre>
                </div>
              </div>
            </div>
          </article>
        </section>

        <section class="mt-8 rounded-[32px] border border-white/70 bg-white/72 p-6 shadow-[0_24px_70px_rgba(15,118,110,0.08)] backdrop-blur dark:border-white/10 dark:bg-slate-900/72">
          <div class="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
            <div>
              <p class="text-xs font-semibold uppercase tracking-[0.24em] text-teal-700 dark:text-teal-300">
                {{ t('installGuide.tipsEyebrow') }}
              </p>
              <h2 class="mt-2 text-2xl font-semibold text-slate-950 dark:text-white">
                {{ t('installGuide.tipsTitle') }}
              </h2>
            </div>
            <div class="grid gap-3 md:grid-cols-3 lg:max-w-3xl">
              <div
                v-for="tip in tips"
                :key="tip"
                class="rounded-2xl border border-slate-200/80 bg-slate-50/80 px-4 py-3 text-sm leading-6 text-slate-600 dark:border-white/10 dark:bg-slate-950/35 dark:text-slate-300"
              >
                {{ tip }}
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>

    <footer class="relative z-10 px-6 py-8">
      <div class="mx-auto flex max-w-6xl flex-col items-center gap-4 border-t border-slate-200/70 pt-6 text-center dark:border-white/10">
        <p class="text-sm text-slate-500 dark:text-slate-400">
          &copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import { useClipboard } from '@/composables/useClipboard'

const { t } = useI18n()
const appStore = useAppStore()
const { copyToClipboard } = useClipboard()

const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'tokenMKT')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const isDark = ref(document.documentElement.classList.contains('dark'))
const currentYear = computed(() => new Date().getFullYear())
const chatgptLogo = '/ChatGPT.png'
const claudeLogo = '/Claude.png'
const geminiLogo = '/Gemini.jpg'

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

const tools = [
  { id: 'codex', name: 'Codex CLI', logo: chatgptLogo },
  { id: 'claude', name: 'Claude Code', logo: claudeLogo },
  { id: 'gemini', name: 'Gemini CLI', logo: geminiLogo }
] as const

const activeTool = ref<'codex' | 'claude' | 'gemini'>('codex')

const guideMap = {
  codex: {
    summaryTitle: t('installGuide.codex.summaryTitle'),
    summary: t('installGuide.codex.summary'),
    highlights: [
      { label: t('installGuide.labels.install'), value: t('installGuide.codex.installValue') },
      { label: t('installGuide.labels.config'), value: t('installGuide.codex.configValue') },
      { label: t('installGuide.labels.verify'), value: t('installGuide.codex.verifyValue') }
    ],
    checklist: [
      t('installGuide.codex.check1'),
      t('installGuide.codex.check2'),
      t('installGuide.codex.check3')
    ],
    primaryBlocks: [
      {
        eyebrow: t('installGuide.blockEyebrows.install'),
        title: t('installGuide.codex.installBlockTitle'),
        description: t('installGuide.codex.installBlockDesc'),
        filename: 'terminal',
        code: 'npm install -g @openai/codex'
      },
      {
        eyebrow: t('installGuide.blockEyebrows.config'),
        title: t('installGuide.codex.configBlockTitle'),
        description: t('installGuide.codex.configBlockDesc'),
        filename: '~/.codex/config.toml',
        code: `model_provider = "OpenAI"
model = "gpt-5.4"
review_model = "gpt-5.4"
model_reasoning_effort = "xhigh"
disable_response_storage = true
network_access = "enabled"
windows_wsl_setup_acknowledged = true
model_context_window = 1000000
model_auto_compact_token_limit = 900000

[model_providers.OpenAI]
name = "OpenAI"
base_url = "https://api.tokenmkt.cc/v1"
wire_api = "responses"
requires_openai_auth = true`
      },
      {
        eyebrow: t('installGuide.blockEyebrows.auth'),
        title: t('installGuide.codex.authBlockTitle'),
        description: t('installGuide.codex.authBlockDesc'),
        filename: '~/.codex/auth.json',
        code: `{
  "OPENAI_API_KEY": "sk-tokenmkt-..."
}`
      },
      {
        eyebrow: t('installGuide.blockEyebrows.verify'),
        title: t('installGuide.codex.verifyBlockTitle'),
        description: t('installGuide.codex.verifyBlockDesc'),
        filename: 'terminal',
        code: 'codex'
      }
    ],
    optionalBlocks: [] as Array<{ eyebrow: string; title: string; description: string; filename: string; code: string }>
  },
  claude: {
    summaryTitle: t('installGuide.claude.summaryTitle'),
    summary: t('installGuide.claude.summary'),
    highlights: [
      { label: t('installGuide.labels.install'), value: t('installGuide.claude.installValue') },
      { label: t('installGuide.labels.config'), value: t('installGuide.claude.configValue') },
      { label: t('installGuide.labels.verify'), value: t('installGuide.claude.verifyValue') }
    ],
    checklist: [
      t('installGuide.claude.check1'),
      t('installGuide.claude.check2'),
      t('installGuide.claude.check3')
    ],
    primaryBlocks: [
      {
        eyebrow: t('installGuide.blockEyebrows.install'),
        title: t('installGuide.claude.installBlockTitle'),
        description: t('installGuide.claude.installBlockDesc'),
        filename: 'terminal',
        code: 'npm install -g @anthropic-ai/claude-code'
      },
      {
        eyebrow: t('installGuide.blockEyebrows.env'),
        title: t('installGuide.claude.envBlockTitle'),
        description: t('installGuide.claude.envBlockDesc'),
        filename: 'terminal',
        code: `export ANTHROPIC_BASE_URL="https://api.tokenmkt.cc/anthropic"
export ANTHROPIC_AUTH_TOKEN="sk-tokenmkt-..."
export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1`
      },
      {
        eyebrow: t('installGuide.blockEyebrows.verify'),
        title: t('installGuide.claude.verifyBlockTitle'),
        description: t('installGuide.claude.verifyBlockDesc'),
        filename: 'terminal',
        code: 'claude'
      }
    ],
    optionalBlocks: [
      {
        eyebrow: t('installGuide.blockEyebrows.ide'),
        title: t('installGuide.claude.ideBlockTitle'),
        description: t('installGuide.claude.ideBlockDesc'),
        filename: '~/.claude/settings.json',
        code: `{
  "env": {
    "ANTHROPIC_BASE_URL": "https://api.tokenmkt.cc/anthropic",
    "ANTHROPIC_AUTH_TOKEN": "sk-tokenmkt-...",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
    "CLAUDE_CODE_ATTRIBUTION_HEADER": "0"
  }
}`
      }
    ]
  },
  gemini: {
    summaryTitle: t('installGuide.gemini.summaryTitle'),
    summary: t('installGuide.gemini.summary'),
    highlights: [
      { label: t('installGuide.labels.install'), value: t('installGuide.gemini.installValue') },
      { label: t('installGuide.labels.config'), value: t('installGuide.gemini.configValue') },
      { label: t('installGuide.labels.verify'), value: t('installGuide.gemini.verifyValue') }
    ],
    checklist: [
      t('installGuide.gemini.check1'),
      t('installGuide.gemini.check2'),
      t('installGuide.gemini.check3')
    ],
    primaryBlocks: [
      {
        eyebrow: t('installGuide.blockEyebrows.install'),
        title: t('installGuide.gemini.installBlockTitle'),
        description: t('installGuide.gemini.installBlockDesc'),
        filename: 'terminal',
        code: 'npm install -g @google/gemini-cli'
      },
      {
        eyebrow: t('installGuide.blockEyebrows.env'),
        title: t('installGuide.gemini.envBlockTitle'),
        description: t('installGuide.gemini.envBlockDesc'),
        filename: 'terminal',
        code: `export GOOGLE_GEMINI_BASE_URL="https://api.tokenmkt.cc/gemini"
export GEMINI_API_KEY="sk-tokenmkt-..."
export GEMINI_MODEL="gemini-2.0-flash"`
      },
      {
        eyebrow: t('installGuide.blockEyebrows.verify'),
        title: t('installGuide.gemini.verifyBlockTitle'),
        description: t('installGuide.gemini.verifyBlockDesc'),
        filename: 'terminal',
        code: 'gemini'
      }
    ],
    optionalBlocks: [] as Array<{ eyebrow: string; title: string; description: string; filename: string; code: string }>
  }
} as const

const currentGuide = computed(() => guideMap[activeTool.value])

const tips = computed(() => [
  t('installGuide.tip1'),
  t('installGuide.tip2'),
  t('installGuide.tip3')
])
</script>
