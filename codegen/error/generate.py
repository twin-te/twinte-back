import pathlib
import re

from definition import definition


def convert_pascal_case_to_snake_case(s: str) -> str:
    """Convert PascalCase to snake_case.

    cf. https://stackoverflow.com/questions/1175208/elegant-python-function-to-convert-camelcase-to-snake-case
    """
    s = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", s)
    s = re.sub("([a-z0-9])([A-Z])", r"\1_\2", s).lower()
    return s


def convert_pascal_case_to_upper_snake_case(s: str) -> str:
    s = convert_pascal_case_to_snake_case(s)
    s = s.upper()
    return s


def generate_go_error_definition(module: str) -> str:
    ret = f"""// Code generated by codegen/error/generate.py; DO NOT EDIT.

package {module}err

import "github.com/twin-te/twinte-back/apperr"

const (
"""

    error_codes = [error_code for error_code, _ in definition[module]]

    max_length = max(map(len, error_codes))

    for error_code in error_codes:
        ret += f'\tCode{error_code.ljust(max_length)} apperr.Code = "{module}.{error_code}"\n'

    ret += """)
"""

    return ret


def generate_go_error_code_map() -> str:
    ret = """// Code generated by codegen/error/generate.py; DO NOT EDIT.

package interceptor

import (
	"github.com/bufbuild/connect-go"
	"github.com/twin-te/twinte-back/apperr"
"""  # noqa: E101, W191

    for module in definition.keys():
        ret += f"""	{module}err "github.com/twin-te/twinte-back/module/{module}/err"
"""

    ret += """)

var AppErrorCodeToConnectErrorCode = map[apperr.Code]connect.Code{
"""
    for i, module in enumerate(definition):
        if i > 0:
            ret += "\n"

        for app_error_code, connect_error_code in definition[module]:
            max_length = max(
                [len(app_error_code) for app_error_code, _ in definition[module]]
            )
            num_spaces = max_length - len(app_error_code)
            ret += f"""	{module}err.Code{app_error_code}:{" " * num_spaces} connect.Code{connect_error_code},
"""

    ret += """}
"""

    return ret


def output_to_file(output_file_path: pathlib.Path, content: str) -> None:
    output_file_path.parent.mkdir(parents=True, exist_ok=True)
    with output_file_path.open("w") as f:
        f.write(content)
    print(f"output to {output_file_path.as_posix()}")


def main() -> None:
    root_path = pathlib.Path(__file__).parents[2]

    # Generate go error definition
    for module in definition.keys():
        output_go_file_path = root_path.joinpath(f"module/{module}/err/code_gen.go")
        content = generate_go_error_definition(module)
        output_to_file(output_go_file_path, content)

    # Generate go error code map
    content = generate_go_error_code_map()
    output_file_path = root_path.joinpath(
        "handler/common/interceptor/error_code_map_gen.go"
    )
    output_to_file(output_file_path, content)


if __name__ == "__main__":
    main()
