# Bowerbird

A lightweight, flexible directory organizer written in Go that automatically sorts your files based on type, date, or both.

> [!WARNING]
> Creation date has not been validated in MacOs, so i recommend not using the flag "creation" in those systems.

> [!WARNING]
> Currently Unix-compatible only. Windows support is not available yet.

## What is Bowerbird?
Bowerbird is a command-line tool that helps you organize cluttered directories by automatically sorting files into structured folder hierarchies.

## Features
- **File Type Organization**: Automatically categorizes files by extension into predefined categories (Images, Videos, Audio, Documents, Code, Archives, Executables)
- **Date-based Sorting**: Organize files by last modification year and/or month
- **Flexible Hierarchy**: Choose whether date or file type comes first in the folder structure
- **Hidden File Control**: Option to ignore or include hidden files (files starting with `.`)
- **Custom Source & Destination**: Specify different source and destination directories
- **Non-destructive**: Moves files into organized folders while preserving file names

## How It Works
Bowerbird operates in two main phases:

### 1. Scanning Phase
The scanner (`internal/scanner`) reads through the specified directory and collects metadata about each file:
- File name and path
- Last modification date
- File extension
- File category (determined by extension)

### 2. Organization Phase
The manipulator (`internal/manipulator`) creates the appropriate folder structure and moves files according to your specified flags:
- Creates nested directories based on your organization preferences
- Moves files into their designated folders
- Preserves file permissions on created directories

### File Categories
Bowerbird recognizes the following file categories:

- **Image**: `.jpg`, `.jpeg`, `.png`, `.gif`, `.svg`, `.webp`, `.bmp`
- **Video**: `.mp4`, `.avi`, `.mkv`, `.mov`, `.wmv`, `.flv`
- **Audio**: `.mp3`, `.wav`, `.flac`, `.aac`, `.ogg`
- **Document**: `.pdf`, `.doc`, `.docx`, `.txt`, `.rtf`, `.pptx`, `.odt`
- **Executable**: `.exe`, `.msi`, `.bat`, `.sh`, `.app`
- **Archive**: `.tar`, `.zip`, `.rar`, `.7z`, `.gz`
- **Code**: `.go`, `.py`, `.js`, `.java`, `.cpp`, `.c`, `.ts`, `.rb`
- **Unknown**: Any file type not in the above categories

## Installation

### Prerequisites
- Go 1.16 or higher

### From Source
```bash
git clone https://github.com/Klithik/bowerbird.git
cd bowerbird
go build -o bowerbird
cp bowerbird /usr/bin
```

## Usage

### Basic Usage
Organize the current directory by file type:
```bash
bowerbird
```

### Common Examples
**Organize a specific directory:**
```bash
bowerbird -source=/path/to/messy/folder -end=/path/to/organized/folder
```

**Organize by year:**
```bash
bowerbird -year
```

**Organize by year and month:**
```bash
bowerbird -month
# Note: -month automatically enables -year
```

**Organize with date priority (date folders first, then type):**
```bash
bowerbird -month -datePrio
# Results in: 2025/October/Image/photo.jpg
```

**Organize with type priority (type folders first, then date):**
```bash
bowerbird -month
# Results in: Image/2025/October/photo.jpg
```

**Include hidden files:**
```bash
bowerbird -ignore_hidden=false
```

## Command-Line Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-source` | string | `.` | The directory to be sorted |
| `-end` | string | `.` | The directory where organized folders should be created |
| `-type` | bool | `true` | Sort elements based on file type (extension) |
| `-year` | bool | `false` | Sort elements by last modification year |
| `-month` | bool | `false` | Sort elements by last modification month (automatically enables `-year`) |
| `-datePrio` | bool | `false` | When `true`, creates date folders first, then type folders. When `false`, creates type folders first, then date folders |
| `-ignore_hidden` | bool | `true` | Ignores hidden files (files starting with `.`) |

## Example Scenarios
### Scenario 1: Organize Downloads Folder
```bash
bowerbird -source=~/Downloads -end=~/Organized
```
This creates an organized copy in `~/Organized` with structure like:
```
Organized/
├── Image/
├── Video/
├── Document/
├── Code/
└── Unknown/
```

### Scenario 2: Organize Photos by Year and Month
```bash
bowerbird -source=~/Photos -month -type=false
```
Creates structure like:
```
Photos/
├── 2024/
│   ├── January/
│   ├── February/
│   └── ...
└── 2025/
    ├── January/
    └── ...
```

### Scenario 3: Complete Organization (Type + Date)
```bash
bowerbird -source=~/Documents -month
```
Creates structure like:
```
Documents/
├── Document/
│   ├── 2024/
│   │   └── December/
│   └── 2025/
│       └── January/
├── Image/
│   └── 2025/
│       └── October/
└── Code/
    └── 2024/
        └── November/
```

## Project Structure

```
bowerbird/
├── main.go                    # Entry point, flag parsing, and directory validation
├── internal/
│   ├── scanner/
│   │   └── scan.go           # File scanning and categorization
│   └── manipulator/
│       └── manipulator.go    # File organization and movement
└── go.mod
```

## How Directory Hierarchies Work
Bowerbird creates nested folder structures based on your flags:

**With `-datePrio=true`:**
```
[Date] → [Type] → File
Example: 2025/October/Image/photo.jpg
```

**With `-datePrio=false` (default):**
```
[Type] → [Date] → File
Example: Image/2025/October/photo.jpg
```

**Without date flags:**
```
[Type] → File
Example: Image/photo.jpg
```

## Contributing
Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

## Acknowledgments
Named after the bowerbird, known for its meticulous organization and arrangement of collected objects.
