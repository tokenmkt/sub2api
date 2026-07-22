export default {
  batchImageGuide: {
    title: 'Batch Image Generation',
    description: 'Submit multiple prompts in one job and download the generated images when complete'
  },
  // Home Page
  home: {
    viewOnGithub: 'View on GitHub',
    viewDocs: 'View Documentation',
    docs: 'Docs',
    switchToLight: 'Switch to Light Mode',
    switchToDark: 'Switch to Dark Mode',
    dashboard: 'Dashboard',
    login: 'Login',
    getStarted: 'Get Started',
    installGuide: 'Install Guide',
    goToDashboard: 'Go to Dashboard',
    badge: 'Official API Access',
    heroTitle: 'Official AI API Gateway',
    // User-focused value proposition
    heroSubtitle: 'Direct access to official model capabilities with a stable, native-grade API layer built for sustained high traffic.',
    heroDescription: 'tokenMKT is designed for production workloads with unified auth, smart routing, realtime billing, and health-aware failover that keeps upstream capability dependable inside your system.',
    heroSecondaryCta: 'View Documentation',
    quickstart: {
      label: 'Quickstart',
      title: 'Almost no integration changes',
      filename: 'quickstart.py',
      comment: '# Only replace base_url, keep the rest unchanged',
      response: 'Hello! I am the tokenMKT official API gateway.'
    },
    metrics: {
      official: { label: 'Official API', value: 'Direct' },
      stability: { label: 'Stable Routing', value: '24/7' },
      purity: { label: 'Native Fidelity', value: 'Raw' }
    },
    trust: {
      title: 'An access layer built for high-traffic production workloads',
      subtitle: 'Not just API availability, but dependable long-term delivery for real business traffic.'
    },
    tags: {
      subscriptionToApi: 'Subscription to API',
      stickySession: 'Session Persistence',
      realtimeBilling: 'Pay As You Go'
    },
    // Pain points section
    painPoints: {
      title: 'Sound Familiar?',
      items: {
        expensive: {
          title: 'High Subscription Costs',
          desc: 'Paying for multiple AI subscriptions that add up every month'
        },
        complex: {
          title: 'Account Chaos',
          desc: 'Managing scattered accounts and API keys across different platforms'
        },
        unstable: {
          title: 'Service Interruptions',
          desc: 'Single accounts hitting rate limits and disrupting your workflow'
        },
        noControl: {
          title: 'No Usage Control',
          desc: "Can't track where your money goes or limit team member usage"
        }
      }
    },
    // Solutions section
    solutions: {
      title: 'We Solve These Problems',
      subtitle: 'Three simple steps to stress-free AI access'
    },
    features: {
      official: {
        title: 'Official API Access',
        desc: 'Stay aligned with upstream auth patterns, interface capability, and model evolution.'
      },
      stability: {
        title: 'Stable Under High Traffic',
        desc: 'Use health-aware routing, multi-node balancing, and automatic switchover to reduce disruption.'
      },
      purity: {
        title: 'Native Model Delivery Path',
        desc: 'Minimize unnecessary middleware interference so results stay close to official output.'
      },
      unifiedGateway: 'One-Click Access',
      unifiedGatewayDesc: 'Get a single API key to call all connected AI models. No separate applications needed.',
      multiAccount: 'Always Reliable',
      multiAccountDesc: 'Smart routing across multiple upstream accounts with automatic failover. Say goodbye to errors.',
      balanceQuota: 'Pay What You Use',
      balanceQuotaDesc: 'Usage-based billing with quota limits. Full visibility into team consumption.'
    },
    // Comparison section
    comparison: {
      title: 'Why Choose Us?',
      headers: {
        feature: 'Comparison',
        official: 'Official Subscriptions',
        us: 'Our Platform'
      },
      items: {
        pricing: {
          feature: 'Pricing',
          official: 'Fixed monthly fee, pay even if unused',
          us: 'Pay only for what you use'
        },
        models: {
          feature: 'Model Selection',
          official: 'Single provider only',
          us: 'Switch between models freely'
        },
        management: {
          feature: 'Account Management',
          official: 'Manage each service separately',
          us: 'Unified key, one dashboard'
        },
        stability: {
          feature: 'Stability',
          official: 'Single account rate limits',
          us: 'Multi-account pool, auto-failover'
        },
        control: {
          feature: 'Usage Control',
          official: 'Not available',
          us: 'Quotas & detailed analytics'
        }
      }
    },
    providers: {
      title: 'Supported AI Models',
      description: 'One API connected to major official model providers',
      supported: 'Supported',
      soon: 'Soon',
      claude: 'Claude',
      gemini: 'Gemini',
      antigravity: 'Antigravity',
      more: 'More'
    },
    // CTA section
    cta: {
      title: 'Bring official AI capability into your product with a steadier API layer',
      description: 'Start integrating after signup, with a setup shaped for production API traffic.',
      button: 'Start Integration'
    },
    footer: {
      allRightsReserved: 'All rights reserved.'
    }
  },

  installGuide: {
    badge: 'Install & Configure',
    title: 'Install Guide',
    subtitle: 'Choose the setup that matches your AI CLI tool and connect to the tokenMKT official API gateway in minutes.',
    backHome: 'Back Home',
    overviewTitle: 'See the flow first',
    stepPrefix: 'Step',
    checklistTitle: 'Before You Start',
    optionalEyebrow: 'Optional',
    optionalTitle: 'Additional Configuration',
    copy: 'Copy',
    copySuccess: 'Copied to clipboard',
    tipsEyebrow: 'Best Practice',
    tipsTitle: 'Recommended Setup Tips',
    labels: { install: 'Install', config: 'Config', verify: 'Verify' },
    blockEyebrows: {
      install: 'Install Command', config: 'Config File', auth: 'Auth File',
      env: 'Environment', ide: 'Editor Config', verify: 'Verify Command'
    },
    tip1: 'Copy the complete configuration together so old local variables do not override the new gateway setup.',
    tip2: 'If you already use an OpenAI-compatible workflow, you usually only need to replace the base URL and API key.',
    tip3: 'After setup, run one minimal verification command before using it in real traffic.',
    codex: {
      summaryTitle: 'Codex CLI Setup', summary: 'Best for OpenAI-compatible workflows. Add one config file and one auth file, then launch with `codex`.',
      installValue: 'Install Codex CLI globally', configValue: 'Write ~/.codex/config.toml and auth.json', verifyValue: 'Run codex in terminal',
      check1: 'Make sure Node.js 18+ is available on your machine.', check2: 'Prepare your tokenMKT API key such as sk-tokenmkt-....', check3: 'Back up any existing ~/.codex files before replacing them.',
      installBlockTitle: 'Install Codex CLI', installBlockDesc: 'Run a global install first if Codex CLI is not already available.',
      configBlockTitle: 'Create the Codex config file', configBlockDesc: 'Point the provider to OpenAI-compatible responses via the tokenMKT gateway.',
      authBlockTitle: 'Create the API key file', authBlockDesc: 'Store your tokenMKT key in auth.json so the CLI can load it automatically.',
      verifyBlockTitle: 'Verify the setup', verifyBlockDesc: 'If the CLI opens normally and can continue the session, the setup is working.'
    },
    claude: {
      summaryTitle: 'Claude Code Setup', summary: 'Best for Claude Code and VSCode Claude Code workflows. The key step is switching Anthropic environment variables to tokenMKT.',
      installValue: 'Install Claude Code globally', configValue: 'Set ANTHROPIC_BASE_URL and token', verifyValue: 'Run claude in terminal',
      check1: 'Make sure Node.js is installed and npm is available.', check2: 'Prepare your tokenMKT API key and confirm the gateway domain is reachable.', check3: 'If you use Claude Code in both terminal and VSCode, configure both with the same values.',
      installBlockTitle: 'Install Claude Code', installBlockDesc: 'Install the official CLI first, then point it to tokenMKT.',
      envBlockTitle: 'Set terminal environment variables', envBlockDesc: 'Base URL, auth token, and nonessential traffic control are the key values for terminal usage.',
      ideBlockTitle: 'Configure VSCode Claude Code', ideBlockDesc: 'If you use Claude Code inside VSCode, inject the same environment variables into settings.json.',
      verifyBlockTitle: 'Verify the setup', verifyBlockDesc: 'If Claude Code starts a session successfully, the connection is using tokenMKT.'
    },
    gemini: {
      summaryTitle: 'Gemini CLI Setup', summary: 'Best for Gemini CLI users. This flow is lightweight: install the CLI and export three environment variables.',
      installValue: 'Install Gemini CLI globally', configValue: 'Set Google Gemini base URL and key', verifyValue: 'Run gemini in terminal',
      check1: 'Make sure Node.js 18+ is available on your machine.', check2: 'Prepare your tokenMKT API key and confirm network access to the tokenMKT gateway.', check3: 'If older Gemini environment variables already exist locally, clear them before applying the new setup.',
      installBlockTitle: 'Install Gemini CLI', installBlockDesc: 'Install the CLI first, then route it through the tokenMKT Gemini gateway endpoint.',
      envBlockTitle: 'Set Gemini environment variables', envBlockDesc: 'Set the base URL, API key, and default model so the CLI uses tokenMKT.',
      verifyBlockTitle: 'Verify the setup', verifyBlockDesc: 'If Gemini CLI opens a normal command session, the setup is complete.'
    }
  },

  // Key Usage Query Page
  keyUsage: {
    title: 'API Key Usage',
    subtitle: 'Enter your API Key to view real-time spending and usage status',
    placeholder: 'sk-ant-mirror-xxxxxxxxxxxx',
    query: 'Query',
    querying: 'Querying...',
    privacyNote: 'Your Key is processed locally in the browser and will not be stored',
    dateRange: 'Date Range:',
    dateRangeToday: 'Today',
    dateRange7d: '7 Days',
    dateRange30d: '30 Days',
    dateRange90d: '90 Days',
    dateRangeCustom: 'Custom',
    apply: 'Apply',
    used: 'Used',
    detailInfo: 'Detail Information',
    tokenStats: 'Token Statistics',
    dailyDetail: 'Daily Detail',
    modelStats: 'Model Usage Statistics',
    // Table headers
    date: 'Date',
    model: 'Model',
    requests: 'Requests',
    inputTokens: 'Input Tokens',
    outputTokens: 'Output Tokens',
    cacheCreationTokens: 'Cache Creation',
    cacheReadTokens: 'Cache Read',
    cacheWriteTokens: 'Cache Write',
    totalTokens: 'Total Tokens',
    cost: 'Cost',
    // Status
    quotaMode: 'Key Quota Mode',
    walletBalance: 'Wallet Balance',
    // Ring card titles
    totalQuota: 'Total Quota',
    limit5h: '5-Hour Limit',
    limitDaily: 'Daily Limit',
    limit7d: '7-Day Limit',
    limitWeekly: 'Weekly Limit',
    limitMonthly: 'Monthly Limit',
    // Detail rows
    remainingQuota: 'Remaining Quota',
    expiresAt: 'Expires At',
    todayExpires: '(expires today)',
    daysLeft: '({days} days)',
    usedQuota: 'Used Quota',
    resetNow: 'Resetting soon',
    subscriptionType: 'Subscription Type',
    subscriptionExpires: 'Subscription Expires',
    // Usage stat cells
    todayRequests: 'Today Requests',
    todayInputTokens: 'Today Input',
    todayOutputTokens: 'Today Output',
    todayTokens: 'Today Tokens',
    todayCacheCreation: 'Today Cache Creation',
    todayCacheRead: 'Today Cache Read',
    todayCost: 'Today Cost',
    rpmTpm: 'RPM / TPM',
    totalRequests: 'Total Requests',
    totalInputTokens: 'Total Input',
    totalOutputTokens: 'Total Output',
    totalTokensLabel: 'Total Tokens',
    totalCacheCreation: 'Total Cache Creation',
    totalCacheRead: 'Total Cache Read',
    totalCost: 'Total Cost',
    avgDuration: 'Avg Duration',
    // Messages
    enterApiKey: 'Please enter an API Key',
    querySuccess: 'Query successful',
    queryFailed: 'Query failed',
    queryFailedRetry: 'Query failed, please try again later',
    noDailyUsage: 'No daily usage data',
  },

  // Setup Wizard
  setup: {
    title: 'tokenMKT Setup',
    description: 'Configure your tokenMKT instance',
    database: {
      title: 'Database Configuration',
      description: 'Connect to your PostgreSQL database',
      host: 'Host',
      port: 'Port',
      username: 'Username',
      password: 'Password',
      databaseName: 'Database Name',
      sslMode: 'SSL Mode',
      passwordPlaceholder: 'Password',
      ssl: {
        disable: 'Disable',
        require: 'Require',
        verifyCa: 'Verify CA',
        verifyFull: 'Verify Full'
      }
    },
    redis: {
      title: 'Redis Configuration',
      description: 'Connect to your Redis server',
      host: 'Host',
      port: 'Port',
      username: 'Username (optional)',
      password: 'Password (optional)',
      database: 'Database',
      usernamePlaceholder: 'Leave empty for default user',
      passwordPlaceholder: 'Password',
      enableTls: 'Enable TLS',
      enableTlsHint: 'Use TLS when connecting to Redis (public CA certs)'
    },
    admin: {
      title: 'Admin Account',
      description: 'Create your administrator account',
      email: 'Email',
      password: 'Password',
      confirmPassword: 'Confirm Password',
      passwordPlaceholder: 'Min 8 characters',
      confirmPasswordPlaceholder: 'Confirm password',
      passwordMismatch: 'Passwords do not match'
    },
    ready: {
      title: 'Ready to Install',
      description: 'Review your configuration and complete setup',
      database: 'Database',
      redis: 'Redis',
      adminEmail: 'Admin Email'
    },
    status: {
      testing: 'Testing...',
      success: 'Connection Successful',
      testConnection: 'Test Connection',
      installing: 'Installing...',
      completeInstallation: 'Complete Installation',
      completed: 'Installation completed!',
      redirecting: 'Redirecting to login page...',
      restarting: 'Service is restarting, please wait...',
      timeout: 'Service restart is taking longer than expected. Please refresh the page manually.'
    }
  },

  // Common
}
