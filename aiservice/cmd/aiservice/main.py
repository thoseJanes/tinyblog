import asyncio
from concurrent import futures
import grpc
from pkg.proto.aiservice.v1 import *
from internal.aiserver.chain import *
from internal.aiserver.aiserver import AIService

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