# addon-gen

## support language

- [x] python [pypi](https://pypi.org/project/addongen/)
    - [x] binding
    - [x] event
    - [ ] component
    - [ ] object


## example

### python

```python
from addongen import server, serverevent

system = server.registerSystem(0, 0)


def helloworld():
    b = system.createEventData(serverevent.displayChatEvent)
    b.data.message = "hello world"
    system.broadcastEvent(serverevent.displayChatEvent, b)


system.update(helloworld)
```

## more infomation

- [wiki](https://github.com/KAIYOHUGO/addon-gen/wiki)