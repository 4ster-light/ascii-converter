import type React from "react";

export type CardProps = {
	title: string;
	description: string;
	items: string[];
};

export function Card({
	title,
	description,
	items,
}: CardProps): React.JSX.Element {
	return (
		<div className="card bg-base-200 shadow-sm mb-6">
			<div className="card-body">
				<h2 className="card-title text-primary">{title}</h2>
				<p className="text-base-content">{description}</p>
				<ul className="list-disc list-inside space-y-1">
					{items.map((item) => (
						<li key={item} className="text-base text-base-content">
							{item}
						</li>
					))}
				</ul>
			</div>
		</div>
	);
}
