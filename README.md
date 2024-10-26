# ASCII-Art generator
This is a simple website that generates ASCII art from a provided image. [Check it out here](https://artscii.onrender.com/).

> [!NOTE]
> This website is still in development and it's continuosly adding new features.

### How to run the website locally
#### Clone the repository:
- Using SSH:
```bash
git clone git@github.com:4ster-light/ascii-converter.git
```
- Using HTTPS:
```bash
git clone https://github.com/4ster-light/ascii-converter.git 
```
- Change the working directory to the project folder:
```bash
cd ascii-converter
```

#### Run the website locally, for this you have three options:

- If you have have Go, Bun and Make installed on your machine, you can use the following command directly:
```bash
make run
```
This will build and run the project for you, look at the makefile for other available commands.

- If you have Docker installed on your machine, you can use the following commands:
```bash
docker build -t ascii-converter .
docker run -p 8080:8080 ascii-converter
```

- If you have have Nix installed on your machine, you can use the following commands:
```bash
nix develop
```
This will open a nix shell with all the required dependencies installed so you can use make as previously described.

---
This are all equivalent, just use whichever is more convinient for you.

#### Open your browser and navigate to `http://localhost:8080`.
