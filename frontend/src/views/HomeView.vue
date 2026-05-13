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

  <div v-else class="nex-home">
    <nav class="nex-nav">
      <router-link to="/home" class="logo">
        <span class="logo-dot"></span>
        {{ siteName }}
      </router-link>
      <ul class="nav-links">
        <li><a href="#features">功能特性</a></li>
        <li><a href="#models">支持模型</a></li>
        <li><a href="#integration">快速接入</a></li>
        <li><a :href="docUrl" target="_blank" rel="noopener noreferrer">查看文档</a></li>
        <li>
          <router-link :to="isAuthenticated ? dashboardPath : '/login'" class="nav-cta">
            {{ isAuthenticated ? '进入控制台' : '开始使用' }}
          </router-link>
        </li>
      </ul>
    </nav>

    <section class="hero">
      <div class="badge">
        <span class="badge-dot"></span>
        官方渠道 · 全球节点 · 企业级稳定
      </div>

      <h1 class="hero-title">
        企业级<em>官方 API</em><br />统一接入网关
      </h1>

      <p class="hero-sub">
        一个端点，接入 OpenAI、Claude、Gemini 等全球顶级大模型。99.99% SLA 保障，毫秒级响应，兼容原生 SDK。
      </p>

      <div class="hero-actions">
        <router-link :to="isAuthenticated ? dashboardPath : '/login'" class="btn-primary">
          开始使用 →
        </router-link>
        <a :href="docUrl" target="_blank" rel="noopener noreferrer" class="btn-ghost">查看文档</a>
      </div>

      <div class="terminal">
        <div class="terminal-bar">
          <span class="tbar-dot t-red"></span>
          <span class="tbar-dot t-yellow"></span>
          <span class="tbar-dot t-green"></span>
          <span class="tbar-title">~/project — apia8 request</span>
        </div>
        <div class="terminal-body">
          <div><span class="t-comment"># 将 base_url 替换为 apia8 端点，其余代码无需改动</span></div>
          <div>&nbsp;</div>
          <div>
            <span class="t-kw">from</span>
            <span class="t-punct"> openai </span>
            <span class="t-kw">import</span>
            <span> OpenAI</span>
          </div>
          <div>&nbsp;</div>
          <div><span>client </span><span class="t-punct">= </span><span>OpenAI(</span></div>
          <div><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="t-key">api_key</span><span class="t-punct"> = </span><span class="t-str">"sk-apia8-xxxxxxxx"</span><span class="t-punct">,</span></div>
          <div><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="t-key">base_url</span><span class="t-punct"> = </span><span class="t-str">"https://api.apia8.com/v1"</span></div>
          <div><span class="t-punct">)</span></div>
          <div>&nbsp;</div>
          <div><span>response </span><span class="t-punct">= </span><span>client.chat.completions.create(</span></div>
          <div><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="t-key">model</span><span class="t-punct"> = </span><span class="t-str">"gpt-5.5"</span><span class="t-punct">,</span></div>
          <div><span>&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="t-key">messages</span><span class="t-punct"> = [</span><span class="t-str">{"role": "user", "content": "Hello!"}</span><span class="t-punct">]</span></div>
          <div><span class="t-punct">)</span></div>
          <div>&nbsp;</div>
          <div><span class="t-comment"># ✓ 已连接至 apia8 官方 API · 延迟 38ms · 余额 ¥286.40</span></div>
          <div>&nbsp;<span class="t-cursor"></span></div>
        </div>
      </div>
    </section>

    <div class="stats">
      <div class="stats-inner">
        <div class="stat-item"><div class="stat-num"><span>99.99</span><span>%</span></div><div class="stat-label">可用性 SLA</div></div>
        <div class="stat-item"><div class="stat-num"><span>38</span><span>ms</span></div><div class="stat-label">平均响应延迟</div></div>
        <div class="stat-item"><div class="stat-num"><span>20</span><span>+</span></div><div class="stat-label">支持模型数量</div></div>
        <div class="stat-item"><div class="stat-num"><span>8</span><span>万+</span></div><div class="stat-label">全球开发者用户</div></div>
      </div>
    </div>

    <section class="features" id="features">
      <p class="section-label">// 核心优势</p>
      <h2 class="section-title">为什么选择 apia8</h2>
      <p class="section-desc">我们不是简单的转发，而是在保障官方质量的前提下，为你消除接入障碍。</p>

      <div class="features-grid">
        <div v-for="feature in features" :key="feature.title" class="feat-card">
          <div class="feat-icon">{{ feature.icon }}</div>
          <h3 class="feat-title">{{ feature.title }}</h3>
          <p class="feat-desc">{{ feature.desc }}</p>
        </div>
      </div>
    </section>

    <section class="models-section" id="models">
      <div class="models-inner">
        <div class="models-header">
          <div>
            <p class="section-label">// 支持模型</p>
            <h2 class="section-title">覆盖主流大模型</h2>
          </div>
          <p class="models-note">GPT 模型已支持，Claude、Gemini、DeepSeek 模型即将推出。</p>
        </div>

        <div class="models-grid">
          <div v-for="model in models" :key="model.name" class="model-chip" :class="{ pending: !model.available }">
            <div class="model-provider">{{ model.provider }}</div>
            <div class="model-name">{{ model.name }}</div>
            <span class="model-tag" :class="model.available ? 'tag-official' : 'tag-new'">
              {{ model.available ? '已支持' : '即将推出' }}
            </span>
          </div>
        </div>
      </div>
    </section>

    <section class="integration" id="integration">
      <div class="integration-inner">
        <div>
          <p class="section-label">// 快速接入</p>
          <h2 class="section-title">三步完成接入</h2>
          <p class="section-desc">无需任何运维配置，分钟级完成集成，立即投入生产。</p>

          <div class="steps">
            <div class="step-item">
              <div class="step-num">01</div>
              <div><div class="step-title">注册并获取 API Key</div><div class="step-desc">登录 apia8 控制台，创建你的专属 API Key。</div></div>
            </div>
            <div class="step-item">
              <div class="step-num">02</div>
              <div><div class="step-title">替换 base_url</div><div class="step-desc">在你的代码中将 OpenAI 的 base_url 替换为 <code>https://api.apia8.com/v1</code>，其他配置完全不变。</div></div>
            </div>
            <div class="step-item">
              <div class="step-num">03</div>
              <div><div class="step-title">上线运行</div><div class="step-desc">发送第一个请求，在控制台实时查看调用日志与费用统计。</div></div>
            </div>
          </div>

          <a :href="docUrl" target="_blank" rel="noopener noreferrer" class="btn-primary">查看文档</a>
        </div>

        <div class="code-block">
          <div class="code-header">
            <div class="code-header-left">
              <span class="tbar-dot t-red"></span>
              <span class="tbar-dot t-yellow"></span>
              <span class="tbar-dot t-green"></span>
            </div>
            <span class="code-lang">Python · Node.js · cURL</span>
          </div>
          <div class="code-body">
            <div><span class="t-comment"># Python — 官方 openai 库</span></div>
            <div>&nbsp;</div>
            <div><span class="t-kw">pip</span><span> install openai</span></div>
            <div>&nbsp;</div>
            <div><span class="t-kw">from</span><span class="t-punct"> openai </span><span class="t-kw">import</span><span> OpenAI</span></div>
            <div>&nbsp;</div>
            <div><span>client <span class="t-punct">= </span>OpenAI(</span></div>
            <div><span>&nbsp;&nbsp;api_key<span class="t-punct">=</span><span class="t-str">"sk-apia8-..."</span><span class="t-punct">,</span></span></div>
            <div><span>&nbsp;&nbsp;base_url<span class="t-punct">=</span><span class="t-str">"https://api.apia8.com/v1"</span></span></div>
            <div><span class="t-punct">)</span></div>
            <div>&nbsp;</div>
            <div><span>resp <span class="t-punct">= </span>client.chat.completions.create(</span></div>
            <div><span>&nbsp;&nbsp;model<span class="t-punct">=</span><span class="t-str">"gpt-5.5"</span><span class="t-punct">,</span></span></div>
            <div><span>&nbsp;&nbsp;messages<span class="t-punct">=[{</span><span class="t-str">"role"</span><span class="t-punct">:</span><span class="t-str">"user"</span><span class="t-punct">,</span></span></div>
            <div><span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="t-str">"content"</span><span class="t-punct">:</span><span class="t-str">"你好"</span><span class="t-punct">}]</span></span></div>
            <div><span class="t-punct">)</span></div>
            <div>&nbsp;</div>
            <div><span class="t-kw">print</span><span class="t-punct">(</span>resp.choices<span class="t-punct">[</span><span class="t-num">0</span><span class="t-punct">]</span>.message.content<span class="t-punct">)</span></div>
            <div>&nbsp;</div>
            <div><span class="t-comment"># → "你好！有什么我可以帮助你的吗？"</span></div>
          </div>
        </div>
      </div>
    </section>

    <section class="trust">
      <div class="trust-inner">
        <div class="trust-item"><span class="trust-icon">✓</span><span>官方 API 质量</span></div>
        <div class="trust-item"><span class="trust-icon">✓</span><span>OpenAI SDK 兼容</span></div>
        <div class="trust-item"><span class="trust-icon">✓</span><span>实时日志与统计</span></div>
        <div class="trust-item"><span class="trust-icon">✓</span><span>企业级稳定性</span></div>
      </div>
    </section>

    <footer>
      <div class="footer-inner">
        <div class="footer-copy">© {{ currentYear }} {{ siteName }}. 保留所有权利。</div>
        <ul class="footer-links">
          <li><a :href="docUrl" target="_blank" rel="noopener noreferrer">文档</a></li>
          <li><router-link to="/login">控制台</router-link></li>
        </ul>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAuthStore, useAppStore } from '@/stores'

