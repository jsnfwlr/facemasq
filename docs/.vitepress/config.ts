import { defineConfig } from "vitepress"

import fr from "../fr/sidebar.json"
import en from "../en/sidebar.json"
import es from "../es/sidebar.json"
import zh from "../zh/sidebar.json"
import { version } from "../../package.json"

const langs = [
  // {
  //   key: null,
  //   text: "Most Used",
  //   link: null
  // },
  {
    key: "en",
    text: "English",
    link: "/en/",
    sidebar: en
  },
  {
    key: "es",
    text: "Español",
    link: "/es/",
    sidebar: es
  },
  {
    key: "fr",
    text: "Français",
    link: "/fr/",
    sidebar: fr
  },
  {
    key: "zh",
    text: "中文",
    link: "/zh/",
    sidebar: zh
  },
  // {
  //   key: null,
  //   text: "Other",
  //   link: null
  // }
]


export default defineConfig({
  title: "faceMasq",
  description: "Something",
  lang: "en",
  themeConfig: {
    siteTitle: "faceMasq",
    logo: "/logo.png",
    nav: nav(),
    sidebar: sidebar(),
    localeLinks: locales()
  },

  locales: {
    fr: {
      lang: "fr",
      title: "faceMasq",
    }
  },
  cleanUrls: "with-subfolders"
})

function nav() {

  return [
    {
      text: version,
      link: null
    }
  ]
}

function sidebar() {
  const sidebar = {}
  const langList = Object.keys(langs)
  langs.forEach((lang) => {
    if (lang.key !== null) {
      sidebar[lang.link] = lang.sidebar
    }
  })
  return sidebar
}

type localeLink = {
  text: string
  link: string | null
}

function locales() {
  const locales = {
    text: "",
    items: [] as Array<localeLink>
  }
  langs.forEach((lang) => {
    locales.items.push({
      text: lang.text,
      link: lang.link
    })
  })
  return locales

}
