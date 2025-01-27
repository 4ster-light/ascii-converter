/** @type {import("tailwindcss").Config} */
export default {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "./index.html"],
  theme: {
    extend: {
      fontFamily: {
        heading: ["Rubik", "sans-serif"],
        mono: ["Space Mono", "monospace"],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [{
      mocha: {
        "primary": "#89b4fa",       // Catppuccin Blue
        "secondary": "#b4befe",     // Catppuccin Lavender
        "accent": "#f5c2e7",        // Catppuccin Pink
        "neutral": "#181825",       // Catppuccin Mantle
        "base-100": "#1e1e2e",      // Catppuccin Base
        "base-200": "#181825",      // Catppuccin Mantle
        "base-300": "#313244",      // Catppuccin Surface0
        "base-content": "#cdd6f4",  // Catppuccin Text
        "info": "#89dceb",          // Catppuccin Sapphire
        "success": "#a6e3a1",       // Catppuccin Green
        "warning": "#f9e2af",       // Catppuccin Yellow
        "error": "#f38ba8",         // Catppuccin Red
        "--rounded-box": "1rem",
        "--rounded-btn": "0.5rem",
      }
    }],
  },
  darkMode: "class"
}
