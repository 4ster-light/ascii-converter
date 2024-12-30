import { Html } from "@elysiajs/html"

export default function Form() {
  return <form
    hx-post="/ui/spinner"
    hx-trigger="submit"
    hx-target="#ascii-result"
    hx-swap="outerHTML"
  >
    <div class="form-control mb-6">
      <input
        type="file"
        accept="image/*"
        class="file-input file-input-bordered file-input-primary w-full max-w-xs bg-base-100"
      />
    </div>

    <div class="form-control mb-6">
      <label class="label cursor-pointer justify-start gap-4">
        <input
          type="checkbox"
          id="preserve-colors"
          class="checkbox checkbox-primary"
        />
        <span class="label-text text-lg font-bold text-primary">Preserve original colors</span>
      </label>
    </div>

    <div class="flex gap-4 mb-8">
      <button class="btn bg-base-100 border-primary text-primary hover:bg-primary/10 hover:border-primary hover:text-primary">
        Clear Image
      </button>
      <button
        class="btn bg-primary border-primary text-base-100 hover:bg-primary/90 hover:border-primary/90 hover:text-base-100"
        type="submit"
      >
        Generate ASCII
      </button>
    </div>

    <div class="card bg-base-300 duration-200">
      <div class="card-body p-4">
        <pre class="w-auto whitespace-pre-wrap text-[1px] md:text-[2px] font-bold tracking-widest text-base-content">
          <span id="ascii-result"></span>
        </pre>
      </div>
    </div>
  </form>
}
