<template>
  <div class="install-page">
    <nav class="nex-nav">
      <router-link to="/home" class="logo">
        <span class="logo-dot"></span>
        {{ siteName }}
      </router-link>
      <ul class="nav-links">
        <li><router-link to="/home">首页</router-link></li>
        <li><a href="#tools">工具配置</a></li>
        <li><a href="#steps">安装步骤</a></li>
        <li><router-link to="/home" class="nav-cta">返回首页</router-link></li>
      </ul>
    </nav>

    <main>
      <section class="hero compact-hero">
        <div class="badge">
          <span class="badge-dot"></span>
          {{ t('installGuide.badge') }}
        </div>
        <h1 class="hero-title">{{ t('installGuide.title') }}</h1>
        <p class="hero-sub">{{ t('installGuide.subtitle') }}</p>
      </section>

      <section id="tools" class="tool-switch-wrap">
        <div class="tool-switch">
          <button
            v-for="tool in tools"
            :key="tool.id"
            @click="activeTool = tool.id"
            class="tool-button"
            :class="{ active: activeTool === tool.id }"
          >
            <span class="tool-mark">
              <img :src="tool.logo" :alt="tool.name" />
            </span>
            <span>{{ tool.name }}</span>
          </button>
        </div>
      </section>

      <section id="steps" class="guide-layout">
        <article class="info-card overview-card">
          <p class="section-label">// {{ currentGuide.summaryTitle }}</p>
          <h2 class="section-title">{{ t('installGuide.overviewTitle') }}</h2>
          <p class="section-desc">{{ currentGuide.summary }}</p>
          <div class="highlight-grid">
            <div v-for="point in currentGuide.highlights" :key="point.label" class="highlight-card">
              <p class="highlight-label">{{ point.label }}</p>
              <p class="highlight-value">{{ point.value }}</p>
            </div>
          </div>
        </article>

        <article class="info-card">
          <div class="step-head">
            <div class="step-num">01</div>
            <div>
              <p class="section-label">// {{ t('installGuide.stepPrefix') }} 01</p>
              <h2 class="card-title">{{ t('installGuide.checklistTitle') }}</h2>
            </div>
          </div>
          <ul class="check-list">
            <li v-for="item in currentGuide.checklist" :key="item">
              <span class="check-dot">✓</span>
              <span>{{ item }}</span>
            </li>
          </ul>
        </article>

        <article v-for="(block, index) in currentGuide.primaryBlocks" :key="block.title" class="code-card">
          <div class="code-card-head">
            <div class="step-head compact">
              <div class="step-num">{{ String(index + 2).padStart(2, '0') }}</div>
              <div>
                <p class="section-label">// {{ t('installGuide.stepPrefix') }} {{ String(index + 2).padStart(2, '0') }}</p>
                <h3 class="card-title">{{ block.title }}</h3>
                <p v-if="block.description" class="card-desc">{{ block.description }}</p>
              </div>
            </div>
            <button @click="copyToClipboard(block.code, t('installGuide.copySuccess'))" class="copy-btn">
              {{ t('installGuide.copy') }}
            </button>
          </div>
          <div class="code-block">
            <div class="code-header">
              <div class="code-header-left">
                <span class="tbar-dot t-red"></span>
                <span class="tbar-dot t-yellow"></span>
                <span class="tbar-dot t-green"></span>
              </div>
              <span class="code-lang">{{ block.filename }}</span>
            </div>
            <pre class="code-body"><code>{{ block.code }}</code></pre>
          </div>
        </article>

        <article v-if="currentGuide.optionalBlocks.length" class="info-card optional-card">
          <p class="section-label">// {{ t('installGuide.optionalEyebrow') }}</p>
          <h2 class="card-title">{{ t('installGuide.optionalTitle') }}</h2>
          <div class="optional-list">
            <div v-for="block in currentGuide.optionalBlocks" :key="block.title" class="code-card nested">
              <div class="code-card-head">
                <div>
                  <h3 class="nested-title">{{ block.title }}</h3>
                  <p v-if="block.description" class="card-desc">{{ block.description }}</p>
                </div>
                <button @click="copyToClipboard(block.code, t('installGuide.copySuccess'))" class="copy-btn">
                  {{ t('installGuide.copy') }}
                </button>
              </div>
              <div class="code-block">
                <div class="code-header">
                  <div class="code-header-left">
                    <span class="tbar-dot t-red"></span>
                    <span class="tbar-dot t-yellow"></span>
                    <span class="tbar-dot t-green"></span>
                  </div>
                  <span class="code-lang">{{ block.filename }}</span>
                </div>
                <pre class="code-body"><code>{{ block.code }}</code></pre>
              </div>
            </div>
          </div>
        </article>

        <article class="info-card tips-card">
          <div>
            <p class="section-label">// {{ t('installGuide.tipsEyebrow') }}</p>
            <h2 class="card-title">{{ t('installGuide.tipsTitle') }}</h2>
          </div>
          <div class="tips-grid">
            <div v-for="tip in tips" :key="tip" class="tip-item">{{ tip }}</div>
          </div>
        </article>
      </section>
    </main>

    <footer>
      <div class="footer-inner">
        <div class="footer-copy">© {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}</div>
        <ul class="footer-links">
          <li><router-link to="/home">首页</router-link></li>
          <li><a href="#tools">工具配置</a></li>
        </ul>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores'
