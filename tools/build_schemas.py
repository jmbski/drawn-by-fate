"""Simple script to run the auto-gen types from the schema"""

import argparse
import json
import os
import re
import shlex
import shutil
import subprocess

from collections import defaultdict
from dataclasses import dataclass, field
from string import Template
from typing import Any, Literal

DIR = os.path.dirname(__file__)
ROOT_DIR = os.path.dirname(DIR)

# Get local Config settings
CONFIG_PATH = os.path.join(DIR, "build-config.json")
CONFIG_FILE = {}
if os.path.exists(CONFIG_PATH):
    with open(CONFIG_PATH, "r", encoding="UTF-8") as fs:
        CONFIG_FILE = json.load(fs)
        if not isinstance(CONFIG_FILE, dict):
            CONFIG_FILE = {}
else:
    print(f"[WARNING]: Config path: '{CONFIG_PATH}' not found")

bname = os.path.basename


def read_file(
    path: str,
    mode: str = "r",
    encoding: str = "UTF-8",
    as_lines: bool = False,
    default_val: Any = None,
) -> str | dict | list:
    """Read data from a file

    Args:
        path (str): The path to the file
        mode (str, optional): IO mode to use. Defaults to "r".
        encoding (str, optional): Encoding format to use.
            Defaults to "latin-1".
        as_lines (bool, optional): If reading a regular text file,
            True will return the value of readlines() instead of read().
            Defaults to False.
        default_val (Any, optional): Value to return if the file is not found.
            Defaults to None.

    Returns:
        str | dict | list: The data read from the file. If the file is
            not found, returns an empty dict
    """

    default_val = default_val or {}

    if not os.path.exists(path):
        print(f"Warning: Path '{path}' does not exist")
        return default_val

    _, ext = os.path.splitext(path)

    with open(path, mode, encoding=encoding) as fs:
        match ext.lower():
            case ".json":
                return json.load(fs)
            case _:
                if as_lines:
                    return fs.readlines()
                else:
                    return fs.read()


def to_pascal_case(text: str) -> str:
    """Convert a string to PascalCase."""
    return re.sub(r"(_|-|^|\.)([a-z])", lambda m: m.group(2).upper(), text).replace(
        " ", ""
    )


def to_snake_case(text: str) -> str:
    """Convert a string to snake_case."""
    return (
        re.sub(r"([a-z])([A-Z])", r"\1_\2", text)
        .lower()
        .replace(" ", "_")
        .replace("-", "_")
    )


def parse_template(value: str, replacers: dict[str, str]) -> str:
    template = Template(value)
    return template.safe_substitute(replacers)


@dataclass
class Config:
    output_dir: str = ""
    schema_dir: str = ""
    constr_name: str = ""
    root_schemas: list[str] = field(default_factory=list)
    schema_data: list[tuple[str, str, str, str]] = field(default_factory=list)
    _mapper: dict[str, str] = field(default_factory=dict)

    def __post_init__(self) -> None:
        self._mapper = {
            "root": ROOT_DIR,
        }

        props = ["output_dir", "schema_dir", "constr_name"]
        for prop_name in props:
            prop = getattr(self, prop_name)
            if isinstance(prop, str):
                setattr(self, prop_name, parse_template(prop, self._mapper))

        self._mapper["schemas"] = self.schema_dir

        if self.root_schemas:
            for i, path in enumerate(self.root_schemas):
                parsed_path = parse_template(path, self._mapper)
                self.root_schemas[i] = parsed_path
                self.schema_data.append(self.get_schema_paths(parsed_path))

    def get_schema_paths(self, path: str) -> tuple[str, str, str, str]:
        schema_re = re.compile(r"\.schema\..*")
        schema_name = schema_re.sub("", bname(path))
        # schema_name = bname(path).replace(".json", "").replace(".schema", "")
        schema_data = read_file(path, default_val={})
        if isinstance(schema_data, dict):
            schema_name = schema_data.get("title", schema_name)

        package = to_snake_case(schema_name)
        go_path = os.path.join(self.output_dir, package, f"{package}.go")
        return path, schema_name, package, go_path

    def clear_dirs(self) -> None:
        if os.path.exists(self.output_dir):
            shutil.rmtree(self.output_dir)

        dirs = [self.output_dir] + [paths[3] for paths in self.schema_data]

        for dest_path in dirs:
            os.makedirs(os.path.dirname(dest_path), exist_ok=True)


Cfg = Config(**CONFIG_FILE)


parser = argparse.ArgumentParser()
parser.add_argument("--constructors", "-c", action="store_true")
parser.add_argument("--debug", "-d", action="store_true")
parser.add_argument("--verbose", "-v", action="store_true")
args = parser.parse_args()


def list_paths(path: str) -> list[str]:
    return [os.path.join(path, fname) for fname in os.listdir(path)]


