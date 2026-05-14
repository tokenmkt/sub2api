import { mount } from '@vue/test-utils'
import { defineComponent, h, nextTick } from 'vue'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import AppHeader from '@/components/layout/AppHeader.vue'

const { authState, i18nMessages, logoutMock, replayMock } = vi.hoisted(() => ({
  authState: {
    user: {
      username: 'admin',
      email: 'admin@tokenmkt.cc',
      role: 'admin',
      balance: 499.94,
      avatar_url: ''
    },
    isAdmin: true,
    isSimpleMode: false,
    logout: vi.fn()
  },
  i18nMessages: {
    'nav.dashboard': '控制台',
    'nav.profile': '个人资料',
    'nav.apiKeys': 'API 密钥',
    'nav.github': 'GitHub',
    'nav.logout': '退出登录',
    'common.balance': '余额',
    'common.contactSupport': '联系客服',
    'onboarding.restartTour': '重新查看新手引导'
  } as Record<string, string>,
  logoutMock: vi.fn(),
  replayMock: vi.fn()
}))

vi.mock('@/stores', () => ({
  useAuthStore: () => authState,
  useAppStore: () => ({
    contactInfo: '12345678',
    docUrl: '',
    cachedPublicSettings: {
      custom_menu_items: []
    },
    toggleMobileSidebar: vi.fn()
  }),
  useOnboardingStore: () => ({
    replay: replayMock
  })
}))

vi.mock('@/stores/adminSettings', () => ({
  useAdminSettingsStore: () => ({
    customMenuItems: []
  })
}))

vi.mock('vue-router', () => ({
  useRouter: () => ({
    push: vi.fn()
  }),
  useRoute: () => ({
    name: 'Dashboard',
    meta: {
      titleKey: 'nav.dashboard'
    },
    params: {}
  })
}))

vi.mock('vue-i18n', async (importOriginal) => {
  const actual = await importOriginal<typeof import('vue-i18n')>()

  return {
    ...actual,
    useI18n: () => ({
      t: (key: string) => i18nMessages[key] ?? key
    })
  }
})

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

const translate = (key: string) => {
  return i18nMessages[key] ?? key
}

describe('AppHeader', () => {
  beforeEach(() => {
    authState.logout = logoutMock
    logoutMock.mockReset()
    replayMock.mockReset()
  })

  it('does not show the GitHub link in the user dropdown for admins', async () => {
    const wrapper = mount(AppHeader, {
      global: {
        stubs: {
          AnnouncementBell: true,
          LocaleSwitcher: true,
          SubscriptionProgressMini: true,
          Icon: {
            props: ['name'],
            template: '<span class="icon-stub" :data-icon="name" />'
          },
          RouterLink: RouterLinkStub
        },
        mocks: {
          $t: translate
        }
      }
    })

    await wrapper.get('button[aria-label="User Menu"]').trigger('click')
    await nextTick()

    expect(wrapper.text()).toContain('个人资料')
    expect(wrapper.text()).toContain('API 密钥')
    expect(wrapper.text()).not.toContain('GitHub')
    expect(wrapper.find('a[href="https://github.com/Wei-Shaw/sub2api"]').exists()).toBe(false)
  })
})
