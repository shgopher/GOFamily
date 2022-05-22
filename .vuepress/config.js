module.exports = {
  // 站点配置

  title: 'GOFamily - go 后端程序员宝典',
  description: '🔥 go 后端程序员宝典，包含了：算法，数据库，网络操作系统，分布式，系统设计等一揽子知识体系',
  head: [
    ['link', {rel: 'shortcut icon', type: "image/x-icon", href: `/favicon.ico`}],
  ],
  host: 'localhost',
  base:'/GOFamily/',
  port: 8080,
  dest: '.vuepress/dist',
  plugins: [
    ['vuepress-plugin-container',
      {
        type: 'right',
        defaultTitle: ''
      }
    ],
    ['vuepress-plugin-container',
      {
        type: 'center',
        defaultTitle: ''
      }
    ],
    ['vuepress-plugin-container',
      {
        type: 'quote',
        before: info => `<div class="quote"><p class="title">${info}</p>`,
        after: '</div>'
      },
    ],
    ['vuepress-plugin-container',
      {
        type: 'not-print',
        defaultTitle: ''
      },
    ],
    [
      '@vuepress/google-analytics',
      {
        'ga': 'G-GFKQEFHX3B'
      }
    ],
    [
      'vuepress-plugin-comment',
      {
        choosen: 'gitalk',
        options: {
          clientID: '4865a323784f55266cab',
          clientSecret: '16cadc307e5b93bcdc81cf7285098e08342d792e',
          repo: 'GOFamily',
          owner: 'shgopher',
          admin: ['shgopher'],
          labels: ["Gitalk", "Comment"],
          id: '<%- ("gofamily" + (frontmatter.to.path || window.location.pathname)).slice(-50) %>',
          title: '「Comment」<%- window.location.origin + (frontmatter.to.path || window.location.pathname) %>',
          body: '<%- window.location.origin + (frontmatter.to.path || window.location.pathname) %>',
          distractionFreeMode: false,
          pagerDirection: 'last',
        }
      }
    ],
    ['@vuepress/back-to-top'],
    ['@vuepress/nprogress'],
    'vuepress-plugin-baidu-autopush',
    ['vuepress-plugin-baidu-tongji-analytics', {
      key: '45951f610a1fa82985715b79291a8de9'
    }],
  ],
  markdown: {
    anchor: {permalink: false},
    toc: {includeLevel: [2, 3]},
  },
  // 主题和它的配置
  theme: '@vuepress/theme-default',
  themeConfig: {
    logo: 'https://avatars.githubusercontent.com/u/42873232',
    lastUpdated: '最后更新',
    smoothScroll: true,
    editLinks: true,
    repo: 'https://github.com/shgopher/GOFamily',
    docsBranch: 'release',
    editLinkText: '在GitHub中编辑',
    // 添加导航栏
    nav: [
      {
        text: '首页', link: '/'
      }, 
      {
        text:'Github',link:'https://github.com/shgopher/GOFamily',
      },
      {
        text:'微信公众号',link:'/#wechat.png',
      },
      {
        text:'作者',link:'https://github.com/shgopher',
      },
    ], 
    sidebar:[
      {
        title: 'GO',
        collapsable: false,
        children: [
          {
            title:"基础",
            collapsable: true,
            path:'/go/基础/',
            children: [
              {
                title: '变量声明',
                path: '/go/基础/变量声明/',
                collapsable: false,
              },
              {
                title: '常量声明',
                path: '/go/基础/常量声明/',
                collapsable: false,
              },
              {
                title: '零值',
                path: '/go/基础/零值/',
                collapsable: false,
              },
              {
                title: '复合字面量',
                path: '/go/基础/复合字面量/',
                collapsable: false,
              },
              {
                title: '数字类型',
                path: '/go/基础/数字类型/',
                collapsable: false,
              },
              {
                title: 'slice',
                path: '/go/基础/slice/',
                collapsable: false,
              },
              {
                title: 'string',
                path: '/go/基础/string/',
                collapsable: false,
              },
              {
                title: 'map',
                path: '/go/基础/map/',
                collapsable: false,
              },
              {
                title: '函数方法',
                path: '/go/基础/函数方法/',
                collapsable: false,
              },
              {
                title: 'interface',
                path: '/go/基础/interface/',
                collapsable: false,
              },
              {
                title: '逻辑和判断语句',
                path: '/go/基础/逻辑和判断语句/',
                collapsable: false,
              },
              {
                title: '关键字',
                path: '/go/基础/关键字/',
                collapsable: false,
              },
              {
                title: '泛型',
                path: '/go/基础/泛型/',
                collapsable: false,
              },
              {
                title: '其他内容',
                path: '/go/基础/其他内容/',
                collapsable: false,
              }
            ]
          },
          {
            title: '并发',
            collapsable: true,
            path:'/go/并发/',
            children:[
                {
                  title: '并发原语',
                  collapsable: false,
                  path:'/go/并发/并发原语/',
                },
                {
                  title: 'channel',
                  collapsable: false,
                  path:'/go/并发/channel/',
                },
                {
                  title: 'atomic',
                  collapsable: false,
                  path:'/go/并发/atomic/',
                },
                {
                  title: 'context',
                  collapsable: false,
                  path:'/go/并发/context/',
                },
                {
                  title: '并发模型',
                  collapsable: false,
                  path:'/go/并发/并发模型/',
                },
                {
                  title: '内存模型',
                  collapsable: false,
                  path:'/go/并发/内存模型/',
                },

            ],
          },
          {
            title: 'runtime',
            collapsable: true,
            path:'/go/runtime/',
            children:[
              {
                title: '三色gc算法',
                collapsable: false,
                path:'/go/runtime/三色gc算法/',
              },
              {
                title: '堆内存分配',
                collapsable: false,
                path:'/go/runtime/堆内存分配/',
              },
              {
                title: '栈内存管理',
                collapsable: false,
                path:'/go/runtime/栈内存管理/',
              },
              {
                title: '系统监控',
                collapsable: false,
                path:'/go/runtime/系统监控/',
              },
              {
                title: 'G:M:P',
                collapsable: false,
                path:'/go/runtime/gmp/',
              },
              {
                title: '定时器',
                collapsable: false,
                path:'/go/runtime/定时器/',
              },
              {
                title: 'netpool',
                collapsable: false,
                path:'/go/runtime/netpool/',
              }
            ],
          },
          {
            title: '工程',
            collapsable: true,
            path:'/go/工程/',
            children:[
              {
                title: '测试',
                collapsable: false,
                path:'/go/工程/测试/',
              },
              {
                title: '性能剖析',
                collapsable: false,
                path:'/go/工程/性能剖析/',
              },
              {
                title: '包管理工具',
                collapsable: false,
                path:'/go/工程/包管理工具/',
              },
              {
                title: '动态调试',
                collapsable: false,
                path:'/go/工程/动态调试/',
              },
              {
                title: '错误处理',
                collapsable: false,
                path:'/go/工程/错误处理/',
              }, 
              {
                title: 'cgo',
                collapsable: false,
                path:'/go/工程/cgo/',
              },
              {
                title: 'golint',
                collapsable: false,
                path:'/go/工程/golint/',
              },
              {
                title: '反射',
                collapsable: false,
                path:'/go/工程/反射/',
              },
              {
                title: 'web',
                collapsable: false,
                path:'/go/工程/web/',
              },
              {
                title: 'wasm',
                collapsable: false,
                path:'/go/工程/wasm/',
              },
              {
                title: '命令',
                collapsable: false,
                path:'/go/工程/命令/',
              },
              {
                title: '优秀第三方包',
                collapsable: false,
                path:'/go/工程/优秀第三方包/',
              },
              {
                title: 'go标准库',
                collapsable: false,
                path:'/go/工程/go标准库/',
              },
              {
                title: '项目组织形式',
                collapsable: false,
                path:'/go/工程/项目组织形式/',
              },
              {
                title: 'go命名惯例',
                collapsable: false,
                path:'/go/工程/go命名惯例/',
              },

            ],
          },
        ]
      },
      {
        title: '408',
        collapsable: false,
        children: [
          {
            title:'算法',
            // path:'/408/并发/',
            collapsable:true,
            children:[],
          },
          {
            title:'网络',
            collapsable:true,
            //path:'/408/网络/',
            children:[],
          },
          {
            title:'数据库',
            collapsable:true,
            // path:'/408/数据库/',
            children:[],
          },
          {
            title:'操作系统',
            collapsable:true,
            // path:'/408/操作系统/',
            children:[],
          },
          {
            title:'设计模式',
            collapsable:true,
            // path:'/408/设计模式/',
            children:[],
          },
          {
            title:'组成原理',
            collapsable:true,
            // path:'/408/组成原理/',
            children:[],
          },
        ],
      },
      {
        title: '架构设计',
        collapsable: false,
        children: [],
      },
      {
        title: '技术方向',
        collapsable: false,
        children: [],
      },
      {
        title: '其它',
        collapsable: false,
        children: [],
      },
    ],
  },
}
