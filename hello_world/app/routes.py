from app import app
from app.forms import NoteForm
from flask import render_template, flash, redirect, url_for

@app.route('/')
@app.route('/index')
def index():
    return render_template('index.html', title='boop', noun='world')

@app.route('/easter')
def easter():
    return render_template('easter.html', body="If I write something here will it print?")

@app.route('/note', methods=['GET', 'POST'])
def submit_note():
    form = NoteForm()
    if form.validate_on_submit():
        flash('Message {} received! Redirecting ...'.format(
            form.message.data))
        return redirect(url_for('index'))
    return render_template('note.html', title='Leave a note', form=form)
