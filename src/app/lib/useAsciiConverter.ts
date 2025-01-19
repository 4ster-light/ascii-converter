import { useState } from 'react';
import { getAsciiChar, calculateDimensions, getPixelBrightness, getImageData, loadImage } from './utils';
import { DEFAULT_CONFIG, AsciiConfig } from './data';

export const useAsciiConverter = (config: AsciiConfig = DEFAULT_CONFIG) => {
  const [result, setResult] = useState<string>('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const createAsciiArt = (
    imageData: ImageData,
    width: number,
    height: number,
    preserveColors: boolean
  ): string => {
    let ascii = '';

    // Calculate step sizes for better performance
    const widthStep = Math.max(1, Math.floor(width / 500));
    const heightStep = Math.max(1, Math.floor(widthStep));

    for (let y = 0; y < height; y += heightStep) {
      let line = '';
      for (let x = 0; x < width; x += widthStep) {
        // Sample pixel colors using averaged area
        const { r: avgR, g: avgG, b: avgB } = sampleArea(
          imageData,
          x,
          y,
          widthStep,
          heightStep,
          width,
          height
        );

        const brightness = getPixelBrightness(avgR, avgG, avgB);
        const char = getAsciiChar(brightness);

        line += preserveColors
          ? `<span style="color: rgb(${Math.round(avgR)},${Math.round(avgG)},${Math.round(avgB)})">${char}</span>`
          : char;
      }
      ascii += line + '\n';
    }

    return ascii;
  };

  const sampleArea = (
    imageData: ImageData,
    startX: number,
    startY: number,
    widthStep: number,
    heightStep: number,
    totalWidth: number,
    totalHeight: number
  ) => {
    let totalR = 0, totalG = 0, totalB = 0;
    let sampleCount = 0;

    for (let sy = 0; sy < heightStep && startY + sy < totalHeight; sy++) {
      for (let sx = 0; sx < widthStep && startX + sx < totalWidth; sx++) {
        const i = ((startY + sy) * totalWidth + (startX + sx)) * 4;
        totalR += imageData.data[i];
        totalG += imageData.data[i + 1];
        totalB += imageData.data[i + 2];
        sampleCount++;
      }
    }

    return {
      r: totalR / sampleCount,
      g: totalG / sampleCount,
      b: totalB / sampleCount
    };
  };

  const convertImageToAscii = async (file: File): Promise<void> => {
    setLoading(true);
    setError(null);

    try {
      const imageUrl = URL.createObjectURL(file);
      const img = await loadImage(imageUrl);
      
      const { width, height } = calculateDimensions(
        img.width,
        img.height,
        config.maxWidth || DEFAULT_CONFIG.maxWidth!,
        config.maxHeight || DEFAULT_CONFIG.maxHeight!
      );

      const imageData = getImageData(img, width, height);
      const ascii = createAsciiArt(imageData, width, height, config.preserveColors);
      
      setResult(ascii);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to convert image');
    } finally {
      setLoading(false);
    }
  };

  return {
    result,
    loading,
    error,
    convertImageToAscii,
    clearResult: () => setResult('')
  };
}
