# server-bridge

# API Documentation

## Overview

This API allows users to interact with a remote server using SSH keys for authentication. It provides endpoints for uploading files, downloading files, and retrieving directory details. The API is built using Golang and the Gin framework.

## Requirements

- Golang 1.18+
- Gin framework
- SSH keys for authentication
- Remote server access

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/p-shubh/server-bridge.git
   ```
2. Navigate to the project directory:
   ```bash
   cd your-api
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### 1. Login and Upload File

**Endpoint:** `POST /upload`

**Description:** Logs in to the remote server using the provided SSH key and server address, and uploads a file (max 5 MB).

**Request:**
```json
{
  "server_address": "string",
  "username": "string",
  "ssh_key": "string",
  "file": "file"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "File uploaded successfully"
}
```

### 2. Login and Download File

**Endpoint:** `POST /download`

**Description:** Logs in to the remote server using the provided SSH key and server address, and downloads a file within 1.5 seconds.

**Request:**
```json
{
  "server_address": "string",
  "username": "string",
  "ssh_key": "string",
  "file_path": "string"
}
```

**Response:**
- File content as binary data

### 3. Login and Show Directory Details

**Endpoint:** `POST /directory`

**Description:** Logs in to the remote server using the provided SSH key and server address, and retrieves details of the directory from the root folder.

**Request:**
```json
{
  "server_address": "string",
  "username": "string",
  "ssh_key": "string",
  "directory_path": "string"
}
```

**Response:**
```json
{
  "status": "success",
  "directory_details": [
    {
      "name": "string",
      "type": "string",
      "size": "int",
      "mod_time": "string"
    }
  ]
}
```

## Example Usage

### Request to Upload File
```bash
curl -X POST http://localhost:8080/upload \
  -F "server_address=example.com" \
  -F "username=user" \
  -F "ssh_key=@/path/to/ssh_key" \
  -F "file=@/path/to/file"
```

### Request to Download File
```bash
curl -X POST http://localhost:8080/download \
  -d '{"server_address":"example.com", "username":"user", "ssh_key":"ssh_key_content", "file_path":"/path/to/file"}' \
  -H "Content-Type: application/json"
```

### Request to Show Directory Details
```bash
curl -X POST http://localhost:8080/directory \
  -d '{"server_address":"example.com", "username":"user", "ssh_key":"ssh_key_content", "directory_path":"/path/to/directory"}' \
  -H "Content-Type: application/json"
```

Replace the placeholder implementation with the actual logic for handling SSH connections and file operations.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
