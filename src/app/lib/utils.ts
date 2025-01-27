import { ASCII_CHARS, type ImageDimensions } from "./data";

export function getPixelBrightness(r: number, g: number, b: number): number {
	return (0.299 * r + 0.587 * g + 0.114 * b) / 255; // Perceived brightness formula
}

export function getAsciiChar(brightness: number): string {
	const charIndex = Math.floor(brightness * (ASCII_CHARS.length - 1));
	return ASCII_CHARS[charIndex];
}

export function calculateDimensions(
	imgWidth: number,
	imgHeight: number,
	maxWidth: number,
	maxHeight: number,
): ImageDimensions {
	const widthRatio = maxWidth / imgWidth;
	const heightRatio = maxHeight / imgHeight;
	const ratio = Math.min(widthRatio, heightRatio);

	return {
		width: Math.floor(imgWidth * ratio),
		height: Math.floor(imgHeight * ratio),
	};
}

export function loadImage(url: string): Promise<HTMLImageElement> {
	return new Promise((resolve, reject) => {
		const img = new Image();
		img.onload = () => resolve(img);
		img.onerror = () => reject(new Error("Failed to load image"));
		img.src = url;
	});
}

export function getImageData(
	img: HTMLImageElement,
	width: number,
	height: number,
): ImageData {
	const canvas = document.createElement("canvas");
	const ctx = canvas.getContext("2d");

	if (!ctx) {
		throw new Error("Could not get canvas context");
	}

	canvas.width = width;
	canvas.height = height;
	ctx.drawImage(img, 0, 0, width, height);

	return ctx.getImageData(0, 0, width, height);
}
