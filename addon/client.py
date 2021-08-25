from addon.__eventbinding__ import __eventbinding__
from addon.__systemevent__ import __systemevent__

__js__ = window


class registerSystem(__systemevent__, __eventbinding__):
    def __init__(self, majorVersion: int, minorVersion: int):
        self.__js__ = __js__.client.registerSystem(majorVersion, minorVersion)

    def listenForEvent(self, event: str, fn: function):
        self.__js__.listenForEvent(event, fn)


# https://bedrock.dev/docs/stable/Scripting#log(Message)
def log(Message: str):
    __js__.client.log(Message)