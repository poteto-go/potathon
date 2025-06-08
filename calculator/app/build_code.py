import subprocess
import os


def build_cpp(cppcode_path):

    # Define the output binary path
    binary_path = os.path.splitext(cppcode_path)[0] + ".out"

    try:
        # Compile the C++ code
        result = subprocess.run(
            ["g++", cppcode_path, "-o", binary_path],
            check=True,
            capture_output=True,
            text=True
        )
        print(f"Compilation successful: {result.stdout}")
        return binary_path
    except subprocess.CalledProcessError as e:
        print(f"Compilation failed: {e.stderr}")
        return None
    
if __name__ == "__main__":
    sample_cpp = "_fixture/sample_code.cpp"
    binary = build_cpp(sample_cpp)
    if binary:
        print(f"Binary created at: {binary}")
    else:
        print("Failed to create binary.")