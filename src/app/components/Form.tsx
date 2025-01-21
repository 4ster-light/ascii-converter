import type React from "react";
import { useRef, useState } from "react";
import { AsciiPreview } from "../components/AsciiPreview";
import type { AsciiConfig } from "../lib/data";
import { useAsciiConverter } from "../lib/useAsciiConverter";

export const Form: React.FC = () => {
	const [preserveColors, setPreserveColors] = useState(false);
	const fileInputRef = useRef<HTMLInputElement>(null);

	const config: AsciiConfig = {
		preserveColors,
		maxWidth: 800,
		maxHeight: 600,
	};

	const { result, loading, error, convertImageToAscii, clearResult } =
		useAsciiConverter(config);

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();

		if (!fileInputRef.current?.files?.length) {
			alert("Please select an image file");
			return;
		}

		await convertImageToAscii(fileInputRef.current.files[0]);
	};

	const handleClear = () => {
		clearResult();
		if (fileInputRef.current) {
			fileInputRef.current.value = "";
		}
	};

	return (
		<form onSubmit={handleSubmit} className="space-y-6">
			<div className="form-control">
				<input
					ref={fileInputRef}
					type="file"
					accept="image/*"
					className="file-input file-input-bordered file-input-primary w-full max-w-xs bg-base-100"
				/>
			</div>

			<div className="form-control">
				<label className="label cursor-pointer justify-start gap-4">
					<input
						type="checkbox"
						checked={preserveColors}
						onChange={(e) => setPreserveColors(e.target.checked)}
						className="checkbox checkbox-primary"
					/>
					<span className="label-text text-lg font-bold text-primary">
						Preserve original colors
					</span>
				</label>
			</div>

			<div className="flex gap-4">
				<button
					type="button"
					onClick={handleClear}
					className="btn bg-base-100 border-primary text-primary hover:bg-primary/10"
				>
					Clear Image
				</button>
				<button
					type="submit"
					className="btn bg-primary border-primary text-base-100 hover:bg-primary/90"
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
};
