import click
import socket
import threading
import json
from rich import print as rprint
from rich.console import Console
import click_completion
import logging
from urllib.parse import urlparse
from .utils.function import *

click_completion.init()

console = Console()


@click.command("scan")
@click.argument("url", type=click.STRING)
def scan(url):
    try:
        ip_address = socket.gethostbyname(url)
        rprint(
            f'Scanning ports for [bold]{url}[/bold] ([cyan]{ip_address}[/cyan])...')
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
            logname = f"logs/scan_{url}.json"
            logging.basicConfig(filename=logname,
                                filemode='a',
                                format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                                datefmt='%H:%M:%S',
                                level=logging.INFO)
            logging.info(f"open ports: {open_ports}")
            rprint('[bold]Open ports:[/bold]', open_ports)

    except socket.gaierror:
        rprint(
            '[bold red]Hostname could not be resolved. Please enter a valid hostname.[/bold red]')
    except socket.error as e:
        rprint('[bold red]Could not connect to server.[/bold red]')
        rprint(f'[italic red]{e}[/italic red]')


@click.command("atk")
@click.argument("url", type=click.STRING)
@click.option("--method",
              "-m",
               type=click.Choice(["GET", "POST", "PUT"]),
               help="The HTTP method used in the targeted url (GET, PUT, POST)"
)
@click.option(
    "-p",
    "--payload",
    help="The request payload (JSON string or dictionary)",
)
@click.option("--threads", 
              "-t", 
              default=100, 
              help="Number of threads for DDoS attack"
)
def attack(url, method, payload, threads):
    rprint(
        f'Initiating DDoS attack on [bold]{url}[/bold] with [green]{method} method[/green] and [cyan]{threads}[/cyan] threads...')
    headers = {"Content-Type": "application/json"}  # Set default for JSON

    # Try converting payload to a dictionary if it's a JSON string
    try:
        # Check if payload is not None before conversion
        if payload is not None:
            payload = json.loads(payload)
    except json.JSONDecodeError:
        pass
    for _ in range(threads):
        thread = threading.Thread(target=ddos, args=(method, url, payload, headers))
        thread.start()


@click.command()
def version():
    rprint(f"[bold blue]version 0.3.3[/bold blue]")


@click.command()
@click.argument("url", type=click.STRING)
def map(url):
    # Step 1: Discover URLs on the website
    parsed_url = urlparse(url)
    domain = parsed_url.netloc
    discovered_urls = discover_urls(url)
    logname = f"logs/map_{domain}.json"
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
