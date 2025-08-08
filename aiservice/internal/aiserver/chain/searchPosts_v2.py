
from langchain.callbacks.stdout import StdOutCallbackHandler
from langchain.agents import AgentType, initialize_agent, load_tools, AgentExecutor, create_react_agent
#from langgraph.prebuilt import create_react_agent
from langchain.tools import tool
from langchain.tools import StructuredTool
from langchain.prompts import PromptTemplate
from langchain.schema.runnable import RunnableLambda, RunnablePassthrough, RunnableBranch
from langchain.output_parsers import PydanticOutputParser
from pydantic import field_validator, BaseModel, Field
from typing import Dict, List, Any

from internal.pkg.runnable.llm import GetLanguageModel
from pymysql.cursors import Cursor
from internal.pkg.mysql.mysql import MysqlPool
import asyncio
    
def query(sentence) -> str:
    with MysqlPool.connection() as conn:
        with conn.cursor() as cursor:
            try:
                cursor:Cursor
                cursor.execute(sentence)
                return f"{cursor.fetchall()}"
            except Exception as err:
                return f"{err}"


class SearchPostsModel:
    def __init__(self):
        # tools = load_tools(["ddg-search"], llm=llm.llm)
        tools = []
        tools.append(StructuredTool.from_function(
            func=query,
            name="查询数据库",
            description="输入查询语句（注意加分号），输出字符串作为结果"
        ))

        agent = initialize_agent(
            tools,
            GetLanguageModel(streaming=True),
            agent=AgentType.CHAT_ZERO_SHOT_REACT_DESCRIPTION,  # 专为 Chat 模型优化的 Agent
            handle_parsing_errors=True,
            verbose=True,
            max_iterations=3
        )

        template = PromptTemplate(
            input_variables=["prompt", "memory"],
            template="""你是一个搜索助手，你需要根据用户的描述在数据库中搜索相关博客，返回相关博客的Id，并给出一个简要的搜索结果与用户需求贴合程度的分析。
            存储博客的数据库名称为post_table，其中包含title, content, postId, Id, updatedAt, createdAt等字段。
            注意，你只拥有数据库中表post_table的select权限，且每次搜索时必须限制搜索结果的数量最大为100，
            如非必要，不要搜索content字段，防止占用太多数据库资源。
            
            你可以进行多轮迭代使用工具来解决问题，当你决定继续获取信息时，在输出开头使用<continue>;当你决定输出结果给用户时，在输出开头使用<output>。在4轮迭代内需要返回最终结果。
            输出的格式为：<output><贴合程度分析>postid:postid:postid...\n
            输出的示例为：
            <example>
            输入：给我最让人震惊的一篇博客。
            输出：<output><你要求找到最能让你震惊的博客，但是没有说出哪类博客最能让你震惊，以下结果是一般而言能让人震惊的博客>34:43:56:88:204
            </example>
            用户的要求是：<request>{request}</request>\n
            当前操作的历史是：<memory>{memory}</memory>
            """,
        )

        self.chain = template | agent #| agent #| outputParser

    async def astream(self, request):
        memory = []
        present:str = ""
        outputLoop = False
        for i in range(5):
            if outputLoop:
                yield "<end>"
                break
            async for chunk in self.chain.astream({"request":request, "memory":memory}):
                if outputLoop:
                    yield chunk
                else:
                    present += chunk
                    if len(present) > 10 and not outputLoop:
                        if present.startswith("<output>"):
                            outputLoop = True
                            yield present
            memory.append([i+1, present])
    def invoke(self, request):
        return self.chain.invoke({"request":request})
        # for chunk in :
        #     print(chunk.content, flush=True)
        # async for chunk in self.chain.astream({"prompt":prompt, "content":content}):
        #     print(chunk.content, flush=True)
            # await asyncio.sleep(0)

SearchPostsChain = SearchPostsModel()