import { useClipboard } from '@/composables/useClipboard'

const { t } = useI18n()
const appStore = useAppStore()
const { copyToClipboard } = useClipboard()

const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'apia8')
const currentYear = computed(() => new Date().getFullYear())

const tools = [
  { id: 'codex', name: 'Codex CLI', logo: '/ChatGPT.png' },
  { id: 'claude', name: 'Claude Code', logo: '/Claude.png' },
  { id: 'gemini', name: 'Gemini CLI', logo: '/Gemini.jpg' }
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
base_url = "https://api.apia8.com/v1"
wire_api = "responses"
requires_openai_auth = true`
      },
      {
        eyebrow: t('installGuide.blockEyebrows.auth'),
        title: t('installGuide.codex.authBlockTitle'),
        description: t('installGuide.codex.authBlockDesc'),
        filename: '~/.codex/auth.json',
        code: `{
  "OPENAI_API_KEY": "sk-apia8-..."
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
        code: `export ANTHROPIC_BASE_URL="https://api.apia8.com/anthropic"
export ANTHROPIC_AUTH_TOKEN="sk-apia8-..."
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
    "ANTHROPIC_BASE_URL": "https://api.apia8.com/anthropic",
    "ANTHROPIC_AUTH_TOKEN": "sk-apia8-...",
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
        code: `export GOOGLE_GEMINI_BASE_URL="https://api.apia8.com/gemini"
export GEMINI_API_KEY="sk-apia8-..."
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


<style scoped>
.install-page {
  --bg: #eaf4f8;
  --bg2: rgba(255, 255, 255, 0.78);
  --bg3: rgba(238, 248, 252, 0.9);
  --accent: #008eb2;
  --accent2: #005f86;
  --gold: #b47b1f;
  --text: #102235;
  --text-muted: #526c82;
  --border: rgba(0, 113, 145, 0.14);
  --border2: rgba(0, 142, 178, 0.28);
  min-height: 100vh;
  background:
    radial-gradient(circle at top, rgba(0, 180, 216, 0.18), transparent 34%),
    linear-gradient(180deg, #f7fbfd 0%, var(--bg) 52%, #f8fbfc 100%);
  color: var(--text);
  font-family: 'DM Sans', ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  overflow-x: hidden;
  position: relative;
}

.install-page::before {
  content: '';
  position: fixed;
  inset: 0;
  background-image:
    linear-gradient(rgba(0,113,145,0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0,113,145,0.055) 1px, transparent 1px);
  background-size: 48px 48px;
  pointer-events: none;
  z-index: 0;
}

main, footer, .nex-nav { position: relative; z-index: 1; }

.nex-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 5%;
  height: 64px;
  background: rgba(247, 251, 253, 0.86);
  backdrop-filter: blur(18px);
  border-bottom: 1px solid var(--border);
}

.logo {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 800;
  font-size: 1.35rem;
  color: var(--text);
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
}

.logo-dot, .badge-dot {
  border-radius: 50%;
  background: var(--accent);
  box-shadow: 0 0 10px var(--accent);
  animation: pulse 2s ease-in-out infinite;
}

.logo-dot { width: 8px; height: 8px; }
.badge-dot { width: 6px; height: 6px; }

.nav-links {
  display: flex;
  align-items: center;
  gap: 2rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

.nav-links a {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.875rem;
  transition: color 0.2s;
}

.nav-links a:hover { color: var(--accent2); }

.nav-cta {
  border: 1px solid var(--border2);
  color: var(--accent) !important;
  padding: 6px 18px;
  border-radius: 6px;
}

.hero {
  min-height: 54vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 118px 5% 54px;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: rgba(0,180,216,0.08);
  border: 1px solid var(--border2);
  border-radius: 99px;
  padding: 5px 14px;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.75rem;
  color: var(--accent);
  margin-bottom: 2rem;
}

.hero-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 800;
  font-size: clamp(2.4rem, 6vw, 4.8rem);
  line-height: 1.05;
  letter-spacing: -0.03em;
  color: var(--text);
  margin: 0 0 1.5rem;
}

.hero-sub {
  font-size: 1.05rem;
  color: var(--text-muted);
  max-width: 680px;
  margin: 0;
  font-weight: 300;
  line-height: 1.8;
}

.tool-switch-wrap, .guide-layout {
  max-width: 1100px;
  margin: 0 auto;
  padding: 0 5%;
}

.tool-switch {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 12px;
  padding: 20px;
  background: rgba(255,255,255,0.72);
  border: 1px solid var(--border);
  border-radius: 12px;
  box-shadow: 0 22px 60px rgba(8, 47, 73, 0.08);
}

.tool-button {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  border: 1px solid var(--border);
  background: var(--bg3);
  color: var(--text-muted);
  border-radius: 10px;
  padding: 12px 18px;
  font-weight: 600;
  transition: border-color 0.2s, background 0.2s, color 0.2s;
}

.tool-button.active, .tool-button:hover {
  border-color: var(--border2);
  background: rgba(0,142,178,0.08);
  color: var(--text);
}

.tool-mark {
  display: inline-flex;
  width: 28px;
  height: 28px;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #fff;
  overflow: hidden;
  box-shadow: 0 0 0 1px rgba(255,255,255,0.12);
}

.tool-mark img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.guide-layout {
  padding-top: 24px;
  padding-bottom: 80px;
}

.info-card, .code-card {
  background: var(--bg2);
  border: 1px solid var(--border);
  border-radius: 14px;
  margin-top: 18px;
  overflow: hidden;
  box-shadow: 0 22px 60px rgba(8, 47, 73, 0.08);
}

.info-card {
  padding: 30px;
}

.info-card:hover, .code-card:hover { border-color: var(--border2); }

.section-label, .highlight-label {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.72rem;
  color: var(--accent);
  letter-spacing: 0.15em;
  text-transform: uppercase;
  margin-bottom: 1rem;
}

.section-title, .card-title, .nested-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 700;
  color: var(--text);
  letter-spacing: -0.02em;
  margin: 0;
}

.section-title { font-size: clamp(1.8rem, 4vw, 2.8rem); }
.card-title { font-size: 1.35rem; }
.nested-title { font-size: 1.05rem; }

.section-desc, .card-desc, .highlight-value, .check-list, .tip-item {
  color: var(--text-muted);
  font-size: 0.92rem;
  line-height: 1.75;
  font-weight: 300;
}

.highlight-grid, .tips-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-top: 24px;
}

