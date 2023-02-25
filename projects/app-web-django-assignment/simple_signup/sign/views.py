from django.contrib.auth.models import User
from .models import BaseRegisterForm
from django.views.generic.edit import CreateView
from django.shortcuts import render
from django.views.generic import TemplateView
from django.contrib.auth.mixins import LoginRequiredMixin

# Create your views here.


class IndexView(LoginRequiredMixin, TemplateView):
    template_name = 'protect/index.html'


class BaseRegisterView(CreateView):
    model = User
    form_class = BaseRegisterForm
    success_url = '/'