const authStore = useAuthStore()
const appStore = useAppStore()

const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'apia8')
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '/install-guide')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')

const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => isAdmin.value ? '/admin/dashboard' : '/dashboard')
const currentYear = computed(() => new Date().getFullYear())

const features = [
  { icon: '🔑', title: '官方渠道直连', desc: '对接官方 API 能力，减少中间层不确定性，确保模型能力稳定还原。' },
  { icon: '⚡', title: '全球加速节点', desc: '多地区节点自动就近分配，智能负载均衡，大幅降低网络抖动。' },
  { icon: '🛡️', title: '99.99% 高可用', desc: '多线路冗余备份，单节点故障自动切换，提供企业级稳定性保障。' },
  { icon: '🔌', title: '零改造接入', desc: '兼容 OpenAI SDK 接口规范，仅需修改 base_url，迁移成本极低。' },
  { icon: '📊', title: '实时用量监控', desc: '控制台展示 Token 消耗、请求量、错误率等核心指标。' },
  { icon: '💰', title: '透明按量计费', desc: '按实际调用使用，支持配额与明细查询，团队成本更可控。' },
]

const models = [
  { provider: 'OpenAI', name: 'GPT-5.5', available: true },
  { provider: 'OpenAI', name: 'GPT-5.4', available: true },
  { provider: 'OpenAI', name: 'GPT-5.2', available: true },
  { provider: 'Anthropic', name: 'Claude Opus 4.1', available: false },
  { provider: 'Anthropic', name: 'Claude Sonnet 4', available: false },
  { provider: 'Google', name: 'Gemini 2.5 Pro', available: false },
  { provider: 'Google', name: 'Gemini 2.5 Flash', available: false },
  { provider: 'DeepSeek', name: 'DeepSeek V3', available: false },
]

