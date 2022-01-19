module.exports = {
  "parser": "@typescript-eslint/parser",
  "plugins": [
    "@typescript-eslint"
  ],
  " ": [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/typescript/recommended",
    "@vue/prettier",
    "@vue/prettier/@typescript-eslint",
  ],
  "rules": {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
    "arrow-parens": "off",
    "eol-last": "error",
    "semi": [
      "error",
      "never"
    ],
    "quotes": [
      2,
      "double"
    ],
    "vue/max-attributes-per-line": ["error", {
      "singleline": {
        "max": 100
      },
      "multiline": {
        "max": 100
      }
    }],
    "vue/max-len": [
      "error",
      {
        "code": 600,
        "template": 600,
      }
    ],
    "vue/no-multiple-template-root": "off"
  }
}