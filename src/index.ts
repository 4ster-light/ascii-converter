import { Elysia, t } from "elysia"
import { html } from "@elysiajs/html"
import { tailwind } from "@gtramontina.com/elysia-tailwind";
import { staticPlugin } from "@elysiajs/static";
import { swagger } from "@elysiajs/swagger";
import { Index } from "../templates/views";
import ApiRoutes from "./routes/api";
import UiRoutes from "./routes/ui";

const TITLE = "✰ArtSCII✰"
const PORT = process.env.PORT || 3000

new Elysia()
  .use(html())
  .use(staticPlugin())
  .use(tailwind({
    path: "/public/css/index.css",
    source: "./templates/index.css",
    config: "./tailwind.config.js",
    options: { minify: true }
  }))
  .use(swagger({
    path: "/docs",
    exclude: ["/", /^\/ui\/.*/],
    documentation: {
      info: {
        title: TITLE + " API",
        description: "From image to ascii art converter",
        version: "1.0.0"
      }
    },
  }))
  .get("/", () => Index())
  .use(ApiRoutes)
  .use(UiRoutes)
  .listen(PORT)

console.log("Server running at http://localhost:" + PORT)