onMounted(() => {
  authStore.checkAuth()
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>

<style scoped>
*, *::before, *::after { box-sizing: border-box; }

.nex-home {
  --bg: #eaf4f8;
  --bg2: rgba(255, 255, 255, 0.78);
  --bg3: rgba(238, 248, 252, 0.9);
  --accent: #008eb2;
  --accent2: #005f86;
  --gold: #b27705;
  --text: #102235;
  --text-muted: #526c82;
  --border: rgba(0, 113, 145, 0.14);
  --border2: rgba(0, 142, 178, 0.28);
  --shadow: 0 22px 60px rgba(8, 47, 73, 0.08);
  min-height: 100vh;
  background:
    radial-gradient(circle at top, rgba(0, 180, 216, 0.18), transparent 34%),
    linear-gradient(180deg, #f7fbfd 0%, var(--bg) 52%, #f8fbfc 100%);
  color: var(--text);
  font-family: 'DM Sans', ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  font-size: 16px;
  line-height: 1.6;
  overflow-x: hidden;
  position: relative;
}

.nex-home::before {
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
  box-shadow: 0 10px 30px rgba(8, 47, 73, 0.06);
}

.logo {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 800;
  font-size: 1.35rem;
  letter-spacing: 0;
  color: var(--text);
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
}

.logo-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--accent);
  box-shadow: 0 0 10px var(--accent);
  animation: pulse 2s ease-in-out infinite;
}

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
  font-weight: 400;
  letter-spacing: 0.01em;
  transition: color 0.2s;
}