def join_paths(root: str, paths: list[str]) -> list[str]:
    return [os.path.join(root, path) for path in paths]


SubReturn = Literal["out", "err", "both"]
""" String literal type representing the output choices for cmdx """


# pylint: disable=C0103
@dataclass
class SubReturns:
    """Enum class for SubReturn values"""

    OUT: SubReturn = "out"
    ERR: SubReturn = "err"
    BOTH: SubReturn = "both"


# pylint: enable=C0103


def run_command(
    cmd: list[str] | str, rtrn: SubReturn = "out", print_out: bool = True
) -> str | tuple[str, str]:
    """Executes a command and returns the output or error

    Args:
        cmd (list[str] | str): - A list of strings that make up the command or a string
            that will be split by spaces
        rtrn (SubReturn, optional): What outputs to return. If both, it will return a
            tuple of (stdout, stderr)Defaults to 'out'.

    Returns:
        str | tuple[str, str]: The output of the command or a tuple of (stdout, stderr)
    """

    if isinstance(cmd, str):
        cmd = shlex.split(cmd)

    try:
        process = subprocess.run(
            cmd,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            encoding="utf-8",
            check=True,
        )
    except subprocess.CalledProcessError as e:
        # Print and handle the errors here if needed
        process = e

    stdout = process.stdout
    stderr = process.stderr
    if print_out:
        if stdout:
            print(stdout)
        if stderr:
            print("\nERROR:\n", stderr)

    if rtrn == "out":
        return process.stdout
    if rtrn == "err":
        return process.stderr

    return process.stdout, process.stderr


def write_file(
    path: str,
    data: Any,
    mode: str = "w",
    encoding: str = "UTF-8",
    indent: int = 4,
) -> None:
    """Write text to a file

    Args:
        path (str): The path to the file
        data (Any): The data to write
        mode (str, optional): Read/write mode. Defaults to "w".
        encoding (str, optional): Encoding to write with. Defaults to ENCODING.
        indent (int, optional): Indent to apply if JSON. Defaults to 4.
        sort_keys (bool, optional): Whether to sort the output keys for YAML.
            Defaults to False.
    """

    _, ext = os.path.splitext(path)

    with open(path, mode, encoding=encoding) as fs:
        match ext.lower():
            case ".json":
                json.dump(data, fs, indent=indent)
            case _:
                if isinstance(data, list):
                    fs.writelines(data)
                else:
                    fs.write(data)
        print(f"Output written to {path}")


def flatten_json(data: object, parent_key: str = "") -> dict[str, object]:
    """Recursively flattens a nested JSON-like structure (dicts and lists) into dot-separated keys.

    Args:
        data: The JSON data (typically parsed into dicts, lists, and primitives).
        parent_key: The base key used for recursion (used internally).

    Returns:
        A flattened dictionary where keys are dot-separated paths to primitive values.
    """
    items: dict[str, object] = {}

    if isinstance(data, dict):
        for key, value in data.items():
            full_key = f"{parent_key}.{key}" if parent_key else key
            items.update(flatten_json(value, full_key))
    elif isinstance(data, list):
        for index, value in enumerate(data):
            full_key = f"{parent_key}.{index}" if parent_key else str(index)
            items.update(flatten_json(value, full_key))
    else:
        items[parent_key] = data

    return items


def write_log(stdout: str = "", stderr: str = "") -> None:
    if stdout:
        with open(os.path.join(DIR, "stdout"), "w", encoding="UTF-8") as fs:
            fs.write(stdout)
    if stderr:
        with open(os.path.join(DIR, "stderr"), "w", encoding="UTF-8") as fs:
            fs.write(stderr)


def compile_go() -> None:

    for schema_path, schema_name, package, go_path in Cfg.schema_data:
        print(f"Generating types for {schema_path}..")

        debug = " --debug=print-schema-resolving" if args.debug else ""

        cmd = f"quicktype --package {package}{debug} --top-level {to_pascal_case(schema_name)}  -s schema {schema_path} -o {go_path}"
        if args.verbose:
            print(cmd)
        if args.debug:
            out, err = run_command(cmd, SubReturns.BOTH, False)
            write_log(out, err)
        else:
            run_command(cmd)


def log_error(msg: str):
    print(msg)
    write_log(stderr=msg)


def l2str(items: list[str]) -> str:
    return "\n".join(items)


def compile_schemas() -> None:

    Cfg.clear_dirs()

    compile_go()

    # run_command(f"quicktype -s schema {schema_path} -o {ts_path} ")
    print("Types generated successfully")

    if args.constructors:
        print("Generating constructors")
        # run_command(f"go run {Cfg.constructors_path} {Cfg.output_dir}")


def main():
    compile_schemas()

    # print(json.dumps(TEMPLATES, indent=4))


if __name__ == "__main__":
    main()
