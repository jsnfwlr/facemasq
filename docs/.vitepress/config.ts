import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'faceMasq',
  description: 'Something',
  themeConfig: {
    siteTitle: 'faceMasq',

    nav: [
      { text: 'About', link: '/' },
      { text: 'User Guide', link: '/guide/introduction' },
      { text: 'Developer Docs', link: '/dev/introduction' },
      { text: 'Changelog', link: 'https://github.com/jsnfwlr/facemasq' }
    ],
    sidebar: [
      {
        text: 'User Guide',
        items: [
          { text: 'Introduction', link: '/guide/introduction' },
          { text: 'Installation', link: '/guide/installation' },
          { text: 'Getting Started', link: '/guide/getting-started' },
          { text: 'Managing Records', link: '/guide/managing-records' },
          { text: 'Advanced Installation', link: '/guide/advanced-installation' },
          { test: 'Understanding Errors', link: '/errors/'}
        ]
      },
      {
        text: 'Developer Docs',
        items: [
          { text: 'Introduction', link: '/dev/introduction' },
          { text: 'Dependencies', link: '/dev/dependencies' },
          { text: 'Contributing', link: '/dev/contributing' },
          { text: 'API Documentation', link: '/dev/api-documentation' },
          { text: 'UI Documentation', link: '/dev/ui-documentation' },
        ]
      }
    ]
  },
})
