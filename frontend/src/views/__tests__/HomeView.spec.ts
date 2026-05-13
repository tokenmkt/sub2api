import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import HomeView from '../HomeView.vue'

vi.mock('vue-i18n', async (importOriginal) => ({
  ...(await importOriginal<typeof import('vue-i18n')>()),
  useI18n: () => ({
    t: (key: string) => ({
      'home.providers.supported': '已支持',
      'home.providers.soon': '即将推出',
    })[key] ?? key,
  }),
}))

vi.mock('@/stores', () => ({
  useAppStore: () => ({
    siteName: 'apia8',
    siteLogo: '',
    docUrl: '',
    cachedPublicSettings: null,
    publicSettingsLoaded: true,
    fetchPublicSettings: vi.fn(),
  }),
  useAuthStore: () => ({
    isAuthenticated: false,
    isAdmin: false,
    user: null,
    checkAuth: vi.fn(),
  }),
}))

describe('HomeView', () => {
  beforeEach(async () => {
    Object.defineProperty(window, 'matchMedia', {
      writable: true,
      value: vi.fn().mockImplementation((query: string) => ({
        matches: false,
        media: query,
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        addListener: vi.fn(),
        removeListener: vi.fn(),
        dispatchEvent: vi.fn(),
      })),
    })
  })

  it('shows apia8 branding, the current model statuses, and opens docs in a new page', async () => {
    const wrapper = mount(HomeView, {
      global: {
        stubs: {
          LocaleSwitcher: true,
          Icon: true,
          RouterLink: {
            props: ['to'],
            template: '<a :href="typeof to === \'string\' ? to : to.path"><slot /></a>',
          },
        },
      },
    })

    expect(wrapper.text()).toContain('apia8')
    expect(wrapper.text()).toContain('开始使用 →')
    expect(wrapper.text()).not.toContain('免费开始使用')
    expect(wrapper.text()).toContain('GPT-5.5')
    expect(wrapper.text()).toContain('Claude Opus 4.1')
    expect(wrapper.text()).toContain('Claude Sonnet 4')
    expect(wrapper.text()).toContain('Gemini 2.5 Pro')
    expect(wrapper.text()).toContain('Gemini 2.5 Flash')
    expect(wrapper.text()).toContain('DeepSeek V3')
    expect(wrapper.text()).toContain('GPT 模型已支持，Claude、Gemini、DeepSeek 模型即将推出。')
    expect(wrapper.findAll('.tag-new')).toHaveLength(5)

    const newPageDocLinks = wrapper
      .findAll('a[href="/install-guide"]')
      .filter((link) => link.attributes('target') === '_blank')

    expect(newPageDocLinks.length).toBeGreaterThan(0)
    newPageDocLinks.forEach((link) => {
      expect(link.attributes('rel')).toContain('noopener noreferrer')
    })
  })
})
