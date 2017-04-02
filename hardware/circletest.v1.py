import RPi.GPIO as GPIO  
import time  



def go(pin):
        print("hello")
        GPIO.output(pin,GPIO.HIGH)
        time.sleep(5)  
        GPIO.output(pin,GPIO.LOW)

# to use Raspberry Pi board pin numbers  
# set up GPIO output channel

GPIO.setmode(GPIO.BOARD)

GPIO.setup(18, GPIO.OUT) 
GPIO.setup(13, GPIO.OUT) 
GPIO.output(18,GPIO.HIGH)
GPIO.output(13,GPIO.HIGH)
time.sleep(5)  
GPIO.output(18,GPIO.LOW)
GPIO.output(16,GPIO.LOW)



GPIO.cleanup()