.nav-links a:hover { color: var(--accent2); }

.nav-cta {
  background: transparent;
  border: 1px solid var(--border2);
  color: var(--accent) !important;
  padding: 6px 18px;
  border-radius: 6px;
  font-size: 0.85rem !important;
  font-weight: 500 !important;
  transition: background 0.2s, border-color 0.2s !important;
}

.nav-cta:hover {
  background: rgba(0,180,216,0.08) !important;
  border-color: var(--accent) !important;
  color: var(--accent2) !important;
}

section,
.stats,
footer { position: relative; z-index: 1; }

.hero {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 120px 5% 80px;
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
  animation: fadeUp 0.8s ease both;
}

.badge-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: var(--accent);
  animation: pulse 1.5s ease-in-out infinite;
}

.hero-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 800;
  font-size: clamp(2.6rem, 7vw, 5rem);
  line-height: 1.05;
  letter-spacing: 0;
  color: var(--text);
  max-width: 900px;
  margin: 0 0 1.5rem;
  animation: fadeUp 0.8s 0.1s ease both;
}

.hero-title em {
  font-style: normal;
  color: var(--accent2);
}

.hero-sub {
  font-size: 1.1rem;
  color: var(--text-muted);
  max-width: 560px;
  margin: 0 0 2.5rem;
  font-weight: 300;
  animation: fadeUp 0.8s 0.2s ease both;
}

.hero-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 4rem;
  animation: fadeUp 0.8s 0.3s ease both;
  flex-wrap: wrap;
  justify-content: center;
}

.btn-primary,
.btn-ghost {
  font-size: 0.925rem;
  padding: 12px 28px;
  border-radius: 8px;
  cursor: pointer;
  text-decoration: none;
  display: inline-block;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent), var(--accent2));
  color: #fff;
  font-weight: 600;
  border: none;
  letter-spacing: 0.01em;
  box-shadow: 0 14px 30px rgba(0, 113, 145, 0.18);
  transition: opacity 0.2s, transform 0.15s, box-shadow 0.2s;
}

.btn-primary:hover {
  opacity: 0.94;
  transform: translateY(-1px);
  box-shadow: 0 18px 36px rgba(0, 113, 145, 0.22);
}

.btn-ghost {
  background: transparent;
  color: var(--text);
  border: 1px solid var(--border2);
  transition: border-color 0.2s, background 0.2s;
}

.btn-ghost:hover {
  border-color: rgba(0,180,216,0.5);
  background: rgba(0,180,216,0.05);
}

.terminal,
.code-block {
  background: #092033;
  border: 1px solid var(--border2);
  border-radius: 12px;
  overflow: hidden;
  text-align: left;
  box-shadow: var(--shadow);
}

.terminal {
  width: 100%;
  max-width: 720px;
  animation: fadeUp 0.8s 0.4s ease both;
}

.terminal-bar,
.code-header {
  background: #102b42;
  padding: 10px 16px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--border);
}

.terminal-bar { gap: 7px; }
.code-header { justify-content: space-between; }
.code-header-left { display: flex; gap: 7px; }

.tbar-dot {
  width: 11px;
  height: 11px;
  border-radius: 50%;
}

