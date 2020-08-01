from app import app, db
from app.models import Note
from app.forms import NoteForm
from flask import render_template, flash, redirect, url_for

@app.route('/')
@app.route('/index')
def index():
    query = db.session.execute("SELECT name, message, timestamp FROM note ORDER BY timestamp desc")
    entries = [dict(name=row[0], message=row[1], timestamp=row[2]) for row in query]
    return render_template('index.html', title='boop the world', noun='guestbook', entries=entries)

@app.route('/easter')
def easter():
    return render_template('easter.html', body="If I write something here will it print?")

@app.route('/note', methods=['GET', 'POST'])
def submit_note():
    form = NoteForm()
    if form.validate_on_submit():
        flash('Message {} received! Redirecting ...'.format(
            form.message.data))
        note = Note(name=form.name.data, message=form.message.data)
        db.session.add(note)
        db.session.commit()
        return redirect(url_for('index'))
    return render_template('note.html', title='Leave a note', form=form)
