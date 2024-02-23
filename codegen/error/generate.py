import csv
import pathlib

module_name_to_error_codes: dict[str, list[str]] = {}
error_code_to_description: dict[str, str] = {}


def generate_go_code(module_name: str) -> str:
    error_codes = module_name_to_error_codes[module_name]

    ret = f"""// Code generated by codegen/error/generate.py; DO NOT EDIT.

package {module_name}err

import "github.com/twin-te/twinte-back/apperr"

const (
"""

    max_length = max(map(len, error_codes))
    for error_code in error_codes:
        ret += f'\tCode{error_code.ljust(max_length)} apperr.Code = "{module_name}.{error_code}"\n'

    ret += """)
"""

    return ret


def generate_markdown_list() -> str:
    ret = """<!-- Code generated by codegen/error/generate.py; DO NOT EDIT. -->

## Error Code List

"""

    for module_name, error_codes in module_name_to_error_codes.items():
        ret += f"""### {module_name}
| Error Code | Description |
| :---: | :---: |
"""
        for error_code in error_codes:
            description = error_code_to_description[error_code]
            ret += f"| `{module_name}.{error_code}` | {description} |\n"
        ret += "\n"

    return ret


def load_definition() -> None:
    definition_file_path = pathlib.Path(__file__).parent.joinpath("definition.csv")

    with definition_file_path.open("r") as f:
        reader = csv.reader(f)
        rows = [row for row in reader][1:]  # exclude header

    for module_name, error_code, description in rows:
        if module_name not in module_name_to_error_codes:
            module_name_to_error_codes[module_name] = []
        module_name_to_error_codes[module_name].append(error_code)
        error_code_to_description[error_code] = description

    print(f"load definition from {definition_file_path.as_posix()}")


def output_to_file(output_file_path: pathlib.Path, content: str) -> None:
    output_file_path.parent.mkdir(parents=True, exist_ok=True)
    with output_file_path.open("w") as f:
        f.write(content)
    print(f"output to {output_file_path.as_posix()}")


def main() -> None:
    load_definition()

    # Go Code
    for module_name in module_name_to_error_codes:
        output_go_file_path = (
            pathlib.Path(__file__)
            .parents[2]
            .joinpath(f"module/{module_name}/err/code.gen.go")
        )
        content = generate_go_code(module_name)
        output_to_file(output_go_file_path, content)

    # Markdown List
    output_markdown_list_path = pathlib.Path(__file__).parent.joinpath("list.gen.md")
    content = generate_markdown_list()
    output_to_file(output_markdown_list_path, content)


if __name__ == "__main__":
    main()
