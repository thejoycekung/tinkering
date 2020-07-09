from app import app
from flask import render_template

@app.route('/')
@app.route('/index')
def index():
    return render_template('index.html', title='boop', noun='world')

@app.route('/easter')
def easter():
    return render_template('easter.html', body="If I write something here will it print?")
