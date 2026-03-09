# webappManager

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?logo=go)](https://golang.org/)

A lightweight Go application for managing and running multiple web services concurrently. Perfect for developers who need to run several local development servers simultaneously.

## 🚀 Features

- **Concurrent Execution**: Run multiple web services simultaneously using Go routines
- **Flexible Service Support**: Compatible with any command-line executable service
- **Caddy Reverse Proxy Support**: Route local domains to your dev services (e.g., `app.localhost` -> `localhost:3000`)
- **Simple Configuration**: Easy-to-modify Go code for adding/removing services
- **Output Streaming**: Real-time stdout/stderr output from all managed services
- **Cross-Platform**: Works on Windows, macOS, and Linux

## 📋 Prerequisites

- Go 1.16 or higher
- The command-line tools you want to run (e.g., `npm`, `python3`, etc.)
- [Caddy](https://caddyserver.com/) installed if you want local domain reverse proxy support

## 🔧 Installation

### From Source

1. Clone the repository:

```bash
git clone https://github.com/ethancroll/webappManager.git
cd webappManager
```

2. Build the application:

```bash
go build -o webappManager main.go
```

3. Run the executable:

```bash
./webappManager
```

## 💻 Usage

### Reverse Proxy with Caddy (Local Domains)

This project can start Caddy alongside your app servers so you can use friendly local domains instead of raw ports.

Example flow:

1. Add reverse proxy entries in your `Caddyfile` (for example, `admin.localhost` -> `localhost:3000`, `client.localhost` -> `localhost:3001`).
2. Add local host mappings where needed:
   - `127.0.0.1 admin.localhost`
   - `127.0.0.1 client.localhost`
3. Start everything with `webappManager`; Caddy and your services run together.

From the current configuration, Caddy is launched like this:

```go
go mustRun(`C:\code\caddy\caddy.exe`, `C:\code\caddy`, "run", "--config", "Caddyfile")
```

Once running, use your local domains in the browser instead of `localhost:<port>`.

### Basic Configuration

Edit the `main.go` file to configure your services. Each service is defined using the `mustRun()` function:

```go
func main() {
    go mustRun("npm", "C:\\code\\localFile", "start")                          // Run Node.js server
    mustRun("python3", "C:\\code\\secretsHolder", "-m", "http.server", "9000") // Run Python HTTP server
}
```

### Function Signature

```go
func mustRun(cmd, dir string, args ...string)
```

**Parameters:**

- `cmd`: The command to execute (e.g., `"npm"`, `"python3"`, `"go"`)
- `dir`: The working directory for the command
- `args`: Variable number of command arguments

### Examples

#### Running an npm Development Server

```go
go mustRun("npm", "/home/user/my-react-app", "run", "dev")
```

#### Running Multiple Python HTTP Servers

```go
go mustRun("python3", "/var/www/site1", "-m", "http.server", "8000")
go mustRun("python3", "/var/www/site2", "-m", "http.server", "8001")
```

#### Running a Go Server

```go
go mustRun("go", "/home/user/my-go-api", "run", ".")
```

#### Mixed Services

```go
func main() {
    go mustRun("npm", "/path/to/frontend", "start")              // Frontend on port 3000
    go mustRun("npm", "/path/to/backend", "run", "dev")          // Backend API
    go mustRun("python3", "/path/to/docs", "-m", "http.server", "8080") // Documentation server
    mustRun("go", "/path/to/microservice", "run", "main.go")     // Go microservice (blocking call)
}
```

**Note:** All `mustRun()` calls except the last one should be prefixed with `go` to run concurrently. The last call should be blocking to keep the program alive.

## 🛠️ How It Works

The application uses Go's `os/exec` package to manage external processes:

1. **Command Building**: Constructs commands with specified arguments
2. **Directory Setting**: Sets the working directory for each service
3. **Output Routing**: Routes stdout and stderr to the parent process
4. **Concurrent Execution**: Uses goroutines for parallel service execution

## 📝 Configuration Tips

- **Path Formats**: Use forward slashes (`/`) on Unix-like systems and backslashes (`\\`) on Windows, or use `filepath.Join()` for cross-platform compatibility
- **Port Conflicts**: Ensure each service uses a unique port to avoid conflicts
- **Blocking Call**: Always have at least one non-concurrent `mustRun()` call to prevent the program from exiting immediately
- **Error Handling**: Monitor the console output for any service errors

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Commit your changes**: `git commit -m 'Add amazing feature'`
4. **Push to the branch**: `git push origin feature/amazing-feature`
5. **Open a Pull Request**

### Development Guidelines

- Keep the code simple and readable
- Add comments for complex logic
- Test on multiple platforms if possible
- Update documentation for new features

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Ethan Carroll** ([@ethancroll](https://github.com/ethancroll))

## 🐛 Issues and Support

If you encounter any issues or have questions:

- Open an issue on [GitHub Issues](https://github.com/ethancroll/webappManager/issues)
- Provide detailed information about your environment and the problem

## 🌟 Acknowledgments

- Built with Go's powerful concurrency model
- Inspired by the need for simple multi-service development workflows

## 📊 Project Status

This project is actively maintained. Feel free to use it, fork it, and contribute!

---

**Happy Coding!** 🚀
