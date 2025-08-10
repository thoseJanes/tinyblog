
from langchain.callbacks.stdout import StdOutCallbackHandler
from langgraph.prebuilt import create_react_agent
#from langgraph.prebuilt import create_react_agent
from langchain.tools import tool
from langchain.tools import StructuredTool
from langchain.prompts import ChatPromptTemplate, PromptTemplate
from langchain.schema.runnable import RunnableLambda, RunnablePassthrough, RunnableBranch
from langchain.output_parsers import PydanticOutputParser
from pydantic import field_validator, BaseModel, Field
from typing import Dict, List, Any

from langchain_core.messages import SystemMessage,HumanMessage

from internal.pkg.runnable.llm import GetLanguageModel
from pymysql.cursors import Cursor
from internal.pkg.mysql.mysql import MysqlPool
import asyncio

class SearchPostsOutput(BaseModel):
    evaluation: str = Field(description="简要的用户需求贴合程度的分析")
    ids: List[str] = Field(description="相关博客的id")


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
            description="输入查询语句（注意加分号），输出字符串作为结果。存储博客的数据库名称为post_table，其中包含title, content, id, updatedAt, createdAt等字段。"
        ))

        outputParser = PydanticOutputParser(pydantic_object=SearchPostsOutput)
        instruction = outputParser.get_format_instructions()
        self.template = PromptTemplate(
            template="""你是一个有认知能力的搜索助手，你需要根据用户的描述在数据库中搜索相关博客，返回相关博客的id，并给出一个简要的用户需求贴合程度的分析。
            注意，你只拥有数据库中表post_table的select权限，且每次搜索时必须限制搜索结果的数量最大为100，
            如非必要，不要搜索content字段，防止占用太多数据库资源。
            
            你可以进行多轮迭代使用工具来解决问题
            当你决定输出最终结果给用户时，输出的格式为：<贴合程度分析>id,id,id...，注意不要有其它任何最终结果输出\n
            如果没有强调，你需要从意义而非字面上理解用户的需求，同时运用工具和你的认知能力，查看工具获得的内容来返回结果，例子是：
            <example>
            输入：内容包含中文的博客。
            输出：<我查看了这些博客的内容，它们的内容中使用了中文>id1,id2,id3...
            输入：内容包含“中文”的博客。
            输出：<我查看了这些博客的内容，它们的内容中包含有“中文”字符>id1,id2,id3...
            </example>
            用户的要求是：<request>{input}</request>\n
            """,
            input_variables=["input"]
        )

        agent = create_react_agent(
            GetLanguageModel(streaming=True),
            tools,
            debug=True
        )
        self.chain = agent
    async def ainvoke(self, request):
        content=self.template.invoke({"input":request}).text
        print(content)
        return await self.chain.ainvoke({"messages":{"role":"human", "content":content}})

SearchPostsChain = SearchPostsModel()
