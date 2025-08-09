import os
import yaml

config_file = open(os.path.join(os.path.dirname(__file__), "aiservice.yaml"), encoding="UTF-8")
Config = yaml.load(config_file, yaml.FullLoader)

config_file.close()
