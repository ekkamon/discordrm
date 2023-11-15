import type { Config } from "tailwindcss"

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        sky: "#4BA5F5",
        blurple: {
          DEFAULT: "#5865F2",
          600: "#4951BD",
          700: "#4147A3",
        },
        grayish: "#23272A",
        almost: "#1E1F22",
        dark: "#313338",
      },
    },
  },
  plugins: [],
}
export default config
