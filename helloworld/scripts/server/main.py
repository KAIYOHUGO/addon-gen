from addon import server, serverevent

system = server.registerSystem()


def helloworld():
    b = system.createEventData("minecraft:display_chat_event")
    b.data.message = "helloworld"
    system.broadcastEvent("minecraft:display_chat_event", b)


system.update(helloworld)
