import RPi.GPIO as gpio
import time
from flask import Flask
app=Flask(__name__)

left = 16
right = 11

gpio.setmode(gpio.BOARD)

def go(duration):
	gpio.setup(left, gpio.OUT)
	gpio.setup(right, gpio.OUT)
	gpio.output(left, gpio.HIGH)
	gpio.output(right, gpio.HIGH)

	time.sleep(duration)

	gpio.output(left, gpio.LOW)
	gpio.output(right, gpio.LOW)

	gpio.cleanup()

if __name__ == "__main__":
	go(2)
