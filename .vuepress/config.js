module.exports = {
  // 站点配置

  title: 'GOFamily - go 程序员宝典',
  description: '🔥 go 程序员宝典，包含了：算法，数据库，网络操作系统，分布式，系统设计等一揽子知识体系',
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
    // [
    //   'vuepress-plugin-comment',
    //   {
    //     choosen: 'gitalk',
    //     options: {
    //       clientID: '4865a323784f55266cab',
    //       clientSecret: '16cadc307e5b93bcdc81cf7285098e08342d792e',
    //       repo: 'GOFamily',
    //       owner: 'shgopher',
    //       admin: ['shgopher'],
    //       labels: ["Gitalk", "Comment"],
    //       id: '<%- ("gofamily" + (frontmatter.to.path || window.location.pathname)).slice(-50) %>',
    //       title: '「Comment」<%- window.location.origin + (frontmatter.to.path || window.location.pathname) %>',
    //       body: '<%- window.location.origin + (frontmatter.to.path || window.location.pathname) %>',
    //       distractionFreeMode: false,
    //       pagerDirection: 'last',
    //     }
    //   }
    // ],
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
      // {
      //   text:"推荐项目",
      //   ariaLabel: 'project',
      //   items:[
      //     {
      //       text:"hui 【web框架】",
      //       link:"https://github.com/shgopher/hui",
      //     },
      //     {
      //       text:"key  【服务授权系统】",
      //       link:"https://github.com/shgopher/key",
      //     },
      //     {
      //       text:"ka  【秒杀服务】",
      //       link:"https://github.com/shgopher/ka",
      //     },
      //     {
      //       text:"go-short 【短链接服务】",
      //       link:"https://github.com/shgopher/go-key",
      //     },
      //   ]
      // } 
      // ,
      {
        text:"系列教程",
        ariaLabel: 'Menu',
        items:[
          {
            text:"GOFamily 【go语言教程】",
            link:"https://shgopher.github.io/GOFamily/",
          },
          {
            text:"408  【基础408知识教程】",
            link:"https://shgopher.github.io/408/",
          },
          {
            text:"luban  【系统设计教程】",
            link:"https://shgopher.github.io/luban/",
          },
          {
            text:"dingdang  【工具教程】",
            link:"https://shgopher.github.io/dingdang/",
          },
          {
            text:"god  【给程序员写的书】",
            link:"https://shgopher.github.io/god/",
          },
        ]
      },
      {
        text:'作者',link:'https://shgopher.github.io/',
      },
    ], 
    sidebar:'auto'
  },
}
