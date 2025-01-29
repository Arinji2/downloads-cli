import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          background: "#161616",
          primary: "#2C9E05",
          offWhite: "#C3C3C3",
          darkWhite: "#656363",
          darkBlue: "#448293",
          lightBlue: "#1F74CF",
          darkYellow: "#D5B169",
          lightYellow: "#CEAF10",
          grey: "#929292",
        },
      },
      boxShadow: {
        "terminal-shadow": "7px 7px 3px #1B1B1B",
      },
    },
  },
  plugins: [],
} satisfies Config;
