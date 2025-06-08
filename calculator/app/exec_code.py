import subprocess


def exec_code(binary_path, input_path, output_path, err_path):
    with (
        open(input_path, "r") as input_file,
        open(output_path, "w") as output_file,
        open(err_path, "w") as err_file,
    ):
        try:
            # Execute the binary with the input file
            result = subprocess.run(
                [binary_path], stdin=input_file, stdout=output_file, stderr=err_file
            )
            return result.returncode
        except Exception as e:
            err_file.write(str(e))
            return -1

if __name__ == "__main__":
    sample_binary = "_sample/a.out"
    sample_input = "_sample/input.txt"
    sample_output = "_sample/output.txt"
    sample_error = "_sample/error.txt"
    exec_code(sample_binary, sample_input, sample_output, sample_error)