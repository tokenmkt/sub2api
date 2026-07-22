export default {
  batchImageGuide: {
    title: '图片批量生成',
    description: '一次提交多条提示词，任务完成后可统一下载图片结果'
  },
  // Home Page
  home: {
    viewOnGithub: '在 GitHub 上查看',
    viewDocs: '查看文档',
    docs: '文档',
    switchToLight: '切换到浅色模式',
    switchToDark: '切换到深色模式',
    dashboard: '控制台',
    login: '登录',
    getStarted: '立即开始',
    installGuide: '安装教程',
    goToDashboard: '进入控制台',
    badge: '官方 API 接入',
    heroTitle: '官方 AI API 网关',
    // 新增：面向用户的价值主张
    heroSubtitle: '直连官方模型能力，交付稳定、纯正、可长期承载高流量的 API 接入层。',
    heroDescription: 'tokenMKT 面向正式业务场景，提供统一鉴权、智能路由、实时计费与健康切换，让上游能力以更稳的方式进入你的系统。',
    heroSecondaryCta: '查看文档',
    quickstart: { label: '快速接入', title: '接入方式几乎不变', filename: 'quickstart.py', comment: '# 只需替换 base_url，其余不变', response: '你好！我是 tokenMKT 官方 API 网关。' },
    metrics: {
      official: { label: '官方 API', value: '直连' },
      stability: { label: '稳定路由', value: '24/7' },
      purity: { label: '纯正上游', value: '原生' }
    },
    trust: { title: '为高流量业务准备的接入层', subtitle: '不只是能调用，而是让正式业务长期稳定地调用。' },
    tags: {
      subscriptionToApi: '订阅转 API',
      stickySession: '会话保持',
      realtimeBilling: '按量计费'
    },
    // 用户痛点区块
    painPoints: {
      title: '你是否也遇到这些问题？',
      items: {
        expensive: {
          title: '订阅费用高',
          desc: '每个 AI 服务都要单独订阅，每月支出越来越多'
        },
        complex: {
          title: '多账号难管理',
          desc: '不同平台的账号、密钥分散各处，管理起来很麻烦'
        },
        unstable: {
          title: '服务不稳定',
          desc: '单一账号容易触发限制，影响正常使用'
        },
        noControl: {
          title: '用量无法控制',
          desc: '不知道钱花在哪了，也无法限制团队成员的使用'
        }
      }
    },
    // 解决方案区块
    solutions: {
      title: '我们帮你解决',
      subtitle: '简单三步，开始省心使用 AI'
    },
    features: {
      official: { title: '官方 API 接入', desc: '保持接口能力、鉴权方式与上游模型演进节奏同步。' },
      stability: { title: '稳定承载高流量', desc: '通过多节点路由、健康探测与自动切换降低抖动和中断。' },
      purity: { title: '纯正模型输出链路', desc: '减少非必要中间层干预，让返回结果更接近官方原始能力。' },
      unifiedGateway: '一键接入',
      unifiedGatewayDesc: '获取一个 API 密钥，即可调用所有已接入的 AI 模型，无需分别申请。',
      multiAccount: '稳定可靠',
      multiAccountDesc: '智能调度多个上游账号，自动切换和负载均衡，告别频繁报错。',
      balanceQuota: '用多少付多少',
      balanceQuotaDesc: '按实际使用量计费，支持设置配额上限，团队用量一目了然。'
    },
    // 优势对比
    comparison: {
      title: '为什么选择我们？',
      headers: {
        feature: '对比项',
        official: '官方订阅',
        us: '本平台'
      },
      items: {
        pricing: {
          feature: '付费方式',
          official: '固定月费，用不完也付',
          us: '按量付费，用多少付多少'
        },
        models: {
          feature: '模型选择',
          official: '单一服务商',
          us: '多模型随意切换'
        },
        management: {
          feature: '账号管理',
          official: '每个服务单独管理',
          us: '统一密钥，一站管理'
        },
        stability: {
          feature: '服务稳定性',
          official: '单账号易触发限制',
          us: '多账号池，自动切换'
        },
        control: {
          feature: '用量控制',
          official: '无法限制',
          us: '可设配额、查明细'
        }
      }
    },
    providers: {
      title: '已支持的 AI 模型',
      description: '一个 API，连接主流官方模型',
      supported: '已支持',
      soon: '即将推出',
      claude: 'Claude',
      gemini: 'Gemini',
      antigravity: 'Antigravity',
      more: '更多'
    },
    // CTA 区块
    cta: {
      title: '把官方 AI 能力稳定接进你的产品',
      description: '注册后即可开始接入，适合面向生产环境的 API 调用场景。',
      button: '立即接入'
    },
    footer: {
      allRightsReserved: '保留所有权利。'
    }
  },

  installGuide: {
    badge: '安装与配置', title: '安装教程', subtitle: '按照你常用的 AI CLI 工具选择对应配置，几分钟内接入 tokenMKT 官方 API 网关。',
    backHome: '返回首页', overviewTitle: '安装流程', stepPrefix: '步骤', checklistTitle: '安装前准备', optionalEyebrow: '可选配置', optionalTitle: '补充配置',
    copy: '复制', copySuccess: '已复制到剪贴板', tipsEyebrow: '使用建议', tipsTitle: '接入时建议这样做',
    labels: { install: '安装', config: '配置', verify: '验证' },
    blockEyebrows: { install: '安装命令', config: '配置文件', auth: '鉴权文件', env: '环境变量', ide: '编辑器配置', verify: '验证命令' },
    tip1: '优先复制完整配置，不要只改部分环境变量，避免本地残留旧值。', tip2: '如果你已经接过 OpenAI 兼容接口，通常只需要替换 Base URL 和 API Key。', tip3: '接入完成后先跑一次最小命令确认链路畅通，再开始正式业务调用。',
    codex: {
      summaryTitle: 'Codex CLI 接入', summary: '适合需要 OpenAI 兼容工作流的场景。按照配置文件和鉴权文件两步完成接入，之后直接用 `codex` 启动即可。',
      installValue: '全局安装 Codex CLI', configValue: '写入 ~/.codex/config.toml 与 auth.json', verifyValue: '终端执行 codex 验证',
      check1: '确保本机已安装 Node.js 18+。', check2: '准备好 tokenMKT API Key，例如 sk-tokenmkt-...。', check3: '建议先备份已有 ~/.codex 配置，再覆盖为新的网关接入参数。',
      installBlockTitle: '安装 Codex CLI', installBlockDesc: '如果你尚未安装 Codex CLI，先执行一次全局安装。', configBlockTitle: '写入 Codex 配置文件', configBlockDesc: '将模型提供方切到 OpenAI，并把请求发到 tokenMKT 网关。', authBlockTitle: '写入 API Key', authBlockDesc: '把你的 tokenMKT Key 放进 auth.json，CLI 会自动读取。', verifyBlockTitle: '验证是否安装完成', verifyBlockDesc: '执行后如果 CLI 可以正常启动并继续会话，就说明接入成功。'
    },
    claude: {
      summaryTitle: 'Claude Code 接入', summary: '适合使用 Claude Code 或 VSCode Claude Code 插件的场景。核心是把 Anthropic 相关环境变量切到 tokenMKT 网关。', installValue: '全局安装 Claude Code', configValue: '设置 ANTHROPIC_BASE_URL 与 Token', verifyValue: '终端执行 claude 验证',
      check1: '确保本机已安装 Node.js，并能执行 npm 命令。', check2: '准备好 tokenMKT API Key，并确认网关域名可访问。', check3: '如果你同时在终端和 VSCode 使用 Claude Code，建议两边都写入同一组配置。', installBlockTitle: '安装 Claude Code', installBlockDesc: '先安装官方 CLI，再接入 tokenMKT 网关。', envBlockTitle: '配置终端环境变量', envBlockDesc: '终端模式下最关键的是 Base URL、Token，以及关闭非必要流量。', ideBlockTitle: '配置 VSCode Claude Code', ideBlockDesc: '如果你在 VSCode 里使用 Claude Code，也可以直接在 settings.json 注入相同环境变量。', verifyBlockTitle: '验证是否安装完成', verifyBlockDesc: '执行后如果 Claude Code 能正常打开会话，即表示链路已经切到 tokenMKT。'
    },
    gemini: {
      summaryTitle: 'Gemini CLI 接入', summary: '适合使用 Gemini CLI 的场景。配置方式最轻，只需要安装 CLI 并写入三项环境变量。', installValue: '全局安装 Gemini CLI', configValue: '设置 Google Gemini Base URL 与 Key', verifyValue: '终端执行 gemini 验证',
      check1: '确保本机已安装 Node.js 18+。', check2: '准备好 tokenMKT API Key，并确认你的网络可以访问 tokenMKT 网关。', check3: '如果你本机已有旧的 Gemini 环境变量，建议先清掉再重新设置。', installBlockTitle: '安装 Gemini CLI', installBlockDesc: '先安装 CLI 工具本体，再接入 tokenMKT 的 Gemini 网关地址。', envBlockTitle: '配置 Gemini 环境变量', envBlockDesc: '设置 Base URL、API Key 和默认模型后，CLI 就会走 tokenMKT 通道。', verifyBlockTitle: '验证是否安装完成', verifyBlockDesc: '执行后若 Gemini CLI 能正常进入命令行会话，即表示接入成功。'
    }
  },

  // Key Usage Query Page
  keyUsage: {
    title: 'API Key 用量查询',
    subtitle: '输入您的 API Key 以查看实时消费金额与使用状态',
    placeholder: 'sk-ant-mirror-xxxxxxxxxxxx',
    query: '查询',
    querying: '查询中...',
    privacyNote: '您的 Key 仅在浏览器本地处理，不会被存储',
    dateRange: '统计范围:',
    dateRangeToday: '今日',
    dateRange7d: '7 天',
    dateRange30d: '30 天',
    dateRange90d: '90 天',
    dateRangeCustom: '自定义',
    apply: '应用',
    used: '已使用',
    detailInfo: '详细信息',
    tokenStats: 'Token 统计',
    dailyDetail: '按日明细',
    modelStats: '模型用量统计',
    // Table headers
    date: '日期',
    model: '模型',
    requests: '请求数',
    inputTokens: '输入 Tokens',
    outputTokens: '输出 Tokens',
    cacheCreationTokens: '缓存创建',
    cacheReadTokens: '缓存读取',
    cacheWriteTokens: '缓存写入',
    totalTokens: '总 Tokens',
    cost: '费用',
    // Status
    quotaMode: 'Key 限额模式',
    walletBalance: '钱包余额',
    // Ring card titles
    totalQuota: '总额度',
    limit5h: '5 小时限额',
    limitDaily: '日限额',
    limit7d: '7 天限额',
    limitWeekly: '周限额',
    limitMonthly: '月限额',
    // Detail rows
    remainingQuota: '剩余额度',
    expiresAt: '过期时间',
    todayExpires: '(今日到期)',
    daysLeft: '({days} 天)',
    usedQuota: '已用额度',
    resetNow: '即将重置',
    subscriptionType: '订阅类型',
    subscriptionExpires: '订阅到期',
    // Usage stat cells
    todayRequests: '今日请求',
    todayInputTokens: '今日输入',
    todayOutputTokens: '今日输出',
    todayTokens: '今日 Tokens',
    todayCacheCreation: '今日缓存创建',
    todayCacheRead: '今日缓存读取',
    todayCost: '今日费用',
    rpmTpm: 'RPM / TPM',
    totalRequests: '累计请求',
    totalInputTokens: '累计输入',
    totalOutputTokens: '累计输出',
    totalTokensLabel: '累计 Tokens',
    totalCacheCreation: '累计缓存创建',
    totalCacheRead: '累计缓存读取',
    totalCost: '累计费用',
    avgDuration: '平均耗时',
    // Messages
    enterApiKey: '请输入 API Key',
    querySuccess: '查询成功',
    queryFailed: '查询失败',
    queryFailedRetry: '查询失败，请稍后重试',
    noDailyUsage: '暂无按日用量数据',
  },

  // Setup Wizard
  setup: {
    title: 'tokenMKT 安装向导',
    description: '配置您的 tokenMKT 实例',
    database: {
      title: '数据库配置',
      description: '连接到您的 PostgreSQL 数据库',
      host: '主机',
      port: '端口',
      username: '用户名',
      password: '密码',
      databaseName: '数据库名称',
      sslMode: 'SSL 模式',
      passwordPlaceholder: '密码',
      ssl: {
        disable: '禁用',
        require: '要求',
        verifyCa: '验证 CA',
        verifyFull: '完全验证'
      }
    },
    redis: {
      title: 'Redis 配置',
      description: '连接到您的 Redis 服务器',
      host: '主机',
      port: '端口',
      username: '用户名（可选）',
      password: '密码（可选）',
      database: '数据库',
      usernamePlaceholder: '默认用户留空',
      passwordPlaceholder: '密码',
      enableTls: '启用 TLS',
      enableTlsHint: '连接 Redis 时使用 TLS（公共 CA 证书）'
    },
    admin: {
      title: '管理员账户',
      description: '创建您的管理员账户',
      email: '邮箱',
      password: '密码',
      confirmPassword: '确认密码',
      passwordPlaceholder: '至少 8 个字符',
      confirmPasswordPlaceholder: '确认密码',
      passwordMismatch: '密码不匹配'
    },
    ready: {
      title: '准备安装',
      description: '检查您的配置并完成安装',
      database: '数据库',
      redis: 'Redis',
      adminEmail: '管理员邮箱'
    },
    status: {
      testing: '测试中...',
      success: '连接成功',
      testConnection: '测试连接',
      installing: '安装中...',
      completeInstallation: '完成安装',
      completed: '安装完成！',
      redirecting: '正在跳转到登录页面...',
      restarting: '服务正在重启，请稍候...',
      timeout: '服务重启时间超出预期，请手动刷新页面。'
    }
  },

  // Common
}
