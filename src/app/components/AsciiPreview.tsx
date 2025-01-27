import type { ReactElement } from "react";

export function AsciiPreview({ ascii }: { ascii: string }): ReactElement {
	return (
		<div className="card bg-base-200 overflow-x-auto">
			<div className="card-body p-4">
				<pre className="ascii-art font-mono text-[6px] leading-[6px] sm:text-xs sm:leading-3 md:text-sm md:leading-4">
					{ascii}
				</pre>
			</div>
		</div>
	);
}
