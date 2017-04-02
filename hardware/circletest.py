import Rpi.GPIO as GPIO
import time

def spin(pin, t):
    GPIO.output(pin, GPIO.HIGH)
    time.sleep(t)
    GPIO.output(pin, GPIO.LOW)

GPIO.setmode(GPIO.BOARD)

GPIO.setup(11, GPIO.OUT) #left wheel forward
GPIO.setup(12, GPIO.OUT) #left wheel back
GPIO.setup(13, GPIO.OUT) #right wheen forward`
GPIO.setup(15, GPIO.OUT) #right wheel backward

for e in range(0,10):
    for i in [11,12,13,15]:
        spin(i, 1);

GPIO.cleanup()


