from addon import server

system = server.registerSystem()


def helloworld():
    system.executeCommand("/say helloworld")


system.update(helloworld)
