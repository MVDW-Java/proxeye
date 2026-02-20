# proxeye

![Jenkins Pipeline](https://img.shields.io/badge/Jenkins-Pipeline-blue)

Proxeye is a lightweight HTTP proxy server designed to bypass Gelbooru's new referer header requirement for accessing media files. It fetches post metadata from the Gelbooru API and streams the media file.

## Quick Start

### Using Docker (Recommended)

1.  Create a `.env` file in the project directory with the following content:
    ```bash
    GELBOORU_API_KEY="your_api_key_here"
    GELBOORU_USER_ID="your_user_id_here"
    PROXEYE_PORT="8080"  # Optional, defaults to 8080
    ```
2.  Run `docker-compose up -d`

### Manual Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/yourusername/proxeye.git
    cd proxeye
    ```
2.  Set environment variables:
    ```bash
    export GELBOORU_API_KEY="your_api_key_here"
    export GELBOORU_USER_ID="your_user_id_here"
    export PROXEYE_PORT="8080"  # Optional, defaults to 8080
    ```
3.  Build and run:
    ```bash
    go build -o proxeye main.go
    ./proxeye
    ```

## Usage

Once running, access Gelbooru media by visiting:

```
http://localhost:8080/<post_id>[.<extension>]
```

For example, to access post `1234567`:
The file extension at the end of the URL is optional for services that don't read mime types like Discord.

```
http://localhost:8080/1234567
```

The server will:

1.  Fetch post metadata from Gelbooru API
2.  Extract the media URL
3.  Proxy the media with correct referer headers
4.  Stream the content to your client

## Building from Source

### Prerequisites

- Go 1.22 or higher
- Git

### Build Script

Use the included build script to create binaries for multiple platforms:

```bash
chmod +x build.sh
./build.sh
```

Binaries will be created in the `dist/` directory with names following the pattern:

```
Proxeye_<OS>-<ARCH>[.exe]
```

### Supported Platforms

| OS      | Architectures                |
| ------- | ---------------------------- |
| Windows | amd64, arm64                 |
| Linux   | amd64, arm64, ppc64, riscv64 |
| FreeBSD | amd64, arm64, riscv64        |
| OpenBSD | amd64, arm64, ppc64, riscv64 |
| macOS   | amd64, arm64                 |

## Configuration

### Environment Variables

| Variable           | Description               | Default  |
| ------------------ | ------------------------- | -------- |
| `PROXEYE_PORT`     | Port for the proxy server | `8080`   |
| `GELBOORU_API_KEY` | Your Gelbooru API key     | Required |
| `GELBOORU_USER_ID` | Your Gelbooru user ID     | Required |

### Getting API Credentials

1.  Log in to your Gelbooru account
2.  Go to Account Settings → API Key
3.  Generate or copy your API key and user ID

## Docker Build

To build your own Docker image:

```bash
docker build -t proxeye .
docker run -p 8080:8080 -e GELBOORU_API_KEY=xxx -e GELBOORU_USER_ID=yyy proxeye
```

## Disclaimer

This tool is for educational purposes only. Use it responsibly and in accordance with Gelbooru's Terms of Service. The developers are not responsible for any misuse of this software.
