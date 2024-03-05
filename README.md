# Jin: Your DDOS CLI Toolkit

```bash


         ___   ___   __      __
        /  /  /  /  /  \    /  /
       /  /  /  /  /    \  /  /
      /  /  /  /  /   \  \/  /
  ___/  /  /  /  /  /  \    /
 |_____/  /__/  /__/    \__/


Usage: main.py [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  attack
  check
  map
  scan
  version

```

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

## **Contributing:**

We encourage contributions from the community! If you have an idea for a new tool or want to improve existing ones, feel free to fork the repository and submit a pull request.

## **License:**

This project is licensed under the [Apache 2 LICENSE](LICENSE). See the LICENSE file for details.

## **Disclaimer:**

`The tools provided in this project are intended for ethical and informative purposes only. Please use them responsibly and avoid any actions that could harm others or violate their privacy.`
