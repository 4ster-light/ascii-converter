import React, { useState, useRef } from 'react';

export const Form: React.FC = () => {
  const [preserveColors, setPreserveColors] = useState(false);
  const [result, setResult] = useState("");
  const [loading, setLoading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Simplified ASCII characters for better alignment
  const ASCII_CHARS = '$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,"^`\'. ';

  const getPixelBrightness = (r: number, g: number, b: number): number => {
    // Using perceived brightness formula
    return (0.299 * r + 0.587 * g + 0.114 * b) / 255;
  };

  const getAsciiChar = (brightness: number): string => {
    // Map brightness (0-1) to ASCII character
    const charIndex = Math.floor(brightness * (ASCII_CHARS.length - 1));
    return ASCII_CHARS[charIndex];
  };

  const createAsciiArt = (
    imageData: ImageData,
    width: number,
    height: number,
    preserveColors: boolean
  ): string => {
    let ascii = '';

    // Much higher resolution with smaller steps
    const widthStep = Math.max(1, Math.floor(width / 500));
    const heightStep = Math.max(1, Math.floor(widthStep));

    // Calculate actual output dimensions
    const outputWidth = Math.floor(width / widthStep);
    const outputHeight = Math.floor(height / heightStep);

    for (let y = 0; y < height; y += heightStep) {
      let line = '';
      for (let x = 0; x < width; x += widthStep) {
        // Sample pixel colors
        let totalR = 0, totalG = 0, totalB = 0;
        let sampleCount = 0;

        // Average colors in the step area for better representation
        for (let sy = 0; sy < heightStep && y + sy < height; sy++) {
          for (let sx = 0; sx < widthStep && x + sx < width; sx++) {
            const i = ((y + sy) * width + (x + sx)) * 4;
            totalR += imageData.data[i];
            totalG += imageData.data[i + 1];
            totalB += imageData.data[i + 2];
            sampleCount++;
          }
        }

        const avgR = totalR / sampleCount;
        const avgG = totalG / sampleCount;
        const avgB = totalB / sampleCount;

        const brightness = getPixelBrightness(avgR, avgG, avgB);
        const char = getAsciiChar(brightness);

        if (preserveColors) {
          // Use CSS RGB colors for colored output
          line += `<span style="color: rgb(${Math.round(avgR)},${Math.round(avgG)},${Math.round(avgB)})">${char}</span>`;
        } else {
          line += char;
        }
      }
      ascii += line + '\n';
    }

    return ascii;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!fileInputRef.current?.files?.length) {
      alert('Please select an image file');
      return;
    }

    const file = fileInputRef.current.files[0];
    setLoading(true);

    try {
      // Create image element
      const img = new Image();
      const imageUrl = URL.createObjectURL(file);

      img.onload = () => {
        // Create canvas to get image data
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');

        if (!ctx) {
          throw new Error('Could not get canvas context');
        }

        // Calculate dimensions maintaining aspect ratio
        const maxWidth = 800;
        const maxHeight = 600;
        let width = img.width;
        let height = img.height;

        if (width > maxWidth) {
          height = Math.round((height * maxWidth) / width);
          width = maxWidth;
        }

        if (height > maxHeight) {
          width = Math.round((width * maxHeight) / height);
          height = maxHeight;
        }

        // Set canvas size
        canvas.width = width;
        canvas.height = height;

        // Draw image to canvas
        ctx.drawImage(img, 0, 0, width, height);

        // Get image data
        const imageData = ctx.getImageData(0, 0, width, height);

        // Convert to ASCII
        const ascii = createAsciiArt(imageData, width, height, preserveColors);

        // Set result
        setResult(ascii);
        setLoading(false);

        // Clean up
        URL.revokeObjectURL(imageUrl);
      };

      img.onerror = () => {
        throw new Error('Failed to load image');
      };

      img.src = imageUrl;
    } catch (error) {
      console.error('Error converting image:', error);
      alert('Error converting image to ASCII art');
      setLoading(false);
    }
  };

  const handleClear = () => {
    setResult('');
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="form-control mb-6">
        <input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          className="file-input file-input-bordered file-input-primary w-full max-w-xs bg-base-100"
        />
      </div>

      <div className="form-control mb-6">
        <label className="label cursor-pointer justify-start gap-4">
          <input
            type="checkbox"
            id="preserve-colors"
            name="preserve-colors"
            checked={preserveColors}
            onChange={(e) => setPreserveColors(e.target.checked)}
          />
          <span className="label-text text-lg font-bold text-primary">Preserve original colors</span>
        </label>
      </div>

      <div className="flex gap-4 mb-8">
        <button
          type="button"
          onClick={handleClear}
          className="btn bg-base-100 border-primary text-primary hover:bg-primary/10 hover:border-primary hover:text-primary"
        >
          Clear Image
        </button>
        <button
          type="submit"
          className="btn bg-primary border-primary text-base-100 hover:bg-primary/90 hover:border-primary/90 hover:text-base-100"
          disabled={loading}
        >
          {loading ? 'Converting...' : 'Generate ASCII'}
        </button>
      </div>

      <div className="card bg-base-300 duration-200">
        <div className="card-body p-4">
          {result && (
            <div className="flex justify-center">
              <pre
                className="font-mono text-xs md:text-sm lg:text-base whitespace-pre text-base-content"
                dangerouslySetInnerHTML={{ __html: result }}
                style={{
                  lineHeight: '0.5',
                  letterSpacing: '0',
                  textAlign: 'center',
                  fontFamily: 'Consolas, Monaco, "Courier New", monospace',
                  fontSize: '2px',
                  transform: 'scale(1.2)',
                  transformOrigin: 'center',
                  margin: '10px 0',
                  display: 'inline-block',
                  maxWidth: '100%',
                  overflow: 'hidden'
                }}
              />
            </div>
          )}
        </div>
      </div>
    </form>
  );
}
