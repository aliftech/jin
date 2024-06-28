import click
from rich import print as rprint
import src.commands as commands
import sys

banner = f"""
         ___   ___   __      __
        /  /  /  /  /  \    /  /
       /  /  /  /  /    \  /  /
      /  /  /  /  /   \  \/  /
  ___/  /  /  /  /  /  \    /
 |_____/  /__/  /__/    \__/
  
"""

@click.group(invoke_without_command=True)
@click.pass_context
def cli(context):
    rprint(f"[bold][red]{banner}[/red][/bold]")
    if context.invoked_subcommand is None:
        click.echo(context.get_help())

        while True:
            try:
                # Prompt user for input
                user_input = input("Enter command: ").strip()
                
                # Handle input and exit if command is 'exit'
                if user_input.lower() in ["exit", "quit"]:
                    click.echo("Exiting...")
                    break

                # Split input into arguments and invoke the CLI
                args = user_input.split()
                if args:
                    # Simulate running command
                    sys.argv = ['main.py'] + args
                    cli.main(prog_name='main.py', standalone_mode=False)
            except (EOFError, KeyboardInterrupt):
                click.echo("\nExiting...")
                break

cli.add_command(commands.scan)
cli.add_command(commands.attack)
cli.add_command(commands.map)
cli.add_command(commands.version)
cli.add_command(commands.logs)

if __name__ == "__main__":
    cli()
