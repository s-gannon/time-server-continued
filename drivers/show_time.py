#!/usr/bin/env python3

import re
import serial

sp = serial.Serial('/dev/serial0', 9600, timeout=1.0)
sp.flushInput()
while True:
    try:
        data = sp.read_until()
        print(re.sub('[0-9]', 'X', data.decode('ascii')))
    except serial.SerialTimeoutException:
        print('Timeout.')
        continue
    except KeyboardInterrupt:
        print('Keyboard Interrupt.')
        break
print('Exiting...')
sp.close()
