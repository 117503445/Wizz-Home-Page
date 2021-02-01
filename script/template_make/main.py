from htutil import file
import json
from pathlib import Path
import toml


def template_make(raw: str, cfg: dict) -> str:
    for d in cfg:
        raw = raw.replace(cfg[d], '${'+d+'}')
    return raw


def main():
    config = toml.loads(file.read_all_text('config.toml'))

    path_in = Path('in') / config['in']['path']

    raw = file.read_all_text(path_in / 'raw.txt')
    cfg = json.loads(file.read_all_text(path_in / 'cfg.json'))

    template = template_make(raw, cfg)

    path_out = Path('out')
    file.write_all_text(path_out / 'template.txt', template)


if __name__ == '__main__':
    main()
