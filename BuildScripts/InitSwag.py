import os
import subprocess
import sys
import call_cmd


def init_swag():
    call_cmd.call('swag i')


if __name__ == "__main__":
    init_swag()
