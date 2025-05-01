# ioManager 🧠⚙️

**ioManager** is a modular AWS Lambda-based system for processing and transforming files as part of automation workflows. Designed to operate via command-style JSON payloads, this project supports uploading files, converting them, splitting and merging them, and more.

---
# 📋 To-Do List 
📁 Project Infrastructure

- [x] Configure AWS Lambda with container image

- [x] Define modular folder structure (commands/, controllers/, providers/, utils/)

- [x] Create unified input payload structure (command, args)

- [x] Create standard response structure (success, error)

- [x] Get task settings

🚀 Commands

- [x] Save-Input-File: Save xlsx file (main batch)

- [x] Explode-Input-File: Split an Excel file into multiple .json files (tasks of batch)

- [ ] Implode-Outputs: Merge multiple .json files into a final Excel file in main-output/

🔥 Commands Under Development

- [ ] set-new-task-configuration: Create a configuration file for each task in S3 (/$taskname/input/exploded/$idtask/config.json)

📊 Suggested Future Commands (File Management Focus)

- [ ] parquet-files: Create a .parquet  archive containing multiple S3 files (glue athena applications)


/$taskname/output/
    |- output-final.xlsx
    |- output-merged.json

📖 Documentation

- [ ] Document all available commands

- [ ] Provide examples of input and output payloads

- [ ] Document local development setup (sam local, Docker)

- [ ] Explain expected S3 folder structures

🔬 Optional Future Improvements

- [ ] Implement batch operations to automate implode/merge over full S3 prefixes

- [ ] Add schema validation for incoming JSON files

- [ ] Add operational logs saved as log.txt inside related S3 folders

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

| Command           | Description                                                                  |
| ----------------- | ---------------------------------------------------------------------------- |
| `save-input`      | Saves a file (XLSX, JSON, etc.) to S3                                        |
| `explode-input`   | Splits an XLSX file into multiple JSON files (one per row) and uploads to S3 |
| `implode`         | Merges multiple JSON files from an S3 prefix into an Excel file              |
| `download-output` | Returns the final generated file as bytes                                    |

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