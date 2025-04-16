# ioManager 🧠⚙️

**ioManager** is a modular AWS Lambda-based system for processing and transforming files as part of automation workflows. Designed to operate via command-style JSON payloads, this project supports uploading files, converting them, splitting and merging them, and more.

---

## 🚀 Deployment on AWS Lambda

This project is designed to run as an AWS Lambda function. It can be deployed as a Docker container or using AWS SAM.

### Requirements

- Configured AWS CLI
- Docker or `sam-cli`
- IAM permissions for Lambda, S3, and CloudWatch Logs

### Example Deployment with SAM

```bash
sam build
sam deploy --guided
```

---

## 📦 Input Payload

The Lambda expects commands in JSON format. General structure:

```json
{
  "command": "save-input",
  "args": {
    "format": "base64",
    "body": "<base64_string>",
    "filename": "file.xlsx",
    "encoded": "base64",
    "account_name": "company123",
    "task_name": "task1",
    "id": "uuid"
  }
}
```

### Supported Commands

| Command          | Description                                                                 |
|------------------|-----------------------------------------------------------------------------|
| `save-input`     | Saves a file (XLSX, JSON, etc.) to S3                                       |
| `explode-input`  | Splits an XLSX file into multiple JSON files (one per row) and uploads to S3|
| `implode`        | Merges multiple JSON files from an S3 prefix into an Excel file             |
| `download-output`| Returns the final generated file as bytes                                   |

> Each command uses specific `args`.

---

## 📤 Example Usage

### 1. Save a file to S3

```json
{
  "command": "save-input",
  "args": {
    "format": "base64",
    "body": "TWFuIGlzIGRpc3Rpbmd1aXNoZWQ=",
    "filename": "file.xlsx",
    "encoded": "base64",
    "account_name": "company123",
    "task_name": "taskX",
    "id": "unique-id"
  }
}
```

### 2. Explode: Split Excel file into JSONs

```json
{
  "command": "explode-input",
  "args": {
    "account_name": "company123",
    "task_name": "taskX",
    "id": "unique-id",
    "filename": "file.xlsx"
  }
}
```

### 3. Implode: Merge JSONs into Excel

```json
{
  "command": "implode",
  "args": {
    "account_name": "company123",
    "task_name": "taskX",
    "id": "unique-id"
  }
}
```

---

## ✅ Example Response

```json
{
  "success": true,
  "message": "File processed successfully",
  "filename": "file.xlsx",
  "s3_key": "company123/taskX/input/unique-id/file.xlsx"
}
```

On error:

```json
{
  "success": false,
  "message": "Invalid or corrupted file"
}
```

---

## 🧱 Project Structure

```
.
├── cmd/                    # Lambda entrypoints
├── controllers/           # Command handlers: save, explode, implode...
├── models/                # Data structures
├── providers/             # External logic: S3, XLSX, etc.
├── utils/                 # Utility functions
└── main.go
```

---

## 🛠 Tech Stack

- Golang
- AWS Lambda
- S3 (input/output)
- Excel (via [tealeg/xlsx](https://github.com/tealeg/xlsx) or similar)

---

## 📬 Contact

Built with care by [Patricio Labarca](https://github.com/patriciojlg) 🧙‍♂️  
For questions, suggestions, or improvements, open a PR or Issue on this repo!

---