import { useState } from "react";
import { type AsciiConfig, DEFAULT_CONFIG } from "./data.ts";
import { calculateDimensions, getImageData, loadImage } from "./utils";

export function useAsciiConverter(config: AsciiConfig) {
	const [result, setResult] = useState<string>("");
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);

	function createAsciiArt(
		imageData: ImageData,
		width: number,
		height: number,
	): string {
		const { preserveColors, resolution, charSet, colorIntensity } = config;
		const step = Math.max(1, Math.floor(15 / resolution)); // More balanced scaling
		const chars = charSet;

		let ascii = "";
		const rgbCache = new Map<string, string>();

		for (let y = 0; y < height; y += step) {
			let line = "";
			for (let x = 0; x < width; x += step) {
				// Get single pixel instead of area average
				const idx = (y * width + x) * 4;
				const r = imageData.data[idx];
				const g = imageData.data[idx + 1];
				const b = imageData.data[idx + 2];

				// Improved brightness calculation
				const brightness = (0.299 * r + 0.587 * g + 0.114 * b) / 255;
				const charIndex = Math.floor(brightness * (chars.length - 1));
				const char = chars[charIndex] || " ";

				if (preserveColors) {
					const key = `${r},${g},${b}`;
					if (!rgbCache.has(key)) {
						const adjusted = adjustColorIntensity(r, g, b, colorIntensity);
						rgbCache.set(key, `rgb(${adjusted.r},${adjusted.g},${adjusted.b})`);
					}
					line += `<span style="color:${rgbCache.get(key)}">${char}</span>`;
				} else {
					line += char;
				}
			}
			ascii += `${line}\n`;
		}
		return ascii;
	}

	function adjustColorIntensity(
		r: number,
		g: number,
		b: number,
		intensity: number,
	) {
		const factor = 1 + intensity / 10;
		return {
			r: Math.min(255, r * factor),
			g: Math.min(255, g * factor),
			b: Math.min(255, b * factor),
		};
	}

	async function convertImageToAscii(file: File): Promise<void> {
		setLoading(true);
		setError(null);

		try {
			const imageUrl = URL.createObjectURL(file);
			const img = await loadImage(imageUrl);

			const { width, height } = calculateDimensions(
				img.width,
				img.height,
				config.maxWidth || DEFAULT_CONFIG.maxWidth,
				config.maxHeight || DEFAULT_CONFIG.maxHeight,
			);

			const imageData = getImageData(img, width, height);
			const ascii = createAsciiArt(
				imageData,
				width,
				height,
				// config.preserveColors,
			);

			setResult(ascii);
		} catch (err) {
			setError(err instanceof Error ? err.message : "Failed to convert image");
		} finally {
			setLoading(false);
		}
	}

	return {
		result,
		loading,
		error,
		convertImageToAscii,
		clearResult: () => setResult(""),
	};
}
