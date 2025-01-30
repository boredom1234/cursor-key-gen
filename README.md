# 🚀 Cursor Key Generation API

This is a simple Go-based API that generates unique telemetry keys. It creates a UUID-based key and returns it as a JSON response when you call the `/generate` endpoint.

## 🛠️ Features
- Generates a **UUID-based telemetry key** 🔑
- Returns structured JSON output 📦
- Logs API requests and responses 📜
- Lightweight and fast ⚡

## 📌 Prerequisites
- **Go 1.18+** installed 🐹

## 🚀 How to Run
Clone the repository and navigate to the project directory:

```sh
 git clone <repo-url>
 cd <repo-directory>
```

Then, run the Go server:

```sh
go run main.go
```

By default, the server starts on **port 8080**.

## 🔥 Usage
Once the server is running, you can generate a cursor key using **cURL**:

```sh
curl -X GET http://localhost:8080/generate
```

### 🎯 Example Response
```json
{
    "telemetry.machineId": "123e4567-e89b-12d3-a456-426614174000",
    "telemetry.macMachineId": "123e4567-e89b-12d3-a456-426614174000",
    "telemetry.devDeviceId": "123e4567-e89b-12d3-a456-426614174000",
    "telemetry.sqmId": "123e4567-e89b-12d3-a456-426614174000",
    "lastModified": "2025-01-30T12:34:56Z",
    "version": "1.0.1"
}
```

## 📝 API Details
| Method | Endpoint     | Description |
|--------|-------------|-------------|
| GET    | `/generate` | Returns a new UUID-based telemetry key |

## 🤖 Logs
The server logs all activity:
- Request received ✅
- UUID generated 🔄
- Response sent 📩
- Error handling 🚨

## ⚙️ Configuration
- The server **runs on port 8080** by default.
- Change the port in `main.go`:
  ```go
  http.ListenAndServe(":8080", nil)
  ```

## 🛑 Error Handling
If something goes wrong, the server logs the error and returns:
```json
{
    "error": "Internal Server Error"
}
```

## 💡 Future Improvements
- ✅ Add Docker support 🐳
- ✅ Implement authentication 🔒
- ✅ Rate limiting ⏳

---
🔗 **Happy Coding!** 🖥️🚀

