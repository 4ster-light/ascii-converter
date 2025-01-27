import type { FormEvent, ReactElement } from "react";
import { useRef, useState } from "react";
import { ASCII_CHARS } from "../lib/data";
import { useAsciiConverter } from "../lib/useAsciiConverter";
import { AsciiPreview } from "./AsciiPreview";

export function Form(): ReactElement {
	const fileInputRef = useRef<HTMLInputElement>(null);
	const [config, setConfig] = useState({
		preserveColors: false,
		resolution: 5,
		charSet: ASCII_CHARS,
		colorIntensity: 5,
	});

	const { result, loading, error, convertImageToAscii, clearResult } =
		useAsciiConverter({
			...config,
			maxWidth: 800,
			maxHeight: 600,
		});

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

	return (
		<form onSubmit={handleSubmit} className="space-y-6">
			<div className="form-control">
				<input
					ref={fileInputRef}
					id="file-input"
					type="file"
					accept="image/*"
					className="file-input file-input-bordered file-input-primary w-full bg-base-100"
				/>
			</div>

			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div className="form-control">
					<label htmlFor="resolution-range" className="label">
						<span className="label-text">Resolution</span>
					</label>
					<input
						id="resolution-range"
						type="range"
						min="1"
						max="10"
						value={config.resolution}
						onChange={(e) =>
							setConfig((prev) => ({
								...prev,
								resolution: Number(e.target.value),
							}))
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
						value={config.charSet}
						onChange={(e) =>
							setConfig((prev) => ({ ...prev, charSet: e.target.value }))
						}
						className="select select-primary bg-base-100"
					>
						<option value={ASCII_CHARS}>Default</option>
						<option value="@%#*+=-:. ">Simple</option>
						<option value="▓▒░ ">Block</option>
						<option value="01">Binary</option>
					</select>
				</div>

				<div className="form-control">
					<label className="label cursor-pointer justify-start gap-2">
						<input
							type="checkbox"
							checked={config.preserveColors}
							onChange={(e) =>
								setConfig((prev) => ({
									...prev,
									preserveColors: e.target.checked,
								}))
							}
							className="checkbox checkbox-primary checkbox-sm"
						/>
						<span className="label-text">Preserve Colors</span>
					</label>
				</div>

				{config.preserveColors && (
					<div className="form-control">
						<label htmlFor="intensity-range" className="label">
							<span className="label-text">Color Intensity</span>
						</label>
						<input
							id="intensity-range"
							type="range"
							min="1"
							max="10"
							value={config.colorIntensity}
							onChange={(e) =>
								setConfig((prev) => ({
									...prev,
									colorIntensity: Number(e.target.value),
								}))
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
					className="btn bg-base-100 border-primary text-primary hover:bg-primary/10 flex-1"
				>
					Clear Image
				</button>
				<button
					type="submit"
					className="btn bg-primary border-primary text-base-100 hover:bg-primary/90 flex-1"
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

			{result && <AsciiPreview ascii={result} />}
		</form>
	);
}
