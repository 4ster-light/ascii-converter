# ASCII-Art generator
This is a simple website that generates ASCII art from a provided image.
[Check it out here](https://artscii.onrender.com/)

## The technology used to build this website is:
- **[Go](https://golang.org/)**: A programming language created at Google.
- **[Chi](https://go-chi.io/#/)**: A lightweight, idiomatic and composable router for building Go HTTP services.
- **[Templ](https://templ.guide/)**: A Go template engine.
- **[HTMX](https://htmx.org/)**: A JavaScript library for making AJAX "just work" in your HTML.

> [!NOTE]
> This website is still in development and it's adding new features every day.

## How to run the website
### Using Docker
If you want to run the website using Docker, you can follow these steps:

1. Clone the repository:
```bash
git clone https://github.com/07devAdvb/ascii-converter.git 
```
2. Change the working directory to the project folder:
```bash
cd ascii-converter
```
3. Build the Docker image:
```bash
docker build -t ascii .
```
4. Run the Docker container:
```bash
docker run -p 8080:8080 ascii
```
5. Open your browser and navigate to `http://localhost:8080`.
---
### Using Go
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
