
import os
import yaml
from langchain_openai import ChatOpenAI

config_path = os.path.join(os.path.curdir, "..", "configs", "aiservice.yaml")
with open(config_path, encoding="UTF-8") as config_file:
    config = yaml.load(config_file, yaml.FullLoader)
llm_config = config['llm']

llm = ChatOpenAI(
    model= llm_config['model'],
    openai_api_key= llm_config['api-key'],
    base_url= llm_config['url'],
    streaming=True,  # 启用流式输出
    # callbacks=[FormattedStreamingCallback()],
)