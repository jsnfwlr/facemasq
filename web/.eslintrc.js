module.exports = {
  root: true,
  env: {
    node: true,
  },
  parser: "vue-eslint-parser",
  parserOptions: {
    parser: "@typescript-eslint/parser",
    ecmaVersion: 2020,
    sourceType: "module",
  },
  plugins: ["prettier", "@typescript-eslint"],
  extends: [
    // "plugin:vue/vue3-essential",
    // "eslint:recommended",
    // "@vue/typescript/recommended",
    // "@vue/prettier",
    // "@vue/prettier/@typescript-eslint",
    "plugin:@typescript-eslint/recommended",
    "eslint:recommended",
    "plugin:prettier/recommended",
    "plugin:vue/vue3-essential",
  ],
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
    "arrow-parens": "off",
    "eol-last": "error",
    semi: ["error", "never"],
    quotes: [2, "double"],
    "vue/max-attributes-per-line": [
      "error",
      {
        singleline: {
          max: 100,
        },
        multiline: {
          max: 100,
        },
      },
    ],
    "vue/max-len": [
      "error",
      {
        code: 600,
        template: 600,
      },
    ],
    "vue/no-multiple-template-root": "off",
    "vue/multi-word-component-names": "off",
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": "error",
  },
  ignorePatterns: ["dist"],
}
