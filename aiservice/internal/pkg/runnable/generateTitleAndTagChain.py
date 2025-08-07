from langchain.callbacks.streaming_stdout import StreamingStdOutCallbackHandler
from langchain.callbacks.stdout import StdOutCallbackHandler
from langchain.agents import AgentType, initialize_agent, load_tools, AgentExecutor
from langchain.tools import tool
from langchain.tools import StructuredTool
from langchain.prompts import PromptTemplate
from langchain.schema.runnable import RunnableLambda, RunnablePassthrough, RunnableBranch
import os
import yaml
from langchain_openai import ChatOpenAI
from langchain.output_parsers import PydanticOutputParser
from pydantic import field_validator, BaseModel, Field

from internal.pkg.runnable import llm


from typing import Dict, List, Any

import asyncio
# ChatOpenAI、agent、链等，本身是线程安全的，但是如果附带记忆机制，或含跨线程callbacks，则不是线程安全的。

class OutputModel(BaseModel):
    thoughts:str = Field(description="思考过程")
    title:str = Field(description="拟定的博客标题")
    tags:List[str] = Field(description="拟定的博客标签")
    @field_validator('title')
    @classmethod
    def ensure_title(cls, v:Any):
        if len(v) > 100 or len(v) < 1:
            raise ValueError("(╯‵□′)╯︵┻━┻博客标题长度需要在1～100之间")
        return v

class GenerateTitleAndTagModel:
    def __init__(self):
        # tools = load_tools(["ddg-search"], llm=llm.llm)
        # tools = []

        # agent = initialize_agent(
        #     tools,
        #     llm.llm,
        #     agent=AgentType.CHAT_ZERO_SHOT_REACT_DESCRIPTION,  # 专为 Chat 模型优化的 Agent
        #     handle_parsing_errors=True,
        #     verbose=True
        # )
        config_path = os.path.join(os.path.curdir, "..", "configs", "aiservice.yaml")
        with open(config_path, encoding="UTF-8") as config_file:
            config = yaml.load(config_file, yaml.FullLoader)
        llm_config = config['llm']

        llm = ChatOpenAI(
            model= llm_config['model'],
            openai_api_key= llm_config['api-key'],
            base_url= llm_config['url'],
            temperature=0.7
            # streaming=True,  # 启用流式输出
            # callbacks=[FormattedStreamingCallback()],
        )

        outputParser = PydanticOutputParser(pydantic_object=OutputModel)
        instruction = outputParser.get_format_instructions()

        print(instruction)
        template = PromptTemplate(
            input_variables=["prompt", "content"],
            template="""你是一个博客助手，现在需要为用户拟定一个1～100字符的博客标题，并为博客添加合适的标签。给出你的思考过程。
            用户的要求是：<request>{prompt}</request>，用户的博客内容是：<content>{content}</content>。
            {instruction}""",
            partial_variables={"instruction":instruction}
        )

        self.chain = template | llm | outputParser #| agent #| outputParser

    async def ainvoke(self, prompt, content):
        return await self.chain.ainvoke({"prompt":prompt, "content":content})
    def invoke(self, prompt, content):
        return self.chain.invoke({"prompt":prompt, "content":content})
        # for chunk in :
        #     print(chunk.content, flush=True)
        # async for chunk in self.chain.astream({"prompt":prompt, "content":content}):
        #     print(chunk.content, flush=True)
            # await asyncio.sleep(0)

GenerateTitleAndTagChain = GenerateTitleAndTagModel()
