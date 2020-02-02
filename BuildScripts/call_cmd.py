import os
import subprocess
import sys


def call(cmd):
    p = os.path.split(sys.path[0])[0]
    print(p)
    s = subprocess.run(cmd, cwd=str(p))
    print(s)
