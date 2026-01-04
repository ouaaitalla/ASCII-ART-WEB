# ASCII Art Web Generator (Go)

This project is a simple **Go web application** that generates **ASCII art** from user input text using selectable banner styles. It runs a local HTTP server and renders results using HTML templates.

---

## Features

* Web interface for ASCII art generation
* Supports multiple banner styles
* Input validation (ASCII only, size limits)
* Uses Go’s standard `net/http` package
* HTML templates for clean output rendering

---

## Project Structure

```
.
├── main.go
├──Handler/
│  ├── ascii_art.go
│  ├── generateascii.go
│  ├── Home.go
│  └── validascii.go
├── go.mod
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   ├── index.html
│   └── result.html
└── README.md
```

---

## Requirements

* Go **1.18+**
* A web browser

---

## How to Run

1. **Clone the repository**

   ```bash
   git clone <your-repo-url>
   cd <your-project-folder>
   ```

2. **Run the server**

   ```bash
   go run main.go
   ```

3. **Open your browser**

   ```
   http://localhost:8080
   ```

---

## How It Works

### Routes

| Endpoint     | Method | Description          |
| ------------ | ------ | -------------------- |
| `/`          | GET    | Serves the home page |
| `/ascii-art` | POST   | Generates ASCII art  |

---

### ASCII Art Generation

* Text input is limited to **1024 characters**
* Only **valid ASCII characters** are accepted
* Newlines are supported
* Banner files are loaded from the `banners/` directory
* Output is rendered using HTML templates

---

## Error Handling

The server responds with appropriate HTTP status codes:

* `400 Bad Request` – invalid input or request
* `404 Not Found` – invalid route
* `405 Method Not Allowed` – unsupported HTTP method
* `500 Internal Server Error` – server or file errors

---

## Example Usage

1. Enter text in the input field
2. Choose a banner style
3. Submit the form
4. View the generated ASCII art result

---

## Technologies Used

* **Go (Golang)**
* `net/http`
* `html/template`
* Standard Go file handling

---

## Notes

* Banner files must exist in the `banners/` directory
* Template files must exist in the `templates/` directory
* Server runs on port **8080** by default