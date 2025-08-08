from langchain.prompts import PromptTemplate
from internal.pkg.runnable.llm import GetLanguageModel
import asyncio

class PolishContentModel:
    def __init__(self):
        template = PromptTemplate(
            input_variables=["prompt", "content"],
            template="""你是一个博客助手，现在需要给用户写的博客润色，你只能尽力润色然后输出润色结果，不要输出其它内容。
            用户的要求是：<request>{prompt}</request>，用户的博客内容是：<content>{content}</content>。""",
        )

        llm = GetLanguageModel(streaming=True)

        self.chain = template | llm

    async def astream(self, prompt, content):
        async for chunk in self.chain.astream({"prompt":prompt, "content":content}):
            yield chunk.content
    def invoke(self, prompt, content):
        return self.chain.invoke({"prompt":prompt, "content":content})

PolishContentChain = PolishContentModel()
