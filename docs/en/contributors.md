<script setup>
import { VPTeamMembers } from 'vitepress/theme'

const members = [
  {
    avatar: 'https://www.github.com/jsnfwlr.png',
    name: 'Jason Fowler',
    title: 'Creator',
    links: [
      { icon: 'github', link: 'https://github.com/jsnfwlr' },
      { icon: 'twitter', link: 'https://twitter.com/jsnfwlr' }
    ]
  }
]
</script>

# Contributors

<VPTeamMembers size="small" :members="members" />
