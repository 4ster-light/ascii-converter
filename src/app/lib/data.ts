export interface AsciiConfig {
	preserveColors: boolean;
	maxWidth: number;
	maxHeight: number;
	resolution: number;
	charSet: string;
	colorIntensity: number;
}

export interface ImageDimensions {
	width: number;
	height: number;
}

export const ASCII_CHARS = " .,:;i1tfLCG08@"; // Reordered for better B&W gradient

export const CHAR_SET_OPTIONS = [
	{ name: "Default", value: ASCII_CHARS },
	{ name: "High Contrast", value: "@%#*+=-:. " }, // Renamed from "Simple"
	{ name: "Blocks", value: "▄▓▒░█" }, // Improved block characters
	{ name: "Binary", value: "01 " },
];

export const DEFAULT_CONFIG: AsciiConfig = {
	preserveColors: false,
	maxWidth: 500, // Reduced for better density
	maxHeight: 500, // Reduced for better density
	resolution: 5,
	charSet: ASCII_CHARS,
	colorIntensity: 5,
};
