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
        shades: {
          lightBlack: "#1D1D1D",
        },
        brand: {
          background: "#161616",
          primary: "#1D6704",
          primaryLight: "#35C804",
          offWhite: "#C3C3C3",
          darkWhite: "#A7A5A5",
          darkBlue: "#73AEBF",
          darkYellow: "#D5B169",
        },
      },
      boxShadow: {
        brand: "6px 6px 0 #101010",
      },
    },
  },
  plugins: [],
} satisfies Config;
