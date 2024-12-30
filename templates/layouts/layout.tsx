import { Html } from "@elysiajs/html"

export default function Layout({ children }: { children: JSX.Element }) {
  return <html lang="en" data-theme="artscii">
    <head>
      <meta charset="UTF-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />

      {/* SEO */}
      <meta name="description" content="From image to ascii art converter" />
      <meta name="keywords" content="ascii, converter, artscii, ascii art, ascii art converter, 4ster, Aster, David, David Vivar, David Vivar Bogónez" />
      <meta name="author" content="4ster-light" />
      <meta property="og:title" content="ArtSCII" />
      <meta property="og:description" content="From image to ascii art converter" />
      <meta property="og:image" content="/public/img/banner.jpeg" />

      {/* Twitter */}
      <meta property="twitter:card" content="/public/img/banner.jpeg" />
      <meta property="twitter:site" content="@4ster_light" />
      <meta property="twitter:creator" content="@4ster_light" />
      <meta property="twitter:title" content="ArtSCII" />
      <meta property="twitter:description" content="From image to ascii art converter" />
      <meta property="twitter:image" content="/public/img/banner.jpeg" />

      {/* HTMX */}
      <script src="/public/js/htmx.min.js"></script>

      {/* Tailwind CSS */}
      <link rel="stylesheet" href="/public/css/index.css" />

      <title>✰ArtSCII✰</title>
    </head>
    <body>
      <div class="min-h-screen p-8 bg-base-200 transition-colors duration-800">
        <div class="container max-w-4xl mx-auto">
          <header class="relative flex justify-between items-center mb-8">
            <h1 class="text-5xl font-bold text-primary">✰ArtSCII✰</h1>
            <label class="swap swap-rotate m-1">
              <input
                type="checkbox"
                class="theme-controller"
                value="dark"
              />
              {/* sun icon */}
              <svg
                class="swap-on h-10 w-10 fill-current text-primary"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24">
                <path d="M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z" />
              </svg>
              {/* moon icon */}
              <svg
                class="swap-off h-10 w-10 fill-current text-primary"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24">
                <path d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z" />
              </svg>
            </label>
          </header>
          <main class="card bg-base-100 shadow-xl">
            <div class="card-body">
              {children}
            </div>
          </main>
          <footer class="footer footer-center p-4 mt-8 text-base-content">
            <div>
              <p class="text-lg">Made by David Vivar Bogónez, aka: Aster</p>
              <div class="flex gap-4">
                <a href="https://x.com/4ster_light" class="link link-primary hover:text-secondary transition-colors">Twitter</a>
                <a href="https://github.com/4ster-light" class="link link-primary hover:text-secondary transition-colors">GitHub</a>
                <a href="https://k4r.dev" class="link link-primary hover:text-secondary transition-colors">Website</a>
              </div>
            </div>
          </footer>
        </div>
      </div>
    </body>
  </html>
}
