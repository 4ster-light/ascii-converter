/** @type {import("tailwindcss").Config} */
export default {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "./index.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        artscii: {
          "primary": "#fb923c",           // orange-400 (vibrant orange)
          "secondary": "#fdba74",         // orange-300 (soft secondary)
          "accent": "#fff7ed",            // orange-50 (very soft accent)
          "neutral": "#44403c",           // stone-700 (softer neutral)
          "base-100": "#fff7ed",          // orange-50 (warm white)
          "base-200": "#ffedd5",          // orange-100 (soft bg)
          "base-300": "#fed7aa",          // orange-200 (deeper elements)
          "base-content": "#292524",      // stone-800 (softer text)
          "--rounded-box": "0.75rem",     // border radius
          "--rounded-btn": "0.5rem",      // button radius
        },
        dark: {
          "primary": "#fb923c",           // orange-400 (brighter for dark mode)
          "secondary": "#ea580c",         // orange-600
          "accent": "#3f3f46",            // zinc-700
          "neutral": "#18181b",           // zinc-900
          "base-100": "#27272a",          // zinc-800
          "base-200": "#18181b",          // zinc-900
          "base-300": "#3f3f46",          // zinc-700
          "base-content": "#fafafa",      // zinc-50
          "--rounded-box": "0.75rem",
          "--rounded-btn": "0.5rem",
        }
      },
    ],
  },
  darkMode: "class"
}
