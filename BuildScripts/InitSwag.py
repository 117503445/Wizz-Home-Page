import os
import subprocess
import sys

def init_swag():
    p=os.path.split(sys.path[0])[0]
    print(p)
    cmd = 'swag i'
    s=subprocess.run(cmd,cwd=str(p))
    print(s)
   
    

if __name__ == "__main__":
    init_swag()
