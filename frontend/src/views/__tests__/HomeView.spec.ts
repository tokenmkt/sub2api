import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import { defineComponent, h } from 'vue'

import HomeView from '../HomeView.vue'

const checkAuth = vi.fn()
const fetchPublicSettings = vi.fn()

vi.mock('vue-i18n', async () => {
  const actual = await vi.importActual<typeof import('vue-i18n')>('vue-i18n')

  return {
    ...actual,
    useI18n: () => ({
      t: (key: string) => {
        const messages: Record<string, string> = {
          'home.viewDocs': '查看文档',
          'home.switchToLight': '切换到浅色模式',
          'home.switchToDark': '切换到深色模式',
          'home.dashboard': '控制台',
          'home.login': '登录',
          'home.getStarted': '立即开始',
          'home.installGuide': '安装教程',
          'home.goToDashboard': '进入控制台',
          'home.docs': '文档',
          'home.badge': 'Official API Access',
          'home.heroTitle': '官方 AI API 网关',
          'home.heroSubtitle': '直连官方模型能力，交付稳定、纯正、可长期承载高流量的 API 接入层。',
          'home.heroDescription': 'tokenMKT 面向正式业务场景，提供统一鉴权、智能路由、实时计费与健康切换，让上游能力以更稳的方式进入你的系统。',
          'home.heroSecondaryCta': '查看文档',
          'home.quickstart.label': '快速接入',
          'home.quickstart.title': '接入方式几乎不变',
          'home.quickstart.filename': 'quickstart.py',
          'home.quickstart.comment': '# 只需替换 base_url，其余不变',
          'home.quickstart.response': '你好！我是 tokenMKT 官方 API 网关。',
          'home.metrics.official.label': '官方 API',
          'home.metrics.official.value': 'Direct',
          'home.metrics.stability.label': '稳定路由',
          'home.metrics.stability.value': '24/7',
          'home.metrics.purity.label': '纯正上游',
          'home.metrics.purity.value': 'Native',
          'home.trust.title': '为高流量业务准备的接入层',
          'home.trust.subtitle': '不只是能调用，而是让正式业务长期稳定地调用。',
          'home.features.official.title': '官方 API 接入',
          'home.features.official.desc': '保持接口能力、鉴权方式与上游模型演进节奏同步。',
          'home.features.stability.title': '稳定承载高流量',
          'home.features.stability.desc': '通过多节点路由、健康探测与自动切换降低抖动和中断。',
          'home.features.purity.title': '纯正模型输出链路',
          'home.features.purity.desc': '减少非必要中间层干预，让返回结果更接近官方原始能力。',
          'home.providers.title': 'AI 模型',
          'home.providers.description': '一个 API，连接主流官方模型',
          'home.providers.supported': '已支持',
          'home.providers.soon': '即将推出',
          'home.providers.claude': 'Claude',
          'home.providers.gemini': 'Gemini',
          'home.providers.more': '更多',
          'home.cta.title': '把官方 AI 能力稳定接进你的产品',
          'home.cta.description': '注册后即可开始接入，适合面向生产环境的 API 调用场景。',
          'home.cta.button': '立即接入',
          'home.footer.allRightsReserved': '保留所有权利。'
        }

        return messages[key] ?? key
      }
    })
  }
})

vi.mock('@/stores', () => ({
  useAuthStore: () => ({
    isAuthenticated: false,
    isAdmin: false,
    user: null,
    checkAuth
  }),
  useAppStore: () => ({
    siteName: 'tokenMKT',
    siteLogo: '',
    docUrl: 'https://docs.example.com',
    publicSettingsLoaded: false,
    fetchPublicSettings,
    cachedPublicSettings: {
      site_name: 'tokenMKT',
      site_logo: '',
      site_subtitle: 'AI API Gateway Platform',
      doc_url: 'https://docs.example.com',
      home_content: ''
    }
  })
}))

describe('HomeView', () => {
  beforeEach(() => {
    checkAuth.mockReset()
    fetchPublicSettings.mockReset()
    document.documentElement.classList.remove('dark')
    localStorage.clear()
    Object.defineProperty(window, 'matchMedia', {
      writable: true,
      value: vi.fn().mockImplementation(() => ({
        matches: false,
        media: '',
        onchange: null,
        addListener: vi.fn(),
        removeListener: vi.fn(),
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        dispatchEvent: vi.fn()
      }))
    })
  })

  it('突出官方 API、稳定路由和纯正上游卖点', async () => {
    const RouterLinkStub = defineComponent({
      name: 'RouterLink',
      props: {
        to: {
          type: [String, Object],
          required: false,
          default: ''
        }
      },
      setup(props, { slots }) {
        const toValue = typeof props.to === 'string' ? props.to : JSON.stringify(props.to)
        return () => h('a', { 'data-to': toValue }, slots.default?.())
      }
    })

    const wrapper = mount(HomeView, {
      global: {
        stubs: {
          LocaleSwitcher: { template: '<div />' },
          Icon: { template: '<span />' },
          RouterLink: RouterLinkStub
        }
      }
    })

    await nextTick()

    const text = wrapper.text()

    expect(text).toContain('Official API Access')
    expect(text).toContain('官方 API 接入')
    expect(text).toContain('稳定承载高流量')
    expect(text).toContain('纯正模型输出链路')
    expect(text).toContain('快速接入')
    expect(text).toContain('只需替换 base_url')
    expect(text).toContain('quickstart.py')
    expect(text).toContain('https://api.tokenmkt.cc/v1')
    expect(text).toContain('安装教程')
    expect(wrapper.html()).toContain('data-to="/install-guide"')
    expect(checkAuth).toHaveBeenCalledTimes(1)
    expect(fetchPublicSettings).toHaveBeenCalledTimes(1)
  })
})
