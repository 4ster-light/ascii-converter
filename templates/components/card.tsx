import { Html } from "@elysiajs/html"

export interface CardProps {
  title: string
  description: string
  items: string[]
}

export default function Card(cardProps: CardProps) {
  const { title, description, items } = cardProps

  return <div class="card bg-base-200 shadow-sm mb-6">
    <div class="card-body">
      <h2 class="card-title text-primary">{title}</h2>
      <p class="text-base-content">{description}</p>
      <ul class="list-disc list-inside space-y-1">
        {items.map(item => <li class="text-base text-base-content">{item}</li>)}
      </ul>
    </div>
  </div>
}
