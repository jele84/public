from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
	return 'Hello, anonymous!'


@app.route('/<string:name>')
def greeting(name: str):
	return f'Hello, {name.capitalize()}!'

if __name__ == '__main__':
	app.run()
