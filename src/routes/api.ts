import { Elysia, t } from "elysia";
import generateAsciiArt from "../lib/ascii_text";
import generateImageFromAsciiArt from "../lib/ascii_image";
import { BodyData, bodySchema } from "../validation/types";

export default function ApiRoutes(app: Elysia) {
  app.group("/api", app =>
    app
      .post("/convert-to-ascii", convertToAscii, bodySchema)
      .post("/convert-to-image", convertToImage, bodySchema)
  )
  return app
}

async function convertToAscii(body: BodyData) {
  const { image, preserveColors } = body

  try {
    const asciiArt = await generateAsciiArt(image, preserveColors)
    return {
      success: true,
      asciiArt
    }
  } catch (error) {
    return {
      success: false,
      error
    }
  }
}

async function convertToImage(body: BodyData) {
  const { image, preserveColors } = body

  try {
    const asciiArt = await generateAsciiArt(image, preserveColors)
    const newImage = await generateImageFromAsciiArt(asciiArt)

    return {
      success: true,
      newImage
    }
  } catch (error) {
    return {
      success: false,
      error
    }
  }
}
