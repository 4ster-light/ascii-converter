import React from 'react';
import { Card, CardProps } from './components/Card';
import { Form } from './components/Form';
import { Layout } from './Layout';

export const Home: React.FC = () => {
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
  ];

  return <Layout>
    <div>
      <section>
        {cards.map((card, index) => (
          <Card key={index} {...card} />
        ))}
      </section>
      <br />
      <section>
        <p className="mb-4 -mt-4">
          This is a website to convert any image to ASCII art. It allows you to upload an image and convert it to an ASCII text format in a heartbeat.
        </p>
        <p className="mb-6">
          At the moment, it supports common image formats like PNG, JPEG, JPG, and WEBP.
        </p>
      </section>
      <section>
        <Form />
      </section>
    </div>
  </Layout>
}
