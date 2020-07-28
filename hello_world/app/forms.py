from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField
from wtforms.validators import DataRequired

class NoteForm(FlaskForm):
    name = StringField('Name')
    message = StringField('Message', validators=[DataRequired()])
    submit = SubmitField('Submit')
