#!/usr/bin/env python3
#
# Author : easytojoin@163.com (Jok)
# Date   : Thu May 30 14:58:08 CST 2024
#

import schedule
import time

def hello():
    print("Hello, world!")

schedule.every(5).seconds.do(hello)

while True:
    print("Run")
    schedule.run_pending()
    time.sleep(1)
