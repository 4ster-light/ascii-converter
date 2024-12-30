import { Elysia, t } from "elysia";
import { BodyData, bodySchema } from "../validation/types";

export default function UiRoutes(app: Elysia) {
  app.group("/ui", app =>
    app
      .get("/spinner", () => Spinner, bodySchema)
      .get("/ascii-result", () => AsciiResult, bodySchema)
  )
  return app
}

function Spinner(body: BodyData) {
  return "Spinner"
}

function AsciiResult(body: BodyData) {
  return "Ascii Result"
}