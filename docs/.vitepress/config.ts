import { defineConfig } from "vitepress"

import fr from "../fr/definition"
import en from "../en/definition"
import es from "../es/definition"
import zh from "../zh/definition"
import { version } from "../../package.json"

const langs = [
  en,
  fr,
  es,
  zh,
]


export default defineConfig({
  title: "faceMasq",
  description: "Something",
  lang: "en",
  outDir: "./dist/docs/",
  themeConfig: {
    siteTitle: "faceMasq",
    logo: "/logo.png",
    nav: nav(),
    sidebar: sidebar(),
    localeLinks: localeLinks(),
    outlineTitle: langs[detectLocale()].outlineTitle,
    editLink: {
      pattern: "https://github.com/jsnfwlr/facemasq/edit/main/docs/:path",
      text: langs[detectLocale()].editLinkText
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/jsnfwlr/facemasq' },
    ]
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
      link: ""
    }
  ]
}

function sidebar() {
  const sidebar = {}
  langs.forEach((lang) => {
    if (lang.key !== null) {
      sidebar[lang.link] = lang.sidebar
    }
  })
  return sidebar
}

type LocaleLink = {
  text: string
  link: string
}

type Locales = {
  lang: string
  text: string
}

function localeLinks() {
  const locales = {
    text: "",
    items: [] as LocaleLink[]
  }
  langs.forEach((lang) => {
    locales.items.push({
      text: lang.text,
      link: lang.link
    })
  })
  return locales

}

function locales() {
  const locales = [] as Locales[]
  // {
  //   lang: "",
  //   title: ""
  // }
  langs.forEach((lang) => {
    locales.push({
      lang: lang.key,
      text: "faceMasq " + lang.text
    })
  })
  return locales
}

function detectLocale() {
  const index = 0
  // const preference = navigator.language.substring(0, 1)
  // let index = langs.findIndex(lang => lang.key === preference)
  return index
}