.highlight-card, .tip-item, .check-list li {
  background: var(--bg3);
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 16px;
}

.step-head {
  display: flex;
  align-items: flex-start;
  gap: 18px;
}

.step-head.compact { flex: 1; }

.step-num {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 1px solid var(--border2);
  background: rgba(0,180,216,0.05);
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  color: var(--accent);
  flex-shrink: 0;
}

.check-list {
  list-style: none;
  padding: 0;
  margin: 24px 0 0;
  display: grid;
  gap: 10px;
}

.check-list li {
  display: flex;
  gap: 12px;
}

.check-dot { color: var(--accent); }

.code-card-head {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  padding: 24px;
  border-bottom: 1px solid var(--border);
}

.copy-btn {
  align-self: flex-start;
  border: 1px solid var(--border2);
  background: transparent;
  color: var(--accent);
  border-radius: 8px;
  padding: 8px 14px;
  font-size: 0.78rem;
  font-weight: 700;
  cursor: pointer;
}

.copy-btn:hover { background: rgba(0,180,216,0.08); }

.code-block {
  background: #092033;
}

.code-header {
  background: #102b42;
  padding: 10px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--border);
}

.code-header-left { display: flex; gap: 7px; }

.tbar-dot { width: 11px; height: 11px; border-radius: 50%; }
.t-red { background: #ff5f57; }
.t-yellow { background: #febc2e; }
.t-green { background: #28c840; }

.code-lang {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.72rem;
  color: var(--text-muted);
}

.code-body {
  margin: 0;
  padding: 22px 24px;
  overflow-x: auto;
  white-space: pre-wrap;
  word-break: break-word;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.82rem;
  line-height: 1.85;
  color: #d8edf7;
}

.optional-list { margin-top: 18px; }
.code-card.nested { margin-top: 14px; background: var(--bg3); }

.tips-card {
  display: grid;
  grid-template-columns: 0.8fr 1.2fr;
  gap: 28px;
  align-items: start;
}

footer {
  padding: 48px 5% 32px;
  border-top: 1px solid var(--border);
}

.footer-inner {
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 1rem;
}

.footer-copy { font-size: 0.8rem; color: var(--text-muted); }
.footer-links { display: flex; gap: 2rem; list-style: none; margin: 0; padding: 0; }
.footer-links a { font-size: 0.8rem; color: var(--text-muted); text-decoration: none; }
.footer-links a:hover { color: var(--text); }

@keyframes pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.6; transform: scale(0.85); }
}

@media (max-width: 900px) {
  .nav-links { display: none; }
  .hero { min-height: 46vh; }
  .highlight-grid, .tips-grid, .tips-card { grid-template-columns: 1fr; }
  .code-card-head { flex-direction: column; }
}
</style>
