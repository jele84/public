from flask import Flask
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def hello():
	return '<h1>Hello, anonymous!</h1>'

@app.route('/time')
def time():
	return str(datetime.utcnow())

@app.route('/<string:name>')
def greeting(name: str):
	return f'<h1>Hello, {name.capitalize()}!</h1>'

if __name__ == '__main__':
	app.run()
