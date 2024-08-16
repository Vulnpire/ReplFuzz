# Replace & Fuzz

replaces parameter values in URLs with payloads from a wordlist. It reads URLs from standard input, applies payloads from a wordlist to parameters specified in a parameter file, and outputs the modified URLs.

## Features

- Replaces values of specified parameters with payloads from a wordlist.
- Preserves special characters (`&`, `?`) in URLs.
- Filters out URLs that do not contain any of the specified parameters.

## Install

`go install -v github.com/Vulnpire/replfuzz@latest`

## Usage

1. **Prepare Files**:
   - **`wordlist.txt`**: A list of payloads, one per line.
   - **`params.txt`**: A list of parameters to be replaced in URLs. Format: `key=` (e.g., `file=`, `redirect=`).


![image](https://github.com/user-attachments/assets/f6cf1ce9-1154-4c06-a484-a99af73f2d4b)
