export interface AsciiConfig {
	preserveColors: boolean;
	maxWidth: number;
	maxHeight: number;
}

export interface ImageDimensions {
	width: number;
	height: number;
}

export const ASCII_CHARS =
	"$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ";

export const DEFAULT_CONFIG: AsciiConfig = {
	preserveColors: false,
	maxWidth: 800,
	maxHeight: 600,
};
