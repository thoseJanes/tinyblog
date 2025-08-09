import subprocess
import os
import inspect

def printCmakeError(result):
    print("output:")
    print(result.stdout)
    if result.returncode != 0:
        print("Command failed.")
        print("error:")
        print(result.returncode)
        print(result.stderr)
        exit(1)


proto_file = "aiservice.proto"
proto_dir = os.path.dirname(__file__)

if os.path.exists(os.path.join(proto_dir, proto_file)):
    proto_file_name = proto_file.rsplit(".", maxsplit=1)[0]
    proto_file_suffix = proto_file.rsplit(".", maxsplit=1)[1]

    gen_command = ["python", 
                        "-m", "grpc_tools.protoc", 
                        "-I", proto_dir, 
                        f"--python_out={proto_dir}",
                        f"--grpc_python_out={proto_dir}",
                        f"--mypy_out={proto_dir}",
                        proto_file]# -S选项用于指定源代码目录。
    sed_import = ["sed",
                "-i",
                f"s/import {proto_file_name}_pb2/import pkg.proto.aiservice.v1.{proto_file_name}_pb2/",
                f"{proto_dir}/{proto_file_name}_pb2_grpc.py"]

    result = subprocess.run(gen_command, capture_output=True, text=True)
    printCmakeError(result)
    result = subprocess.run(sed_import, capture_output=True, text=True)
    printCmakeError(result)
    