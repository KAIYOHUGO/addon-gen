from addon import server, serverevent

system = server.registerSystem(0, 0)


def helloworld():
    b = system.createEventData(serverevent.displayChatEvent)
    b.data.message = "hello world"
    system.broadcastEvent(serverevent.displayChatEvent, b)


system.update(helloworld)
