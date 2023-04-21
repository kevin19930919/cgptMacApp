# OpenAI ChatGPT API Tool

A simple Go-based tool to interact with OpenAI's ChatGPT API, featuring a web-based user interface created with Lorca.

## Features

- Web-based UI for sending messages to the ChatGPT API
- Asynchronous sending of multiple requests
- Local storage and management of API keys

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system (version 1.17 or later)
- [Lorca](https://github.com/zserge/lorca) library installed

### Installation

1. Clone the repository:


2. Change into the project directory:


3. Build the project:


4. Run the compiled executable:


### Configuration

Create a `config.yaml` file in the project root directory, and add your OpenAI API key as follows:

```yaml
chatGPTAPIKey: "your_api_key_here"
```
Replace your_api_key_here with your actual API key.

### Usage
After running the compiled executable, a new window will open with the user interface. 
You can interact with the ChatGPT API by sending messages through the provided form. 
The application will send requests asynchronously, and the API responses will be displayed in the output area.

### Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.
