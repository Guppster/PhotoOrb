import RPi.GPIO as GPIO
import time
from flask import Flask
app=Flask(__name__)

def go(pin):
    GPIO.setmode(GPIO.BOARD)
    GPIO.setup(pin, GPIO.OUT)
    GPIO.output(pin, GPIO.HIGH)
    print("Running")
    time.sleep(1)
    GPIO.output(pin, GPIO.LOW)
    GPIO.cleanup()

move = {"right_back":11,
        "right_front":13,
        "left_back":16,
        "left_front":18}

@app.route('/servo/<string:movement>')
def servo(movement):
    for i in move.values():
        go(i)
    return "Yes?"

if __name__ == "__main__":
    app.run(debug=True)
