import os
import yaml
from langchain_openai import ChatOpenAI
from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from config.config import Config
# ChatOpenAI、agent、链等，本身是线程安全的，但是如果附带记忆机制，或含跨线程callbacks，则不是线程安全的。



def GetLanguageModel(temperature=0.7, streaming=False) -> ChatOpenAI:
    llm_config = Config['llm']

    llm = ChatOpenAI(
        model= llm_config['model'],
        openai_api_key= llm_config['api-key'],
        base_url= llm_config['url'],
        temperature=temperature,
        streaming=streaming,  # 启用流式输出
        callbacks=[StreamingStdOutCallbackHandler()],
    )
    
    return llm
