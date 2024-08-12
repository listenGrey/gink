Gink P2P File Transfer
Gink is a command-line based Peer-to-Peer (P2P) file transfer application that allows users to easily send and receive files directly between devices over the internet. Using a websocket-based communication protocol, Gink ensures a reliable and efficient file transfer experience.

Features
P2P File Transfer: Direct file transfers between peers without the need for intermediate servers.
WebSocket Protocol: Utilizes websockets for real-time data transmission.
CLI Based: Simple and intuitive command-line interface.
Cross-platform: Supports both Windows and Linux operating systems.
Transfer Metadata: Sends file metadata including the filename prior to the file content.
Installation
To get started with Gink, follow these installation steps. Ensure that you have Go installed on your system (version 1.13 or higher is recommended).

Building from Source
Clone the repository and build the application:

bash
复制代码
git clone https://github.com/yourusername/gink.git
cd gink
go build -o gink main.go
Executables
Alternatively, you can download pre-compiled executables for Windows and Linux from the Releases page.

Usage
Once installed, you can run Gink using the following commands in your command-line interface.

Running the Program
To start the application, navigate to the directory containing the executable and run:

bash
复制代码
./gink run
Commands
gink help: Displays all available commands.
gink local <file directory>: Sets the local directory for saving received files.
gink add <IP address> <port>: Adds a peer’s address for sending files.
gink des: Lists all configured peer destinations.
gink send <filename> <destination index>: Sends a file to a specified peer.
Configuration
Edit the config.json file to set up initial parameters, including the default directory for received files and known peers.

Contributing
Contributions to Gink are welcome! Please read our Contributing Guide for details on our code of conduct, and the process for submitting pull requests.

License
This project is licensed under the MIT License - see the LICENSE.md file for details.

Acknowledgments
Thanks to the Go community for the comprehensive libraries and tools.
Special thanks to gorilla/websocket for the WebSocket support.