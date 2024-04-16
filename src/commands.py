import click
import socket
import threading
from rich import print as rprint
from rich.console import Console
import click_completion
import logging
from urllib.parse import urlparse

from .utils.function import *
from .utils.exploit import *

click_completion.init()

console = Console()


@click.command("scan")
@click.argument("target", type=click.STRING)
def scan(target):
    try:
        ip_address = socket.gethostbyname(target)
        rprint(
            f'Scanning ports for [bold]{target}[/bold] ([cyan]{ip_address}[/cyan])...')
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
            logname = f"logs/scan_{target}.json"
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
@click.argument("target", type=click.STRING)
@click.argument("port", type=click.INT)
@click.option("--threads", default=100, help="Number of threads for DDoS attack")
def attack(target, port, threads):
    rprint(
        f'Initiating DDoS attack on [bold]{target}[/bold] on port [cyan]{port}[/cyan] with [cyan]{threads}[/cyan] threads...')
    for _ in range(threads):
        thread = threading.Thread(target=ddos, args=(target, port))
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


@click.command("check")
@click.argument("url", type=click.STRING)
@click.option("--sql", is_flag=True, default=False, help="Check for SQL injection vulnerability")
@click.option("--xss", is_flag=True, default=False, help="Check for XSS vulnerability")
@click.option("--conf", is_flag=True, default=False, help="Check for insecure configuration vulnerability")
@click.option("--dirtv", is_flag=True, default=False, help="Check for directory traversal vulnerability")
@click.option("--rcev", is_flag=True, default=False, help="Check for remote code execution vulnerability")
@click.option("--fuv", is_flag=True, default=False, help="Check for file upload vulnerability")
def check(url, sql, xss, conf, dirtv, rcev, fuv):
    if url:
        check_vulnerabilities(url, sql, xss, conf, dirtv, rcev, fuv)
    else:
        rprint("[white bold]No URL provided.[/white bold]")


@click.command("sinject")
@click.argument("url", type=click.STRING)
@click.option("--method", default='GET', help="Request method")
@click.option("--p", default='GET', help="The parameter to do SQL injection")
def sql_injection(url, method, p):
    if url:
        exploit_sql_injection(url, method, p)

    else:
        rprint("[white bold]No URL provided.[/white bold]")


@click.command("gdep")
@click.argument("url", type=click.STRING)
def get_dependencies(url):
    try:
        dependencies = discover_dependencies(url)

        rprint("[white bold]Dependencies Lists...[/white bold]")
        rprint(f"[red]{dependencies}[/red]")

    except requests.exceptions.RequestException as e:
        rprint(f"[red bold]Error fetching website: {e}[/red bold]")
        return None
