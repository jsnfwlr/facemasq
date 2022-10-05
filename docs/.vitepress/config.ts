import { defineConfig } from "vitepress"

import en from "../en/sidebar.json"
import fr from "../fr/sidebar.json"
import { version } from "../../package.json"

const langs = {
  "en": {
    text: "English",
    link: "/en/",
    sidebar: en
  },
  "fr": {
    text: "French",
    link: "/fr/",
    sidebar: fr
  }
}


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
  const lang = "en"
  return [
    {
      text: version,
      items: [
        {
          text: "Changelog",
          link: "https://github.com/jsnfwlr/facemasq/blob/main/CHANGELOG.md"
        },
        {
          text: "Contributing",
          link: "https://github.com/jsnfwlr/facemasq/blob/main/.github/contributing.md"
        },
        {
          text: "License",
          link: "https://github.com/jsnfwlr/facemasq/blob/main/LICENSE"
        }
      ]
    }
  ]
}

function sidebar() {
  const sidebar = {}
  const langList = Object.keys(langs)
  langList.forEach((lang) => {
    sidebar[langs[lang].link] = langs[lang].sidebar
  })
  return sidebar
}


function locales() {
  const langList = Object.keys(langs)
  const locales = {
    text: "",
    items: []
  }
  langList.forEach((lang) => {
    locales.items.push({
      text: langs[lang].text,
      link: langs[lang].link
    })
  })
  return locales

}
