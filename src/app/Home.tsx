import type React from "react";
import { Layout } from "./Layout";
import { Card } from "./components/Card";
import { Form } from "./components/Form";

export function Home(): React.JSX.Element {
	return (
		<Layout>
			<div>
				<section>
					{[
						{
							title: "SOME FEATURES ARE MISSING",
							description: "But they'll be coming back soon:",
							items: ["Downloading images."],
						},
						{
							title: "THERE'S MORE COMING",
							description: "I'm working on more features, like:",
							items: [
								"A progress bar.",
								"Better peformance.",
								"Custom characters.",
								"AND SUPPORT FOR VIDEOS.",
							],
						},
					].map((card) => (
						<Card key={card.title} {...card} />
					))}
				</section>
				<br />
				<section>
					<p className="mb-4 -mt-4">
						This is a website to convert any image to ASCII art. It allows you
						to upload an image and convert it to an ASCII text format in a
						heartbeat.
					</p>
					<p className="mb-6">
						At the moment, it supports common image formats like PNG, JPEG, JPG,
						and WEBP.
					</p>
				</section>
				<section>
					<Form />
				</section>
			</div>
		</Layout>
	);
}
