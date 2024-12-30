import { Html } from "@elysiajs/html"
import Layout from "../layouts/layout"
import Form from "../components/form"
import Card, { CardProps } from "../components/card"

export function Index() {
  const cards: CardProps[] = [
    {
      title: "SOME FEATURES ARE MISSING",
      description: "But they'll be coming back soon:",
      items: ["Downloading images."]
    },
    {
      title: "THERE'S MORE COMING",
      description: "I'm working on more features, like:",
      items: [
        "A progress bar.",
        "Better peformance.",
        "Custom characters.",
        "AND SUPPORT FOR VIDEOS."
      ]
    }
  ]

  return <Layout>
    <div>
      <section>
        {cards.map(cardProps => <Card {...cardProps} />)}
      </section>
      <br />
      <section>
        <p class="mb-4 -mt-4">
          This is a website to convert any image to ASCII art. It allows you to upload an image and convert it to an ASCII text format in a heartbeat.
        </p>
        <p class="mb-6">
          At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
        </p>
      </section>
      <section>
        <Form />
      </section>
    </div>
  </Layout>
}
