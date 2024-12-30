import { t } from "elysia"

export interface BodyData {
  image: File,
  preserveColors: boolean
}

export const bodySchema = {
  body: t.Object({
    file: t.File(),
    preserveColors: t.Boolean()
  })
} as const
