/** @type {import("tailwindcss").Config} */
export default {
    content: ["./src/**/*.{js,jsx,ts,tsx}", "./index.html"],
    theme: {
        extend: {
            fontFamily: {
                heading: ["Rubik", "sans-serif"],
                mono: ["Space Mono", "monospace"],
            },
            colors: {
                primary: "#89b4fa",
                secondary: "#b4befe",
                accent: "#f5c2e7",
                neutral: "#181825",
                base: {
                    100: "#1e1e2e",
                    200: "#181825",
                    300: "#313244",
                    content: "#cdd6f4",
                },
                info: "#89dceb",
                success: "#a6e3a1",
                warning: "#f9e2af",
                error: "#f38ba8",
            },
            borderRadius: {
                box: "1rem",
                btn: "0.5rem",
            },
        },
    },
    plugins: [],
}
