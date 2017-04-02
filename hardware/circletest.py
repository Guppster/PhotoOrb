import Rpi.GPIO as GPIO
import time

def go(pin):
    GPIO.output(pin, GPIO.HIGH)
    time.sleep(1)
    GPIO.output(pin, GPIO.LOW)



GPIO.setmode(GPIO.BOARD)

GPIO.setup(11, GPIO.OUT) #left wheel forward
GPIO.setup(13, GPIO.OUT) #left wheel back
go(11)
go(13)

GPIO.cleanup()

##We have to reset after each pair

GPIO.setmode(GPIO.BOARD)
GPIO.setup(16, GPIO.OUT)
GPIO.setup(18, GPIO.OUT)

go(16)
go(18)

GPIO.cleanup()


