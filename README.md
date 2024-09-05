# ASCII-Art generator
This is a simple website that generates ASCII art from a provided image.

## The technology used to build this website is:
- **[Go](https://golang.org/)**: A programming language created at Google.
- **[Chi](https://go-chi.io/#/)**: A lightweight, idiomatic and composable router for building Go HTTP services.
- **[Templ](https://templ.guide/)**: A Go template engine.
- **[HTMX](https://htmx.org/)**: A JavaScript library for making AJAX "just work" in your HTML.

> [!WARNING]
> This website is still in development and is not yet ready for production.
> At the moment, it can only convert png images, to a black and white output.

## How to run the website
In case you want to run the website locally, you can follow these steps:

1. Clone the repository:
```bash
git clone https://github.com/07devAdvb/ascii-converter.git 
```

2. Change the working directory to the project folder:
```bash
cd ascii-converter
```

3. Run the website:
```bash
go run ./cmd/server/main.go
```

4. Open your browser and navigate to `http://localhost:8080`.