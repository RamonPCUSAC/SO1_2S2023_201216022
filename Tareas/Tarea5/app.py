from flask import Flask
import os
import signal
import sys
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

@app.route('/')
def hello_world():
    return 'Hola Mundo 201216022'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
