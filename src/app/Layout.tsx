import type React from "react";
import type { ReactElement } from "react";
import { useState } from "react";

export function Layout({
	children,
}: { children: React.ReactNode }): ReactElement {
	const [stars] = useState(() =>
		Array.from({ length: 50 }, () => ({
			x: Math.random() * 100,
			y: Math.random() * 100,
			size: Math.random() * 4 + 2,
			delay: Math.random() * 2,
		})),
	);

	return (
		<div className="min-h-screen bg-base-100 relative overflow-hidden">
			{/* Animated star background */}
			<div className="absolute inset-0 z-0">
				{stars.map((star, i) => (
					<div
						key={i.toString()}
						className="absolute bg-primary rounded-full star"
						style={{
							left: `${star.x}%`,
							top: `${star.y}%`,
							width: `${star.size}px`,
							height: `${star.size}px`,
							animationDelay: `${star.delay}s`,
							animation: "star-pulse 2s infinite ease-in-out",
						}}
					/>
				))}
			</div>

			<div className="relative z-10 container mx-auto px-4 py-8">
				<header className="mb-12 text-center">
					<h1 className="text-6xl font-heading font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
						✰ArtSCII✰
					</h1>
				</header>

				<main className="backdrop-blur-lg bg-base-200/50 rounded-2xl shadow-xl p-6">
					{children}
				</main>

				{/* Updated footer */}
				<footer className="mt-12 text-center text-base-content/80">
					<div className="space-y-2">
						<p className="text-sm">Crafted with ❤️ by David Vivar Bogónez</p>
						<div className="flex justify-center gap-4">
							{/* Social links with hover effects */}
						</div>
					</div>
				</footer>
			</div>
		</div>
	);
}
