# Jin: Your Hacking CLI Toolkit

![Screenshot_20](https://github.com/aliftech/jin/assets/47414125/3745d4ad-bb89-4e3c-b4a4-9c593d33a16d)

Jin is a hacking command-line tools designed to make your scan port, gathering urls, check vulnerability and sending DDOS attack to your target. This tools is made for ethical and education purpose. I recommend you not to use this tools for harmfull action.

## **Current Tools:**

- **scan**: This tool scans a target host for open ports, providing valuable information for network analysis and troubleshooting.

- **attack**: This command is used to send DDOS target to target.

- **version**: This command will show the current version of
  this cli tool.

- **map**: Mapping and gathering the related urls in the target website.

- **check**: Check website vulnerability.

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

## **Contributing:**

We encourage contributions from the community! If you have an idea for a new tool or want to improve existing ones, feel free to fork the repository and submit a pull request.

## **License:**

This project is licensed under the [Apache 2 LICENSE](LICENSE). See the LICENSE file for details.

## **Disclaimer:**

`The tools provided in this project are intended for ethical and informative purposes only. Please use them responsibly and avoid any actions that could harm others or violate their privacy.`
