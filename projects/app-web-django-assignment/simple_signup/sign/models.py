from django.db import models
from django.contrib.auth.forms import UserCreationForm
from django.contrib.auth.models import User
from django.forms import EmailField, CharField

# Create your models here.


class BaseRegisterForm(UserCreationForm):
    email = EmailField(label="Email")
    first_name = CharField(label="Имя")
    last_name = CharField(label="Фамилия")

    class Meta:
        model = User
        fields = ("username",
                  "first_name",
                  "last_name",
                  "email",
                  "password1",
                  "password2")
