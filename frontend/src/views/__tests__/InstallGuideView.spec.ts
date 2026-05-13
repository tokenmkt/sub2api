import { mount } from '@vue/test-utils'
import { describe, expect, it, vi } from 'vitest'
import InstallGuideView from '../InstallGuideView.vue'

vi.mock('vue-i18n', async (importOriginal) => ({
  ...(await importOriginal<typeof import('vue-i18n')>()),
  useI18n: () => ({
    t: (key: string) => ({
      'home.viewDocs': '查看文档',
      'home.switchToLight': '切换到浅色模式',
      'home.switchToDark': '切换到深色模式',
      'home.footer.allRightsReserved': '保留所有权利。',
      'installGuide.badge': '安装与配置',
      'installGuide.title': 'apia8 安装指南',
      'installGuide.subtitle': '按照你常用的 AI CLI 工具选择对应配置，几分钟内接入 apia8 官方 API 网关。',
      'installGuide.backHome': '返回首页',
      'installGuide.overviewTitle': '安装流程',
      'installGuide.stepPrefix': '步骤',
      'installGuide.checklistTitle': '安装前准备',
      'installGuide.copy': '复制',
      'installGuide.copySuccess': '已复制到剪贴板',
      'installGuide.optionalEyebrow': '可选配置',
      'installGuide.optionalTitle': '补充配置',
      'installGuide.tipsEyebrow': '使用建议',
      'installGuide.tipsTitle': '接入时建议这样做',
      'installGuide.labels.install': '安装',
      'installGuide.labels.config': '配置',
      'installGuide.labels.verify': '验证',
      'installGuide.blockEyebrows.install': '安装命令',
      'installGuide.blockEyebrows.config': '配置文件',
      'installGuide.blockEyebrows.auth': '鉴权文件',
      'installGuide.blockEyebrows.env': '环境变量',
      'installGuide.blockEyebrows.ide': '编辑器配置',
      'installGuide.blockEyebrows.verify': '验证命令',
      'installGuide.tip1': '优先复制完整配置。',
      'installGuide.tip2': '通常只需要替换 Base URL 和 API Key。',
      'installGuide.tip3': '先跑一次最小命令确认链路。',
      'installGuide.codex.summaryTitle': 'Codex CLI 接入',
      'installGuide.codex.summary': '适合需要 OpenAI 兼容工作流的场景。',
      'installGuide.codex.installValue': '全局安装 Codex CLI',
      'installGuide.codex.configValue': '写入配置文件',
      'installGuide.codex.verifyValue': '终端执行 codex 验证',
      'installGuide.codex.check1': '确保本机已安装 Node.js 18+。',
      'installGuide.codex.check2': '准备好 apia8 API Key，例如 sk-apia8-...。',
      'installGuide.codex.check3': '建议先备份已有配置。',
      'installGuide.codex.installBlockTitle': '安装 Codex CLI',
      'installGuide.codex.installBlockDesc': '先执行一次全局安装。',
      'installGuide.codex.configBlockTitle': '写入 Codex 配置文件',
      'installGuide.codex.configBlockDesc': '将模型提供方切到 OpenAI。',
      'installGuide.codex.authBlockTitle': '写入 API Key',
      'installGuide.codex.authBlockDesc': '把你的 apia8 Key 放进 auth.json。',
      'installGuide.codex.verifyBlockTitle': '验证是否安装完成',
      'installGuide.codex.verifyBlockDesc': '执行后如果 CLI 可以正常启动即可。',
      'installGuide.claude.summaryTitle': 'Claude Code 接入',
      'installGuide.claude.summary': '适合使用 Claude Code 的场景。',
      'installGuide.claude.installValue': '全局安装 Claude Code',
      'installGuide.claude.configValue': '设置 ANTHROPIC_BASE_URL 与 Token',
      'installGuide.claude.verifyValue': '终端执行 claude 验证',
      'installGuide.claude.check1': '确保本机已安装 Node.js。',
      'installGuide.claude.check2': '准备好 apia8 API Key。',
      'installGuide.claude.check3': '建议两边都写入同一组配置。',
      'installGuide.claude.installBlockTitle': '安装 Claude Code',
      'installGuide.claude.installBlockDesc': '先安装官方 CLI。',
      'installGuide.claude.envBlockTitle': '配置终端环境变量',
      'installGuide.claude.envBlockDesc': '配置 Base URL、Token。',
      'installGuide.claude.ideBlockTitle': '配置 VSCode Claude Code',
      'installGuide.claude.ideBlockDesc': '注入相同环境变量。',
      'installGuide.claude.verifyBlockTitle': '验证是否安装完成',
      'installGuide.claude.verifyBlockDesc': '执行 claude 验证。',
      'installGuide.gemini.summaryTitle': 'Gemini CLI 接入',
      'installGuide.gemini.summary': '适合使用 Gemini CLI 的场景。',
      'installGuide.gemini.installValue': '全局安装 Gemini CLI',
      'installGuide.gemini.configValue': '设置 Google Gemini Base URL 与 Key',
      'installGuide.gemini.verifyValue': '终端执行 gemini 验证',
      'installGuide.gemini.check1': '确保本机已安装 Node.js 18+。',
      'installGuide.gemini.check2': '准备好 apia8 API Key。',
      'installGuide.gemini.check3': '建议先清掉旧变量。',
      'installGuide.gemini.installBlockTitle': '安装 Gemini CLI',
      'installGuide.gemini.installBlockDesc': '先安装 CLI 工具本体。',
      'installGuide.gemini.envBlockTitle': '配置 Gemini 环境变量',
      'installGuide.gemini.envBlockDesc': '设置 Base URL、API Key。',
      'installGuide.gemini.verifyBlockTitle': '验证是否安装完成',
      'installGuide.gemini.verifyBlockDesc': '执行 gemini 验证。',
    })[key] ?? key,
  }),
}))

vi.mock('@/stores/app', () => ({
  useAppStore: () => ({
    siteName: 'apia8',
    siteLogo: '',
    docUrl: '',
    cachedPublicSettings: null,
    showSuccess: vi.fn(),
    showError: vi.fn(),
  }),
}))

describe('InstallGuideView', () => {
  it('renders a public install guide for apia8', () => {
    const wrapper = mount(InstallGuideView, {
      global: {
        stubs: {
          RouterLink: {
            props: ['to'],
            template: '<a :href="typeof to === \'string\' ? to : to.path"><slot /></a>'
          },
          LocaleSwitcher: true,
          Icon: true,
        }
      }
    })

    expect(wrapper.text()).toContain('apia8 安装指南')
    expect(wrapper.text()).toContain('Codex CLI')
    expect(wrapper.text()).toContain('Claude Code')
    expect(wrapper.text()).toContain('安装前准备')
    expect(wrapper.text()).toContain('写入 Codex 配置文件')
    expect(wrapper.text()).toContain('https://api.apia8.com/v1')
    expect(wrapper.text()).toContain('sk-apia8')
    expect(wrapper.text()).toContain('使用建议')
  })
})
