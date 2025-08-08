from concurrent import futures
import grpc
import pkg.proto.aiservice.v1.aiservice_pb2_grpc as pb 
from pkg.proto.aiservice.v1.aiservice_pb2 import PromptContentRequest, GenerateTitleAndTagResponse, PromptRequest, PolishContentResponse, SearchPostsResponse
from internal.aiserver.chain.generateTitleAndTag import GenerateTitleAndTagChain, GenerateTitleAndTagOutput
from internal.aiserver.chain.polishContent import PolishContentChain, PolishContentModel
from internal.aiserver.chain.searchPosts import SearchPostsChain, SearchPostsModel

import asyncio

class AIService(pb.AIService):
    async def generateTitleAndTag(self, request:PromptContentRequest, context:grpc.ServicerContext):
        print("generateTitleAndTag requested")
        resp:GenerateTitleAndTagOutput = await GenerateTitleAndTagChain.ainvoke(request.prompt, request.content)
        return GenerateTitleAndTagResponse(title=resp.title, tags=resp.tags)
    async def polishContent(self, request:PromptContentRequest, context:grpc.ServicerContext):
        print("polishContent requested")
        async for chunk in PolishContentChain.astream(request.prompt, request.content):
            yield PolishContentResponse(contentChunk=chunk)
    async def searchPosts(self, request:PromptRequest, context:grpc.ServicerContext):
        print("searchPosts requested")
        messages = await SearchPostsChain.ainvoke(request.prompt)
        result:str = messages["messages"][-1].content
        try:
            slices = result.split(">")
            evaluation = slices[0][1:]
            id_list = slices[1].split(",")
            return SearchPostsResponse(evaluation=evaluation, ids=id_list)
        except:
            return SearchPostsResponse(evaluation="查找失败！最终信息：{result}", ids=[])

async def serve():
    #service中的方法需要是线程安全的。
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    pb.add_AIServiceServicer_to_server(AIService(), server)
    server.add_insecure_port('[::]:50051')
    await server.start()
    print("Server started, listening on port 50051")
    await server.wait_for_termination()

if __name__ == '__main__':
    asyncio.run(serve())