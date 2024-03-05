import click
from rich import print as rprint
import src.commands as commands

banner = f"""
         ___   ___   __      __
        /  /  /  /  /  \    /  /
       /  /  /  /  /    \  /  /
      /  /  /  /  /   \  \/  /
  ___/  /  /  /  /  /  \    /
 |_____/  /__/  /__/    \__/
  
"""


@click.group()
def cli():
    pass


cli.add_command(commands.scan)
cli.add_command(commands.attack)
cli.add_command(commands.map)
cli.add_command(commands.check)
cli.add_command(commands.version)

if __name__ == "__main__":
    rprint(f"[bold][red]{banner}[/red][/bold]")
    cli()
