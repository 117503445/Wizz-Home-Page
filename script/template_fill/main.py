from htutil import file
import json
from pathlib import Path
import toml


def make_template(template: str, value: dict) -> str:
    for d in value:
        template = template.replace('${'+d+'}', value[d])
    return template


def main():

    config = file.read_all_text('config.toml')
    config = toml.loads(config)

    path_in = Path('in') / config['in']['path']
    template = file.read_all_text(path_in / 'template.txt')
    path_out = Path('out')
    for p in path_in.glob('*.json'):
        in_value = json.loads(file.read_all_text(p))
        result = make_template(template, in_value)
        # print(result)
        file.write_all_text(
            path_out / ('.'.join(p.name.split('.')[:-1]) + '.txt'), result)


if __name__ == '__main__':
    main()
