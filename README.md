
---

# Nebula Pack

Nebula Pack is a package manager for Lua, written in Go. It provides functionality similar to `package.json` in Node.js, allowing you to manage dependencies, scripts, and other project configurations.

## Features

- Parse and process `nebula.json` configuration files.
- Handle HTTP POST requests to deduct dependencies.
- Basic error handling and logging.

## Getting Started

### Prerequisites

- Go (version 1.13 or higher)

### Installation

Clone the repository:

```bash
git clone https://github.com/your-username/nebula-pack.git
cd nebula-pack
```

### Running the Server

Run the server:

```bash
go run main.go
```

The server will start on port `7777` by default. You can access it at `http://localhost:7777/nebula-config`.

### Usage

#### Sending a POST Request

You can send a POST request to `http://localhost:7777/nebula-config` with a `nebula.json` payload:

```json
{
    "name": "my-lua-project",
    "version": "1.0.0",
    "description": "Description of your Lua project",
    "main": "main.lua",
    "scripts": {
        "start": "lua main.lua",
        "test": "echo \"Error: no test specified\" && exit 1"
    },
    "dependencies": {
        "lua-socket": "^3.0.0"
    },
    "devDependencies": {
        "luacheck": "^0.24.0"
    },
    "author": "Your Name",
    "license": "MIT"
}
```

#### Example using cURL:

```bash
curl -X POST http://localhost:7777/nebula-config -H "Content-Type: application/json" -d '{
    "name": "my-lua-project",
    "version": "1.0.0",
    "description": "Description of your Lua project",
    "main": "main.lua",
    "scripts": {
        "start": "lua main.lua",
        "test": "echo \"Error: no test specified\" && exit 1"
    },
    "dependencies": {
        "lua-socket": "^3.0.0"
    },
    "devDependencies": {
        "luacheck": "^0.24.0"
    },
    "author": "Your Name",
    "license": "MIT"
}'
```

### Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- Keagan Gilmore
- Built with ❤️ and ☕.

---