.t-red { background: #ff5f57; }
.t-yellow { background: #febc2e; }
.t-green { background: #28c840; }

.tbar-title,
.code-lang {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.72rem;
  color: var(--text-muted);
}

.tbar-title { margin-left: 8px; }
.code-lang { letter-spacing: 0.05em; }

.terminal-body,
.code-body {
  padding: 20px 24px;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.82rem;
  line-height: 1.8;
}

.code-body { font-size: 0.8rem; line-height: 1.85; }

.t-comment { color: #3d5a78; }
.t-key { color: #e9c46a; }
.t-str { color: #52d9a8; }
.t-num { color: #f4a261; }
.t-kw { color: var(--accent); }
.t-punct { color: #7fa8c8; }
.t-cursor {
  display: inline-block;
  width: 8px;
  height: 14px;
  background: var(--accent);
  vertical-align: middle;
  animation: blink 1.1s step-end infinite;
}

.stats {
  padding: 64px 5%;
  background: rgba(255, 255, 255, 0.38);
  border-top: 1px solid var(--border);
  border-bottom: 1px solid var(--border);
}

.stats-inner {
  max-width: 1100px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
}

.stat-item {
  padding: 32px 24px;
  border-right: 1px solid var(--border);
  text-align: center;
}

.stat-item:last-child { border-right: none; }

.stat-num {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 800;
  font-size: 2.4rem;
  color: var(--text);
  letter-spacing: 0;
  margin-bottom: 6px;
}

.stat-num span { color: var(--accent); }

.stat-label {
  font-size: 0.82rem;
  color: var(--text-muted);
  letter-spacing: 0.04em;
  text-transform: uppercase;
}

.features,
.integration {
  padding: 100px 5%;
}

.features {
  max-width: 1200px;
  margin: 0 auto;
}

.section-label {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.72rem;
  color: var(--accent);
  letter-spacing: 0.15em;
  text-transform: uppercase;
  margin-bottom: 1rem;
}

.section-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 700;
  font-size: clamp(1.9rem, 4vw, 2.6rem);
  color: var(--text);
  letter-spacing: 0;
  line-height: 1.1;
  margin-bottom: 1rem;
  max-width: 600px;
}

.section-desc {
  color: var(--text-muted);
  font-size: 1rem;
  max-width: 540px;
  margin-bottom: 4rem;
  font-weight: 300;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 14px;
  background: rgba(255,255,255,0.72);
  border: 1px solid var(--border);
  border-radius: 12px;
  overflow: visible;
  padding: 14px;
  box-shadow: var(--shadow);
  backdrop-filter: blur(12px);
}

.feat-card {
  background: var(--bg2);
  padding: 36px 32px;
  border: 1px solid rgba(255, 255, 255, 0.72);
  border-radius: 10px;
  transition: background 0.2s, transform 0.2s, border-color 0.2s;
  position: relative;
  overflow: hidden;
}

.feat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, transparent, var(--accent), transparent);
  opacity: 0;
  transition: opacity 0.3s;
}

.feat-card:hover::before { opacity: 1; }
.feat-card:hover {
  background: var(--bg3);
  border-color: var(--border2);
  transform: translateY(-2px);
}

.feat-icon {
  width: 44px;
  height: 44px;
  background: rgba(0,180,216,0.08);
  border: 1px solid var(--border2);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  font-size: 1.25rem;
}

.feat-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 700;
  font-size: 1.05rem;
  color: var(--text);
  margin-bottom: 10px;
}

.feat-desc {
  color: var(--text-muted);
  font-size: 0.875rem;
  line-height: 1.7;
  font-weight: 300;
}

.models-section {
  padding: 80px 5%;
  background: rgba(255, 255, 255, 0.52);
  border-top: 1px solid var(--border);
  border-bottom: 1px solid var(--border);
}

.models-inner,
.integration-inner,
.trust-inner,
.footer-inner {
  max-width: 1100px;
  margin: 0 auto;
}

.models-header {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 3rem;
  flex-wrap: wrap;
  gap: 1rem;
}

.models-note {
  color: var(--text-muted);
  font-size: 0.85rem;
  font-weight: 300;
  max-width: 280px;
  text-align: right;
}

.models-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 12px;
}

.model-chip {
  background: var(--bg2);
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 16px 20px;
  box-shadow: 0 12px 28px rgba(8, 47, 73, 0.05);
  transition: border-color 0.2s, background 0.2s, transform 0.2s;
}

.model-chip.pending { opacity: 0.72; }
.model-chip:hover {
  border-color: var(--border2);
  background: #fff;
  transform: translateY(-2px);
}

.model-provider {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.7rem;
  color: var(--text-muted);
  letter-spacing: 0.08em;
  text-transform: uppercase;
  margin-bottom: 6px;
}

.model-name {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 600;
  font-size: 0.9rem;
  color: var(--text);
  margin-bottom: 4px;
}

.model-tag {
  display: inline-block;
  font-size: 0.68rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
}

.tag-official {
  background: rgba(0,180,216,0.12);
  color: var(--accent);
  border: 1px solid rgba(0,180,216,0.25);
}

.tag-new {
  background: rgba(233,196,106,0.12);
  color: var(--gold);
  border: 1px solid rgba(233,196,106,0.25);
}

.integration-inner {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: center;
}

.integration {
  background: linear-gradient(180deg, rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 0.44) 100%);
}

.steps { counter-reset: step; }

.step-item {
  display: flex;
  gap: 20px;
  margin-bottom: 2.5rem;
}

.step-num {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 1px solid var(--border2);
  background: rgba(0,180,216,0.05);
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.8rem;
  color: var(--accent);
  flex-shrink: 0;
  margin-top: 2px;
}

.step-title {
  font-family: 'Syne', ui-sans-serif, system-ui, sans-serif;
  font-weight: 600;
  color: var(--text);
  font-size: 1rem;
  margin-bottom: 6px;
}

.step-desc {
  font-size: 0.875rem;
  color: var(--text-muted);
  font-weight: 300;
  line-height: 1.65;
}

.step-desc code {
  font-family: 'JetBrains Mono', ui-monospace, monospace;
  font-size: 0.8em;
  color: var(--accent);
}

.trust {
  padding: 60px 5%;
  border-top: 1px solid var(--border);
}

.trust-inner {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 48px;
  flex-wrap: wrap;
}

.trust-item {
  display: flex;
  align-items: center;
  gap: 10px;
  color: var(--text-muted);
  font-size: 0.85rem;
  font-weight: 300;
}

.trust-icon { color: var(--accent); }

footer {
  padding: 48px 5% 32px;
  border-top: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.38);
}

.footer-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 1rem;
}

