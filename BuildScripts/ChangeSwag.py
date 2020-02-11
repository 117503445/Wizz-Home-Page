# encoding:utf8
import os
import InitSwag


def read_all_lines(path: str):
    with open(path, 'r', encoding='utf-8')as f:
        lines = f.readlines()
        for i in range(len(lines)):
            lines[i] = lines[i].replace('\n', '')
            lines[i] = lines[i].replace('\r', '')
        return lines


def write_all_texts(path: str, content: list):
    with open(path, 'w', encoding='utf-8')as f:
        text = '\n'.join(content)
        f.write(text)


if __name__ == "__main__":

    c = input(
        'Input 0 -> Change to localhost:8080\nInput 1 -> Change to ali.117503445.top:8080\nInput 2 -> other\n')

    if c == '0':
        s = '// @host localhost:8080'
    elif c == '1':
        s = '// @host ali.117503445.top:8080'
    elif c == '2':
        s = '// @host ' + input()

    with open('../main.go', 'r', encoding='utf-8') as f:
        lines = f.readlines()
    lines = read_all_lines('../main.go')
    for i in range(len(lines)):
        if '// @host' in lines[i]:
            lines[i] = s
    write_all_texts('../main.go', lines)

    InitSwag.init_swag()
