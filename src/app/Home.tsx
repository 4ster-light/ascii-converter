import type { ReactElement } from "react";
import { Layout } from "./Layout";
import { Form } from "./components/Form";

export function Home(): ReactElement {
	return (
		<Layout>
			<div className="space-y-6">
				<section className="prose prose-invert max-w-none text-center">
					<h2 className="text-xl font-semibold">Image to ASCII Converter</h2>
					<p className="text-base-content/80">
						Transform any image into ASCII art instantly. Upload your image and
						customize the conversion settings below.
					</p>
				</section>
				<section>
					<Form />
				</section>
			</div>
		</Layout>
	);
}