.footer-copy {
  font-size: 0.8rem;
  color: var(--text-muted);
  font-weight: 300;
}

.footer-links {
  display: flex;
  gap: 2rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

.footer-links a {
  font-size: 0.8rem;
  color: var(--text-muted);
  text-decoration: none;
  transition: color 0.2s;
}

.footer-links a:hover { color: var(--text); }

@keyframes pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.6; transform: scale(0.85); }
}

@keyframes blink { 0%, 100% { opacity: 1; } 50% { opacity: 0; } }

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 900px) {
  .stats-inner { grid-template-columns: repeat(2, 1fr); }
  .stat-item { border-right: none; border-bottom: 1px solid var(--border); }
  .features-grid { grid-template-columns: 1fr; }
  .integration-inner { grid-template-columns: 1fr; gap: 40px; }
  .models-note { text-align: left; }
  .nav-links { display: none; }
}

@media (max-width: 640px) {
  .hero { min-height: auto; padding-top: 112px; }
  .hero-title { font-size: 2.55rem; }
  .hero-sub { font-size: 1rem; }
  .terminal-body,
  .code-body {
    padding: 18px;
    font-size: 0.72rem;
    overflow-x: auto;
  }
  .stats-inner { grid-template-columns: 1fr; }
  .stat-item { padding: 24px 16px; }
  .features,
  .integration { padding: 72px 5%; }
  .features-grid { padding: 10px; }
  .feat-card { padding: 28px 24px; }
  .footer-inner { align-items: flex-start; }
}
</style>
