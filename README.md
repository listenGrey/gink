# Gink P2P File Transfer
****
Gink is a command-line based Peer-to-Peer (P2P) file transfer application that allows users to easily send and receive files directly between devices over the internet. Using some protocols like TCP, WebSocket... Gink ensures a reliable and efficient file transfer experience.

### Features
- **Direct P2P File Transfers**: Enables files to be sent directly between users without intermediate storage.
- **CLI Support**: Easy-to-use command-line interface.
- **Cross-platform**: Available for both Windows and Linux OS.
- **Automatic File Reception**: Files are automatically received and saved according to user preferences.
### Installation
- **Linux:**
  ``` bash
    curl 
    ```
- **Windows:**
    ``` bash
  curl
  ```
### Configuration
Edit the config.json file to set up initial parameters, including the default directory for received files and known peers:
``` json
{
    "local_save_path": "/your/download/direction",
    "destinations": ["ip:port"],
    "history_file_path": "/your/direction/history.json",
    "protocols": ["tcp","websocket"] 
    // Now only support tcp and websocket protocol.
    // First string is in use transfer protocol.
}
```
### Commands
``` bash
gink run: Starts the application to receive files.
gink help:  Displays all available commands.
gink local <file directory>: Sets the local directory for saving received files.
gink add <IP:port>: Adds a peer IP address for sending files.
gink des: Lists all configured peer destinations.
gink send -f <filepath> -d <destination index>: Sends a file to a specified peer.
gink history: Displays the history of file transfers
gink protocol: List all protocols
gink stop: Stop the application
```
### Contributing
Issues and suggestions for improvements are welcome! Please submit Pull Requests to improve this project.