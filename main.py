import click
from rich import print as rprint
import src.commands as commands

banner = f"""

      /\__/
     / _ \ 
    ( (_| |
     \___/ 
      / _ \ 
     ( (_| | 
      \___/ 

   Jin - Your DDOS CLI Tools 

"""


@click.group()
def cli():
    pass


cli.add_command(commands.scan_ports)
cli.add_command(commands.attack)
cli.add_command(commands.version)

if __name__ == "__main__":
    rprint(f"[bold][red]{banner}[/red][/bold]")
    cli()
