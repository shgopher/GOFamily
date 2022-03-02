import { defineUserConfig } from 'vuepress'
import type { DefaultThemeOptions } from 'vuepress'

export default defineUserConfig<DefaultThemeOptions>({
  // 站点配置
  lang: 'zh-CN',
  title: 'GOFamily - go 后端程序员宝典',
  description: '🔥 go 后端程序员宝典，包含了：算法，数据库，网络操作系统，分布式，系统设计等一揽子知识体系',
  head: [
    ['link', {rel: 'shortcut icon', type: "image/x-icon", href: `/favicon.ico`}]
  ],
  host: 'localhost',
  base:'/GOFamily/',
  port: 8080,
  dest: '.vuepress/dist',
  plugins:[
    ['@vuepress/container',
      {
        type: 'right',
        defaultTitle: ''
      }
    ],
    ['@vuepress/container',
      {
        type: 'center',
        defaultTitle: ''
      }
    ],
    ['@vuepress/container',
      {
        type: 'quote',
        before: info => `<div class="quote"><p class="title">${info}</p>`,
        after: '</div>'
      },
    ],
    ['@vuepress/container',
      {
        type: 'not-print',
        defaultTitle: ''
      },
    ],
    ['@vuepress/plugin-prismjs',
      {
        preloadLanguages:['markdown', 'jsdoc', 'yaml'],
      }
    ],
  ],
  markdown: {
    anchor: false,
    toc: {level: [2, 3]},
  },
  // 主题和它的配置
  theme: '@vuepress/theme-default',
  themeConfig: {
    logo: 'https://avatars.githubusercontent.com/u/42873232',
    lastUpdated: true,
    smoothScroll: true,
    // repo: 'fenixsoft/awesome-fenix',
    editLinks: true,
    editLinkText: '在GitHub中',
    // 添加导航栏
    navbar: [
      {
        text: '首页', link: '/'
      }, 
    ],
    sidebar:[
      {
        text:'GO',
        children:[
            {
              text:'基础',
              link:'/go/基础',
              collapsible:false,
              children:[
                {
                  text:'数字类型',
                  link:'/go/基础/数字类型',
                },
                {
                  text:'slice',
                  link:'/go/基础/slice',
                },
                {
                  text:'string',
                  link:'/go/基础/string',
                },
                {
                  text:'map',
                  link:'/go/基础/map',
                },
                {
                  text:'其它类型',
                  link:'/go/基础/其它类型',
                },
                {
                  text:'关键字',
                  link:'/go/基础/关键字',
                },
                {
                  text:'函数方法',
                  link:'/go/基础/函数方法',
                },
                {
                  text:'接口',
                  link:'/go/基础/interface',
                },
                {
                  text:'逻辑和判断语句',
                  link:'/go/基础/逻辑和判断语句',
                },
                {
                  text:'泛型',
                  link:'/go/基础/泛型',
                },
              ]
            },
            {
              text:'并发',
              link:'/go/并发',
              collapsible:false,
              children:[
                {
                  text:'并发模型',
                  link:'/go/并发/并发模型',
                },
                {
                  text:'并发模型',
                  link:'/go/并发/并发模型',
                },
                {
                  text:'并发原语',
                  link:'/go/并发/并发原语',
                },
                {
                  text:'内存模型',
                  link:'/go/并发/内存模型',
                },
                {
                  text:'atomic',
                  link:'/go/并发/atomic',
                },
                {
                  text:'channel',
                  link:'/go/并发/channel',
                },
                {
                  text:'context',
                  link:'/go/并发/context',
                },
              ]
            },
            {
              text:'runtime',
              link:'/go/runtime',
              collapsible:false,
              children:[
                  {
                    text:'G:M:P',
                    link:'/go/runtime/gmp',
                  },
                  {
                    text:'netpool',
                    link:'/go/runtime/netpool',
                  },
                  {
                    text:'栈内存管理',
                    link:'/go/runtime/栈内存管理',
                  },
                  {
                    text:'堆内存分配',
                    link:'/go/runtime/堆内存分配',
                  },
                  {
                    text:'系统监控',
                    link:'/go/runtime/系统监控',
                  },
                  {
                    text:'三色gc算法',
                    link:'/go/runtime/三色gc算法',
                  },
                  {
                    text:'定时器',
                    link:'/go/runtime/定时器',
                  },
                  
                ]
            },
            {
              text:'工程',
              link:'/go/工程',
              collapsible:false,
              children:[
                {
                  text:'包管理工具',
                  link:'/go/工程/包管理工具',
                },
                {
                  text:'测试',
                  link:'/go/工程/测试',
                },
                {
                  text:'错误处理',
                  link:'/go/工程/错误处理',
                },
                {
                  text:'动态调试',
                  link:'/go/工程/动态调试',
                },
                {
                  text:'反射',
                  link:'/go/工程/反射',
                },
                {
                  text:'go自带命令',
                  link:'/go/工程/命令',
                },
                {
                  text:'性能剖析',
                  link:'/go/工程/性能剖析',
                },
                {
                  text:'优秀第三方包',
                  link:'/go/工程/优秀第三方包',
                },
                {
                  text:'cgo',
                  link:'/go/工程/cgo',
                },
                {
                  text:'golint',
                  link:'/go/工程/golint',
                },
                {
                  text:'wasm in go',
                  link:'/go/工程/wasm',
                },
                {
                  text:'go web包',
                  link:'/go/工程/web',
                },
              ]
            }
        ]
    },
    {
      text:'408',
      children:[
          {
            text:'算法和数据结构',
              link:'/408/算法',
              collapsible:false,
              children:[
                
              ]
          },
          {
            text:'网络',
              link:'/408/网络',
              collapsible:false,
              children:[
                
              ]
          },
          {
            text:'操作系统',
              link:'/408/操作系统',
              collapsible:false,
              children:[
                
              ]
          },
          {
            text:'数据库',
              link:'/408/数据库',
              collapsible:false,
              children:[
                
              ]
          },
          {
            text:'设计模式',
              link:'/408/设计模式',
              collapsible:false,
              children:[
                
              ]
          },
          {
            text:'计算机组成原理',
              link:'/408/组成原理',
              collapsible:false,
              children:[
                
              ]
          },
      ]
    },
    {
      text:'系统设计',
      children:[
        {
          text:'架构设计基础',
            link:'/system/架构设计基础',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'容器',
            link:'/system/容器',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'网关',
            link:'/system/网关',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'消息队列',
            link:'/system/消息队列',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'rpc',
            link:'/system/rpc',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'搜索引擎',
            link:'/system/搜索引擎',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'实战',
            link:'/system/实战',
            collapsible:false,
            children:[
              
            ]
        },
      ]
    },
    {
      text:'技术方向',
      children:[
        {
          text:'后端',
            link:'/tech/后端开发',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'前端',
            link:'/tech/前端开发',
            collapsible:false,
            children:[
              
            ]
        },
      ]
    },
    {
      text:'其它',
      children:[
        {
          text:'测试',
            link:'/other/测试',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'正则表达式',
            link:'/other/正则表达式',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'git',
            link:'/other/git',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'auth',
            link:'/other/auth',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'devOps',
            link:'/other/devOps',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'重构技术',
            link:'/other/重构技术',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'程序员',
            link:'/other/程序员',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'ddd',
            link:'/other/ddd',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'okr',
            link:'/other/okr',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'性能优化技术',
            link:'/other/性能优化技术',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'debug 技术',
            link:'/other/debug',
            collapsible:false,
            children:[
              
            ]
        },
        {
          text:'线上快速排障',
            link:'/other/线上快速排障',
            collapsible:false,
            children:[
              
            ]
        },
      ]
    }
    ],
  },
})

