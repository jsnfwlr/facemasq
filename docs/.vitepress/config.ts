import { defineConfig } from 'vitepress'

import { version } from '../../web/package.json'


export default defineConfig({
  title: 'faceMasq',
  description: 'Something',
  themeConfig: {
    siteTitle: 'faceMasq',
    logo: "/logo.png",
    nav: nav(),
    sidebar: {
      '/guide/': userSidebar(),
      '/dev/': devSidebar()
    }
  },
})

function nav() {
  return [
      { 
        text: 'About', 
        link: '/' 
      },
      { 
        text: 'User Guide', 
        link: '/guide/introduction' 
      },
      { 
        text: 'Developer Docs', 
        link: '/dev/introduction'
      },
      { 
        text: version, 
        items: [
          { 
            text: "Changelog",
            link: 'https://github.com/jsnfwlr/facemasq/blob/main/CHANGELOG.md' 
          },
          { 
            text: "Contributing",
            link: 'https://github.com/jsnfwlr/facemasq/blob/main/.github/contributing.md' 
          },
          { 
            text: "License",
            link: 'https://github.com/jsnfwlr/facemasq/blob/main/LICENSE' 
          }
        ]
        
      }

  ]
}

function userSidebar() {
 return [
    {
      text: 'User Guide',
      collapsible: true,
      items: [
        { text: 'Introduction', link: '/guide/introduction' },
        { text: 'Installation', link: '/guide/installation' },
        { text: 'Getting Started', link: '/guide/getting-started' },
        { text: 'Managing Records', link: '/guide/managing-records' },
        { text: 'Advanced Installation', link: '/guide/advanced-installation' },
        { test: 'Understanding Errors', link: '/errors/'}
      ]
    }
  ]
}

function devSidebar() {
  return [
    {
      text: 'Developer Docs',
      collapsible: true,
      items: [
        { text: 'Introduction', link: '/dev/introduction' },
        { text: 'Dependencies', link: '/dev/dependencies' },
        { text: 'Contributing', link: '/dev/contributing' },
        { text: 'API Documentation', link: '/dev/api-documentation' },
        { text: 'UI Documentation', link: '/dev/ui-documentation' },
      ]
    }
  ]
}
