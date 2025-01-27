import { useState } from "react";
import { ASCII_CHARS, type AsciiConfig } from "./data.ts";
import {
	calculateDimensions,
	getImageData,
	getPixelBrightness,
	loadImage,
} from "./utils";

const DEFAULT_CONFIG: AsciiConfig = {
	preserveColors: false,
	maxWidth: 100,
	maxHeight: 100,
	resolution: 100,
	charSet: ASCII_CHARS,
	colorIntensity: 0,
};

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
		const step = Math.max(1, Math.ceil((width * height) / (resolution * 1000)));
		const chars = charSet;

		let ascii = "";
		const rgbCache = new Map<string, string>();

		for (let y = 0; y < height; y += step) {
			let line = "";
			for (let x = 0; x < width; x += step) {
				const { r, g, b } = samplePixelArea(
					imageData,
					x,
					y,
					step,
					width,
					height,
				);
				const brightness = getPixelBrightness(r, g, b);
				const char = chars[Math.floor(brightness * (chars.length - 1))];

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
			ascii += `${line}\r\n`;
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

	function samplePixelArea(
		imageData: ImageData,
		x: number,
		y: number,
		step: number,
		width: number,
		height: number,
	) {
		// Improved sampling with weighted average
		let totalR = 0;
		let totalG = 0;
		let totalB = 0;
		let count = 0;

		for (let dy = 0; dy < step && y + dy < height; dy++) {
			for (let dx = 0; dx < step && x + dx < width; dx++) {
				const idx = ((y + dy) * width + (x + dx)) * 4;
				totalR += imageData.data[idx];
				totalG += imageData.data[idx + 1];
				totalB += imageData.data[idx + 2];
				count++;
			}
		}

		return {
			r: Math.round(totalR / count),
			g: Math.round(totalG / count),
			b: Math.round(totalB / count),
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
