import type { ReactElement } from "react";
import type { AsciiConfig } from "../lib/data.ts";

export function AsciiPreview({
	ascii,
	config,
}: {
	ascii: string;
	config: AsciiConfig;
}): ReactElement {
	return (
		<div className="card bg-base-200 overflow-x-auto">
			<div className="bg-base-200 rounded-box overflow-x-auto p-4">
				<pre
					className="ascii-art font-mono whitespace-pre text-center mx-auto"
					style={{
						fontSize: `${Math.max(8, 16 - config.resolution)}px`, // Unified size
						lineHeight: 1,
						transform: `scale(${0.5 + config.resolution / 10})`,
						transformOrigin: "center",
						letterSpacing: "-0.1em",
						wordSpacing: "-0.2em",
					}}
					dangerouslySetInnerHTML={{ __html: ascii }}
				/>
			</div>
		</div>
	);
}
