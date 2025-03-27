# Conti Victims Scraper

![Language](https://img.shields.io/badge/Language-Go-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Status](https://img.shields.io/badge/Status-Archived-red.svg)
![Platform](https://img.shields.io/badge/Platform-Tor%20Network-orange.svg)

## Overview

This repository contains a Golang-based web scraper designed to automate the extraction of victim-related information from the Conti ransomware group's former onion site on the Tor network. The scraper systematically collects data about companies targeted by the Conti group, including company names and associated information released by the attackers.

## Context

During 2022, the Conti ransomware group operated a Tor-based platform where they publicly listed their victims, along with ransom demands and leaked data excerpts. This scraper was created to efficiently capture and structure that information for subsequent analysis.

## Technical Features

- Developed using Golang, optimized for performance and concurrency.
- Integrates with Tor via SOCKS5 proxy to securely access onion sites.
- Parses HTML content to retrieve structured victim information.
- Outputs results into structured JSON format (`conti_db.json`).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ramonfure/contiscraper_victims.git
   cd contiscraper_victims
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Ensure Tor service is running and properly configured for SOCKS5 proxy.

## Usage

Execute the scraper using:

```bash
go run contilinks.go
```

Data extracted from the scraper will be saved in the `conti_db.json` file in the project root.

## Disclaimer

- The Conti onion domain is no longer active; thus, this scraper is intended solely for educational purposes, historical data analysis, or cybersecurity research.
- Ethical use and legal compliance are mandatory.

## License

This project is distributed under the MIT License. See [LICENSE](LICENSE) for details.
