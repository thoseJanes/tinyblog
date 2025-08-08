from langchain.tools import tool
from internal.pkg.mysql.mysql import MysqlPool
from pymysql.cursors import Cursor
from internal.pkg.runnable.llm import GetLanguageModel
from langgraph.graph import MessagesState, StateGraph, START, END
from langchain_core.messages import SystemMessage, ToolMessage, HumanMessage
from typing import Literal
@tool("数据库查询工具")
def query(sentence) -> str:
    '''
    这是用于查询数据库内容的工具。输入Mysql查询语句（注意加分号），输出字符串作为结果。
    '''
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
        self.llm = GetLanguageModel(streaming=True).bind_tools([query])

        agent_builder = StateGraph(MessagesState)
        agent_builder.add_node("llmCall", self.llmCall)
        agent_builder.add_node("environment", self.toolNode)

        agent_builder.add_edge(START, "llmCall")
        agent_builder.add_conditional_edges(
            "llmCall",
            self.shouldContinue,
            {
                "Action": "environment",
                "End": END,
            },
        )
        agent_builder.add_edge("environment", "llmCall")
        self.agent = agent_builder.compile()

    def llmCall(self, state:MessagesState):
        print("\n====llmCall======\n", state, "\n==========\n")
        return {
            "messages":[
                self.llm.invoke(
                    [
                        SystemMessage(
                            content = """
                                你是一个搜索助手，你需要<important>使用工具</important>，查询mysql数据库，找到相关博客的Id，并给出一个简要的用户需求贴合程度的分析。
                                存储博客的数据库名称为post_table，其中包含title, content, id, updatedAt, createdAt等字段。
                                你只拥有数据库中表post_table的select权限，且每次搜索时必须限制搜索结果的数量最大为100，
                                如果能从title获取结果，可以减少对content字段的搜索来节省资源。
                                当查询结束，返回最终结果时，输出的格式为：<用户需求贴合程度分析>id,id,id...:<查询使用语句1><查询使用语句2>...，不要输出其它任何内容。\n
                                """
                        )
                    ] + state["messages"]
                )
            ]
        }
    def toolNode(self, state: dict):
        print("\n====toolNode======\n", state, "\n==========\n")
        result = []
        for tool_call in state["messages"][-1].tool_calls:
            observation = query.invoke(tool_call["args"])
            result.append(ToolMessage(content=observation, tool_call_id=tool_call["id"]))
        return {"messages": result}
    
    def shouldContinue(self, state: MessagesState) -> Literal["Action", "End"]:
        print("\n====shouldContinue======\n", state, "\n==========\n")
        messages = state["messages"]
        last_message = messages[-1]
        # If the LLM makes a tool call, then perform an action
        if last_message.tool_calls:
            return "Action"
        # Otherwise, we stop (reply to the user)
        print("---------stop----------")
        return "End"
    
    async def astream(self, prompt):
        messages = [HumanMessage(content=prompt)]
        async for chunk in self.agent.astream({"messages": messages}):
            yield chunk
    async def ainvoke(self, prompt):
        messages = [HumanMessage(content=prompt)]
        result = await self.agent.ainvoke({"messages": messages})
        return result
SearchPostsGraph = SearchPostsModel()