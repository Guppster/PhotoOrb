import RPi.GPIO as GPIO  
import time  
from flask import Flask
app=Flask(__name__)


def go(pin,t):
        GPIO.setup(pin, GPIO.OUT) 
        GPIO.output(pin,GPIO.HIGH)
        time.sleep(t)
        GPIO.output(pin,GPIO.LOW)
        
def front(t):
        for i in range(t):
                go(16, 0.1)
                go(11, 0.1)

def doturn(rev,diff,movement):
        diff += 0.1
        if movement == "left":
                for i in range(rev):
                        go(16, diff)
                        go(11, 0.1)
        elif movement == "right":
                for i in range(rev):
                        go(16, 0.01)
                        go(11, diff)
                
moves = {"right":[11,13],
        "left":[16,18]
        }

@app.route('/move/<string:movement>/<int:t>')
def move(movement,t):
        GPIO.setmode(GPIO.BOARD)
        if movement == 'front':
                front(t)
                GPIO.cleanup()
                return "moving"
        elif movement in moves:
                for i in moves[movement]:
                        go(i,t)
                GPIO.cleanup()
                return "moving"
        else:
                return "couldn't find command"


@app.route('/turn/<string:movement>/<int:rev>/<float:diff>')
def turn(movement, rev, diff):
        GPIO.setmode(GPIO.BOARD)
        doturn(rev,diff,movement)
        GPIO.cleanup()
        return "turn"

if __name__ == "__main__":
        app.run(debug=True)

