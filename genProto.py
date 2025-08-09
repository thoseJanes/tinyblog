import subprocess
import os
import inspect
from typing import List

def printCmakeError(result):
    print("output:")
    print(result.stdout)
    if result.returncode != 0:
        print("Command failed.")
        print("error:")
        print(result.returncode)
        print(result.stderr)
        exit(1)

def copyTo(f, t):
    copy_cmd = ["cp", f, t]
    result = subprocess.run(copy_cmd, capture_output=True, text=True)
    printCmakeError(result)

def runIfExistsGen(proto_path):
    url = os.path.join(proto_path, "gen")
    if os.path.exists(url + ".py"):
        cmd = ["python", url + ".py"]
        result = subprocess.run(cmd, capture_output=True, text=True)
        printCmakeError(result)
    if os.path.exists(url + ".go"):
        cmd = ["go", "generate", url + ".go"]
        result = subprocess.run(cmd, capture_output=True, text=True)
        printCmakeError(result)


def getProtoFile(item, version):
    return os.path.join(os.path.dirname(__file__), "proto", item, version, f"{item}.proto")

def getProtoPath(proto_item, to_item, version):
    return os.path.join(os.path.dirname(__file__), to_item, "pkg", "proto", proto_item, version)

def copyProtoFileTo(file_item:str, to_items:List[str], version:str):
    proto_file = getProtoFile(file_item, version)
    for item in to_items:
        proto_path = getProtoPath(file_item, item, version)
        copyTo(proto_file, proto_path)
        runIfExistsGen(proto_path)

if __name__ == "__main__":
    item = "aiservice"
    to_items = ["aiservice", "tinyblog"]
    version = "v1"
    copyProtoFileTo(item, to_items, version)

    item = "tinyblog"
    to_items = ["tinyblog"]
    version = "v1"
    copyProtoFileTo(item, to_items, version)