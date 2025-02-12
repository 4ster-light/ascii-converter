import type { FormEvent, ReactElement } from "react";
import { useRef, useState } from "react";
import { ASCII_CHARS, CHAR_SET_OPTIONS } from "../lib/data";
import { useAsciiConverter } from "../lib/useAsciiConverter";
import { AsciiPreview } from "./AsciiPreview";

export function Form(): ReactElement {
	const fileInputRef = useRef<HTMLInputElement>(null);
	const configRef = useRef({
		preserveColors: false,
		resolution: 5,
		charSet: ASCII_CHARS,
		colorIntensity: 5,
		maxWidth: 120,
		maxHeight: 80,
	});

	const [currentConfig, setCurrentConfig] = useState(configRef.current);
	const { result, loading, error, convertImageToAscii, clearResult } =
		useAsciiConverter(configRef.current);

	function handleClear() {
		clearResult();
		if (fileInputRef.current) fileInputRef.current.value = "";
	}

	async function handleSubmit(e: FormEvent) {
		e.preventDefault();
		if (!fileInputRef.current?.files?.length) {
			alert("Please select an image file");
			return;
		}
		await convertImageToAscii(fileInputRef.current.files[0]);
	}

	const updateConfig = (updates: Partial<typeof configRef.current>) => {
		configRef.current = { ...configRef.current, ...updates };
		setCurrentConfig(configRef.current);
	};

	return (
		<form onSubmit={handleSubmit} className="space-y-6">
			<div className="form-control">
				<input
					ref={fileInputRef}
					id="file-input"
					type="file"
					accept="image/*"
					className="w-full p-2 border-2 border-primary rounded-btn text-primary bg-base-100"
				/>
			</div>

			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div className="form-control">
					<label htmlFor="resolution-range" className="label">
						<span className="label-text">Resolution (Detail Level)</span>
						<span className="label-text-alt">
							{currentConfig.resolution}/10
						</span>
					</label>
					<input
						id="resolution-range"
						type="range"
						min="1"
						max="10"
						value={currentConfig.resolution}
						onChange={(e) =>
							updateConfig({ resolution: Number(e.target.value) })
						}
						className="range range-primary"
					/>
				</div>

				<div className="form-control">
					<label htmlFor="char-set-select" className="label">
						<span className="label-text">Character Set</span>
					</label>
					<select
						id="char-set-select"
						value={currentConfig.charSet}
						onChange={(e) => updateConfig({ charSet: e.target.value })}
						className="w-full p-2 border-2 border-primary rounded-btn bg-base-100"
					>
						{CHAR_SET_OPTIONS.map((opt) => (
							<option key={opt.name} value={opt.value}>
								{opt.name}
							</option>
						))}
					</select>
				</div>

				<div className="form-control">
					<label className="label cursor-pointer justify-start gap-2">
						<input
							type="checkbox"
							checked={currentConfig.preserveColors}
							onChange={(e) =>
								updateConfig({ preserveColors: e.target.checked })
							}
							className="checkbox checkbox-primary checkbox-sm"
						/>
						<span className="label-text">Preserve Colors</span>
					</label>
				</div>

				{currentConfig.preserveColors && (
					<div className="form-control">
						<label htmlFor="intensity-range" className="label">
							<span className="label-text">Color Intensity</span>
							<span className="label-text-alt">
								{currentConfig.colorIntensity}/10
							</span>
						</label>
						<input
							id="intensity-range"
							type="range"
							min="1"
							max="10"
							value={currentConfig.colorIntensity}
							onChange={(e) =>
								updateConfig({ colorIntensity: Number(e.target.value) })
							}
							className="range range-primary"
						/>
					</div>
				)}
			</div>

			<div className="flex flex-col sm:flex-row gap-4">
				<button
					type="button"
					onClick={handleClear}
					className="w-full p-2 border-2 border-primary rounded-btn text-primary hover:bg-primary/10"
				>
					Clear Image
				</button>
				<button
					type="submit"
					className="w-full p-2 border-2 border-primary rounded-btn bg-primary text-base-100 hover:bg-primary/90"
					disabled={loading}
				>
					{loading ? "Converting..." : "Generate ASCII"}
				</button>
			</div>

			{error && (
				<div className="alert alert-error">
					<span>{error}</span>
				</div>
			)}

			{result && (
				<AsciiPreview
					ascii={result}
					config={currentConfig}
					key={`${currentConfig.resolution}-${currentConfig.charSet}`}
				/>
			)}
		</form>
	);
}
