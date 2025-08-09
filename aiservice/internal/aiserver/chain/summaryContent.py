from langchain.prompts import PromptTemplate
from internal.pkg.runnable.llm import GetLanguageModel
import asyncio

class SummaryContentModel:
    def __init__(self):
        template = PromptTemplate(
            input_variables=["content"],
            template="""你是一个博客助手，现在需要给用户写的博客生成一个300字之内的内容描述或摘要。
            只输出最终摘要或者描述，不要输出其它任何内容。
            用户的博客内容是：<content>{content}</content>。""",
        )

        llm = GetLanguageModel(streaming=True)

        self.chain = template | llm

    async def astream(self, content):
        async for chunk in self.chain.astream({"content":content}):
            yield chunk.content
    def invoke(self, content):
        return self.chain.invoke({"content":content})

SummaryContentChain = SummaryContentModel()
