It sounds like you're wrapping up a great project! Transitioning a CLI tool into a web application is a classic "rite of passage" in software development.

Since you need to document this specifically for a Go-based web server, I’ve structured this README to be professional, clear, and easy for any collaborator (or recruiter) to follow.

---

# ASCII Art Web

A robust Go-based web application that provides a Graphical User Interface (GUI) for generating ASCII art. This project extends the functionality of the original `ascii-art` CLI by allowing users to input text, select a banner style, and view the rendered result directly in their browser.

## Authors

* **[Favor Owuor](https://learn.zone01kisumu.ke/git/fcharles)**
* **[Richard Ochola](https://learn.zone01kisumu.ke/git/riotieno)**
* **[Peter Gati](https://learn.zone01kisumu.ke/git/piregi)**

## Description

ASCII Art Web is a server-side application that handles HTTP requests to transform plain text into stylized ASCII characters using pre-defined banner templates. The project emphasizes:

* **Dynamic Content:** Using Go `html/template` to render data.
* **Request Handling:** Processing `POST` data from web forms.
* **Error Management:** Implementing standard HTTP status codes for a reliable user experience.

## Usage

To run the server locally, ensure you have [Go](https://go.dev/) installed, then follow these steps:

1. **Clone the repository:**
```bash
git clone <your-repo-url>
cd ascii-art-web

```


2. **Run the server:**
```bash
go run main.go

```


3. **Access the GUI:**
Open your browser and navigate to: `http://localhost:8080` (or your designated port).

## Implementation Details

### The Algorithm

The core logic follows a specific pipeline to ensure the ASCII art is rendered correctly:

1. **Input Parsing:** The server receives the text and banner choice (shadow, standard, or thinkertoy) via a `POST` request.
2. **Banner Loading:** The application reads the corresponding `.txt` file from the server's filesystem. If the file is missing, it triggers a `500 Internal Server Error` or `404 Not Found`.
3. **Map Building:** The banner file is split into blocks (each character in these banners is typically 8 lines high). These are mapped to their respective ASCII values.
4. **String Construction:** The input text is processed line-by-line. For each character, the algorithm retrieves the 8-line representation and appends it to a buffer.
5. **Output:** The final string is sent back to the client and displayed within a `<pre>` tag to preserve formatting.

### HTTP Endpoints

| Endpoint | Method | Description | Success Code |
| --- | --- | --- | --- |
| `/` | `GET` | Renders the main GUI with the input form. | `200 OK` |
| `/ascii-art` | `POST` | Processes form data and returns the ASCII result. | `200 OK` |

### Error Handling

The server is designed to be resilient, returning:

* **400 Bad Request:** For malformed inputs or empty requests.
* **404 Not Found:** If the user attempts to access a non-existent route or template.
* **500 Internal Server Error:** If there is a failure in reading banner files or processing the logic server-side.

---

Would you like me to help you write the specific Go code for the `http.HandleFunc` logic to match these endpoints?