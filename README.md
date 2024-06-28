# Jin: Your Hacking CLI Toolkit

![Screenshot_23](https://github.com/aliftech/jin/assets/47414125/318b6b70-54af-4fa8-8d1a-ce3a64929bb6)

Jin is a hacking command-line tools designed to make your scan port, gathering urls, check vulnerability and sending DDOS attack to your target. This tools is made for ethical and education purpose. I recommend you not to use this tools for harmfull action.

## **Current Tools:**

- **scan**: This tool scans a target host for open ports, providing valuable information for network analysis and troubleshooting.

- **atk**: This command is used to send DDOS target to target.

- **version**: This command will show the current version of
  this cli tool.

- **map**: Mapping and gathering the related urls in the target website.

### **Installation:**

- **Clone the repository:**

```Bash
git clone https://github.com/aliftech/jin.git
```

**`Use code with caution`**

- **Navigate to the project directory:**

```Bash
cd jin
```

- **Create Virtual env**

```bash
py -m venv env
```

```bash
.\env\Scripts\activate
```

- **Install dependencies:**

```Bash
pip install -r requirements.txt
```

### **Usage:**

- Run the script with the desired command:

```Bash
python main.py <command> [options]
```

- Replace <command> with the specific tool you want to use.

## **How To Use**

### **Port Scanning**

The port scanning is used to scanning the open port of the targeted website. This function is called using the following command:

```bash
python main.py scan example.com
```

- The `example.com` above mean the targeted website domain.

- **Naming a log file**

  When you running both scan or map command, the system will automaticly store your scanning or mapping result into a log file. The filename automaticly generate by system, but now, you can naming the log file by yourself.

  ```bash
    scan example.com -wl example_test_scan
  ```

  The command above will generate a log file name `scan_example_test_scan.json`. The `-wl` parameter is used to renamed the logs file from the default name generate by JIN into user expected name.

### **DDOS Attack**

To launch a ddos attack to the targeted website, you can use the following command:

```bash
python main.py atk example.com -m [GET, PUT, POST] -p [payload of request] -t [number of thread (default 100)]
```

- The `example.com` above mean the targeted website domain.
- The `80` is the port number of targeted website. To gathering information about the opening port, you need to run the port scanning function to your targeted website.
- `[option] --threads 100` is the optional parameter, the number of threads to use for processing.

### **Mapping Related URL**

This function is designed to mapping all the related urls of the target. Here is the command to call this function:

```bash
python main.py map https://www.example.com
```

- `https://www.example.com` is the targeted url.

- **Naming a log file**

  When you running both scan or map command, the system will automaticly store your scanning or mapping result into a log file. The filename automaticly generate by system, but now, you can naming the log file by yourself.

  ```bash
    map https://www.example.com -wl example_test_map
  ```

  The command above will generate a log file name `map_example_test_map.json`. The `-wl` parameter is used to renamed the logs file from the default name generate by JIN into user expected name.

### **Logs Management**

This function is used to manage the application logs. You can list all logs file, read, and even delete the logs file.

- **List all logs**

  This command is used to list all logs files.

  ```bash
  logs -ls [scan || map]
  ```

  If you set scan as the param -ls value, the application will shows you all of the scanning logs files. Meanwhile, if you input map as the -ls value, it will return all mapping logs files.

- **Read logs file**

  This command is used to read the log content.

  ```bash
  logs -r filename.json
  ```

- **Delete a single los file**

  This command is used to delete a single log file.

  ```bash
  logs -rm filename.json
  ```

- **Delete all logs files**

  This command will help you to delete all logs files.

  ```bash
  logs -rm all
  ```

## **Run JIN CLI App Using Docker**

You can also running the cli app using docker by the following command:

```bash
docker compose run jin
```

Then, you will see the following result.

![Screenshot_23](https://github.com/aliftech/jin/assets/47414125/318b6b70-54af-4fa8-8d1a-ce3a64929bb6)

## **Run JIN Through Docker Image**

You can also build and run JIN using our existing docker image.

- Pull our latest JIN image using the following command

  ```bash
  docker pull wahyouka/jin
  ```

- Run the JIN docker container using the following command.

  ```bash
  docker run -it wahyouka/jin:latest
  ```

## **Contributing:**

We encourage contributions from the community! If you have an idea for a new tool or want to improve existing ones, feel free to fork the repository and submit a pull request.

## **License:**

This project is licensed under the [Apache 2 LICENSE](LICENSE). See the LICENSE file for details.

## **Disclaimer:**

`The tools provided in this project are intended for ethical and informative purposes only. Please use them responsibly and avoid any actions that could harm others or violate their privacy.`
