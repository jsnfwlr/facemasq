import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/pages/dashboard.vue'
import { defineAsyncComponent } from 'vue'

const routes = [
  {
    meta: {
      title: 'Dashboard',
      layout: 'Standard'
    },
    path: '/',
    name: 'home',
    component: Home
  },
  {
    meta: {
      title: 'Devices',
      layout: 'Standard'
    },
    path: '/devices',
    name: 'Devices',
    component: defineAsyncComponent(() => import('@/pages/devices.vue'))
  },
  {
    meta: {
      title: 'Manage params',
      layout: 'Standard'
    },
    path: '/manage/:param',
    name: 'Manage',
    component: defineAsyncComponent(() => import('@/pages/params.vue'))
  },
  {
    meta: {
      title: 'Manage params',
      layout: 'Standard'
    },
    path: '/admin/:param',
    name: 'Admin',
    component: defineAsyncComponent(() => import('@/pages/params.vue'))
  },
  {
    meta: {
      title: 'Information',
      layout: 'Standard'
    },
    path: '/about/info',
    name: 'Info',
    component: defineAsyncComponent(() => import('@/pages/info.vue'))
  }
]


const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { top: 0 }
  },
  linkActiveClass: "active",
  linkExactActiveClass: "exact"
})

const defaultDocumentTitle = 'faceMasq'


router.afterEach(to => {
  /* Set document title from route meta */
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

export default router
