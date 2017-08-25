start a server
#!/bin/bash

nohup ../spidertest >../spider.log 2>&1 &


kill a server
#!/bin/bash

ps axw | grep "../spidertest" | grep -v grep |awk '{print("kill", $1)}' | bash
